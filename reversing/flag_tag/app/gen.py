flag = b'SECCON{wh4ts_Ur_r3c0mm3nd3d_w4y_2_d3c0mp1l3_WASM?}'
flag += b'\x00'*(8 - (len(flag) % 8))

def ROTL(a, b):
    return ((a<<b) | (a>>(8-b))) & 0xff

def QR(s, a, b, c, d):
    s[b] ^= ROTL((s[a] + s[d]) & 0xff, 1)
    s[c] ^= ROTL((s[b] + s[a]) & 0xff, 2)
    s[d] ^= ROTL((s[c] + s[b]) & 0xff, 3)
    s[a] ^= ROTL((s[d] + s[c]) & 0xff, 4)

enc = ""
state = [
    ord('N'), ord('e'), ord('k'), ord('o'),
    ord('P'), ord('u'), ord('n'), ord('c'),
    0, 0, 0, 0,
    0, 0, 0, 0,
]
for j in range(0, len(flag), 8):
    block = flag[j:j+8]
    for i in range(8):
        state[8+i] = block[i]

    for rnd in range(128):
        # odd
        QR(state, 0, 4, 8, 12)
        QR(state, 5, 9, 13, 1)
        QR(state, 10, 14, 2, 6)
        QR(state, 15, 3, 7, 11)
        # even
        QR(state, 0, 1, 2, 3)
        QR(state, 5, 6, 7, 4)
        QR(state, 10, 11, 8, 9)
        QR(state, 15, 12, 13, 14)

    enc += ''.join(map(lambda x: f'{x:02x}', state))

print(enc)
