import sys


input = sys.stdin.readline


def first_false(completed: list) -> int:
    for i in range(len(completed)):
        if not completed[i]:
            return i
    return -1


def make_team(completed: list, friends: list) -> int:
    first = first_false(completed)
    if first == -1:
        return 1

    ret = 0

    for i in range(first+1, len(completed)):
        if not completed[i] and i in friends[first]:
            completed[first] = completed[i] = True
            ret += make_team(completed, friends)
            completed[first] = completed[i] = False

    return ret


C = int(input())

for _ in range(C):
    n, m = map(int, input().split())
    friends = [[] for _ in range(n)]
    completed = [False] * n
    tmp = list(map(int, input().split()))
    for _ in range(m):
        a, b = tmp.pop(0), tmp.pop(0)
        friends[a].append(b)
        friends[b].append(a)

    print(make_team(completed, friends))


