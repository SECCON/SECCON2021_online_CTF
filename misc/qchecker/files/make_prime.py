from Crypto.Util.number import *

def next_prime(n):
  n += 1
  while not isPrime(n):
    n += 1
  return n

m1 = next_prime(2**64)
m2 = next_prime(m1)
m3 = next_prime(m2)
m4 = next_prime(m3)
print(hex(m1))
print(hex(m2))
print(hex(m3))
print(hex(m4))
