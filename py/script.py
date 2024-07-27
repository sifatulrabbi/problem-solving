class Solution:
    def isValid(self, s: str) -> bool:
        if len(s) < 1 or len(s) % 2 != 0:
            return False

        valid = True
        counterpart = {
            "(": ")",
            "{": "}",
            "[": "]",
        }
        matches: list[str] = []
        for p in s:
            if p in "[{(":
                matches.append(counterpart[p])
            elif p in ")}]":
                if len(matches) < 1 or matches.pop() != p:
                    valid = False
                    break

        return valid if len(matches) < 1 else False


class MinStack:
    _stack: list[int]
    _min_val_idx: int

    def __init__(self):
        self._stack = []
        self._min_val_idx = -1

    def push(self, val) -> None:
        if not isinstance(val, int):
            # self._stack.append(None)
            return

        self._stack.append(val)

        if self._min_val_idx < 0:
            self._min_val_idx = len(self._stack) - 1
        elif val < self.getMin():
            self._min_val_idx = len(self._stack) - 1

    def pop(self) -> None:
        self._stack.pop()

        if self._min_val_idx <= len(self._stack) - 1:
            return
        minid = 0
        for i in range(len(self._stack)):
            if self._stack[i] < self._stack[minid]:
                minid = i
        self._min_val_idx = minid

    def top(self) -> int:
        return self._stack[len(self._stack) - 1]

    def getMin(self) -> int:
        if self._min_val_idx < 0:
            return 0
        return self._stack[self._min_val_idx]


if __name__ == "__main__":
    s = Solution()
    # print(f"expected False for ''. got '{s.isValid('')}'")
    # print(f"expected True for '{'[({})]'}'. got '{s.isValid('[({})]')}'")
    # print(f"expected False for '{'[(])'}'. got '{s.isValid('[({})]')}'")
    # print(f"expected True for '{'()[]{}'}'. got '{s.isValid('()[]{}')}'")

    # minstack = MinStack()
    # print(minstack.push(1))
    # print(minstack.push(2))
    # print(minstack.push(0))
    # print(minstack.getMin())
    # print(minstack.pop())
    # print(minstack.top())
    # print(minstack.getMin())
