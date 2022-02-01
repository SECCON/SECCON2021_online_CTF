import os
import binascii

def gen_filename():
    base = './prog/'
    if 'BASE' in os.environ:
        base = os.environ['BASE'] + '/'
    return base + binascii.hexlify(os.urandom(16)).decode('utf-8')


dirname =  gen_filename()
os.system(f'mkdir -p {dirname} && chmod 777 {dirname}')
cmd = f'docker run -i -v "$PWD/{dirname}:/prog" seccon_tree'
os.system(cmd)
