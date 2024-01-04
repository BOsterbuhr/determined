import os

import _data
from datasets.distributed import split_dataset_by_node

def main():
    rank = os.environ["RANK"]
    world_size = os.environ["WORLD_SIZE"]
    ds = split_dataset_by_node(_data.get_dataset(), rank=rank, world_size=world_size)
    ds_len = sum(1 for _ in ds)
    print(f"Dataset split size: {ds_len}")


if __name__ == "__main__":
    main()
