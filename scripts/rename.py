import os

for i in range(1,14):

    path = f'/Users/emilioziniades/adventofcode/2021/day{i}/'

    if i <= 9:
        oldex = path + f'day{i}-example.txt'
        oldin = path + f'day{i}-input.txt'
    else:
        oldex = path + f'{i}-example.txt'
        oldin = path + f'{i}-input.txt'

    newex = path + f'{i}.ex'
    newin = path + f'{i}.in'
    
        
    os.rename(oldex, newex)
    os.rename(oldin, newin)
    #print(oldex, oldin)
    #print(newex, newin)
