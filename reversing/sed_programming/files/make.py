# python3 make.py > checker.sed

import string
import random
import sys

FLAG = "SECCON{mARkOV_4Lg0Ri7hM}"
N = 13
# Rule 150 (0b10010110)
RULES = [
  "0", # 000
  "1", # 001
  "1", # 010
  "0", # 011
  "1", # 100
  "0", # 101
  "0", # 110
  "1", # 111
]
REPLACE_SYMBOLS = True
SHUFFLE = True

random.seed("J8fXEm5OOYx3wDcY")

"""
Overview

1-dimensional cellular automaton is used as encryption
(and implementing cellular automaton with Markov algorithm is proof of Turing completeness).
If N-th generation of input and the FLAG are equal, this script output "CONGRATS" and "WRONG" otherwise.

Step 1
  Replace characters of input with their binary representations.

  input:  FLAG
  output: 01000110 01001100 01000001 01000111

Step 2
  Add "N<>".

  input:  01000110010011000100000101000111
  output: 1101<>01000110010011000100000101000111

Step 3
  Move ">" to the end.

  input:  1101<>01000110010011000100000101000111
  output: 1101<01000110010011000100000101000111>

for (int i=n; i>0; i--) {

Step 4
  Pop up "-" and "@0".

  input:  1101<01000110010011000100000101000111>
  output: 1101-<@001000110010011000100000101000111>

Step 5
  i--.

  input:  1101-<@001000110010011000100000101000111>
  output: 1100<@001000110010011000100000101000111>

for (int j=0; j<len(input)*8; j++) {

Step 6
  Apply cellular automaton rules and move "@".
  If the result is 0 output "x" and "y" otherwise.

  input:  1100<@001000110010011000100000101000111>
  output: 1100<@y01000110010011000100000101000111>

Step 7
  Move "x" or "y" to the end.

  input:  1100<@y01000110010011000100000101000111>
  output: 1100<@01000110010011000100000101000111>y

Step 8
  Change "x" to "a" and "y" to "b".

  input:  1100<@01000110010011000100000101000111>y
  output: 1100<@01000110010011000100000101000111>b

}

Step 9
  Change "<>" to "<#".

  input:  1100<>bbbabaabbbbbaababbbaaabbabbababa
  output: 1100<#bbbabaabbbbbaababbbaaabbabbababa

Step 10
  Move "#" to the end and change "a" to "0" and "b" to "1".

  input:  1100<#bbbabaabbbbbaababbbaaabbabbababa
  output: 1100<11101001111100101110001101101010#

Step 11
  Change "#" to ">".

  input:  1100<11101001111100101110001101101010#
  output: 1100<11101001111100101110001101101010>

}

Step 12
  Change "<" to "t0t".
  "t0t" represents "true" and that n=0.

  input:  <11101001001000000100011000111010>
  output: t0t11101001001000000100011000111010>

Step 13
  Move t[n]t to right changing it to t[n+1]t if correct.

  input:  t0t11101001001000000100011000111010>
  output: t1t1101001001000000100011000111010>

  Change t[n]t to "f" (represents false) if incorrect.

  input:  t0t01101001001000000100011000111010>
  output: f1101001001000000100011000111010>

  Move "f" to the end.

  input:  f1101001001000000100011000111010>
  output: f

Step 14
  Change "t[len]t>" to "t" if no remainders.

  input:  t100000t>
  output: t

  Otherwise change "t[len]t" to "f".
  "f" is move to the end by Step 13.

  input:  t100000t00000000>
  output: f0000000>

Step 15
  Change "t" to "CONGRATS" and "f" to "WRONG".

  input:  t
  output: CONGRATS

  input:  f
  output: WRONG
"""

def int_to_bin(n, l=None):
  if l is None:
    return bin(n).replace("0b", "")
  else:
    return bin(n+(1<<l)).replace("0b1", "")

def make_symbols():
  A = "#-01<>@abftxy234"
  if REPLACE_SYMBOLS:
    S = []
    for i in range(len(A)):
      S += ["1"+int_to_bin(i).replace("0", "l").replace("1", "I")+"1"]
    random.shuffle(S)

    symbols = {}
    for a, s in zip(A, S):
      symbols[a] = s
  else:
    symbols = {}
    for a in A:
      symbols[a] = a
  return symbols

symbols = make_symbols()

# (from, to, constraints, terminal)
# constraints:
#   S[x] = (from, to, terminal, [y, z]) means that S[y] and S[z] must precede S[x].
S = []

# Step 3
pre = len(S)
S += [
  (">0", "0>"),
  (">1", "1>"),
]
step3 = list(range(0, len(S)))

# Step 5
S += [
  ( "0-", "-1"),
  ("01-", "00"),
  ("11-", "10"),
]
S += [("1-", "", [len(S)-2, len(S)-1])]

# Step 7
step7_x = []
step7_y = []
for a in "xy":
  for b in "01>ab":
    S += [(a+b, b+a)]
    if a=="x":
      step7_x += [len(S)-1]
    else:
      step7_y += [len(S)-1]

# Step 8
pre = len(S)
S += [
  ("x", "a", step7_x),
  ("y", "b", step7_y),
]
step8 = list(range(pre, len(S)))

# Step 9
S += [("<>", "<#", step3+step8)]

# Step 10
pre = len(S)
S += [
  ("#a", "0#"),
  ("#b", "1#"),
]
step10 = list(range(pre, len(S)))

# Step 11
S += [("#", ">", step10)]

# Step 6
R = ""
for i in range(8):
  if RULES[i]=="0":
    R += "x"
  else:
    R += "y"

S += [
  ("@000", f"@{R[0]}00", step8),
  ("@001", f"@{R[1]}01", step8),
  ("@010", f"@{R[2]}10", step8),
  ("@011", f"@{R[3]}11", step8),
  ("@100", f"@{R[4]}00", step8),
  ("@101", f"@{R[5]}01", step8),
  ("@110", f"@{R[6]}10", step8),
  ("@111", f"@{R[7]}11", step8),

  ("@00>", f">{R[0]}", step8),
  ("@01>", f">{R[2]}", step8),
  ("@10>", f">{R[4]}", step8),
  ("@11>", f">{R[6]}", step8),
]

# Step 4
S += [
  ("0<", "0-<@0", list(range(len(S)))),
  ("1<", "1-<@0", list(range(len(S)))),
]

# Step 12
S += [("<", "t0t", list(range(len(S))))]

# Step 13
target = "".join(int_to_bin(ord(c), 8) for c in FLAG)
print("flag:  ", target, file=sys.stderr)

for _ in range(N):
  target = "0"+target+"0"
  next = ""
  for i in range(len(FLAG)*8):
    next += RULES[int(target[i:i+3], 2)]
  target = next

print("target:", target, file=sys.stderr)

pre = len(S)
for i in range(len(target)):
  i0 = int_to_bin(i)
  i1 = int_to_bin(i+1)
  if target[i]=="0":
    S += [(f"t{i0}t0", f"t{i1}t")]
    S += [(f"t{i0}t1", "f")]
  else:
    S += [(f"t{i0}t1", f"t{i1}t")]
    S += [(f"t{i0}t0", f"f")]
step13 = list(range(pre, len(S)))

# Step 14
ilen = int_to_bin(len(target))
pre = len(S)
S += [(f"t{ilen}t>", "t")]
S += [(f"t{ilen}t", "f", [len(S)-1])]
S += [
  ("f0", "f"),
  ("f1", "f"),
  ("f>", "f"),
]
step14 = list(range(pre, len(S)))

# Replace symbols
for i in range(len(S)):
  s0 = "".join(symbols[a] for a in S[i][0])
  s1 = "".join(symbols[a] for a in S[i][1])
  if len(S[i])==2:
    S[i] = (s0, s1)
  else:
    S[i] = (s0, s1, S[i][2])

# Step 15
pre = len(S)
step_15_chars = ""
for c, s in [
  ("t", "CONGRATS"),
  ("f", "WRONG"),
]:
  # Each characters should be differ
  assert len(s)==len(set(s))
  c = symbols[c]
  step_15_chars += s
  p = len(S)
  S += [(s[-2]+c, s[-2]+s[-1], [], True)]
  for i in range(len(s)-2):
    S += [(s[i]+c, s[i]+s[i+1]+c)]
  S += [(c, s[0]+c, step13+step14+list(range(p, len(S))))]
step15 = list(range(pre, len(S)))

# Step 1
for c in string.ascii_letters+string.digits+"{}_":
  ok = True
  for k in symbols:
    if c in symbols[k]:
      ok = False
  if ok:
    to = int_to_bin(ord(c), 8)
    to = "".join(symbols[t] for t in to)
    if c in step_15_chars:
      const = step15
    else:
      const = []
    S += [(c, to, const)]

# Step 2
to = f"{int_to_bin(N)}<>"
to = "".join(symbols[t] for t in to)
S += [("^", to, list(range(len(S))))]

print("Number of rules:", len(S), file=sys.stderr)

## Add fake rules
#assert len(S)<1024
#
#for i in range(len(S), 1024):
#  n = random.randint(1, 8)
#  from_ = "".join(symbols[random.choice("234")] for _ in range(n))
#  n = random.randint(1, 8)
#  to = "".join(symbols[random.choice("#-01<>@abftxy234")] for _ in range(n))
#  terminal = random.randint(0, 7)==0
#  S += [(from_, to, [], terminal)]

# Set default valus
for i in range(len(S)):
  s = S[i]
  from_ = s[0]
  to = s[1]
  constraints = s[2] if len(s)>=3 else []
  terminal = s[3] if len(s)>=4 else False
  S[i] = (from_, to, constraints, terminal)

# Shuffle under constraints
if SHUFFLE:
  n = len(S)
  # Number of ancestors + 1
  # A node which have many ancestors should be chosen first for randomness.
  W = [None]*n
  R = [[] for _ in range(n)]
  for i in range(n):
    for c in S[i][2]:
      R[c] += [i]
  for i in range(n)[::-1]:
    W[i] = set([i])
    for c in R[i]:
      W[i] |= W[c]
  for i in range(n):
    W[i] = len(W[i])

  # Since step 7 executes many times, move it to head for speeding up
  for s in step7_x+step7_y:
    W[s] += 99999999

  S2 = []
  used = [False]*n
  while len(S2)<len(S):
    C = []
    for i in range(n):
      if not used[i]:
        ok = True
        for c in S[i][2]:
          if not used[c]:
            ok = False
        if ok:
          C += [i]
    s = 0
    for c in C:
      s += W[c]
    #p = 0
    #p = s-1
    p = random.randint(0, s-1)
    for c in C:
      if p<W[c]:
        S2 += [S[c]]
        used[c] = True
        break
      p -= W[c]
  S = S2

print("""#!/bin/sed -f

# Check flag format
# Some characters are used internally
/^SECCON{[02-9A-HJ-Z_a-km-z]*}$/!{
  cINVALID FORMAT
  b
}

:t""")

for s in S:
  from_, to, _constraints, terminal = s
  t = "t" if terminal else "tt"
  print(f"s/{from_}/{to}/;{t}")
