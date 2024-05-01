import sys
input = sys.stdin.readline

N, M = map(int, input().split())
graph = [list(map(int, input().split())) for _ in range(N)]

tetrominos = [
  # I shape
  [(0,0),(0,1),(0,2),(0,3)],
  [(0,0),(1,0),(2,0),(3,0)],

  # Square
  [(0,0),(0,1),(1,0),(1,1)],

  # L shape
  [(0,0),(1,0),(2,0),(2,1)],
  [(1,0),(1,1),(1,2),(0,2)],
  [(0,0),(0,1),(1,1),(2,1)],
  [(0,0),(0,1),(0,2),(1,0)],

  # reverse L
  [(0,1),(1,1),(2,1),(2,0)],
  [(0,0),(0,1),(0,2),(1,2)],
  [(0,0),(1,0),(2,0),(0,1)],
  [(0,0),(1,0),(1,1),(1,2)],

  # thunder
  [(0,0),(1,0),(1,1),(2,1)],
  [(1,0),(1,1),(0,1),(0,2)],

  # reverse thunder
  [(1,0),(2,0),(1,1),(0,1)],
  [(0,0),(0,1),(1,1),(1,2)],

  # mountain 
  [(0,0),(0,1),(0,2),(1,1)],
  [(0,0),(1,0),(2,0),(1,1)],
  [(0,1),(1,0),(1,1),(1,2)],
  [(1,0),(0,1),(1,1),(2,1)]
]

result = -1
for i in range(N):
    for j in range(M):
        for v in tetrominos:
            a = (v[0][0]+i, v[0][1]+j)
            b = (v[1][0]+i, v[1][1]+j)
            c = (v[2][0]+i, v[2][1]+j)
            d = (v[3][0]+i, v[3][1]+j)

            if 0 <= a[0] < N and 0 <= b[0] < N and 0 <= c[0] < N and 0 <= d[0] < N and 0 <= a[1] < M and 0 <= b[1] < M and 0 <= c[1] < M and 0 <= d[1] < M:
                tmp = graph[a[0]][a[1]]+graph[b[0]][b[1]]+graph[c[0]][c[1]]+graph[d[0]][d[1]]

                if tmp > result:
                    result = tmp
    
print(result)



            