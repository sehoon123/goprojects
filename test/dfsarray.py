# from collections import deque
# import re

# arr = [
#     [1,2,3,4,5],
#     [6,7,8,9,10],
#     [11,12,13,14,15],
#     [16,17,18,19,20]
# ]

# visited = [[0]*len(arr[0]) for _ in range(len(arr))]

# dx = [-1, 0, 1, 0]
# dy = [0, 1, 0, -1]
# values = []

# def dfs(arr, x, y, visited, values):
#     if not (0 <= x < len(arr) and 0 <= y < len(arr[0]) and visited[x][y] == 0):
#         return []

#     visited[x][y] = 1
#     values.append(arr[x][y])

#     for i in range(4):
#         nx = x + dx[i]
#         ny = y + dy[i]

#         dfs(arr, nx, ny, visited, values)

#     return values


# # result = dfs(arr, 0, 0, visited, values)
# # print(result)
# print(values)


# q = deque()
# def bfs(arr, x, y, visited, values):
#     q.append([x, y])
#     values.append(arr[x][y])
#     visited[x][y] = 1

#     while q:
#         x, y= q.popleft()

#         for i in range(4):
#             nx = x + dx[i]
#             ny = y + dy[i]

#             if (0 <= nx < len(arr) and 0 <= ny < len(arr[0]) and visited[nx][ny] == 0):
#                 q.append([nx, ny])
#                 values.append(arr[nx][ny])
#                 visited[nx][ny] = 1


#     return values

# result = bfs(arr, 2, 2, visited, values)
# print(result)

# a = "1234adsf"
# print(a.isalnum())

# b = re.sub(r'[0-9]', '', a)
# print(b)

# calculator.py

class Calculator:
    def add(self, x, y):
        return x + y

    def subtract(self, x, y):
        return x - y

    def multiply(self, x, y):
        return x * y

    def divide(self, x, y):
        if y == 0:
            raise ValueError("Cannot divide by zero!")
        return x / y

# test_calculator.py

import unittest

class TestCalculator(unittest.TestCase):

    def setUp(self):
        # 각 테스트 메서드가 실행되기 전에 실행됩니다.
        self.calc = Calculator()

    def test_add(self):
        self.assertEqual(self.calc.add(3, 5), 8)
        self.assertEqual(self.calc.add(-1, 1), 0)
        self.assertEqual(self.calc.add(-1, -1), -2)

    def test_subtract(self):
        self.assertEqual(self.calc.subtract(5, 3), 2)
        self.assertEqual(self.calc.subtract(-1, -1), 0)

    def test_multiply(self):
        self.assertEqual(self.calc.multiply(2, 3), 6)
        self.assertEqual(self.calc.multiply(-2, 3), -6)

    def test_divide(self):
        self.assertEqual(self.calc.divide(10, 2), 5)
        self.assertRaises(ValueError, self.calc.divide, 10, 0)

    def test_divide_by_zero(self):
        with self.assertRaises(ValueError):
            self.calc.divide(5, 0)

if __name__ == '__main__':
    unittest.main()