arr = [1, 2, 3, 4, 5, 6, 7, 8, 9, 0]
lo, hi = 0, len(arr) - 1

while lo < hi:
    mid = (lo + hi) / 2
    print(f"mid: {mid}, low: {lo}, high: {hi}")
    lo = mid
