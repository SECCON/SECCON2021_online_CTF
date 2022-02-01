import random
from binascii import unhexlify
from output import sigs1, sigs2

def long_to_bytes(x):
    x = int(x)
    if x == 0:
        return b"\x00"
    return x.to_bytes((x.bit_length() + 7) // 8, "big")

p = 39402006196394479212279040100143613805079739270465446667948293404245721771496870329047266088258938001861606973112319
a = -3
b = 27580193559959705877849011840389048093056905856361568521428707301988689241309860865136260764883745107765439761230575
curve = EllipticCurve(GF(p), [a, b])
order = 39402006196394479212279040100143613805079739270465446667946905279627659399113263569398956308152294913554433653942643
Z_n = GF(order)
gx = 26247035095799689268623156744566981891852923491109213387815615900925518854738050089022388053975719786650872476732087
gy = 8325710961489029985546751289520108179287853048861315594709205902480503199884419224438643760392947333078086511627871
G = curve(gx, gy)

sig_cnt = 5
size = sig_cnt*2 + 2
mat = [
    [0 for _ in range(size)] for _ in range(size)
]

for i in range(sig_cnt):
    r,s = sigs1[i]
    s_inv = inverse_mod(s, order)
    mat[i][i] = order
    mat[i+sig_cnt][i+sig_cnt] = 1
    mat[i+sig_cnt][i] = -2^256
    mat[-2][i] = r * s_inv % order
    mat[-1][i] = (s_inv - 2^128)

mat[-2][-2] = 1/(2^256)
mat[-1][-1] = 1

l = matrix(QQ, mat)
llled = l.LLL()

for b in llled:
    d1 = int(b[-2] * 2^256) % order
    msg1 = b[-1]

    valid_msg = True
    for b in long_to_bytes(msg1):
        if b < 0x20 or b > 0x7f:
            valid_msg = False
    if not valid_msg:
        continue

    # print("[+] Found!!")
    # print(long_to_bytes(msg1))
    # print(long_to_bytes(d1))
    found = True
    break

if not found:
    print("?????????")
    exit()

ks = []
for r, s in sigs1:
    s_inv = inverse_mod(s, order)
    k = (msg1 + r*d1) * s_inv % order
    ks.append(int(k))

def untemper(x):
    x ^^= (x >> 18)
    x ^^= ((x << 15) & 0xefc60000)
    x_bottom_14 = (x ^^ (x << 7) & 0x9d2c5680) # & ((1 << 14) - 1)
    x_bottom_21 = (x ^^ (x_bottom_14 << 7) & 0x9d2c5680) # & ((1 << 21) - 1)
    x_bottom_28 = (x ^^ (x_bottom_21 << 7) & 0x9d2c5680) # & ((1 << 28) - 1)
    x ^^= (x_bottom_28 << 7) & 0x9d2c5680
    x_top_22 = x ^^ (x >> 11)
    x ^^= (x_top_22 >> 11)

    return int(x)


def k_split(k):
    ret = []
    for _ in range(4):
        ret.append(k & 0xffffffff)
        k >>= 32

    return ret

state = []
for k in ks:
    k_bottom = k & (2^128 - 1)
    k_top = (k >> 256)
    s1 = list(map(untemper, k_split(k_bottom)))
    s2 = list(map(untemper, k_split(k_top)))

    state = state + s1 + s2

mt_state = tuple(state[:624] + [624])
mt_state = (3, mt_state, None)
random.setstate(mt_state)
assert random.getrandbits(128) == ks[-2] & (2^128 - 1)
assert random.getrandbits(128) == ks[-2] >> 256
assert random.getrandbits(128) == ks[-1] & (2^128 - 1)
assert random.getrandbits(128) == ks[-1] >> 256

k1 = random.getrandbits(384)
k2 = random.getrandbits(384)
r1, s1 = sigs2[0]
r2, s2 = sigs2[1]
rdiff = r1 - r2

d2 = (k1*s1 - k2*s2) * inverse_mod(rdiff, order) % order
# print(long_to_bytes(d2))
flag = long_to_bytes(d1) + long_to_bytes(d2)
print(flag.decode("utf-8"))