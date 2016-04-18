(ns euler.problem62
  "https://projecteuler.net/problem=62

  The cube, 41063625 (345**3), can be permuted to produce two other cubes:
  56623104 (384**3) and 66430125 (405**3). In fact, 41063625 is the smallest
  cube which has exactly three permutations of its digits which are also cube.
  Find the smallest cube for which exactly five permutations of its digits are
  cube."
  (:require [clojure.string :as str]
            [clojure.math.numeric-tower :as math]))

(defn problem62 []
  ;; Add i**3 to a map, the keys of which are an arbitrarily specified
  ;; canonical permutation of its digits, the values of which are are lists of
  ;; all cube numbers seen so far which are permutable to the key.
  (loop [i 0
         cubes {}]
    (let [x (math/expt i 3)
          canonical (str/join (sort (str x)))
          seen (conj (get cubes canonical []) x)]
      (if (= 5 (count seen))
        (apply min seen)
        (recur (inc i)
               (assoc cubes canonical seen))))))
