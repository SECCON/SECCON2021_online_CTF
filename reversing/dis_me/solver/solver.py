import dis
import types
import marshal
F = "../files/chall.cpython-36.pyc"
f = open(F, "rb")
f.seek(12)
c = marshal.load(f)
code = c.co_code[898:]

c = types.CodeType(c.co_argcount, c.co_kwonlyargcount,
        c.co_nlocals, c.co_stacksize, c.co_flags, code, c.co_consts,
        c.co_names, c.co_varnames, "x", "x", 0, b"",
        c.co_freevars, c.co_cellvars)

import marshal
header = b'3\r\r\n\xc8\x8b\x9aa\x91\x04\x00\x00'
code_dump = marshal.dumps(c)
with open("fixed.pyc", "wb") as f:
  f.write(header)
  f.write(code_dump)

'''
$ uncompyle6 fixed.pyc
# uncompyle6 version 3.8.0
# Python bytecode 3.6 (3379)
# Decompiled from: Python 3.6.9 (default, Sep  4 2021, 04:57:41)
# [GCC 9.3.0]
# Embedded file name: x
# Compiled at: 2021-11-21 18:11:20
# Size of source mod 2**32: 1169 bytes
import marshal, base64, types
f = open(__file__, 'rb')
f.seek(12)
s = marshal.loads(f.read()).co_code[4:]
types.FunctionType(marshal.loads(base64.b64decode(bytes([(x - 45) % 256 for x in s[2:2 + s[0] * 256 + s[1]]]))), globals())()
del s
del f
del marshal
del base64
del types
return
# okay decompiling fixed.pyc
'''


import base64
f = open(F, "rb")
f.seek(12)
s = marshal.loads(f.read()).co_code[4:]
code_dump = base64.b64decode(bytes([(x - 45) % 256 for x in s[2:2 + s[0] * 256 + s[1]]]))

with open("fixed-func.pyc", "wb") as f:
  f.write(header)
  f.write(code_dump)


'''
# uncompyle6 version 3.8.0
# Python bytecode 3.6 (3379)
# Decompiled from: Python 3.6.9 (default, Sep  4 2021, 04:57:41)
# [GCC 9.3.0]
# Embedded file name: run.py
# Compiled at: 2021-11-21 18:11:20
# Size of source mod 2**32: 1169 bytes
f = lambda n: n if n <= 1 else (f(n - 1) + f(n - 2)) % 10
flag = input('Input the flag > ')
s = 1
if flag.startswith('SECCON{'):
    if flag.endswith('}'):
        if len(flag) == 40:
            s = sum(abs(ord(c) - ord(str(f(i)))) for i, c in enumerate(flag[7:-1]))
    s or print('Correct! The flag is', flag)
else:
    print('Wrong :(')
# okay decompiling fixed-func.pyc
'''
def fib(n):
    if n <= 1:
      return n
    return (fib(n-1)+fib(n-2)) % 10

flag = "SECCON{"
for i in range(32):
    flag += str(fib(i))
flag += "}"
print(flag)
