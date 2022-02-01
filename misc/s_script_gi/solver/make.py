"""
ulimit -s unlimited
python3 make.py
tar zcvf flag.tar.gz flag.txt
"""

import random
import sys

#small = True
small = False

if small:
  flag = "S3CC0N{dummy_flag>_<_pt>>PT><sCRIp<scr<scr<scr!pt>ipt>ipt>}"
  n = 400
else:
  flag = "SECCON{sanitizing_is_not_so_good><_escaping_is_better_iPt><SCript<ScrIpT<scRIp<scRI<Sc<scr!pt>}"
  n = 8*1024*1024

tree = [[]]

random.seed("m9XbuhM07B0D9YoYbjrwrTBgqs4yTDgqFWXVbvmQYlg7xCkOF3UIGrcLiFh0rvX0")

# expand random branches
print("phase 1")
for i in range(n*4//10):
  if i%10000==0:
    print(i)
  parent = random.randint(0, len(tree)-1)
  l = len(flag) if parent==0 else len("<script>")
  pos = random.randint(0, l)
  tree[parent] += [(pos, len(tree))]
  tree += [[]]

# expand one branch
print("phase 2")
target = len(tree)
if small:
  tree[0] += [(50, len(tree))]
else:
  tree[0] += [(73, len(tree))]
tree += [[]]
for i in range(n*5//10):
  if i%10000==0:
    print(i)
  l = len("<script>")
  pos = random.randint(1, l-1)
  tree[target] += [(pos, len(tree))]
  target = len(tree)
  tree += [[]]

# expand another branch
print("phase 3")
target = len(tree)
if small:
  tree[0] += [(43, len(tree))]
else:
  tree[0] += [(68, len(tree))]
tree += [[]]
for i in range(n*1//10):
  if i%10000==0:
    print(i)
  l = len("<script>")
  pos = random.randint(1, l-1)
  tree[target] += [(pos, len(tree))]
  target = len(tree)
  tree += [[]]

scripts = [""]*(1<<6)
for b in range(1<<6):
  scripts[b] += "<"
  for i in range(6):
    c = "script"[i]
    if b>>i&1:
      c = c.upper()
    scripts[b] += c
  scripts[b] += ">"

sys.setrecursionlimit(n)

max_depth = 0
if small:
  f = open("small.txt", "w")
else:
  f = open("flag.txt", "w")

def bt(node, depth):
  global max_depth
  max_depth = max(max_depth, depth)

  if node==0:
    s = flag
  else:
    s = random.choice(scripts)

  tree[node].sort()
  p = 0
  for i in range(len(s)+1):
    while p<len(tree[node]) and tree[node][p][0]==i:
      bt(tree[node][p][1], depth+1)
      p += 1
    if i<len(s):
      f.write(s[i])
  assert p==len(tree[node])
bt(0, 0)
f.write("\n")

print("max_depth:", max_depth)
