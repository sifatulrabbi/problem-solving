class Solution:
    def longestConsecutive(self, nums: list[int]) -> int:
        if len(nums) < 1:
            return 0

        sorted_nums = sorted(nums)
        seq = 1
        curr_steak = 1

        for i, n in enumerate(sorted_nums[: len(sorted_nums) - 1]):
            diff = sorted_nums[i + 1] - n
            if diff != 1 and diff != 0:
                curr_steak = 1
                continue
            curr_steak += diff
            if curr_steak > seq:
                seq = curr_steak

        return seq


if __name__ == "__main__":
    s = Solution()
    inputs = [
        {
            "input": [9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6],
            "result": 7,
        },
        {
            "input": [2, 20, 4, 10, 3, 4, 5],
            "result": 4,
        },
        {
            "input": [0, 3, 2, 5, 4, 6, 1, 1],
            "result": 7,
        },
    ]
    for t in inputs:
        result = s.longestConsecutive(t["input"])
        print(
            f"PASSED: {t['input']} -> {result}"
            if result == t["result"]
            else f"FAILED: {t['input']} -> expected={t['result']} got={result}"
        )
