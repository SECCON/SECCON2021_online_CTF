key = b'NekoPunch'
enc = bytes.fromhex('6dbf84f73cf6a112268b09525ea550a665e21cb2e3e13af7e3ea0ecb52f5b9cda5b6522b1e978734553f1d7956d4af94bfc3f4d68c8fba9eeecf4035550b9106f70d57d1a6cdaf3211eaaa78d71a9038b71be621241e8b608a43b107f8860f543ab0189aa063800de4bae7d0b11045b8')

def ROTL(a, b):
    return ((a<<b) | (a>>(8-b))) & 0xff

def QRrev(s, a, b, c, d):
    s[a] ^= ROTL((s[d] + s[c]) & 0xff, 4)
    s[d] ^= ROTL((s[c] + s[b]) & 0xff, 3)
    s[c] ^= ROTL((s[b] + s[a]) & 0xff, 2)
    s[b] ^= ROTL((s[a] + s[d]) & 0xff, 1)

state = [
    ord('N'), ord('e'), ord('k'), ord('o'),
    ord('P'), ord('u'), ord('n'), ord('c'),
    0, 0, 0, 0,
    0, 0, 0, 0,
]

flag = ""
for j in range(0, len(enc), 16):
    state = list(enc[j:j+16])

    for rnd in range(128):
        # even
        QRrev(state, 15, 12, 13, 14)
        QRrev(state, 10, 11, 8, 9)
        QRrev(state, 5, 6, 7, 4)
        QRrev(state, 0, 1, 2, 3)
        # odd
        QRrev(state, 15, 3, 7, 11)
        QRrev(state, 10, 14, 2, 6)
        QRrev(state, 5, 9, 13, 1)
        QRrev(state, 0, 4, 8, 12)

    flag += ''.join(map(chr, state[8:]))

print(flag)
