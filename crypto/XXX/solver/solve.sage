import ast

with open("output.txt") as f:
    p = ast.literal_eval(f.readline().strip())
    params = ast.literal_eval(f.readline().strip())

n = len(params) - 1
ts = [(params[i][0] - params[i+1][0]) % p for i in range(n)]
bs = [(params[i][1] - params[i+1][1]) % p for i in range(n)]

M = matrix(ZZ, n+2, n+2)
M.set_block(  0, 0, matrix.identity(n) * p)
M.set_block(  n, 0, matrix(ZZ, 1, n, ts))
M.set_block(n+1, 0, matrix(ZZ, 1, n, bs))
M[  n,   n] = 1
M[n+1, n+1] = p

L = M.LLL()

for row in L:
    try:
        x = int(abs(row[-2])).to_bytes(100, "big")
        print(x.strip(b"\0").decode())
        break
    except:
        pass
