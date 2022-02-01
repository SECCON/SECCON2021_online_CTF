def main():
    f = lambda n: n if n <= 1 else (f(n-1) + f(n-2)) % 10
    flag = input("Input the flag > ")
    s=1
    if flag.startswith("SECCON{") and flag.endswith("}") and len(flag) == 40:
      s = sum(abs(ord(c) - ord(str(f(i)))) for (i, c) in enumerate(flag[7:-1]))
    if not s:
      print(f"Correct! The flag is", flag)
    else:
      print("Wrong :(")

#def unpack():
#    import marshal, base64, types
#    f = open(__file__, "rb")
#    f.seek(12)
#    s = f.read()[4:]
#    exec(marshal.loads(base64.b64decode(
#                bytes([(x-45)%256 for x in s[2:2+s[0]*256+s[1]]])
#                )),
#            globals(), locals())

#print(fib(fib.__code__, 10))

'''
            [Disassembly]
                0       BUILD_LIST              0
                2       LOAD_FAST               0: .0
                4       FOR_ITER                28 (to 34)
                6       UNPACK_SEQUENCE         2
                8       STORE_FAST              1: i
                10      STORE_FAST              2: x
                12      LOAD_FAST               2: x
                14      LOAD_CONST              0: 7
                16      LOAD_FAST               1: i
                18      BINARY_MULTIPLY
                20      BINARY_SUBTRACT
                22      LOAD_CONST              1: 45
                24      BINARY_SUBTRACT
                26      LOAD_CONST              2: 256
                28      BINARY_MODULO
                30      LIST_APPEND             2
                32      JUMP_ABSOLUTE           4
                34      RETURN_VALUE
'''

import dis
o = dis.opmap
start_iter = 8
asm = [
        o["JUMP_ABSOLUTE"], 4,
        0,0,
        o["BUILD_LIST"], 0,
        o["LOAD_FAST"], 0,
        o["FOR_ITER"], 20,
        o["UNPACK_SEQUENCE"], 2,
        o["LOAD_CONST"], 0,
        o["BINARY_MULTIPLY"], 0,
        o["BINARY_SUBTRACT"], 0,
        o["LOAD_CONST"], 1,
        o["BINARY_SUBTRACT"], 0,
        o["LOAD_CONST"], 2,
        o["BINARY_MODULO"], 0,
        o["LIST_APPEND"], 2,
        o["JUMP_ABSOLUTE"], start_iter,
        o["RETURN_VALUE"],
        ]
unpack_lambda = bytes(asm)
dis.dis(unpack_lambda)

import marshal
import base64
import dis
import types
#body = bytes([(x - 45) % 256 for i, x in
#    enumerate(base64.b64encode(marshal.dumps(main.__code__)))])
body = bytes([(x + 7 * i + 45) % 256 for (i, x) in
    enumerate(base64.b64encode(marshal.dumps(main.__code__)))])

body = bytes([len(body) // 256, len(body) % 256]) + body
print(body)

extended_arg = dis.opmap["EXTENDED_ARG"]
jump_abs = dis.opmap["JUMP_ABSOLUTE"]
ret = dis.opmap["RETURN_VALUE"]
load_const = dis.opmap["LOAD_CONST"]

target = len(body) + 4

head = [extended_arg, target // 256, jump_abs, target % 256]
head = bytes(head)

import py_compile
py_compile.compile("x.py", "x.pyc")
f = open("x.pyc", "rb")
f.seek(12)
unpack_code = marshal.load(f)

tail = unpack_code.co_code + bytes([load_const, 1, ret])
#tail = bytes([load_const, 1, ret])

code = head + body + tail

'''
 |  code(argcount, kwonlyargcount, nlocals, stacksize, flags, codestring,
 |        constants, names, varnames, filename, name, firstlineno,
 |        lnotab[, freevars[, cellvars]])
'''

#dis.dis(code)

co_consts = list(unpack_code.co_consts)
co_consts.append(None)
tmp = []
for x in co_consts:
    if isinstance(x, types.CodeType):
        co_varnames = tuple(x for x in x.co_varnames if x not in ("x", "i"))
        co_code = unpack_lambda
        codeobj = types.CodeType(x.co_argcount, x.co_kwonlyargcount,
                x.co_nlocals, x.co_stacksize, x.co_flags, co_code, x.co_consts,
                x.co_names, co_varnames, "x", "x", 0, b"",
                x.co_freevars, x.co_cellvars)
        tmp.append(codeobj)
    else:
        tmp.append(x)
co_consts = tuple(tmp)
co_varnames = unpack_code.co_varnames
codeobj = types.CodeType(unpack_code.co_argcount, unpack_code.co_kwonlyargcount,
        unpack_code.co_nlocals, unpack_code.co_stacksize, unpack_code.co_flags, code, co_consts,
        unpack_code.co_names, co_varnames, "x", "x", 0, b"",
        unpack_code.co_freevars, unpack_code.co_cellvars)
print(unpack_code.co_names)
print(unpack_code.co_varnames)
print(unpack_code.co_nlocals)
print(unpack_code.co_consts)
#codeobj = types.CodeType(0, 0,
#        unpack_code.co_nlocals, 4000, 64, code, co_consts,
#        unpack_code.co_names, unpack_code.co_varnames, "x", "x", 10, b"")

#f = (lambda:x)
#f.__code__=codeobj
#print(f())

import marshal
header = b'3\r\r\n\xc8\x8b\x9aa\x91\x04\x00\x00'
code_dump = marshal.dumps(codeobj)
with open("chall.cpython-36.pyc", "wb") as f:
  f.write(header)
  f.write(code_dump)

