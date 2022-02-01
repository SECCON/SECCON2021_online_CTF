import socket
import base64
from Crypto.Util.strxor import strxor
from Crypto.Util.Padding import pad, unpad
from Crypto.Util.number import bytes_to_long, long_to_bytes
import copy
import os

block_size = 16

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.connect((os.getenv("SECCON_HOST"), int(os.getenv("SECCON_PORT"))))
s.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
msg = s.recv(1024)

msg = msg.split()[-2]  # base64

ref_c = base64.b64decode(msg)

res = b""

prev = copy.deepcopy(ref_c[:16])
for k in range(4):
    ans = []
    for j in range(16):
        for i in range(len(ans)):
            ans[i] ^= (j) ^ (j + 1)

        for i in range(256):
            c = copy.deepcopy(ref_c)

            iv = c[:16]
            iv = iv[: (15 - j)] + bytearray([i]) + bytearray(ans)

            c = c[16:]
            c = [c[i : i + block_size] for i in range(0, len(c), block_size)]

            c.append(c[0])
            c.append(c[1])
            c.append(c[2])
            c.append(c[3])

            c.append(c[k])
            c = b"".join(c)
            s.send(base64.b64encode(iv + c) + b"\r\n")

            r = s.recv(1024)
            if b"Great" in r:
                ans = [i] + ans
                break

    m = strxor(strxor(bytes(ans), b"\x10" * 16), prev)
    res += m

    c = copy.deepcopy(ref_c)

    iv = c[:16]
    c = c[16:]
    c = [c[i : i + block_size] for i in range(0, len(c), block_size)]
    prev = strxor(m, c[k])

print(unpad(res, 16)[16:].decode())
