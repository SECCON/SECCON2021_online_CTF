#!/bin/sh


cd src/
python3.9 setup.py build
cp "build/lib.linux-x86_64-3.9/seccon_tree.cpython-39-x86_64-linux-gnu.so" ../env/
cd -
