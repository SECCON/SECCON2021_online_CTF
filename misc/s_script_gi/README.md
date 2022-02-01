# misc: s/<script>//gi
## Question
Can you figure out why `s/<script>//gi` is insufficient for sanitizing?
This can be bypassed with `<scr<script>ipt>`.

Remove `<script>` (case insensitive) from the input until the input contains no `<script>`.

Note that flag format is `SECCON{[\x20-\x7e]+}`, which means that the flag may contains `<` or `>` as the following examples.

Sample Input 1:

```
S3CC0N{dum<scr<script>ipt>my}
```

Sample Output 1:

```
S3CC0N{dummy}
```

Sample Input 2 (small.txt):

```
S3CC0N{dumm<scrIpT>y_flag>_<_pt>>PT><<SCr<S<<SC<SCRIpT><scRiPT>Ript>sCr<Scri<...
```

Sample Output 2:

```
S3CC0N{dummy_flag>_<_pt>>PT><sCRIp<scr<scr<scr!pt>ipt>ipt>}
```

[flag](files/flag)

## Attachments
- [flag](files/flag)

## Flag
```
SECCON{sanitizing_is_not_so_good><_escaping_is_better_iPt><SCript<ScrIpT<scRIp<scRI<Sc<scr!pt>}
```

## Points
- Initial: `500`
- Last: `95` (solved: `115`)

## Tags
- `author:kusano`
- `ppc`
    