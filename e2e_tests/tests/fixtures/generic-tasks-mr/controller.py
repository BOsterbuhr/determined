import itertools
import os
import pathlib
import time
from typing import Optional

from determined import experimental
from determined.common.api import bindings
from determined.common import context, util

import _data


MAX_WORKERS = 2


def _create_task(**config) -> str:
    client = experimental.Determined()
    sess = client._session

    config_text = util.yaml_safe_dump(config)

    # TODO(ilia): try `inheritContext` instead.
    context_directory = context.read_v1_context(pathlib.Path(os.getcwd()))

    parent_id = os.environ["DET_TASK_ID"]

    req = bindings.v1CreateGenericTaskRequest(
        config=config_text,
        contextDirectory=context_directory,
        parentId=parent_id,
        # TODO(ilia): Make project id optional, inherit from the parent for child tasks.
        projectId=1,
    )

    task_resp = bindings.post_CreateGenericTask(sess, body=req)
    print(f"spawned task {task_resp.taskId}")
    return task_resp.taskId


def _wait(task_id: str, timeout: Optional[float] = 60) -> None:
    print(f"start waiting for {task_id}")
    client = experimental.Determined()
    sess = client._session

    start = time.time()

    for i in itertools.count(0):
        resp = bindings.get_GetTask(sess, taskId=task_id)
        state = resp.task.taskState
        print(state)
        TERMINAL_STATES = [
            bindings.v1GenericTaskState.COMPLETED,
            bindings.v1GenericTaskState.CANCELED,
            bindings.v1GenericTaskState.ERROR,
        ]
        if state in TERMINAL_STATES:
            return state

        if timeout > 0 and (time.time() - start > timeout):
            bindings.post_KillGenericTask(sess, taskId=task_id, body=bindings.v1KillGenericTaskRequest(taskId=task_id))
            print(f"Killed task {task_id}")
            raise RuntimeError(f"timed out waiting for task {task_id} after {i} ticks")

        time.sleep(1.)


def launch_worker(rank: int, world_size: int) -> str:
    config = {
        "entrypoint": ["python", "mapper.py"],
        "resources": {
            "slots": 1,
        },
        "environment": {
            "environment_variables": [
                f"RANK={rank}",
                f"WORLD_SIZE={world_size}",
            ],
        },
    }
    return _create_task(**config)


def main():
    ds = _data.get_dataset()
    world_size = max(getattr(ds, "n_shards", 0), MAX_WORKERS)

    workers = []
    for rank in range(0, world_size):
        worker = launch_worker(rank, world_size)
        workers.append(worker)

    for task_id in workers:
        state = _wait(task_id, timeout=60)
        if state == bindings.v1GenericTaskState.COMPLETED:
            # TODO(ilia): get the results via a run/checkpoint.
            pass
        elif state == bindings.v1GenericTaskState.ERROR:
            raise RuntimeError(f"uh-oh, child task {task_id} has failed")
        elif state == bindings.v1GenericTaskState.CANCELED:
            raise RuntimeError(f"uh-oh, child task {task_id} has been cancelled which is unsupported")
        else:
            raise RuntimeError(f"uh-oh, child task {task_id} has returned unknown state: {state}")
    print("Done")


if __name__ == "__main__":
    main()
