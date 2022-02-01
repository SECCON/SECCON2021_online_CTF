#!/usr/bin/env python3.9
import os

def f(x):
    print(f'value 1: {repr(x)}')
    v = input('value 2: ')
    if len(v) > 8: return
    return eval(f'{x} * {v}', {}, {})

if __name__ == '__main__':
    print("+---------------------------------------------------+")
    print("| The Answer to the Ultimate Question of Life,      |")
    print("|                the Universe, and Everything is 42 |")
    print("+---------------------------------------------------+")

    for x in [6, 6.6, '666', [6666], {b'6':6666}]:
        if f(x) != 42:
            print("Something is fundamentally wrong with your universe.")
            exit(1)
        else:
            print("Correct!")

    print("Congrats! Here is your flag:")
    print(os.getenv("FLAG", "FAKECON{try it on remote}"))
