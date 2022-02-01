cipher = [75, 203, 190, 126, 184, 169, 27, 74, 35, 83, 113, 65, 207, 193, 27, 137, 37, 98, 0, 68, 219, 113, 21, 180, 223, 135, 5, 129, 189, 200, 245, 100, 117, 62, 192, 101, 239, 92, 182, 136, 159, 235, 166, 90, 74, 133, 83, 78, 6, 225, 101, 103, 82, 78, 144, 205, 130, 238, 175, 245, 172, 62, 157, 176]
key = b"SECCON2021"

Sbox = [0xff - i for i in range(0x100)]
j = 0
for i in range(0x100):
    j = (j + Sbox[i] + key[i % 10]) % 0x100
    Sbox[i], Sbox[j] = Sbox[j], Sbox[i]

def FYinv(bits):
    for i in range(63, -1, -1):
        j = (i**3 % 67) % 64
        bits[i], bits[j] = bits[j], bits[i]

def Pinv(data, length):
    for i in range(length // 8):
        bits = []
        for j in range(8):
            bits += [(data[i*8+j] >> k) & 1 for k in range(8)]
        FYinv(bits)
        for j in range(8):
            c = 0
            for k in range(8):
                c |= bits[j*8+k] << k
            data[i*8+j] = c

def Sinv(Sbox, data, length):
    for i in range(length):
        data[i] = Sbox.index(data[i])

for rnd in range(10):
    for i in range(0x40):
        cipher[i] ^= key[9 - rnd]
    Pinv(cipher, 0x40)
    Sinv(Sbox, cipher, 0x40)

print(cipher)
print(''.join(map(chr, cipher)))
