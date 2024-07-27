class Solution:
    operand_fns = {
        "+": lambda l, r: l + r,
        "-": lambda l, r: l - r,
        "*": lambda l, r: l * r,
        "/": lambda l, r: l / r,
    }

    def evalRPN(self, tokens: list[str]) -> int:
        final_arr = self.parse_rpn(tokens)
        while len(final_arr) > 1:
            final_arr = self.parse_rpn(final_arr)

        result = self.isint(final_arr[0])
        return result if result else 0

    def parse_rpn(self, tokens: list[str]) -> list[str]:
        # print("    tokens ->", tokens)
        if len(tokens) < 3:
            return tokens

        l = self.isint(tokens[0])
        r = self.isint(tokens[1])
        fn = self.operand_fns.get(tokens[2])

        if not fn:
            arr = self.parse_rpn(tokens[1:])
            return self.parse_rpn([tokens[0], *arr])

        result = fn(l, r)
        return [str(result), *tokens[3:]]

    def isint(self, v: str):
        try:
            if "." in v:
                return int(v.split(".")[0])
            return int(v)
        except:
            return None


if __name__ == "__main__":
    print("testing evalRPN()")
    s = Solution()
    arr = [
        ["2", "1", "+", "3", "*"],
        ["4", "13", "5", "/", "+"],
        ["4", "2", "3", "+", "*"],
        ["12", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"],
    ]
    for a in arr:
        print("testing:", a, "->", s.evalRPN(a))
