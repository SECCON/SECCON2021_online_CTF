import re

code = open("qchecker.rb").read()
code = code.replace(" ", "").replace("\n", "")

A = re.findall(r'b\[".*?"\]', code)[1]

def decode36base(x):
  n = 0
  for c in x:
    c = ord(c)
    if ord("0")<=c<=ord("9"):
      c -= ord("0")
    else:
      c -= ord("a")
      c += 10
    n = n*36+c
  return n

A = [decode36base(a) for a in A[3:-2].split("+")]
A = A[2:]
M = [2**64+13, 2**64+37, 2**64+51, 2**64+81]

for a, m in zip(A, M):
  print(f"flag % {hex(m)} = {hex(a)}")

def extgcd(a, b):
  if b==0:
    return 1, 0, a
  else:
    x, y, g = extgcd(b, a%b)
    return y, x-a//b*y, g

def CRT(A, M):
  sa = 0
  sm = 1
  for a, m in zip(A, M):
    x, y, g = extgcd(m, sm)
    sa = (sa*x*m+a*y*sm)//g
    sm *= m//g
  return sa%sm, sm

f, _mod = CRT(A, M)

print(f"flag = {hex(f)}")

flag = ""
while f>0:
  flag += chr(f%256)
  f //= 256
print(flag)
