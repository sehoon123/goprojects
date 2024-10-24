A = [8,6,7,5,3,10,9]
S = 15

result = []
def solve(A, i, S, answer=[]):
    if S == 0:
        result.append(answer)
        return 

    if S < 0 or i < 0:
        return 
    

    solve(A, i-1, S-A[i], answer+[A[i]])
    solve(A, i-1, S, answer)

    return 

solve(A, len(A)-1, S, [])
print(result)
