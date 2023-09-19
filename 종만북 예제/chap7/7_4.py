import sys

input = sys.stdin.readline

C = int(input())


def solve(left, right):
    if left == right:
        return fences[left]

    mid = (left + right) // 2
    ret = max(solve(left, mid), solve(mid+1, right))

    lo, hi = mid, mid+1
    height = min(fences[lo], fences[hi])
    ret = max(ret, height * 2)

    while left < lo or hi < right:
        if hi < right and (lo == left or fences[lo-1] < fences[hi+1]):
            hi += 1
            height = min(height, fences[hi])
        else:
            lo -= 1
            height = min(height, fences[lo])

        ret = max(ret, height * (hi - lo + 1))

    return ret


for _ in range(C):
    N = int(input())
    fences = list(map(int, input().split()))
    print(solve(0, N-1))
