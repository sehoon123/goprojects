from functools import cache
n = int(input())

@cache
def fibo(n: int) -> int:
    if n < 2:
        return n
    
    return fibo(n-1) + fibo(n-2)

print(fibo(n))