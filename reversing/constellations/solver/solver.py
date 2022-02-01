# python3 solver.py ../files/constellations

import sys
import struct

data = open(sys.argv[1], "rb").read()

enc = []
for i in range(0xb1):
  t = 0xc95e8+i*0x10
  p, l = struct.unpack("<QQ", data[t:t+0x10])
  p -= 0x400000
  enc += [data[p:p+l].decode()]

C = [
  "Cancer",
  "Aquarius",
  "Pisces",
  "Aries",
  "Leo",
  "Virgo",
  "Capricorn",
  "Gemini",
  "Scorpio",
  "Sagittarius",
  "Libra",
  "Taurus",
]

flag = ""
for e in enc:
  e = e.split("_")
  x = 0
  for c in e:
    x = x*12+C.index(c)
  flag += chr(5*x%256)
print(flag)
