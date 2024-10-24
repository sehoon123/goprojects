s = [2,5,10,5]
p = [40,30,50,10]

k = 16

dp = [[0]*(k+1) for _ in range(len(s)+1)]

for i in range(1,len(s)+1):
    for j in range(1, k+1):
        if s[i-1] <= j:
            dp[i][j] = max(dp[i-1][j], dp[i-1][j-s[i-1]]+p[i-1])
        else:
            dp[i][j] = dp[i-1][j]

# print(dp)

def fractional_kanpsack(s, p, space_left, start):
    n = len(s)
    items = sorted([(p[i], s[i]) for i in range(start, n)], key=lambda x: x[0]/x[1], reverse=True)

    total_value = 0
    for value, weight in items:
        if weight <= space_left:
            space_left -= weight
            total_value += value
        else:
            total_value += (value/weight) * space_left
            break

    return total_value

def knapsack(s, p, space_left, i, current_profit, MP):
    if i < 0 or space_left <= 0:
        return current_profit
    
    upper_bound = fractional_kanpsack(s, p, space_left, i)

    if current_profit + upper_bound <= MP:
        return current_profit
    
    exclude = knapsack(s, p, space_left, i-1,current_profit, MP)

    include = current_profit
    if s[i] <= space_left:
        include = knapsack(s, p, space_left-s[i], i-1, current_profit+p[i], MP)
    
    MP = max(MP, exclude, include)

    return MP

result = knapsack(s, p, k, len(s)-1, 0, 0)
print(result)