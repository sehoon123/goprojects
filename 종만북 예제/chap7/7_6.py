import sys

input = sys.stdin.readline

N = int(input())

for _ in range(N):
    artists = list(input().strip())
    a_num = len(artists)
    fans = list(input().strip())
    f_num = len(fans)

    artists = ['1' if x == 'M' else '0' for x in artists]
    fans = ['1' if x == 'M' else '0' for x in fans]

    artists = ''.join(artists)
    fans = ''.join(fans)

    count = 0

    for i in range(f_num-a_num+1):
        temp = fans[i:i+a_num]

        if int(artists, 2) & int(temp, 2) == 0:
            count += 1

    print(count)



