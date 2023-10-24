from typing import List


def bubble_sort(arr: List[int]):
    for i in range(0, len(arr) - 1):
        for j in range(i + 1, len(arr)):
            if arr[i] > arr[j]:
                arr[i], arr[j] = arr[j], arr[i]


if __name__ == "__main__":
    arr = [30, 50, 20, 10, 40, 90, 80, 60, 70]
    bubble_sort(arr)
    print(arr)
