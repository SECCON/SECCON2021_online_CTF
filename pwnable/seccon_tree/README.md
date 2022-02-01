# pwnable: seccon_tree
## Question
Let's make your own tree!
`nc seccon-tree.quals.seccon.jp 30001`

Here is an example.

```
from seccon_tree import Tree

cat = Tree("cat")
lion = Tree("lion")
tiger = Tree("tiger")

cat.add_child_left(lion)
cat.add_child_right(tiger)

assert(cat.find("lion") is not None)
assert(lion.find("tiger") is None)
```
[seccon_tree](files/seccon_tree)

## Attachments
- [seccon_tree](files/seccon_tree)

## Flag
```
SECCON{h34p_m4n463m3n7_15_h4rd_f0r_hum4n5....}
```

## Points
- Initial: `1`
- Last: `393` (solved: `4`)

## Tags
- `author:moratorium08`
    