import torch
import datasets
import random


class Dataset(torch.utils.data.Dataset):
    def __getitem__(self, idx):
        return random.randint(0, idx)

    def __len__(self):
        return 256


def get_dataset():
    use_hg = False
    if use_hg:
        return datasets.load_dataset("truthful_qa", "generation")["validation"]
    else:
        return Dataset()
