from typing import List


def diagonal_difference(A: List[List[int]]) -> int:
    diagonal_one = sum(A[i][i] for i in range(len(A)))
    diagonal_two = sum(
        A[i][j] for i, j in zip(reversed(range(len(A))), range(len(A)))
    )

    return abs(diagonal_one - diagonal_two)


def main():
    A = [[11, 2, 4], [4, 5, 6], [10, 8, -12]]

    print(diagonal_difference(A))

    print("tests pass.")


if __name__ == "__main__":
    main()
