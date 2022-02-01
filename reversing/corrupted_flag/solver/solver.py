with open("flag.txt.enc", "rb") as f:
    enc = f.read()

def bits(b):
    i = 0
    while len(b) * 8:
        yield (b[i//8] >> (i%8)) & 1
        i += 1

bs = bits(enc)
try:
    output = []
    while True:
        bits = []
        for i in range(7):
            bits.append(next(bs))
        c1 = bits[6] ^ bits[4] ^ bits[2] ^ bits[0]
        c2 = bits[5] ^ bits[4] ^ bits[1] ^ bits[0]
        c3 = bits[3] ^ bits[2] ^ bits[1] ^ bits[0]
        c = (c3 << 2) | (c2 << 1) | c1
        if c:
            bits[7-c] ^= 1
        output += [bits[0], bits[1], bits[2], bits[4]]
except:
    pass

flag= b''
for i in range(len(output) // 8):
    c = 0
    for j in range(8):
        c |= output[i*8+j] << j
    flag += bytes([c])

print(flag)
