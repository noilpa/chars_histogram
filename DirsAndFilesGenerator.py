#!/usr/bin/env python

import os
import random
from string import printable

import sys

FILE_COUNTER = 0
DIR_COUNTER = 0


def next_file_name():
    global FILE_COUNTER
    FILE_COUNTER += 1
    return 'file_{}'.format(FILE_COUNTER)


def next_dir_name():
    global DIR_COUNTER
    DIR_COUNTER += 1
    return 'dir_{}'.format(DIR_COUNTER)


def create_dir(path):
    try:
        os.mkdir(path)
    except OSError:
        print("Creation of the directory %s failed" % path)
    else:
        print("Successfully created the directory %s " % path)


def create_files(path):
    for _ in range(random.randrange(2, 6)):
        file_path = path + '/' + next_file_name()
        with open(file_path, 'w') as f:
            f.write(generate_text())
            print("Successfully created the file %s " % file_path)


def generate_text():
    return ''.join(random.choice(printable) for _ in range(1000))


def main():
    # python DirsAndFilesGenerator.py root_dir depth
    if len(sys.argv) < 2:
        print('Error not enough args')
        return
    root = os.getcwd()
    if len(sys.argv) > 2:
        root = sys.argv[2]

    for _ in range(int(sys.argv[1])):
        root = root + '/' + next_dir_name()
        create_dir(path=root)
        create_files(root)
    print('Total generated files count:', FILE_COUNTER)


if __name__ == '__main__':
    main()
