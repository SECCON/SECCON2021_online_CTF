# python3 solver.py < ../files/checker.sed

import re
import string

rules = []
while True:
  try:
    line = input()
  except EOFError:
    break

  m = re.search("^s/(.*)/(.*)/;(tt?)$", line)
  if m:
    rules += [(m.group(1), m.group(2), m.group(3)=="tt")]

def output():
  with open("checker2.sed", "w") as f:
    f.write(":t\n")
    f.write("p\n")
    for from_, to, terminal in rules:
      t = "tt" if terminal else "t"
      f.write(f"s/{from_}/{to}/;{t}\n")

# Remove single lower letter or digit rules to avoid conflict with the following replaces
rules = [r for r in rules if not (len(r[0])==1 and (r[0].islower() or r[0].isdigit()))]

# Replace "1l1" -> "a", "1I1" -> "b", "1Il1" -> "c", ...
symbols = {}
for i in range(16):
  s = "1"+bin(i).replace("0b", "").replace("0", "l").replace("1", "I")+"1"
  symbols[s] = chr(ord("a")+i)

for i in range(len(rules)):
  from_, to, terminal = rules[i]
  for s in symbols:
    from_ = from_.replace(s, symbols[s])
    to = to.replace(s, symbols[s])
  rules[i] = (from_, to, terminal)

#output()

# Find "0" and "1"
for from_, to, terminal in rules:
  if from_=="S":
    s0 = to[0]
    s1 = to[1]

for i in range(len(rules)):
  from_, to, terminal = rules[i]
  from_ = from_.replace(s0, "0").replace(s1, "1")
  to = to.replace(s0, "0").replace(s1, "1")
  rules[i] = (from_, to, terminal)

#output()

# Find N
for from_, to, terminal in rules:
  if from_=="^":
    N = int(to[:-2], 2)
print("N:", N)

# Find cell automaton rules
S = {
  "e": "0",
  "i": "1",
}
print("Rules:")
for from_, to, terminal in rules:
  if re.search("^j[01]{3}$", from_):
    print(f"  {from_[1:]} -> {S[to[1]]}")

# Find target
t = "d"
target = [""]*1024
for i in range(1024):
  tnt = t+bin(i).replace("0b", "")+t
  for from_, to, terminal in rules:
    if from_==tnt+"0" and to[0]==t:
      target[i] = "0"
    if from_==tnt+"1" and to[0]==t:
      target[i] = "1"
target = "".join(target)
print("target:", target)

# Play back cell automaton

# Gen n  : ... T[n  ][i-1] T[n  ][i  ] T[n  ][i+1] ...
# Gen n+1: ... T[n+1][i-1] T[n+1][i  ] T[n+1][i+1] ...
#
# Since T[n+1][i] = T[n][i-1] ^ T[n][i] ^ T[n][i+1],
# T[n][i+1] = T[n][i-1] ^ T[n][i] ^ T[n+1][i].
# First several bytes of the flag is clear.
# So, this cell automaton can be reversed.

n = len(target)
flag = "SECCON{"+"X"*(n//8-8)+"}"
flag = "".join(bin(ord(f)+0x100).replace("0b1", "") for f in flag)

T = [list(flag)]
for i in range(N):
  t = ["0"]+T[i]+["0"]
  T += [[]]
  for j in range(n):
    T[i+1] += [str(t[j:j+3].count("1")%2)]
T[N] = list(target)

for i in range(N)[::-1]:
  for j in range(1, n-1):
    T[i][j+1] = str((T[i][j-1:j+1]+[T[i+1][j]]).count("1")%2)

flag = "".join(T[0])
flag = "".join(chr(int("".join(flag[i:i+8]), 2)) for i in range(0, n, 8))
print("flag:", flag)
