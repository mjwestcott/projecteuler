#!/usr/bin/env python

import glob
import json
import re
from subprocess import check_output

if __name__ == "__main__":
    with open("answers.json") as f:
        answers = json.load(f)

    for filename in sorted(glob.glob("problem*.py")):
        i = re.search("\d+", filename).group(0)

        attempt = int(check_output(["python", filename]))
        expected = answers[i]

        if attempt == expected:
            print("{} ✓".format(filename), flush=True)
        else:
            print("** FAIL ** {}: attempt={}, exptected={}".format(filename, attempt, expected))
