from ptrlib import *

with open("chall.elf", "rb") as f:
    buf = f.read()

p = Process(["objdump", "-S", "-M", "intel", "chall.elf"])
while True:
    l = p.recvline(timeout=0.5)
    if b'<trap>:' in l:
        break
    if b'<trap>' in l:
        r = l.decode().split('\t')
        size = len(r[1].split())
        addr = int(r[0][:-1], 16)
        buf = buf[:addr] + b'\x90'*size + buf[addr+size:]

with open("patched.elf", "wb") as f:
    f.write(buf)

p.recvline()
