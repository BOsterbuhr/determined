# type: ignore
import os
import pathlib
import random
import sys
import typing

import numpy as np
import pytest
import torch

import determined as det
from determined import gpu, pytorch, workload
from tests.experiment import utils  # noqa: I100
from tests.experiment.fixtures import pytorch_onevar_model

# Apex is included only for GPU trials.
try:
    import apex  # noqa

    HAVE_APEX = True
except ImportError:  # pragma: no cover
    HAVE_APEX = False
    pass


def check_equal_structures(a: typing.Any, b: typing.Any) -> None:
    """
    Check that two objects, consisting of any nested structures of lists and
    dicts, with leaf values of tensors or built-in objects, are equal in
    structure and values.
    """
    if isinstance(a, dict):
        assert isinstance(b, dict)
        assert len(a) == len(b)
        for key in a:
            assert key in b
            check_equal_structures(a[key], b[key])
    elif isinstance(a, list):
        assert isinstance(b, list)
        assert len(a) == len(b)
        for x, y in zip(a, b):
            check_equal_structures(x, y)
    elif isinstance(a, torch.Tensor):
        assert isinstance(b, torch.Tensor)
        assert torch.allclose(a, b)
    else:
        assert a == b


class TestPyTorchTrial:
    def setup_method(self) -> None:
        # This training setup is not guaranteed to converge in general,
        # but has been tested with this random seed.  If changing this
        # random seed, verify the initial conditions converge.
        self.trial_seed = 17
        self.hparams = {
            "hidden_size": 2,
            "learning_rate": 0.5,
            "global_batch_size": 4,
            "lr_scheduler_step_mode": pytorch.LRScheduler.StepMode.MANUAL_STEP.value,
            "dataloader_type": "determined",
            "disable_dataset_reproducibility_checks": False,
        }

    def test_require_global_batch_size(self) -> None:
        bad_hparams = dict(self.hparams)
        del bad_hparams["global_batch_size"]
        with pytest.raises(
            det.errors.InvalidExperimentException, match="is a required hyperparameter"
        ):
            _ = create_trial_and_trial_controller(
                trial_class=pytorch_onevar_model.OneVarTrial,
                hparams=bad_hparams,
            )

    def test_onevar_single(self) -> None:
        """Assert that the training loss and validation error decrease monotonically."""
        trial, trial_controller = create_trial_and_trial_controller(
            trial_class=pytorch_onevar_model.OneVarTrial,
            hparams=self.hparams,
            trial_seed=self.trial_seed,
        )

        train_steps, metrics = trial_controller._train_with_steps(
            training_enumerator=enumerate(trial_controller.training_iterator),
            train_steps=[
                pytorch._TrainStep(step_type=pytorch._TrainStepType.TRAIN, unit=pytorch.Batch(100))
            ],
        )

        assert len(train_steps) == 1, "unexpected train step count"
        assert train_steps[0].limit_reached, "train step did not reach expected limit"
        assert len(metrics) == 100, "metrics length did not match input"

        for i in range(100):
            pytorch_onevar_model.OneVarTrial.check_batch_metrics(
                metrics[i],
                i,
                metric_keyname_pairs=(("loss", "loss_exp"), ("w_after", "w_exp")),
            )
        for older, newer in zip(metrics, metrics[1:]):
            assert newer["loss"] <= older["loss"]

    def test_training_metrics(self) -> None:
        trial, trial_controller = create_trial_and_trial_controller(
            trial_class=pytorch_onevar_model.OneVarTrialWithTrainingMetrics,
            hparams=self.hparams,
            trial_seed=self.trial_seed,
        )

        train_steps, metrics = trial_controller._train_with_steps(
            training_enumerator=enumerate(trial_controller.training_iterator),
            train_steps=[
                pytorch._TrainStep(step_type=pytorch._TrainStepType.TRAIN, unit=pytorch.Batch(100))
            ],
        )

        assert len(train_steps) == 1, "unexpected train step count"
        assert train_steps[0].limit_reached, "train step did not reach expected limit"
        assert len(metrics) == 100, "metrics length did not match input"
        for metric in metrics:
            assert "mse" in metric

    def test_nonscalar_validation(self) -> None:
        trial, trial_controller = create_trial_and_trial_controller(
            trial_class=pytorch_onevar_model.OneVarTrialWithNonScalarValidation,
            hparams=self.hparams,
            expose_gpus=True,
            trial_seed=self.trial_seed,
        )

        val_metrics = trial_controller._validate()

        assert "mse" in val_metrics

    def test_checkpointing_and_restoring(self, tmp_path: pathlib.Path) -> None:
        updated_hparams = {
            "lr_scheduler_step_mode": pytorch.LRScheduler.StepMode.STEP_EVERY_BATCH.value,
            **self.hparams,
        }
        self.checkpoint_and_restore(updated_hparams, tmp_path, (100, 100))

    @pytest.mark.skipif(not torch.cuda.is_available(), reason="no gpu available")
    @pytest.mark.gpu
    @pytest.mark.parametrize(
        "trial_class",
        [
            pytorch_onevar_model.OneVarApexAMPTrial,
            pytorch_onevar_model.OneVarAutoAMPTrial,
            pytorch_onevar_model.OneVarManualAMPTrial,
        ],
        ids=[
            "apex",
            "autocast",
            "manual",
        ],
    )
    def test_scaler_checkpointing_and_restoring(self, trial_class, tmp_path: pathlib.Path) -> None:
        if trial_class is pytorch_onevar_model.OneVarApexAMPTrial and not HAVE_APEX:
            pytest.skip("Apex not available")

        updated_hparams = {
            "global_batch_size": 1,
            **self.hparams,
        }

        tm_a, tm_b = utils.checkpoint_and_restore(
            hparams=updated_hparams, tmp_path=tmp_path, steps=(200, 200)
        )

        amp_metrics_test(trial_class, tm_a)
        amp_metrics_test(trial_class, tm_b)

    def test_restore_invalid_checkpoint(self, tmp_path: pathlib.Path) -> None:
        # Build, train, and save a checkpoint with the normal hyperparameters.
        checkpoint_dir = str(tmp_path.joinpath("checkpoint"))

        class CheckpointingCallback(pytorch.PyTorchCallback):
            def __init__(self):
                self.uuid = None

            def on_checkpoint_upload_end(self, uuid: str) -> None:
                self.uuid = uuid

        checkpoint_callback = CheckpointingCallback()

        def build_callbacks(self) -> typing.Dict[str, pytorch.PyTorchCallback]:
            return {"checkpoint": checkpoint_callback}

        setattr(pytorch_onevar_model.OneVarTrial, "build_callbacks", build_callbacks)

        # Trial A: run with 100 batches and checkpoint
        trial_A = run_pytorch_trial_controller_from_trial_impl(
            trial_class=pytorch_onevar_model.OneVarTrial,
            hparams=self.hparams,
            trial_seed=self.trial_seed,
            max_batches=100,
            min_validation_batches=100,
            min_checkpoint_batches=100,
            checkpoint_dir=checkpoint_dir,
        )

        assert checkpoint_callback.uuid is not None, "trial did not return a checkpoint UUID"

        # Trial A: restore from checkpoint with invalid hparams
        invalid_hparams = {**self.hparams, "features": 2}
        assert invalid_hparams != self.hparams

        with pytest.raises(RuntimeError):
            run_pytorch_trial_controller_from_trial_impl(
                trial_class=pytorch_onevar_model.OneVarTrial,
                hparams=invalid_hparams,
                trial_seed=self.trial_seed,
                max_batches=100,
                min_validation_batches=100,
                min_checkpoint_batches=sys.maxsize,
                checkpoint_dir=checkpoint_dir,
                latest_checkpoint=checkpoint_callback.uuid,
                steps_completed=trial_A._state.batches_trained,
            )

    def test_reproducibility(self) -> None:
        training_metrics = {"A": [], "B": []}
        validation_metrics = {"A": [], "B": []}

        # Trial A
        trial_A = run_pytorch_trial_controller_from_trial_impl(
            trial_class=pytorch_onevar_model.OneVarTrial,
            hparams=self.hparams,
            trial_seed=self.trial_seed,
            max_batches=1000,
            min_validation_batches=100,
            min_checkpoint_batches=sys.maxsize,
        )
        metrics_callback = pytorch_onevar_model.OneVarTrial.metrics_callback
        training_metrics["A"] = metrics_callback.training_metrics
        validation_metrics["A"] = metrics_callback.validation_metrics

        # Trial B
        trial_B = run_pytorch_trial_controller_from_trial_impl(
            trial_class=pytorch_onevar_model.OneVarTrial,
            hparams=self.hparams,
            trial_seed=self.trial_seed,
            max_batches=1000,
            min_validation_batches=100,
            min_checkpoint_batches=sys.maxsize,
        )
        metrics_callback = pytorch_onevar_model.OneVarTrial.metrics_callback
        training_metrics["B"] = metrics_callback.training_metrics
        validation_metrics["B"] = metrics_callback.validation_metrics

        assert len(training_metrics["A"]) == len(training_metrics["B"])
        for A, B in zip(training_metrics["A"], training_metrics["B"]):
            utils.assert_equivalent_metrics(A, B)

        assert len(validation_metrics["A"]) == len(validation_metrics["B"])
        for A, B in zip(validation_metrics["A"], validation_metrics["B"]):
            utils.assert_equivalent_metrics(A, B)

    def test_custom_eval(self) -> None:
        training_metrics = {"A": [], "B": []}  # type: typing.Dict
        validation_metrics = {"A": [], "B": []}  # type: typing.Dict

        trial_A = run_pytorch_trial_controller_from_trial_impl(
            trial_class=pytorch_onevar_model.OneVarTrial,
            hparams=self.hparams,
            trial_seed=self.trial_seed,
            max_batches=900,
            min_validation_batches=100,
            min_checkpoint_batches=sys.maxsize,
        )

        metrics_callback = pytorch_onevar_model.OneVarTrial.metrics_callback

        training_metrics["A"] = metrics_callback.training_metrics
        validation_metrics["A"] = metrics_callback.validation_metrics

        trial_B = run_pytorch_trial_controller_from_trial_impl(
            trial_class=pytorch_onevar_model.OneVarTrialCustomEval,
            hparams=self.hparams,
            trial_seed=self.trial_seed,
            expose_gpus=True,
            max_batches=900,
            min_validation_batches=100,
            min_checkpoint_batches=sys.maxsize,
        )

        metrics_callback = pytorch_onevar_model.OneVarTrial.metrics_callback
        training_metrics["B"] = metrics_callback.training_metrics
        validation_metrics["B"] = metrics_callback.validation_metrics

        for original, custom_eval in zip(training_metrics["A"], training_metrics["B"]):
            assert np.allclose(original["loss"], custom_eval["loss"], atol=1e-6)

        for original, custom_eval in zip(validation_metrics["A"], validation_metrics["B"]):
            assert np.allclose(original["val_loss"], custom_eval["val_loss"], atol=1e-6)

    def test_grad_clipping(self) -> None:
        training_metrics = {}
        validation_metrics = {}

        def make_workloads(tag: str) -> workload.Stream:
            trainer = utils.TrainAndValidate()

            yield from trainer.send(steps=1000, validation_freq=100)
            tm, vm = trainer.result()
            training_metrics[tag] = tm
            validation_metrics[tag] = vm

        controller = utils.make_pytorch_trial_controller_from_trial_impl(
            trial_class=pytorch_onevar_model.OneVarTrialGradClipping,
            hparams=self.hparams,
            trial_seed=self.trial_seed,
            max_batches=100,
            min_validation_batches=10,
            min_checkpoint_batches=sys.maxsize,
        )
        controller.run()

        updated_hparams = {"gradient_clipping_l2_norm": 0.0001, **self.hparams}
        controller = utils.make_pytorch_trial_controller_from_trial_impl(
            trial_class=pytorch_onevar_model.OneVarTrialGradClipping,
            hparams=updated_hparams,
            trial_seed=self.trial_seed,
            max_batches=100,
            min_validation_batches=10,
            min_checkpoint_batches=sys.maxsize,
        )
        controller.run()

        for idx, (original, clipped) in enumerate(
            zip(training_metrics["original"], training_metrics["clipped_by_norm"])
        ):
            if idx < 10:
                continue
            assert original["loss"] != clipped["loss"]

        updated_hparams = {"gradient_clipping_value": 0.0001, **self.hparams}
        controller = utils.make_pytorch_trial_controller_from_trial_impl(
            trial_class=pytorch_onevar_model.OneVarTrialGradClipping,
            hparams=updated_hparams,
            trial_seed=self.trial_seed,
            max_batches=100,
            min_validation_batches=10,
            min_checkpoint_batches=sys.maxsize,
        )
        controller.run()

        for idx, (original, clipped) in enumerate(
            zip(training_metrics["original"], training_metrics["clipped_by_val"])
        ):
            if idx < 10:
                continue
            assert original["loss"] != clipped["loss"]

    def test_per_metric_reducers(self) -> None:
        def make_workloads() -> workload.Stream:
            trainer = utils.TrainAndValidate()
            yield from trainer.send(steps=2, validation_freq=1, scheduling_unit=1)

        utils.run_pytorch_trial_controller_from_trial_impl(
            trial_class=pytorch_onevar_model.OneVarTrialPerMetricReducers,
            hparams=self.hparams,
            trial_seed=self.trial_seed,
            max_batches=2,
            min_validation_batches=1,
            min_checkpoint_batches=sys.maxsize,
        )

    def test_callbacks(self, tmp_path: pathlib.Path) -> None:
        checkpoint_dir = tmp_path.joinpath("checkpoint")
        latest_checkpoint = None
        steps_completed = 0

        controller = None

        hparams1 = dict(self.hparams)
        hparams1["global_batch_size"] = 2
        training_epochs = 2
        num_batches = (
            training_epochs
            * len(pytorch_onevar_model.OnesDataset())
            // hparams1["global_batch_size"]
        )

        def make_workloads1() -> workload.Stream:
            nonlocal controller
            assert controller.trial.counter.trial_startups == 1
            yield workload.train_workload(1, 1, 0, num_batches), workload.ignore_workload_response
            assert controller is not None, "controller was never set!"
            assert controller.trial.counter.__dict__ == {
                "trial_startups": 1,
                "validation_steps_started": 0,
                "validation_steps_ended": 0,
                "checkpoints_written": 0,
                "checkpoints_uploaded": 0,
                "training_started_times": 1,
                "training_epochs_started": 2,
                "training_epochs_ended": 2,
                "training_workloads_ended": 1,
                "trial_shutdowns": 0,
            }
            assert controller.trial.legacy_counter.__dict__ == {
                "legacy_on_training_epochs_start_calls": 2
            }
            yield workload.validation_workload(), workload.ignore_workload_response
            assert controller.trial.counter.__dict__ == {
                "trial_startups": 1,
                "validation_steps_started": 1,
                "validation_steps_ended": 1,
                "checkpoints_written": 0,
                "checkpoints_uploaded": 0,
                "training_started_times": 1,
                "training_epochs_started": 2,
                "training_epochs_ended": 2,
                "training_workloads_ended": 1,
                "trial_shutdowns": 0,
            }
            assert controller.trial.legacy_counter.__dict__ == {
                "legacy_on_training_epochs_start_calls": 2
            }

            interceptor = workload.WorkloadResponseInterceptor()
            yield from interceptor.send(workload.checkpoint_workload())
            nonlocal latest_checkpoint, steps_completed
            latest_checkpoint = interceptor.metrics_result()["uuid"]
            steps_completed = 1
            assert controller.trial.counter.__dict__ == {
                "trial_startups": 1,
                "validation_steps_started": 1,
                "validation_steps_ended": 1,
                "checkpoints_written": 1,
                "checkpoints_uploaded": 1,
                "training_started_times": 1,
                "training_epochs_started": 2,
                "training_epochs_ended": 2,
                "training_workloads_ended": 1,
                "trial_shutdowns": 0,
            }
            assert controller.trial.legacy_counter.__dict__ == {
                "legacy_on_training_epochs_start_calls": 2
            }

        controller = utils.make_pytorch_trial_controller_from_trial_impl(
            trial_class=pytorch_onevar_model.OneVarTrialCallbacks,
            hparams=hparams1,
            checkpoint_dir=str(checkpoint_dir),
            max_batches=100,
            min_validation_batches=10,
            min_checkpoint_batches=sys.maxsize,
        )
        controller.run()
        assert controller.trial.counter.trial_shutdowns == 1

        # Verify the checkpoint loading callback works.
        training_epochs = 1  # Total of 3
        num_batches = (
            training_epochs
            * len(pytorch_onevar_model.OnesDataset())
            // self.hparams["global_batch_size"]
        )

        def make_workloads2() -> workload.Stream:
            yield workload.train_workload(1, 1, 0, num_batches), workload.ignore_workload_response

        controller = utils.make_pytorch_trial_controller_from_trial_impl(
            trial_class=pytorch_onevar_model.OneVarTrialCallbacks,
            hparams=self.hparams,
            checkpoint_dir=str(checkpoint_dir),
            latest_checkpoint=latest_checkpoint,
            steps_completed=steps_completed,
            max_batches=100,
            min_validation_batches=10,
            min_checkpoint_batches=sys.maxsize,
        )
        controller.run()
        assert controller.trial.counter.__dict__ == {
            # Note: trial_startups will get reset by the loading logic.
            "trial_startups": 1,
            "validation_steps_started": 1,
            "validation_steps_ended": 1,
            # Note: checkpoints_written, checkpoints_uploaded, and trial_shutdowns, cannot be
            # persisted, as they are all updated after checkpointing.
            "checkpoints_written": 0,
            "checkpoints_uploaded": 0,
            "training_started_times": 2,
            "training_epochs_started": 3,
            "training_epochs_ended": 3,
            "training_workloads_ended": 2,
            "trial_shutdowns": 1,
        }
        assert controller.trial.legacy_counter.__dict__ == {
            "legacy_on_training_epochs_start_calls": 1
        }

    @pytest.mark.parametrize(
        "lr_scheduler_step_mode", [mode.value for mode in pytorch.LRScheduler.StepMode]
    )
    def test_context(
        self,
        lr_scheduler_step_mode,
    ) -> None:
        def make_workloads() -> workload.Stream:
            trainer = utils.TrainAndValidate()
            yield from trainer.send(steps=1, validation_freq=1, scheduling_unit=1)

        hparams = self.hparams.copy()
        hparams["lr_scheduler_step_mode"] = lr_scheduler_step_mode
        hparams["global_batch_size"] = 64

        controller = utils.make_pytorch_trial_controller_from_trial_impl(
            trial_class=pytorch_onevar_model.OneVarTrialAccessContext,
            hparams=hparams,
            trial_seed=self.trial_seed,
            max_batches=100,
            min_validation_batches=10,
            min_checkpoint_batches=sys.maxsize,
        )
        controller.run()

    def test_variable_workload_size(self) -> None:
        def make_workloads() -> workload.Stream:
            training_metrics = []
            interceptor = workload.WorkloadResponseInterceptor()

            total_steps, total_batches_processed = 10, 0
            for step_id in range(1, total_steps):
                num_batches = step_id
                yield from interceptor.send(
                    workload.train_workload(
                        step_id,
                        num_batches=num_batches,
                        total_batches_processed=total_batches_processed,
                    ),
                )
                metrics = interceptor.metrics_result()
                batch_metrics = metrics["metrics"]["batch_metrics"]
                assert len(batch_metrics) == num_batches, "did not run for expected num_batches"
                training_metrics.extend(batch_metrics)
                total_batches_processed += num_batches

        controller = utils.make_pytorch_trial_controller_from_trial_impl(
            trial_class=pytorch_onevar_model.OneVarTrial,
            hparams=self.hparams,
            trial_seed=self.trial_seed,
            max_batches=100,
            min_validation_batches=10,
            min_checkpoint_batches=sys.maxsize,
        )
        controller.run()

    def test_custom_reducers(self) -> None:
        def make_workloads() -> workload.Stream:
            trainer = utils.TrainAndValidate()

            yield from trainer.send(steps=3, validation_freq=3, scheduling_unit=10)
            training_metrics = trainer.get_avg_training_metrics()
            _, validation_metrics = trainer.result()

            batch_size = self.hparams["global_batch_size"]

            for i, metrics in enumerate(training_metrics):
                expect = pytorch_onevar_model.TriangleLabelSum.expect(
                    batch_size, 10 * i, 10 * (i + 1)
                )
                assert "cls_reducer" in metrics
                assert metrics["cls_reducer"] == expect
                assert "fn_reducer" in metrics
                assert metrics["fn_reducer"] == expect

            for metrics in validation_metrics:
                num_batches = len(pytorch_onevar_model.OnesDataset()) // batch_size
                expect = pytorch_onevar_model.TriangleLabelSum.expect(batch_size, 0, num_batches)
                assert "cls_reducer" in metrics
                assert metrics["cls_reducer"] == expect
                assert "fn_reducer" in metrics
                assert metrics["fn_reducer"] == expect

        controller = utils.make_pytorch_trial_controller_from_trial_impl(
            trial_class=pytorch_onevar_model.OneVarTrial,
            hparams=self.hparams,
            trial_seed=self.trial_seed,
            max_batches=100,
            min_validation_batches=10,
            min_checkpoint_batches=sys.maxsize,
        )
        controller.run()

    def test_reject_unnamed_nondict_metric(self) -> None:
        def make_workloads() -> workload.Stream:
            trainer = utils.TrainAndValidate()
            yield from trainer.send(steps=1, validation_freq=1, scheduling_unit=1)

        controller = utils.make_pytorch_trial_controller_from_trial_impl(
            trial_class=pytorch_onevar_model.OneVarTrial,
            hparams=self.hparams,
            trial_seed=self.trial_seed,
            max_batches=100,
            min_validation_batches=10,
            min_checkpoint_batches=sys.maxsize,
        )

        def reducer_fn(_):
            return 1.0

        # Inject an unnamed metric which returns a non-dict (which is not allowed).
        controller.context.wrap_reducer(reducer_fn)

        with pytest.raises(AssertionError, match="name=None but it did not return a dict"):
            controller.run()

    def test_reject_named_dict_metric(self) -> None:
        # If at some point in the future the webui is able to render scalar metrics inside
        # nested dictionary metrics, this test could go away.

        def make_workloads() -> workload.Stream:
            trainer = utils.TrainAndValidate()
            yield from trainer.send(steps=1, validation_freq=1, scheduling_unit=1)

        controller = utils.make_pytorch_trial_controller_from_trial_impl(
            trial_class=pytorch_onevar_model.OneVarTrial,
            hparams=self.hparams,
            trial_seed=self.trial_seed,
            max_batches=100,
            min_validation_batches=10,
            min_checkpoint_batches=sys.maxsize,
        )

        def reducer_fn(_):
            return {"my_metric": 1.0}

        # Inject a named metric which returns a dict (which is not allowed).
        controller.context.wrap_reducer(reducer_fn, name="my_metric")

        with pytest.raises(AssertionError, match="with name set but it returned a dict anyway"):
            controller.run()

    def test_require_disable_dataset_reproducibility(self) -> None:
        def make_workloads() -> workload.Stream:
            trainer = utils.TrainAndValidate()
            yield from trainer.send(steps=1, validation_freq=1, scheduling_unit=1)
            yield workload.terminate_workload(), [], workload.ignore_workload_response

        hparams = dict(self.hparams)
        hparams["dataloader_type"] = "torch"
        hparams["disable_dataset_reproducibility_checks"] = False

        with pytest.raises(RuntimeError, match="you can disable this check by calling"):
            controller = utils.make_pytorch_trial_controller_from_trial_impl(
                trial_class=pytorch_onevar_model.OneVarTrial,
                hparams=hparams,
                trial_seed=self.trial_seed,
                max_batches=100,
                min_validation_batches=10,
                min_checkpoint_batches=sys.maxsize,
            )
            controller.run()

    def test_custom_dataloader(self) -> None:
        def make_workloads() -> workload.Stream:
            trainer = utils.TrainAndValidate()

            yield from trainer.send(steps=100, validation_freq=10)
            training_metrics, validation_metrics = trainer.result()

            # Check the gradient update at every step.
            for idx, batch_metrics in enumerate(training_metrics):
                pytorch_onevar_model.OneVarTrial.check_batch_metrics(
                    batch_metrics,
                    idx,
                    metric_keyname_pairs=(("loss", "loss_exp"), ("w_after", "w_exp")),
                )

            # We expect the validation error and training loss to be
            # monotonically decreasing.
            for older, newer in zip(training_metrics, training_metrics[1:]):
                assert newer["loss"] <= older["loss"]

        hparams = dict(self.hparams)
        hparams["dataloader_type"] = "torch"
        hparams["disable_dataset_reproducibility_checks"] = True

        controller = utils.make_pytorch_trial_controller_from_trial_impl(
            trial_class=pytorch_onevar_model.OneVarTrial,
            hparams=hparams,
            trial_seed=self.trial_seed,
            max_batches=100,
            min_validation_batches=10,
            min_checkpoint_batches=sys.maxsize,
        )
        controller.run()

    def test_gradient_aggregation(self) -> None:
        AGG_FREQ = 2
        exp_config = utils.make_default_exp_config(
            self.hparams,
            scheduling_unit=1,
            searcher_metric=pytorch_onevar_model.OneVarTrial._searcher_metric,
        )
        exp_config["optimizations"].update(
            {
                "aggregation_frequency": AGG_FREQ,
                "average_aggregated_gradients": True,
            }
        )

        def make_workloads() -> workload.Stream:
            trainer = utils.TrainAndValidate()

            yield from trainer.send(steps=100, validation_freq=10)
            training_metrics, validation_metrics = trainer.result()

            # Check the gradient update at every step.
            for idx, batch_metrics in enumerate(training_metrics):
                if (idx + 1) % AGG_FREQ != 0:
                    # Only test batches which land on aggregation_frequency boundaries.
                    continue
                pytorch_onevar_model.OneVarTrial.check_batch_metrics(
                    batch_metrics,
                    idx,
                    metric_keyname_pairs=(("loss", "loss_exp"), ("w_after", "w_exp")),
                )

            for older, newer in zip(training_metrics, training_metrics[1:]):
                assert newer["loss"] <= older["loss"]

        run_pytorch_trial_controller_from_trial_impl(
            exp_config=exp_config,
            trial_class=pytorch_onevar_model.OneVarTrial,
            hparams=self.hparams,
            trial_seed=self.trial_seed,
            max_batches=100,
            min_validation_batches=10,
            min_checkpoint_batches=sys.maxsize,
        )

    @pytest.mark.skipif(not torch.cuda.is_available(), reason="no gpu available")
    @pytest.mark.gpu
    @pytest.mark.parametrize(
        "trial_class",
        [
            pytorch_onevar_model.OneVarApexAMPTrial,
            pytorch_onevar_model.OneVarAutoAMPTrial,
            pytorch_onevar_model.OneVarManualAMPTrial,
            pytorch_onevar_model.OneVarManualAMPWithNoopApexTrial,
            pytorch_onevar_model.OneVarApexAMPWithNoopScalerTrial,
        ],
        ids=[
            "apex",
            "autocast",
            "manual",
            "manual-with-noop-apex",
            "apex-with-noop-scaler",
        ],
    )
    def test_amp(self, trial_class) -> None:
        """Train a linear model using Determined with Automated Mixed Precision in three ways:
        Using Apex and using PyTorch AMP both "automatically" and "manually". In the "manual" case,
        we use the context manager ``autoscale`` in the model's training and
        evaluating methods; a scaler object is wrapped in a Determined context. The same
        is done under the hood in the first two cases.
        """
        if trial_class is pytorch_onevar_model.OneVarApexAMPTrial and not HAVE_APEX:
            pytest.skip("Apex not available")

        # The assertions logic in make_amp_workloads require a batch size of one
        hparams = dict(self.hparams)
        hparams["global_batch_size"] = 1

        training_metrics = {}

        def make_amp_workloads() -> workload.Stream:
            trainer = utils.TrainAndValidate()
            yield from trainer.send(steps=20, validation_freq=1, scheduling_unit=1)
            nonlocal training_metrics
            training_metrics, _ = trainer.result()

        run_pytorch_trial_controller_from_trial_impl(
            trial_class=trial_class,
            hparams=hparams,
            trial_seed=self.trial_seed,
            expose_gpus=True,
            max_batches=100,
            min_validation_batches=10,
            min_checkpoint_batches=sys.maxsize,
        )

        amp_metrics_test(trial_class, training_metrics)

    @pytest.mark.skipif(not torch.cuda.is_available(), reason="no gpu available")
    @pytest.mark.gpu
    @pytest.mark.parametrize(
        "trial_class",
        [
            pytorch_onevar_model.OneVarApexAMPTrial,
            pytorch_onevar_model.OneVarAutoAMPTrial,
            pytorch_onevar_model.OneVarManualAMPTrial,
        ],
        ids=[
            "apex",
            "autocast",
            "manual",
        ],
    )
    def test_amp_with_gradient_aggregation(self, trial_class) -> None:
        """Similar to test_amp but with gradient aggregation."""
        if trial_class is pytorch_onevar_model.OneVarApexAMPTrial and not HAVE_APEX:
            pytest.skip("Apex not available")

        # The assertions logic in make_amp_workloads require a batch size of one
        hparams = dict(self.hparams)
        hparams["global_batch_size"] = 1

        AGG_FREQ = 2
        exp_config = utils.make_default_exp_config(
            hparams,
            scheduling_unit=1,
            searcher_metric=trial_class._searcher_metric,
        )
        exp_config["optimizations"].update(
            {
                "aggregation_frequency": AGG_FREQ,
                "average_aggregated_gradients": True,
            }
        )

        training_metrics = {}

        def make_amp_workloads() -> workload.Stream:
            trainer = utils.TrainAndValidate()
            yield from trainer.send(steps=20 * AGG_FREQ, validation_freq=1, scheduling_unit=1)
            nonlocal training_metrics
            training_metrics, _ = trainer.result()

        run_pytorch_trial_controller_from_trial_impl(
            exp_config=exp_config,
            trial_class=trial_class,
            hparams=hparams,
            trial_seed=self.trial_seed,
            expose_gpus=True,
            max_batches=100,
            min_validation_batches=10,
            min_checkpoint_batches=sys.maxsize,
        )

        amp_metrics_test(trial_class, training_metrics, agg_freq=AGG_FREQ)

    def checkpoint_and_restore(
        self, hparams: typing.Dict, tmp_path: pathlib.Path, steps: typing.Tuple[int, int] = (1, 1)
    ) -> typing.Tuple[
        typing.Sequence[typing.Dict[str, typing.Any]], typing.Sequence[typing.Dict[str, typing.Any]]
    ]:
        checkpoint_dir = str(tmp_path.joinpath("checkpoint"))
        training_metrics = {"A": [], "B": []}
        validation_metrics = {"A": [], "B": []}

        # Trial A: train 100 batches and checkpoint
        trial_A, trial_controller_A = create_trial_and_trial_controller(
            trial_class=pytorch_onevar_model.OneVarTrialWithLRScheduler,
            hparams=hparams,
            trial_seed=self.trial_seed,
            max_batches=steps[0],
            min_validation_batches=steps[0],
            min_checkpoint_batches=steps[0],
            checkpoint_dir=checkpoint_dir,
        )

        trial_controller_A.run()

        metrics_callback = trial_A.metrics_callback
        checkpoint_callback = trial_A.checkpoint_callback

        training_metrics["A"] = metrics_callback.training_metrics
        assert len(training_metrics["A"]) == steps[0], "training metrics did not match expected length"
        validation_metrics["A"] = metrics_callback.validation_metrics

        assert checkpoint_callback.uuid is not None, "trial did not return a checkpoint UUID"

        # Trial A: restore from checkpoint and train for 100 more batches
        trial_A, trial_controller_A = create_trial_and_trial_controller(
            trial_class=pytorch_onevar_model.OneVarTrialWithLRScheduler,
            hparams=hparams,
            trial_seed=self.trial_seed,
            max_batches=steps[0] + steps[1],
            min_validation_batches=steps[1],
            min_checkpoint_batches=sys.maxsize,
            checkpoint_dir=checkpoint_dir,
            latest_checkpoint=checkpoint_callback.uuid,
            steps_completed=trial_controller_A._state.batches_trained,
        )
        trial_controller_A.run()

        metrics_callback = trial_A.metrics_callback
        training_metrics["A"] += metrics_callback.training_metrics
        validation_metrics["A"] += metrics_callback.validation_metrics

        assert (
            len(training_metrics["A"]) == steps[0] + steps[1]
        ), "training metrics returned did not match expected length"

        # Trial B: run for 200 steps
        trial_B, trial_controller_B = create_trial_and_trial_controller(
            trial_class=pytorch_onevar_model.OneVarTrialWithLRScheduler,
            hparams=hparams,
            trial_seed=self.trial_seed,
            max_batches=steps[0] + steps[1],
            min_validation_batches=steps[0] + steps[1],
            min_checkpoint_batches=sys.maxsize,
            checkpoint_dir=checkpoint_dir,
        )
        trial_controller_B.run()

        metrics_callback = trial_B.metrics_callback

        training_metrics["B"] = metrics_callback.training_metrics
        validation_metrics["B"] = metrics_callback.validation_metrics

        for A, B in zip(training_metrics["A"], training_metrics["B"]):
            utils.assert_equivalent_metrics(A, B)
        print(validation_metrics)
        for A, B in zip(validation_metrics["A"], validation_metrics["B"]):
            utils.assert_equivalent_metrics(A, B)

        return (training_metrics["A"], training_metrics["B"])


@pytest.mark.parametrize(
    "ckpt,istrial",
    [
        ("0.13.13-pytorch-old", False),
        ("0.13.13-pytorch-flex", True),
        ("0.17.6-pytorch", True),
        ("0.17.7-pytorch", True),
    ],
)
def test_checkpoint_loading(ckpt: str, istrial: bool):
    checkpoint_dir = os.path.join(utils.fixtures_path("ancient-checkpoints"), f"{ckpt}")
    trial = pytorch.load_trial_from_checkpoint_path(checkpoint_dir)
    if istrial:
        assert isinstance(trial, pytorch.PyTorchTrial), type(trial)
    else:
        assert isinstance(trial, torch.nn.Module), type(trial)


def amp_metrics_test(trial_class, training_metrics, agg_freq=1):
    loss_prev = None
    GROWTH_INTERVAL = trial_class._growth_interval
    MIN_SCALED_LOSS_TO_REDUCE_SCALE = 32760
    growth_countdown = GROWTH_INTERVAL
    # Only attempt assertions up to and including the penultimate batch, because
    #  we may not have the updated scale from the final batch.
    for idx, (metrics, next_metrics) in enumerate(zip(training_metrics[:-1], training_metrics[1:])):
        if (idx + 1) % agg_freq != 0:
            # Only test batches which land on aggregation_frequency boundaries.
            continue
        # Because the scaler is updated during the optimizer step, which occurs after training from
        #  a batch, the metrics dictionary may not have the updated scale, but we can get it
        #  from the next step.
        scale = next_metrics["scale_before"]
        if "scale" in metrics:
            # In cases where we do know the scale immediately after it's updated, we might as well
            #  do this check. If this fails, something is very wrong.
            assert metrics["scale"] == scale, "scale is inconsistent between batches"
        else:
            metrics["scale"] = scale
        loss = metrics["loss"].item()
        scale_before = metrics["scale_before"]
        scaled_loss = loss * scale_before
        scale = metrics["scale"]
        growth_countdown -= 1
        if scaled_loss >= MIN_SCALED_LOSS_TO_REDUCE_SCALE:
            assert scale < scale_before, (
                f"scale was expected to reduce from {scale_before} but did not "
                f"(scaled_loss={scaled_loss} >= {MIN_SCALED_LOSS_TO_REDUCE_SCALE})) "
            )
            growth_countdown = GROWTH_INTERVAL
        elif growth_countdown == 0:
            assert scale > scale_before, (
                f"scale was expected to grow but did not " f"(growth_countdown={growth_countdown}) "
            )
            growth_countdown = GROWTH_INTERVAL
        else:
            assert scale == scale_before, (
                f"scale changed from {scale_before} to {scale} but not when expected "
                f"(growth_countdown={growth_countdown}) "
                f"(scaled_loss={scaled_loss} < {MIN_SCALED_LOSS_TO_REDUCE_SCALE})) "
            )
            # Check the accuracy of the gradient change.
            metric_keyname_pairs = [("loss", "loss_exp"), ("w_after", "w_exp")]
            if metrics["stage"] in ["small", "zero"]:
                metric_keyname_pairs.append(("w_before", "w_after"))
            trial_class.check_batch_metrics(
                metrics,
                idx,
                metric_keyname_pairs=metric_keyname_pairs,
                atol=1e-4,
            )
            if loss_prev is not None and metrics["stage"] == "one":
                assert loss <= loss_prev, "loss was expected to decrease monotonically"
                loss_prev = loss


def create_trial_and_trial_controller(
    trial_class: pytorch.PyTorchTrial,
    hparams: typing.Dict,
    scheduling_unit: int = 1,
    trial_seed: int = random.randint(0, 1 << 31),
    exp_config: typing.Optional[typing.Dict] = None,
    checkpoint_dir: typing.Optional[str] = None,
    latest_checkpoint: typing.Optional[str] = None,
    steps_completed: int = 0,
    expose_gpus: bool = False,
    max_batches: int = 100,
    min_checkpoint_batches: int = 0,
    min_validation_batches: int = 0,
) -> typing.Tuple[pytorch.PyTorchTrial, pytorch.PyTorchTrialController]:
    assert issubclass(
        trial_class, pytorch.PyTorchTrial
    ), "pytorch test method called for non-pytorch trial"

    if not exp_config:
        assert hasattr(
            trial_class, "_searcher_metric"
        ), "Trial classes for unit tests should be annotated with a _searcher_metric attribute"
        searcher_metric = trial_class._searcher_metric  # type: ignore
        exp_config = utils.make_default_exp_config(
            hparams, scheduling_unit, searcher_metric, checkpoint_dir=checkpoint_dir
        )

    storage_manager = det.common.storage.SharedFSStorageManager(checkpoint_dir or "/tmp")
    with det.core._dummy_init(storage_manager=storage_manager) as core_context:
        core_context.train._trial_id = "1"
        distributed_backend = det._DistributedBackend()
        if expose_gpus:
            gpu_uuids = gpu.get_gpu_uuids()
        else:
            gpu_uuids = []

        pytorch.PyTorchTrialController.pre_execute_hook(trial_seed, distributed_backend)
        trial_context = pytorch.PyTorchTrialContext(
            core_context=core_context,
            trial_seed=trial_seed,
            hparams=hparams,
            slots_per_trial=1,
            num_gpus=len(gpu_uuids),
            exp_conf=exp_config,
            aggregation_frequency=1,
            fp16_compression=False,
            average_aggregated_gradients=True,
            steps_completed=steps_completed,
        )

        trial_inst = trial_class(trial_context)

        trial_controller = pytorch.PyTorchTrialController(
            trial_inst=trial_inst,
            context=trial_context,
            max_length=pytorch.Batch(max_batches),
            min_checkpoint_period=pytorch.Batch(min_checkpoint_batches),
            min_validation_period=pytorch.Batch(min_validation_batches),
            searcher_metric_name=trial_class._searcher_metric,  # type: ignore
            scheduling_unit=scheduling_unit,
            local_training=True,
            latest_checkpoint=latest_checkpoint,
            steps_completed=steps_completed,
        )

        trial_controller._set_data_loaders()
        trial_controller.training_iterator = iter(trial_controller.training_loader)
        return trial_inst, trial_controller
