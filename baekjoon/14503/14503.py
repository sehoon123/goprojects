import sys
input = sys.stdin.readline

n, m = map(int, input().split())
sx, sy, d = map(int, input().split())
MAP = [list(map(int, input().split())) for _ in range(n)]

# 0: north, 1: east, 2: south, 3: west

dx = [-1, 0, 1, 0]
dy = [0, 1, 0, -1]


def dfs(x, y, direction):
    count = 0
    can_clean = False
    if MAP[x][y] == 0:
        MAP[x][y] = 2
        count += 1
    
    for i in range(4):
        nx = x + dx[i]
        ny = y + dy[i]
        if 0 <= nx < n and 0 <= ny < m and MAP[nx][ny] == 0:
            can_clean = True
            break


    if not can_clean:
        nx = x - dx[direction]
        ny = y - dy[direction]
        if 0 <= nx < n and 0 <= ny < m and MAP[nx][ny] != 1:
            count += dfs(nx, ny, direction)
        else:
            return count

    else:
        direction = (direction+3)%4
        nx = x + dx[direction]
        ny = y + dy[direction]

        while not (0 <= nx < n and 0 <= ny < m and MAP[nx][ny] == 0):
            direction = (direction+3)%4
            nx = x + dx[direction]
            ny = y + dy[direction]
        
        count += dfs(nx, ny, direction)
        
    return count     

print(dfs(sx,sy,d))
