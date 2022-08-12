const RESPONSES = {
  multiTrial: {
    getExperimentsDetails: {
      archived: false,
      config: {
        checkpointPolicy: 'best',
        checkpointStorage: {
          bucket: 'det-determined-master-us-west-2-573932760021',
          saveExperimentBest: 0,
          saveTrialBest: 1,
          saveTrialLatest: 1,
          type: 's3',
        },
        dataLayer: { type: 'shared_fs' },
        hyperparameters: {
          global_batch_size: { type: 'const', val: 64 },
          hidden_size: { type: 'const', val: 64 },
          learning_rate: { maxval: 0.1, minval: 0.0001, type: 'double' },
        },
        labels: [],
        maxRestarts: 5,
        name: 'mnist_pytorch_lightning_adaptive',
        profiling: { enabled: false },
        resources: {},
        searcher: {
          bracket_rungs: [],
          divisor: 4,
          max_concurrent_trials: 0,
          max_length: { batches: 937 },
          max_rungs: 5,
          max_trials: 16,
          metric: 'val_loss',
          mode: 'standard',
          name: 'adaptive_asha',
          smaller_is_better: true,
          smallerIsBetter: true,
          source_checkpoint_uuid: null,
          source_trial_id: null,
          stop_once: false,
        },
      },
      configRaw: {
        bind_mounts: [],
        checkpoint_policy: 'best',
        checkpoint_storage: {
          access_key: null,
          bucket: 'det-determined-master-us-west-2-573932760021',
          endpoint_url: null,
          prefix: null,
          save_experiment_best: 0,
          save_trial_best: 1,
          save_trial_latest: 1,
          secret_key: null,
          type: 's3',
        },
        data: {
          url: 'https://s3-us-west-2.amazonaws.com/determined-ai-test-data/' +
            'pytorch_mnist.tar.gz',
        },
        data_layer: { container_storage_path: null, host_storage_path: null, type: 'shared_fs' },
        debug: false,
        description: null,
        entrypoint: 'model_def:MNISTTrial',
        environment: {
          add_capabilities: [],
          drop_capabilities: [],
          environment_variables: { cpu: [], cuda: [], rocm: [] },
          force_pull_image: false,
          image: {
            cpu: 'determinedai/environments:py-3.8-pytorch-1.10-tf-2.8-cpu-ecee7c1',
            cuda: 'determinedai/environments:cuda-11.3-pytorch-1.10-tf-2.8-gpu-ecee7c1',
            rocm: 'determinedai/environments:rocm-4.2-pytorch-1.9-tf-2.5-rocm-ecee7c1',
          },
          pod_spec: null,
          ports: {},
          registry_auth: null,
          slurm: [],
        },
        hyperparameters: {
          global_batch_size: { type: 'const', val: 64 },
          hidden_size: { type: 'const', val: 64 },
          learning_rate: { maxval: 0.1, minval: 0.0001, type: 'double' },
        },
        labels: [],
        max_restarts: 5,
        min_checkpoint_period: { batches: 0 },
        min_validation_period: { batches: 0 },
        name: 'mnist_pytorch_lightning_adaptive',
        optimizations: {
          aggregation_frequency: 1,
          auto_tune_tensor_fusion: false,
          average_aggregated_gradients: true,
          average_training_metrics: true,
          grad_updates_size_file: null,
          gradient_compression: false,
          mixed_precision: 'O0',
          tensor_fusion_cycle_time: 5,
          tensor_fusion_threshold: 64,
        },
        perform_initial_validation: false,
        profiling: { begin_on_batch: 0, enabled: false, end_after_batch: null, sync_timings: true },
        project: '',
        records_per_epoch: 0,
        reproducibility: { experiment_seed: 1659678335 },
        resources: {
          agent_label: '',
          devices: [],
          max_slots: null,
          native_parallel: false,
          priority: null,
          resource_pool: 'compute-pool',
          shm_size: null,
          slots_per_trial: 1,
          weight: 1,
        },
        scheduling_unit: 100,
        searcher: {
          bracket_rungs: [],
          divisor: 4,
          max_concurrent_trials: 0,
          max_length: { batches: 937 },
          max_rungs: 5,
          max_trials: 16,
          metric: 'val_loss',
          mode: 'standard',
          name: 'adaptive_asha',
          smaller_is_better: true,
          source_checkpoint_uuid: null,
          source_trial_id: null,
          stop_once: false,
        },
        workspace: '',
      },
      description: '',
      endTime: '2022-08-05T05:57:21.223917Z',
      forkedFrom: null,
      hyperparameters: {
        global_batch_size: { type: 'const', val: 64 },
        hidden_size: { type: 'const', val: 64 },
        learning_rate: { maxval: 0.1, minval: 0.0001, type: 'double' },
      },
      id: 1249,
      jobId: '0b5c71e7-a4fb-4efc-85da-f5b545d02914',
      jobSummary: null,
      labels: [],
      name: 'mnist_pytorch_lightning_adaptive',
      notes: '',
      numTrials: 16,
      parentArchived: false,
      progress: 1,
      projectId: 102,
      projectName: 'Super Fun Project',
      resourcePool: 'compute-pool',
      searcherType: 'adaptive_asha',
      startTime: '2022-08-05T05:45:35.495201Z',
      state: 'COMPLETED',
      trialIds: [ 3566,
        3567,
        3568,
        3569,
        3570,
        3571,
        3572,
        3573,
        3574,
        3575,
        3576,
        3577,
        3578,
        3579,
        3580,
        3581 ],
      userId: 34,
      workspaceId: 103,
      workspaceName: 'Caleb',
    },
    getExpTrials: {
      pagination: { endIndex: 10, limit: 10, offset: 0, startIndex: 0, total: 16 },
      trials: [ {
        autoRestarts: 0,
        bestAvailableCheckpoint: {
          endTime: '2022-08-05T05:56:47.998811Z',
          resources: {
            'code/': 0,
            'code/adaptive.yaml': 418,
            'code/const.yaml': 349,
            'code/data.py': 3254,
            'code/mnist.py': 2996,
            'code/model_def.py': 1445,
            'code/README.md': 1492,
            'code/startup-hook.sh': 59,
            'load_data.json': 2665,
            'state_dict.pth': 680041,
            'workload_sequencer.pkl': 89,
          },
          state: 'COMPLETED',
          totalBatches: 58,
          uuid: 'b01dfffc-ad3e-4398-b205-b66d920d7b0b',
        },
        bestValidationMetric: {
          endTime: '2022-08-05T05:56:49.692523Z',
          metrics: { accuracy: 0.17780853807926178, val_loss: 2.141766309738159 },
          totalBatches: 58,
        },
        endTime: '2022-08-05T05:56:51.432745Z',
        experimentId: 1249,
        hyperparameters: {
          global_batch_size: 64,
          hidden_size: 64,
          learning_rate: 0.08311236560675993,
        },
        id: 3581,
        latestValidationMetric: {
          endTime: '2022-08-05T05:56:49.692523Z',
          metrics: { accuracy: 0.17780853807926178, val_loss: 2.141766309738159 },
          totalBatches: 58,
        },
        startTime: '2022-08-05T05:56:23.400221Z',
        state: 'COMPLETED',
        totalBatchesProcessed: 58,
        workloads: [],
      },
      {
        autoRestarts: 0,
        bestAvailableCheckpoint: {
          endTime: '2022-08-05T05:56:41.515152Z',
          resources: {
            'code/': 0,
            'code/adaptive.yaml': 418,
            'code/const.yaml': 349,
            'code/data.py': 3254,
            'code/mnist.py': 2996,
            'code/model_def.py': 1445,
            'code/README.md': 1492,
            'code/startup-hook.sh': 59,
            'load_data.json': 2665,
            'state_dict.pth': 680105,
            'workload_sequencer.pkl': 89,
          },
          state: 'COMPLETED',
          totalBatches: 58,
          uuid: 'b3d4ed7e-625c-4083-9e6c-ffba4ac03729',
        },
        bestValidationMetric: {
          endTime: '2022-08-05T05:56:42.958782Z',
          metrics: { accuracy: 0.3433544337749481, val_loss: 1.7270305156707764 },
          totalBatches: 58,
        },
        endTime: '2022-08-05T05:56:49.345991Z',
        experimentId: 1249,
        hyperparameters: {
          global_batch_size: 64,
          hidden_size: 64,
          learning_rate: 0.06607116139078467,
        },
        id: 3580,
        latestValidationMetric: {
          endTime: '2022-08-05T05:56:42.958782Z',
          metrics: { accuracy: 0.3433544337749481, val_loss: 1.7270305156707764 },
          totalBatches: 58,
        },
        startTime: '2022-08-05T05:56:16.181755Z',
        state: 'COMPLETED',
        totalBatchesProcessed: 58,
        workloads: [],
      },
      {
        autoRestarts: 0,
        bestAvailableCheckpoint: {
          endTime: '2022-08-05T05:56:19.781968Z',
          resources: {
            'code/': 0,
            'code/adaptive.yaml': 418,
            'code/const.yaml': 349,
            'code/data.py': 3254,
            'code/mnist.py': 2996,
            'code/model_def.py': 1445,
            'code/README.md': 1492,
            'code/startup-hook.sh': 59,
            'load_data.json': 2664,
            'state_dict.pth': 680105,
            'workload_sequencer.pkl': 89,
          },
          state: 'COMPLETED',
          totalBatches: 58,
          uuid: '3b780651-0250-4307-8ec0-fac09a688970',
        },
        bestValidationMetric: {
          endTime: '2022-08-05T05:56:21.150552Z',
          metrics: { accuracy: 0.1809730976819992, val_loss: 2.1605682373046875 },
          totalBatches: 58,
        },
        endTime: '2022-08-05T05:56:49.524883Z',
        experimentId: 1249,
        hyperparameters: {
          global_batch_size: 64,
          hidden_size: 64,
          learning_rate: 0.0951845051299588,
        },
        id: 3579,
        latestValidationMetric: {
          endTime: '2022-08-05T05:56:21.150552Z',
          metrics: { accuracy: 0.1809730976819992, val_loss: 2.1605682373046875 },
          totalBatches: 58,
        },
        startTime: '2022-08-05T05:55:55.152888Z',
        state: 'COMPLETED',
        totalBatchesProcessed: 58,
        workloads: [],
      },
      {
        autoRestarts: 0,
        bestAvailableCheckpoint: {
          endTime: '2022-08-05T05:56:12.628079Z',
          resources: {
            'code/': 0,
            'code/adaptive.yaml': 418,
            'code/const.yaml': 349,
            'code/data.py': 3254,
            'code/mnist.py': 2996,
            'code/model_def.py': 1445,
            'code/README.md': 1492,
            'code/startup-hook.sh': 59,
            'load_data.json': 2665,
            'state_dict.pth': 680169,
            'workload_sequencer.pkl': 89,
          },
          state: 'COMPLETED',
          totalBatches: 58,
          uuid: '96240de7-d768-4c73-966d-e1aeaf87a534',
        },
        bestValidationMetric: {
          endTime: '2022-08-05T05:56:14.049592Z',
          metrics: { accuracy: 0.20945411920547485, val_loss: 2.0849103927612305 },
          totalBatches: 58,
        },
        endTime: '2022-08-05T05:56:49.741209Z',
        experimentId: 1249,
        hyperparameters: {
          global_batch_size: 64,
          hidden_size: 64,
          learning_rate: 0.07466800470318831,
        },
        id: 3578,
        latestValidationMetric: {
          endTime: '2022-08-05T05:56:14.049592Z',
          metrics: { accuracy: 0.20945411920547485, val_loss: 2.0849103927612305 },
          totalBatches: 58,
        },
        startTime: '2022-08-05T05:55:47.632975Z',
        state: 'COMPLETED',
        totalBatchesProcessed: 58,
        workloads: [],
      },
      {
        autoRestarts: 0,
        bestAvailableCheckpoint: {
          endTime: '2022-08-05T05:55:44.607451Z',
          resources: {
            'code/': 0,
            'code/adaptive.yaml': 418,
            'code/const.yaml': 349,
            'code/data.py': 3254,
            'code/mnist.py': 2996,
            'code/model_def.py': 1445,
            'code/README.md': 1492,
            'code/startup-hook.sh': 59,
            'load_data.json': 2665,
            'state_dict.pth': 680169,
            'workload_sequencer.pkl': 89,
          },
          state: 'COMPLETED',
          totalBatches: 58,
          uuid: '9f0588b0-c441-4243-b1b8-1fe9b20906d3',
        },
        bestValidationMetric: {
          endTime: '2022-08-05T05:55:46.033543Z',
          metrics: { accuracy: 0.17405062913894653, val_loss: 2.138223886489868 },
          totalBatches: 58,
        },
        endTime: '2022-08-05T05:56:49.687014Z',
        experimentId: 1249,
        hyperparameters: {
          global_batch_size: 64,
          hidden_size: 64,
          learning_rate: 0.09915932917592189,
        },
        id: 3577,
        latestValidationMetric: {
          endTime: '2022-08-05T05:55:46.033543Z',
          metrics: { accuracy: 0.17405062913894653, val_loss: 2.138223886489868 },
          totalBatches: 58,
        },
        startTime: '2022-08-05T05:55:19.578336Z',
        state: 'COMPLETED',
        totalBatchesProcessed: 58,
        workloads: [],
      },
      {
        autoRestarts: 0,
        bestAvailableCheckpoint: {
          endTime: '2022-08-05T05:55:13.638730Z',
          resources: {
            'code/': 0,
            'code/adaptive.yaml': 418,
            'code/const.yaml': 349,
            'code/data.py': 3254,
            'code/mnist.py': 2996,
            'code/model_def.py': 1445,
            'code/README.md': 1492,
            'code/startup-hook.sh': 59,
            'load_data.json': 2666,
            'state_dict.pth': 680041,
            'workload_sequencer.pkl': 89,
          },
          state: 'COMPLETED',
          totalBatches: 234,
          uuid: '13253d89-b414-4808-9ea8-a8f5214c953f',
        },
        bestValidationMetric: {
          endTime: '2022-08-05T05:55:14.966958Z',
          metrics: { accuracy: 0.9072389006614685, val_loss: 0.33882153034210205 },
          totalBatches: 234,
        },
        endTime: '2022-08-05T05:57:19.594358Z',
        experimentId: 1249,
        hyperparameters: {
          global_batch_size: 64,
          hidden_size: 64,
          learning_rate: 0.011115759023330877,
        },
        id: 3576,
        latestValidationMetric: {
          endTime: '2022-08-05T05:55:14.966958Z',
          metrics: { accuracy: 0.9072389006614685, val_loss: 0.33882153034210205 },
          totalBatches: 234,
        },
        startTime: '2022-08-05T05:54:44.461327Z',
        state: 'COMPLETED',
        totalBatchesProcessed: 234,
        workloads: [],
      },
      {
        autoRestarts: 0,
        bestAvailableCheckpoint: {
          endTime: '2022-08-05T05:54:45.340228Z',
          resources: {
            'code/': 0,
            'code/adaptive.yaml': 418,
            'code/const.yaml': 349,
            'code/data.py': 3254,
            'code/mnist.py': 2996,
            'code/model_def.py': 1445,
            'code/README.md': 1492,
            'code/startup-hook.sh': 59,
            'load_data.json': 2666,
            'state_dict.pth': 680041,
            'workload_sequencer.pkl': 89,
          },
          state: 'COMPLETED',
          totalBatches: 58,
          uuid: '4c6c59a2-02e7-493c-8471-98f763e0171b',
        },
        bestValidationMetric: {
          endTime: '2022-08-05T05:54:46.761720Z',
          metrics: { accuracy: 0.5215585231781006, val_loss: 1.3084886074066162 },
          totalBatches: 58,
        },
        endTime: '2022-08-05T05:56:49.534376Z',
        experimentId: 1249,
        hyperparameters: {
          global_batch_size: 64,
          hidden_size: 64,
          learning_rate: 0.043852742488714104,
        },
        id: 3575,
        latestValidationMetric: {
          endTime: '2022-08-05T05:54:46.761720Z',
          metrics: { accuracy: 0.5215585231781006, val_loss: 1.3084886074066162 },
          totalBatches: 58,
        },
        startTime: '2022-08-05T05:54:20.408985Z',
        state: 'COMPLETED',
        totalBatchesProcessed: 58,
        workloads: [],
      },
      {
        autoRestarts: 0,
        bestAvailableCheckpoint: {
          endTime: '2022-08-05T05:54:40.971698Z',
          resources: {
            'code/': 0,
            'code/adaptive.yaml': 418,
            'code/const.yaml': 349,
            'code/data.py': 3254,
            'code/mnist.py': 2996,
            'code/model_def.py': 1445,
            'code/README.md': 1492,
            'code/startup-hook.sh': 59,
            'load_data.json': 2665,
            'state_dict.pth': 680105,
            'workload_sequencer.pkl': 89,
          },
          state: 'COMPLETED',
          totalBatches: 58,
          uuid: '73368ce4-c655-4f33-846f-a4f76357dd6b',
        },
        bestValidationMetric: {
          endTime: '2022-08-05T05:54:42.341627Z',
          metrics: { accuracy: 0.26285600662231445, val_loss: 2.150845766067505 },
          totalBatches: 58,
        },
        endTime: '2022-08-05T05:56:49.563417Z',
        experimentId: 1249,
        hyperparameters: {
          global_batch_size: 64,
          hidden_size: 64,
          learning_rate: 0.06524170337897893,
        },
        id: 3574,
        latestValidationMetric: {
          endTime: '2022-08-05T05:54:42.341627Z',
          metrics: { accuracy: 0.26285600662231445, val_loss: 2.150845766067505 },
          totalBatches: 58,
        },
        startTime: '2022-08-05T05:54:16.416854Z',
        state: 'COMPLETED',
        totalBatchesProcessed: 58,
        workloads: [],
      },
      {
        autoRestarts: 0,
        bestAvailableCheckpoint: {
          endTime: '2022-08-05T05:54:16.830546Z',
          resources: {
            'code/': 0,
            'code/adaptive.yaml': 418,
            'code/const.yaml': 349,
            'code/data.py': 3254,
            'code/mnist.py': 2996,
            'code/model_def.py': 1445,
            'code/README.md': 1492,
            'code/startup-hook.sh': 59,
            'load_data.json': 2665,
            'state_dict.pth': 680105,
            'workload_sequencer.pkl': 89,
          },
          state: 'COMPLETED',
          totalBatches: 234,
          uuid: 'c1996b98-bf3c-42e1-a0e0-512dbbfe7226',
        },
        bestValidationMetric: {
          endTime: '2022-08-05T05:54:18.377589Z',
          metrics: { accuracy: 0.688686728477478, val_loss: 0.836029052734375 },
          totalBatches: 234,
        },
        endTime: '2022-08-05T05:54:20.185007Z',
        experimentId: 1249,
        hyperparameters: {
          global_batch_size: 64,
          hidden_size: 64,
          learning_rate: 0.04151271651016727,
        },
        id: 3573,
        latestValidationMetric: {
          endTime: '2022-08-05T05:54:18.377589Z',
          metrics: { accuracy: 0.688686728477478, val_loss: 0.836029052734375 },
          totalBatches: 234,
        },
        startTime: '2022-08-05T05:53:48.862493Z',
        state: 'COMPLETED',
        totalBatchesProcessed: 234,
        workloads: [],
      },
      {
        autoRestarts: 0,
        bestAvailableCheckpoint: {
          endTime: '2022-08-05T05:54:13.406903Z',
          resources: {
            'code/': 0,
            'code/adaptive.yaml': 418,
            'code/const.yaml': 349,
            'code/data.py': 3254,
            'code/mnist.py': 2996,
            'code/model_def.py': 1445,
            'code/README.md': 1492,
            'code/startup-hook.sh': 59,
            'load_data.json': 2665,
            'state_dict.pth': 680105,
            'workload_sequencer.pkl': 89,
          },
          state: 'COMPLETED',
          totalBatches: 58,
          uuid: 'dd9033d5-4203-4623-8db5-fb4083313655',
        },
        bestValidationMetric: {
          endTime: '2022-08-05T05:54:14.777372Z',
          metrics: { accuracy: 0.11194620281457901, val_loss: 2.295724630355835 },
          totalBatches: 58,
        },
        endTime: '2022-08-05T05:56:49.541919Z',
        experimentId: 1249,
        hyperparameters: {
          global_batch_size: 64,
          hidden_size: 64,
          learning_rate: 0.09860315167441819,
        },
        id: 3572,
        latestValidationMetric: {
          endTime: '2022-08-05T05:54:14.777372Z',
          metrics: { accuracy: 0.11194620281457901, val_loss: 2.295724630355835 },
          totalBatches: 58,
        },
        startTime: '2022-08-05T05:53:48.849706Z',
        state: 'COMPLETED',
        totalBatchesProcessed: 58,
        workloads: [],
      } ],
    },
    getExpValidationHistory: [
      {
        endTime: '2022-08-05T05:52:27.682018Z',
        trialId: 3566,
        validationError: 1.2850603,
      },
      {
        endTime: '2022-08-05T05:52:36.612957Z',
        trialId: 3567,
        validationError: 1.0128666,
      },
      {
        endTime: '2022-08-05T05:53:15.559363Z',
        trialId: 3569,
        validationError: 0.5252631,
      },
      {
        endTime: '2022-08-05T05:53:46.631421Z',
        trialId: 3570,
        validationError: 0.26340833,
      },
      {
        endTime: '2022-08-05T05:55:53.389436Z',
        trialId: 3570,
        validationError: 0.23642454,
      },
    ],
    getProject: {
      archived: false,
      description: 'SUPER FUN',
      id: 102,
      immutable: false,
      lastExperimentStartedAt: '2022-08-05T05:45:35.495201Z',
      name: 'Super Fun Project',
      notes: [ { contents: '# Does this work?', name: 'My Custom Notes' } ],
      numActiveExperiments: 0,
      numExperiments: 7,
      userId: 34,
      workspaceId: 103,
      workspaceName: 'Caleb',
    },
    getWorkspace: {
      archived: false,
      id: 103,
      immutable: false,
      name: 'Caleb',
      numExperiments: 7,
      numProjects: 1,
      pinned: true,
      userId: 34,
    },
  },
  singleTrial: {
    getExperimentsDetails: {
      archived: false,
      config: {
        checkpointPolicy: 'best',
        checkpointStorage: {
          bucket: 'det-determined-preview-us-west-2-573932760021',
          saveExperimentBest: 0,
          saveTrialBest: 1,
          saveTrialLatest: 1,
          type: 's3',
        },
        dataLayer: { type: 'shared_fs' },
        hyperparameters: {
          dropout1: { type: 'const', val: 0.25 },
          dropout2: { type: 'const', val: 0.5 },
          global_batch_size: { type: 'const', val: 64 },
          learning_rate: { type: 'const', val: 1 },
          n_filters1: { type: 'const', val: 32 },
          n_filters2: { type: 'const', val: 64 },
          test1: {
            test2: {
              test3: {
                test4: {
                  optimizer_fake: {
                    lr: { type: 'const', val: 0.0001 },
                    momentum: { type: 'const', val: 0.9 },
                  },
                },
              },
            },
          },
        },
        labels: [],
        maxRestarts: 5,
        name: 'mnist_pytorch_const',
        profiling: { enabled: false },
        resources: {},
        searcher: {
          max_length: { batches: 937 },
          metric: 'validation_loss',
          name: 'single',
          smaller_is_better: true,
          smallerIsBetter: true,
          source_checkpoint_uuid: null,
          source_trial_id: null,
        },
      },
      configRaw: {
        bind_mounts: [],
        checkpoint_policy: 'best',
        checkpoint_storage: {
          access_key: null,
          bucket: 'det-determined-preview-us-west-2-573932760021',
          endpoint_url: null,
          prefix: null,
          save_experiment_best: 0,
          save_trial_best: 1,
          save_trial_latest: 1,
          secret_key: null,
          type: 's3',
        },
        data: {
          url: 'https://s3-us-west-2.amazonaws.com/determined-ai-test-data/' +
            'pytorch_mnist.tar.gz',
        },
        data_layer: { container_storage_path: null, host_storage_path: null, type: 'shared_fs' },
        debug: false,
        description: null,
        entrypoint: 'model_def:MNistTrial',
        environment: {
          add_capabilities: [],
          drop_capabilities: [],
          environment_variables: { cpu: [], cuda: [], rocm: [] },
          force_pull_image: false,
          image: {
            cpu: 'determinedai/environments:py-3.8-pytorch-1.10-tf-2.8-cpu-ecee7c1',
            cuda: 'determinedai/environments:cuda-11.3-pytorch-1.10-tf-2.8-gpu-ecee7c1',
            rocm: 'determinedai/environments:rocm-4.2-pytorch-1.9-tf-2.5-rocm-ecee7c1',
          },
          pod_spec: null,
          ports: {},
          registry_auth: null,
          slurm: [],
        },
        hyperparameters: {
          dropout1: { type: 'const', val: 0.25 },
          dropout2: { type: 'const', val: 0.5 },
          global_batch_size: { type: 'const', val: 64 },
          learning_rate: { type: 'const', val: 1 },
          n_filters1: { type: 'const', val: 32 },
          n_filters2: { type: 'const', val: 64 },
          test1: {
            test2: {
              test3: {
                test4: {
                  optimizer_fake: {
                    lr: { type: 'const', val: 0.0001 },
                    momentum: { type: 'const', val: 0.9 },
                  },
                },
              },
            },
          },
        },
        labels: [],
        max_restarts: 5,
        min_checkpoint_period: { batches: 0 },
        min_validation_period: { batches: 0 },
        name: 'mnist_pytorch_const',
        optimizations: {
          aggregation_frequency: 1,
          auto_tune_tensor_fusion: false,
          average_aggregated_gradients: true,
          average_training_metrics: true,
          grad_updates_size_file: null,
          gradient_compression: false,
          mixed_precision: 'O0',
          tensor_fusion_cycle_time: 5,
          tensor_fusion_threshold: 64,
        },
        perform_initial_validation: false,
        profiling: { begin_on_batch: 0, enabled: false, end_after_batch: null, sync_timings: true },
        project: '',
        records_per_epoch: 0,
        reproducibility: { experiment_seed: 1658990898 },
        resources: {
          agent_label: '',
          devices: [],
          max_slots: null,
          native_parallel: false,
          priority: null,
          resource_pool: 'compute-pool',
          shm_size: null,
          slots_per_trial: 1,
          weight: 1,
        },
        scheduling_unit: 100,
        searcher: {
          max_length: { batches: 937 },
          metric: 'validation_loss',
          name: 'single',
          smaller_is_better: true,
          source_checkpoint_uuid: null,
          source_trial_id: null,
        },
        workspace: '',
      },
      description: '',
      endTime: '2022-08-01T16:17:57.473971Z',
      forkedFrom: 1240,
      hyperparameters: {
        'dropout1': { type: 'const', val: 0.25 },
        'dropout2': { type: 'const', val: 0.5 },
        'global_batch_size': { type: 'const', val: 64 },
        'learning_rate': { type: 'const', val: 1 },
        'n_filters1': { type: 'const', val: 32 },
        'n_filters2': { type: 'const', val: 64 },
        'test1.test2.test3.test4.optimizer_fake.lr': { type: 'const', val: 0.0001 },
        'test1.test2.test3.test4.optimizer_fake.momentum': { type: 'const', val: 0.9 },
      },
      id: 1241,
      jobId: '58c84330-69a1-4819-a92a-e4354358c876',
      jobSummary: null,
      labels: [],
      name: 'mnist_pytorch_const',
      notes: '',
      numTrials: 1,
      parentArchived: false,
      progress: 1,
      projectId: 102,
      projectName: 'Super Fun Project',
      resourcePool: 'compute-pool',
      searcherType: 'single',
      startTime: '2022-08-01T16:17:04.155545Z',
      state: 'COMPLETED',
      trialIds: [ 3521 ],
      userId: 34,
      workspaceId: 103,
      workspaceName: 'Caleb',
    },
    getExpTrials: {
      pagination: { endIndex: 1, limit: 2, offset: 0, startIndex: 0, total: 1 },
      trials: [ {
        autoRestarts: 0,
        bestAvailableCheckpoint: {
          endTime: '2022-08-01T16:17:52.922457Z',
          resources: {
            'code/': 0,
            'code/.ipynb_checkpoints/': 0,
            'code/adaptive.yaml': 1067,
            'code/adaptive-fast.yaml': 1099,
            'code/adsf.yaml': 1047,
            'code/bonst.yaml': 434,
            'code/checkpoints/': 0,
            'code/checkpoints/0669b753-fcea-4ccd-b894-18a23d538e27/': 0,
            'code/const.yaml': 1063,
            'code/data.py': 1449,
            'code/distributed.yaml': 499,
            'code/layers.py': 568,
            'code/model_def.py': 3745,
            'code/prof.yaml': 1865,
            'code/README.md': 1407,
            'code/tmp/': 0,
            'code/tmp/MNIST/': 0,
            'code/tmp/MNIST/processed/': 0,
            'code/tmp/MNIST/processed/test.pt': 7920407,
            'code/tmp/MNIST/processed/training.pt': 47520407,
            'code/tmp/MNIST/pytorch_mnist.tar.gz': 11613630,
            'load_data.json': 2898,
            'state_dict.pth': 14419735,
            'workload_sequencer.pkl': 91,
          },
          state: 'COMPLETED',
          totalBatches: 937,
          uuid: '59b63baf-abf2-47ae-af04-3133f8d18a84',
        },
        bestValidationMetric: {
          endTime: '2022-08-01T16:17:55.711568Z',
          metrics: { accuracy: 0.9800955414012739, validation_loss: 0.05823593354751697 },
          totalBatches: 937,
        },
        endTime: '2022-08-01T16:17:57.437644Z',
        experimentId: 1241,
        hyperparameters: {
          'dropout1': 0.25,
          'dropout2': 0.5,
          'global_batch_size': 64,
          'learning_rate': 1,
          'n_filters1': 32,
          'n_filters2': 64,
          'test1.test2.test3.test4.optimizer_fake.lr': 0.0001,
          'test1.test2.test3.test4.optimizer_fake.momentum': 0.9,
        },
        id: 3521,
        latestValidationMetric: {
          endTime: '2022-08-01T16:17:55.711568Z',
          metrics: { accuracy: 0.9800955414012739, validation_loss: 0.05823593354751697 },
          totalBatches: 937,
        },
        startTime: '2022-08-01T16:17:06.046449Z',
        state: 'COMPLETED',
        totalBatchesProcessed: 937,
        workloads: [],
      } ],
    },
    getExpValidationHistory: [
      {
        endTime: '2022-08-01T16:17:55.711568Z',
        trialId: 3521,
        validationError: 0.058235932,
      },
    ],
    getProject: {
      archived: false,
      description: 'SUPER FUN',
      id: 102,
      immutable: false,
      lastExperimentStartedAt: '2022-08-05T05:45:35.495201Z',
      name: 'Super Fun Project',
      notes: [ { contents: '# Does this work?', name: 'My Custom Notes' } ],
      numActiveExperiments: 0,
      numExperiments: 7,
      userId: 34,
      workspaceId: 103,
      workspaceName: 'Caleb',
    },
    getTrialDetails: {
      trial: {
        bestCheckpoint: {
          endTime: '2022-08-01T16:17:52.922457Z',
          resources: {
            'code/': '0',
            'code/.ipynb_checkpoints/': '0',
            'code/adaptive.yaml': '1067',
            'code/adaptive-fast.yaml': '1099',
            'code/adsf.yaml': '1047',
            'code/bonst.yaml': '434',
            'code/checkpoints/': '0',
            'code/checkpoints/0669b753-fcea-4ccd-b894-18a23d538e27/': '0',
            'code/const.yaml': '1063',
            'code/data.py': '1449',
            'code/distributed.yaml': '499',
            'code/layers.py': '568',
            'code/model_def.py': '3745',
            'code/prof.yaml': '1865',
            'code/README.md': '1407',
            'code/tmp/': '0',
            'code/tmp/MNIST/': '0',
            'code/tmp/MNIST/processed/': '0',
            'code/tmp/MNIST/processed/test.pt': '7920407',
            'code/tmp/MNIST/processed/training.pt': '47520407',
            'code/tmp/MNIST/pytorch_mnist.tar.gz': '11613630',
            'load_data.json': '2898',
            'state_dict.pth': '14419735',
            'workload_sequencer.pkl': '91',
          },
          state: 'STATE_COMPLETED',
          totalBatches: 937,
          uuid: '59b63baf-abf2-47ae-af04-3133f8d18a84',
        },
        bestValidation: {
          endTime: '2022-08-01T16:17:55.711568Z',
          metrics: {
            accuracy: 0.9800955414012739,
            validation_loss: 0.05823593354751697,
          },
          numInputs: 0,
          state: 'STATE_COMPLETED',
          totalBatches: 937,
        },
        endTime: '2022-08-01T16:17:57.437644Z',
        experimentId: 1241,
        hparams: {
          dropout1: 0.25,
          dropout2: 0.5,
          global_batch_size: 64,
          learning_rate: 1,
          n_filters1: 32,
          n_filters2: 64,
          test1: { test2: { test3: { test4: { optimizer_fake: { lr: 0.0001, momentum: 0.9 } } } } },
        },
        id: 3521,
        latestTraining: {
          endTime: '2022-08-01T16:17:50.142291Z',
          metrics: { loss: 0.0822979062795639 },
          numInputs: 0,
          state: 'STATE_COMPLETED',
          totalBatches: 937,
        },
        latestValidation: {
          endTime: '2022-08-01T16:17:55.711568Z',
          metrics: { accuracy: 0.9800955414012739, validation_loss: 0.05823593354751697 },
          numInputs: 0,
          state: 'STATE_COMPLETED',
          totalBatches: 937,
        },
        restarts: 0,
        runnerState: 'validating',
        startTime: '2022-08-01T16:17:06.046449Z',
        state: 'STATE_COMPLETED',
        taskId: '1241.1b6b71b5-373c-4f4d-be84-e0287c369b66',
        totalBatchesProcessed: 937,
        totalCheckpointSize: '81491411',
        wallClockTime: 51.286961,
        warmStartCheckpointUuid: '',
      },
      workloads: [
        {
          training: {
            endTime: '2022-08-01T16:17:28.038868Z',
            metrics: { loss: 0.634922206401825 },
            numInputs: 0,
            state: 'STATE_COMPLETED',
            totalBatches: 100,
          },
        },
        {
          training: {
            endTime: '2022-08-01T16:17:30.644858Z',
            metrics: { loss: 0.25738415122032166 },
            numInputs: 0,
            state: 'STATE_COMPLETED',
            totalBatches: 200,
          },
        },
        {
          training: {
            endTime: '2022-08-01T16:17:33.318785Z',
            metrics: { loss: 0.18983736634254456 },
            numInputs: 0,
            state: 'STATE_COMPLETED',
            totalBatches: 300,
          },
        },
        {
          training: {
            endTime: '2022-08-01T16:17:35.960871Z',
            metrics: { loss: 0.1480165719985962 },
            numInputs: 0,
            state: 'STATE_COMPLETED',
            totalBatches: 400,
          },
        },
        {
          training: {
            endTime: '2022-08-01T16:17:38.557971Z',
            metrics: { loss: 0.1530497521162033 },
            numInputs: 0,
            state: 'STATE_COMPLETED',
            totalBatches: 500,
          },
        },
        {
          training: {
            endTime: '2022-08-01T16:17:41.180345Z',
            metrics: { loss: 0.1253373771905899 },
            numInputs: 0,
            state: 'STATE_COMPLETED',
            totalBatches: 600,
          },
        },
        {
          training: {
            endTime: '2022-08-01T16:17:43.825716Z',
            metrics: { loss: 0.1267005056142807 },
            numInputs: 0,
            state: 'STATE_COMPLETED',
            totalBatches: 700,
          },
        },
        {
          training: {
            endTime: '2022-08-01T16:17:46.529609Z',
            metrics: { loss: 0.13082438707351685 },
            numInputs: 0,
            state: 'STATE_COMPLETED',
            totalBatches: 800,
          },
        },
        {
          training: {
            endTime: '2022-08-01T16:17:49.093888Z',
            metrics: { loss: 0.09980783611536026 },
            numInputs: 0,
            state: 'STATE_COMPLETED',
            totalBatches: 900,
          },
        },
        {
          training: {
            endTime: '2022-08-01T16:17:50.142291Z',
            metrics: { loss: 0.0822979062795639 },
            numInputs: 0,
            state: 'STATE_COMPLETED',
            totalBatches: 937,
          },
        },
        {
          checkpoint: {
            endTime: '2022-08-01T16:17:52.922457Z',
            resources: {
              'code/': '0',
              'code/.ipynb_checkpoints/': '0',
              'code/adaptive.yaml': '1067',
              'code/adaptive-fast.yaml': '1099',
              'code/adsf.yaml': '1047',
              'code/bonst.yaml': '434',
              'code/checkpoints/': '0',
              'code/checkpoints/0669b753-fcea-4ccd-b894-18a23d538e27/': '0',
              'code/const.yaml': '1063',
              'code/data.py': '1449',
              'code/distributed.yaml': '499',
              'code/layers.py': '568',
              'code/model_def.py': '3745',
              'code/prof.yaml': '1865',
              'code/README.md': '1407',
              'code/tmp/': '0',
              'code/tmp/MNIST/': '0',
              'code/tmp/MNIST/processed/': '0',
              'code/tmp/MNIST/processed/test.pt': '7920407',
              'code/tmp/MNIST/processed/training.pt': '47520407',
              'code/tmp/MNIST/pytorch_mnist.tar.gz': '11613630',
              'load_data.json': '2898',
              'state_dict.pth': '14419735',
              'workload_sequencer.pkl': '91',
            },
            state: 'STATE_COMPLETED',
            totalBatches: 937,
            uuid: '59b63baf-abf2-47ae-af04-3133f8d18a84',
          },
        },
        {
          validation: {
            endTime: '2022-08-01T16:17:55.711568Z',
            metrics: { accuracy: 0.9800955414012739, validation_loss: 0.05823593354751697 },
            numInputs: 0,
            state: 'STATE_COMPLETED',
            totalBatches: 937,
          },
        },
      ],
    },
    getWorkspace: {
      archived: false,
      id: 103,
      immutable: false,
      name: 'Caleb',
      numExperiments: 7,
      numProjects: 1,
      pinned: true,
      userId: 34,
    },
  },
};

export default RESPONSES;
