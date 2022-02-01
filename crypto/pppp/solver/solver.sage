from Crypto.Util.number import *

with open('output.txt') as f:
    lines = f.readlines()
    n = eval(lines[0].split()[-1])
    e = eval(lines[1].split()[-1])
    c = eval("".join(lines[2].split()[2:]))

c = matrix(Zmod(n), c)

p = int(gcd(int(c[0][0]), n))
assert int(n) % p == 0
q = n // p

phi = (p-1) * (q-1)
d = int(pow(e, -1, phi))
m = c^d

m1 = long_to_bytes(gcd((int(m[1][2]), int(m[1][3]))))
m2 = long_to_bytes(gcd((int(m[2][2]), int(m[2][3]))))
print((m1+m2).decode('utf-8'))