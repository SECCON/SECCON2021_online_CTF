#!/usr/bin/env python3

from flag import flag
import signal
import bcrypt

def check_and_upper(message):
    if len(message) > 24:
        return None
    
    message = message.upper()

    for c in message:
        c = ord(c)
        if ord("A") > c or c > ord("Z"):
            return None
    
    return message

signal.alarm(600)
while True:
    mode = input(
        """1. sign
2. verify
mode: """
    ).strip()

    ## sign mode ##
    if mode == "1":
        message = check_and_upper(input("message: ")) # case insensitive

        if message == None:
            print("invalid")
            continue

        salt = bcrypt.gensalt(5)
        print("mac:", bcrypt.hashpw((message + flag).encode(), salt).decode("utf-8"))

    ## verify mode ##
    else:
        mac = input("mac: ")
        message = check_and_upper(input("message: ")) # case insensitive

        if message is None:
            print("invalid")
            continue

        print("result:", bcrypt.checkpw((message + flag).encode(), mac.encode()))
