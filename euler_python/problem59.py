"""
problem59.py

Each character on a computer is assigned a unique code and the preferred
standard is ASCII (American Standard Code for Information Interchange). For
example, uppercase A = 65, asterisk (*) = 42, and lowercase k = 107.

A modern encryption method is to take a text file, convert the bytes to ASCII,
then XOR each byte with a given value, taken from a secret key. The advantage
with the XOR function is that using the same encryption key on the cipher text,
restores the plain text; for example, 65 XOR 42 = 107, then 107 XOR 42 = 65.

For unbreakable encryption, the key is the same length as the plain text
message, and the key is made up of random bytes. The user would keep the
encrypted message and the encryption key in different locations, and without
both "halves", it is impossible to decrypt the message.

Unfortunately, this method is impractical for most users, so the modified
method is to use a password as a key. If the password is shorter than the
message, which is likely, the key is repeated cyclically throughout the
message. The balance for this method is using a sufficiently long password key
for security, but short enough to be memorable.

Your task has been made easy, as the encryption key consists of three lower
case characters. Using cipher.txt, a file containing the encrypted ASCII codes,
and the knowledge that the plain text must contain common English words,
decrypt the message and find the sum of the ASCII values in the original text.
"""
from collections import Counter
from itertools import cycle

def decrypt(cipher, start=0, n=3):
    """Given a list of ints (cipher), finds the most common item of the
    sublist starting at given index (start), and taking every nth element.
    We assume this item is an int representing the ordinal value of the
    space character, so return item XOR ord(' ')"""
    item, _ = Counter(cipher[start::n]).most_common()[0]
    return item ^ ord(' ')

def problem59():
    with open("data/cipher.txt", "r") as f:
        cipher = [int(x) for x in f.read().split(',')]
    key = [decrypt(cipher, start=i) for i in range(3)]
    return sum(c ^ k for c, k in zip(cipher, cycle(key)))
