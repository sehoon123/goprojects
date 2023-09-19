import sys

input = sys.stdin.readline

n = int(input())

def restore(string: str) -> str:
    # if len(string) == 0:
    #     return ''
    head = string[0]
    tail = string[1:]
    if head == 'w' or head == 'b':
        return head
    else:
        ul = restore(tail)
        ur = restore(tail[len(ul):])
        ll = restore(tail[len(ul)+len(ur):])
        lr = restore(tail[len(ul)+len(ur)+len(ll):])
        return 'x' + ll + lr + ul + ur

for _ in range(n):
    print(restore(input().strip()))



