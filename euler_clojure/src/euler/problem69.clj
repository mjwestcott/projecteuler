(ns euler.problem69
  "https://projecteuler.net/problem=69

  Euler's Totient function, φ(n) [sometimes called the phi function], is used
  to determine the number of numbers less than n which are relatively prime to
  n. For example, as 1, 2, 4, 5, 7, and 8, are all less than nine and
  relatively prime to nine, φ(9)=6.

  It can be seen that n=6 produces a maximum n/φ(n) for n ≤ 10. Find the value
  of n ≤ 1,000,000 for which n/φ(n) is a maximum."
  (:require [euler.toolset :refer [prime? get-primes]]))

;; Note that the phi function multiplies n by (1 - (1/p)) for every p in its
;; unique prime factors. Therefore, phi(n) will diminish as n has a greater
;; number of small unique prime factors. Since we are seeking the largest value
;; for n/phi(n), we want to minimize phi(n). We are therefore looking for the
;; largest number <= 1e6 which is the product of the smallest unique prime
;; factors, i.e successive prime numbers starting from 2.

(defn problem69 []
  (apply max (take-while #(< % 1e6) (reductions * (get-primes)))))
