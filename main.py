import numpy as np
import timeit


def vector_similarity(v1, v2):
    dot = np.dot(v1, np.transpose(v2))
    norm = np.linalg.norm(v1) * np.linalg.norm(v2)
    sim = dot / norm  # between [-1, +1]
    sim = (sim + 1) / 2

    return sim


def main():
    dimension = 128
    number_of_vector = 20000  # how many person in db

    all_members = {}
    for i in range(number_of_vector):
        all_members[i] = np.random.rand(dimension)

    compare_vector = np.random.rand(dimension)

    # compare start
    start = timeit.default_timer()
    for _, vector in all_members.items():
        sim = vector_similarity(vector, compare_vector)
        # print(sim)
    print(f"{timeit.default_timer() - start} sec.")


if __name__ == "__main__":
    main()
