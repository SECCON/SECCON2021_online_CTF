import sys

with open(sys.argv[1]) as f:
  data = f.read()

stack = []
for i in range(len(data)):
  if i%1000000==0:
    print("%.2f %%"%(i/len(data)*100))
  stack += [data[i]]
  if (len(stack)>=8 and
      stack[-8] in "<" and
      stack[-7] in "Ss" and
      stack[-6] in "Cc" and
      stack[-5] in "Rr" and
      stack[-4] in "Ii" and
      stack[-3] in "Pp" and
      stack[-2] in "Tt" and
      stack[-1] in ">"):
    for i in range(8):
      stack.pop()
print("".join(stack)[:-1])
