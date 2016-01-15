// problem58.go
//
// Each character on a computer is assigned a unique code and the preferred
// standard is ASCII (American Standard Code for Information Interchange). For
// example, uppercase A = 65, asterisk (*) = 42, and lowercase k = 107.
//
// A modern encryption method is to take a text file, convert the bytes to ASCII,
// then XOR each byte with a given value, taken from a secret key. The advantage
// with the XOR function is that using the same encryption key on the cipher text,
// restores the plain text; for example, 65 XOR 42 = 107, then 107 XOR 42 = 65.
//
// For unbreakable encryption, the key is the same length as the plain text
// message, and the key is made up of random bytes. The user would keep the
// encrypted message and the encryption key in different locations, and without
// both "halves", it is impossible to decrypt the message.
//
// Unfortunately, this method is impractical for most users, so the modified
// method is to use a password as a key. If the password is shorter than the
// message, which is likely, the key is repeated cyclically throughout the
// message. The balance for this method is using a sufficiently long password key
// for security, but short enough to be memorable.
//
// Your task has been made easy, as the encryption key consists of three lower
// case characters. Using cipher.txt, a file containing the encrypted ASCII codes,
// and the knowledge that the plain text must contain common English words,
// decrypt the message and find the sum of the ASCII values in the original text.

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// Return the int appearing most frequently in a slice of ints.
// In case of tie, return the int occurring first in the slice.
func mostFrequent(xs []int) int {
	if len(xs) == 0 {
		panic("Empty slice passed to mostFrequent")
	}
	freqs := make(map[int]int)
	unique := []int{}
	for _, x := range xs {
		if _, seen := freqs[x]; !seen {
			unique = append(unique, x)
		}
		freqs[x]++
	}
	result, max := 0, 0
	for _, x := range unique { // in order of occurrence in xs
		if count := freqs[x]; count > max {
			result, max = x, count
		}
	}
	return result
}

// Return the encryption key. We assume that the most common element of the
// underlying text data is the space character.
func decrypt(cipher []int, keylen int) []int {
	var key []int
	for i := 0; i < keylen; i++ {
		var items []int
		for j := i; j < len(cipher); j += keylen {
			items = append(items, cipher[j])
		}
		k := mostFrequent(items) ^ int(' ')
		key = append(key, k)
	}
	return key
}

func problem59() int {
	content, err := ioutil.ReadFile("data/cipher.txt")
	if err != nil {
		log.Fatal(err)
	}

	trimmed := strings.TrimSpace(string(content))
	ss := strings.Split(trimmed, ",")
	var cipher []int
	for _, s := range ss {
		x, _ := strconv.Atoi(s)
		cipher = append(cipher, x)
	}

	keylen := 3
	key := decrypt(cipher, keylen)
	var sum int // The sum of the ASCII values of the XOR tranformed data.
	for i, c := range cipher {
		sum += c ^ key[i%keylen]
	}
	return sum
}

func main() {
	ans := problem59()
	fmt.Println(ans)
}
