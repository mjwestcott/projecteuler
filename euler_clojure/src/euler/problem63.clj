(ns euler.problem63
  "https://projecteuler.net/problem=63

  The 5-digit number, 16807=7**5, is also a fifth power. Similarly, the 9-digit
  number, 134217728=8**9, is a ninth power. How many n-digit positive integers
  exist which are also an nth power?"
  (:require [clojure.math.numeric-tower :as math]))

(defn find-powers
  "Return the list of powers to which one can raise n such that
  the result of exponentiation is an integer with number of digits == power"
  ;; (find-powers 6) => [1 2 3 4]
  [n]
  (take-while #(= % (count (str (math/expt n %)))) (iterate inc 1)))

(defn problem63 []
  ;; Retrieve results from find-powers until the empty list indicates no more
  ;; results. Find the length of all items of every resulting list.
  (let [results (take-while not-empty (map find-powers (iterate inc 1)))]
    (reduce + (map count results))))
