from seccon_tree import Tree

s_hoge = "hoge"
s_fuga = "fuga"
s_piyo = "piyo"

hoge = Tree(s_hoge)
fuga = Tree(s_fuga)
piyo = Tree(s_piyo)

hoge.add_child_left(fuga)
fuga.add_child_left(piyo)

assert(id(hoge.find("piyo").get_object()) == id(s_piyo))
assert(id(hoge.find("hoge").get_object()) == id(s_hoge))
assert(id(hoge.find("fuga").get_object()) == id(s_fuga))
assert(fuga.find("hoge") is None)

fuga_obj = hoge.find("fuga")
assert(id(fuga.get_child_left().get_object()) == id(s_piyo))

hoge.add_child_right(Tree(1))
assert(hoge.get_child_right().get_object() == 1)
hoge.add_child_right(Tree(2))

assert(hoge.get_child_right().get_object() == 2)

x = hoge.get_child_right()
del hoge
assert(x.get_object() == 2)


t = Tree(1)
try:
    t.get_object(1)
    assert(False)
except TypeError:
    pass
try:
    t.add_child_right(1, 1, 1)
    assert(False)
except TypeError:
    pass
try:
    t.add_child_left(1, 1, 1)
    assert(False)
except TypeError:
    pass
try:
    t.get_child_left(1, 1, 1)
    assert(False)
except TypeError:
    pass
try:
    t.get_child_right(1, 1, 1)
    assert(False)
except TypeError:
    pass
try:
    s = Tree(1,2,3)
    assert(False)
except TypeError:
    pass
try:
    t.add_child_left(1)
    assert(False)
except TypeError:
    pass

class Z:
    def __repr__(self):
        return 1

#t.add_child_right(Tree(Z()))
#try:
#    assert(t.find(0x100) is None)
#    assert(False)
#except TypeError:
#    pass


print("ok")
