def quicksort(arr):
    quicksort_helper(arr, 0, len(arr) - 1)


def quicksort_helper(arr, first, last):
    if first < last:
        split_point = partition(arr, first, last)

        quicksort_helper(arr, first, split_point - 1)
        quicksort_helper(arr, split_point + 1, last)


def partition(arr, first, last):
    pivot = arr[first]

    left_idx = first + 1
    right_idx = last

    done = False
    while not done:
        while left_idx <= right_idx and arr[left_idx] <= pivot:
            left_idx = left_idx + 1

        while arr[right_idx] >= pivot and right_idx >= left_idx:
            right_idx = right_idx - 1

        if right_idx < left_idx:
            done = True
        else:
            temp = arr[left_idx]
            arr[left_idx] = arr[right_idx]
            arr[right_idx] = temp

    temp = arr[first]
    arr[first] = arr[right_idx]
    arr[right_idx] = temp

    return right_idx


if __name__ == "__main__":
    arr = [30, 50, 20, 10, 40, 90, 80, 60, 70]
    quicksort(arr)
    print(arr)
