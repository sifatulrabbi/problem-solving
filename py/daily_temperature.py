def dailyTemperature(temperatures: list[int]) -> list[int]:
    arr: list[int] = []
    for i, cur_temp in enumerate(temperatures):
        # print("    i=", i, " value=", cur_temp)
        arr.append(0)
        for j, next_temp in enumerate(temperatures):
            if j <= i:
                continue
            # print("    j=", j, " value=", next_temp)
            if next_temp > cur_temp:
                arr[i] = j - i
                break
    return arr


if __name__ == "__main__":
    tests = [
        [30, 38, 30, 36, 35, 40, 28],
        # [22, 21, 20],
    ]
    for t in tests:
        print("for:", t)
        print("result:", dailyTemperature(t))
        print()
