import random

def random_adjacency_matrix(n):   
    matrix = [[random.randint(1, 10) for i in range(n)] for j in range(n)]

    # No vertex connects to itself
    for i in range(n):
        matrix[i][i] = 0

    # If i is connected to j, j is connected to i
    for i in range(n):
        for j in range(n):
            matrix[j][i] = matrix[i][j]

    return matrix

N = 5
f = open(f"data", "w")
mat = random_adjacency_matrix(N)
for i in range(N) :
    for j in range(N) :
        f.write(f"{mat[i][j]} ")
    f.write("\n")