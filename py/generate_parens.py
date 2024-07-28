class Solution:
    def generateParenthesis(self, n: int) -> list[str]:
        arr: list[str] = []
        for i in range(n):
            pass
        return arr


if __name__ == "__main__":
    s = Solution()
    for v in [2, 4, 6]:
        print("n =", v, "result =", s.generateParenthesis(v))
