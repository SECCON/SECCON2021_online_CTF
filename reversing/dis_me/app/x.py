import marshal, base64, types
f = open(__file__, "rb")
f.seek(12)
s = marshal.loads(f.read()).co_code[4:]

types.FunctionType(marshal.loads(base64.b64decode(
            bytes([(x- 7 *i -45)%256 for i, x in
                enumerate(s[2:2+s[0]*256+s[1]])])
            )), globals())()

#types.FunctionType(marshal.loads(base64.b64decode(bytes(tmp))), globals())()

del s
del f
del marshal
del base64
del types
