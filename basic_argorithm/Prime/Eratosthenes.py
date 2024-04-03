import sys

input = sys.stdin.readline

prime = [1] * 1000001

for i in range(len(prime)):
    if i <= 1:
        prime[i] = 0
    
    if prime[i] != 0:
        tmp = i
        mul = 2
        while tmp*mul < len(prime):
            prime[tmp*mul] = 0
            mul+=1
        
        mul = 2

print(prime[919])


