(ns euler.problem59
  "Each character on a computer is assigned a unique code and the preferred
  standard is ASCII (American Standard Code for Information Interchange). For
  example, uppercase A = 65, asterisk (*) = 42, and lowercase k = 107.

  A modern encryption method is to take a text file, convert the bytes to ASCII,
  then XOR each byte with a given value, taken from a secret key. The advantage
  with the XOR function is that using the same encryption key on the cipher text,
  restores the plain text; for example, 65 XOR 42 = 107, then 107 XOR 42 = 65.

  For unbreakable encryption, the key is the same length as the plain text
  message, and the key is made up of random bytes. The user would keep the
  encrypted message and the encryption key in different locations, and without
  both 'halves', it is impossible to decrypt the message.

  Unfortunately, this method is impractical for most users, so the modified
  method is to use a password as a key. If the password is shorter than the
  message, which is likely, the key is repeated cyclically throughout the
  message. The balance for this method is using a sufficiently long password key
  for security, but short enough to be memorable.

  Your task has been made easy, as the encryption key consists of three lower
  case characters. Using cipher.txt, a file containing the encrypted ASCII codes,
  and the knowledge that the plain text must contain common English words,
  decrypt the message and find the sum of the ASCII values in the original text."
  (:require [clojure.string :as str]))

(defn most-frequent [items]
  ;; [3 7 7 7 4 4] => 7
  (->> items
       frequencies
       (apply max-key val)
       key))

(defn xor-transform
  "Use given encryption key (vector of ints) to XOR the collection."
  [encrypt-key coll]
  (for [[c k] (map vector coll (cycle encrypt-key))]
    (bit-xor c k)))

(defn decrypt
  "Return the encryption key of the given cipher (vector of ints), assuming key
  length x. We assume that the most common element of the underlying text data
  is the space character."
  [cipher x]
  (for [i (range x)]
    (let [item (->> (subvec cipher i)
                    (take-nth x)
                    (most-frequent))]
      (bit-xor item (int \space)))))

(defn problem59 []
  (let [file (slurp "resources/cipher.txt")
        cipher (->> (str/split file #",")
                    (map str/trim)
                    (map #(Integer/parseInt %))
                    vec)
        encrypt-key (decrypt cipher 3)]
    (reduce + (xor-transform encrypt-key cipher))))
