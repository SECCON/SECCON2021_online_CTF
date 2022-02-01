def f(data, i, j):
    t = data[i]
    data[i] = data[j]
    data[j] = t

def S(Sbox, data, length):
    for i in range(length):
        data[i] = Sbox[data[i]]

def FY(bitarray):
    for i in range(64):
        j = ((i * i * i) % 67) % 64
        f(bitarray, j, i)

def P(data, length):
    for i in range(length // 8):
        bitarray = array(64)
        for j in range(8):
            p = data[i*8+j]
            for k in range(8):
                bitarray[j*8+k] = p % 2
                p //= 2
        FY(bitarray)
        for j in range(8):
            p = 0
            b = 1
            for k in range(8):
                p |= bitarray[j*8+k] * b
                b *= 2
            data[i*8+j] = p

def check(flag, length):
    # Generate SBox (just like RC4)
    key = array(10)
    key[0] = 83
    key[1] = 69
    key[2] = 67
    key[3] = 67
    key[4] = 79
    key[5] = 78
    key[6] = 50
    key[7] = 48
    key[8] = 50
    key[9] = 49    
    Sbox = array(0x100)
    for i in range(0x100):
        Sbox[i] = 0xff - i
    j = 0
    for i in range(0x100):
        j = (j + Sbox[i] + key[i % 10]) % 0x100
        f(Sbox, i, j)

    if length % 8 != 0:
        length += 8 - (length % 8)
    for rnd in range(10):
        # Substitution
        S(Sbox, flag, length)
        # Permutation
        P(flag, length)
        # XOR
        for i in range(length):
            flag[i] = flag[i] ^ key[rnd]
    return compare(flag)

def compare(flag):
    ans = array(0x40)
    ans[0] = 75
    ans[1] = 203
    ans[2] = 190
    ans[3] = 126
    ans[4] = 184
    ans[5] = 169
    ans[6] = 27
    ans[7] = 74
    ans[8] = 35
    ans[9] = 83
    ans[10] = 113
    ans[11] = 65
    ans[12] = 207
    ans[13] = 193
    ans[14] = 27
    ans[15] = 137
    ans[16] = 37
    ans[17] = 98
    ans[18] = 0
    ans[19] = 68
    ans[20] = 219
    ans[21] = 113
    ans[22] = 21
    ans[23] = 180
    ans[24] = 223
    ans[25] = 135
    ans[26] = 5
    ans[27] = 129
    ans[28] = 189
    ans[29] = 200
    ans[30] = 245
    ans[31] = 100
    ans[32] = 117
    ans[33] = 62
    ans[34] = 192
    ans[35] = 101
    ans[36] = 239
    ans[37] = 92
    ans[38] = 182
    ans[39] = 136
    ans[40] = 159
    ans[41] = 235
    ans[42] = 166
    ans[43] = 90
    ans[44] = 74
    ans[45] = 133
    ans[46] = 83
    ans[47] = 78
    ans[48] = 6
    ans[49] = 225
    ans[50] = 101
    ans[51] = 103
    ans[52] = 82
    ans[53] = 78
    ans[54] = 144
    ans[55] = 205
    ans[56] = 130
    ans[57] = 238
    ans[58] = 175
    ans[59] = 245
    ans[60] = 172
    ans[61] = 62
    ans[62] = 157
    ans[63] = 176
    for i in range(0x40):
        if flag[i] != ans[i]:
            return 1
    return 0

def main():
    # Input flag
    s_flag = array(6)
    s_flag[0] = 70
    s_flag[1] = 76
    s_flag[2] = 65
    s_flag[3] = 71
    s_flag[4] = 58
    s_flag[5] = 32
    s_correct = array(9)
    s_correct[0] = 67
    s_correct[1] = 111
    s_correct[2] = 114
    s_correct[3] = 114
    s_correct[4] = 101
    s_correct[5] = 99
    s_correct[6] = 116
    s_correct[7] = 33
    s_correct[8] = 10
    s_wrong = array(9)
    s_wrong[0] = 87
    s_wrong[1] = 114
    s_wrong[2] = 111
    s_wrong[3] = 110
    s_wrong[4] = 103
    s_wrong[5] = 46
    s_wrong[6] = 46
    s_wrong[7] = 46
    s_wrong[8] = 10
    write(s_flag, 6)
    flag = array(0x40)
    length = read(flag, 0x40)

    # Check flag
    if check(flag, length) == 0:
        write(s_correct, 9)
    else:
        write(s_wrong, 9)

def write(data, length):
    for i in range(length):
        putc(data[i])

def read(data, length):
    for i in range(length):
        c = getc()
        if c == 0x0A:
            break
        data[i] = c
    return i
