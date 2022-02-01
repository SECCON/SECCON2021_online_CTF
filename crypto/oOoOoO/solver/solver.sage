import os
import socket

def solve(FLAG, S, M, v, counter):
    mat = []
    for i in range(len(v)):
        vv = []
        for j in range(len(v)+1):
            vv.append(2 if i == j else 0)
        vv[-1] = v[i] * boost
        mat.append(vv)
    vv = []
    for i in range(len(v)+1):
        vv.append(-1)
    vv[-1] = -(S + counter * M) * boost
    # print("value:", S + counter * M)
    # vv[-1] = -cheet * boost
    mat.append(vv)

    mat = matrix(ZZ, mat).LLL()
    # print(mat)

    ans = ""
    for v in mat:
        res = []
        if v[-1] != 0:
            continue
        f = True

        for i in range(len(v)-1):
            if v[i] != 1 and v[i] != -1:
                f = False
            if v[i] == -1:
                res.append(i)

        if not f:
            continue

        ans = ""
        for j in reversed(range(len(FLAG))):
            if j in res:
                ans += chr(FLAG[j]).upper()
            else:
                ans += chr(FLAG[j]).lower()
    return ans

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.connect((os.getenv('SECCON_HOST'), int(os.getenv('SECCON_PORT'))))
s.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
msg = s.recv(1024).split()

M = int(msg[2])
S = int(msg[5])
FLAG = msg[8]

N = len(FLAG)
boost = 100

S = ((S - (int.from_bytes(FLAG, 'little') % M)) % M)

v = [1]
while len(v) < len(FLAG):
    v.append((v[-1] * 256) % M)

v = list(map(lambda x: (x * (ord('o') - ord('O')) % M), v))

for i in range(N):
    ans = solve(FLAG, S, M, v, i)
    if ans != "":
        s.send(ans.encode() + b"\n")
        break

print(s.recv(1024).decode('utf-8')[:-1])
