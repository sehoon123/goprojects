import sys
input = sys.stdin.readline

N, K = map(int, input().split())

is_prime = [1] * (N+1)

idx = 0
found = False
for i in range(2, N+1):
    if is_prime[i] != 0:
        mul = 1
        while i*mul < N+1:
            if is_prime[i*mul] != 0:
                is_prime[i*mul] = 0
                idx += 1
            if idx == K:
                print(i*mul)
                found = True
                break
            mul += 1
        if found:
            break
        
        mul = 1
    
