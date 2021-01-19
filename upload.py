#!/usr/bin/env python3
import requests
import sys


def upload(filenames):
    files = []
    i = 0
    for filename in filenames:

        files.append(
            ("file[]" , (filename, open(filename, "rb"))),
        )

    res = requests.post(f"{sys.argv[1]}",files=files)
    print(res.text)

upload(sys.argv[2:])
