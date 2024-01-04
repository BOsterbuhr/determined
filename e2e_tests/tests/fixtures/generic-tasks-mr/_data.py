import datasets

def get_dataset():
    return datasets.load_dataset("truthful_qa", "generation")["validation"]
