import heapq

def dijkstra(graph, start):
    distances = [float('inf')] * len(graph)
    distances[start] = 0
    queue = []
    heapq.heappush(queue, (0, start))

    while queue:
        current_distance, current_node = heapq.heappop(queue)

        if current_distance > distances[current_node]:
            continue

        for adjacent, weight in graph[current_node]:
            distance = current_distance + weight

            if distance < distances[adjacent]:
                distances[adjacent] = distance
                heapq.heappush(queue, (distance, adjacent))

    return distances


lists = [[0, 2, 3], [0, 1, 2], [2, 3, 4],[ 3, 1, -6]]

graph = [[] for _ in range(len(lists))]

for i in range(len(lists)):
    graph[lists[i][0]].append((lists[i][1], lists[i][2]))

print(graph)

print(dijkstra(graph, 0))

def bellman_ford(start):
    distances = [float('inf')] * len(roads)
    distances[start] = 0

    for i in range(n):
        for road in roads:
            start, end, cost = road

            distance = distances[start] + cost
            if distance < distances[end]:
                distances[end] = distance
                if i == n-1:
                    return -1
    
    return distances
