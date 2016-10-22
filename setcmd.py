#!/usr/bin/env python

####################################################
# setcmd.py
# Little script to set the command for lonely-shell.
# https://github.com/vesche/lonely-shell
####################################################

import readline
import sys


def main():
    while True:
        command = raw_input("lonely-shell> ")
        if command == '':
            continue
        elif command.lower() == "quit":
            sys.exit(0)
        elif command.lower() == "status":
            with open("static/command.html") as f:
                print f.read()
        else:
            with open("static/command.html", 'w') as f:
                f.write(command)


if __name__ == "__main__":
    main()
