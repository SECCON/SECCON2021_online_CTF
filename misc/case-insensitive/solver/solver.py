import os
import socket 
import string
import bcrypt

def make_letters(cnt):
    res = 'ﬃ' * (cnt // 3)
    if cnt % 3 == 1:
        res += "s"
    elif cnt % 3 == 2:
        res += "ß"
    return (res.encode(), res.upper().encode())

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.connect((os.getenv('SECCON_HOST'), int(os.getenv('SECCON_PORT'))))
s.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)

msg = s.recv(1024) # ~~~ mode:

ans = b""

for i in range(71, -1, -1):
    s.send(b"1" + b"\n")
    s.recv(1024) # message: 

    msg1 = ''
    msg2 = ''

    msg1, msg2 = make_letters(i)
    s.send(msg1 + b"\n")
    res = s.recv(1024).split()
    mac = res[1]

    res = False

    for c in string.printable[:-5]:
        c = c.encode()
        if bcrypt.checkpw(msg2 + ans + c, mac):
            res = True
            ans += c
            break

    if not res:
        break

print(ans.decode('utf-8'))
