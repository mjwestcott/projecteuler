(ns euler.problem58
  "Starting with 1 and spiralling anticlockwise in the following way,
  a square spiral with side length 7 is formed.

    37 36 35 34 33 32 31
    38 17 16 15 14 13 30
    39 18  5  4  3 12 29
    40 19  6  1  2 11 28
    41 20  7  8  9 10 27
    42 21 22 23 24 25 26
    43 44 45 46 47 48 49

  It is interesting to note that the odd squares lie along the bottom
  right diagonal, but what is more interesting is that 8 out of the 13
  numbers lying along both diagonals are prime; that is, a ratio of
  8/13 ≈ 62%.

  If one complete new layer is wrapped around the spiral above, a
  square spiral with side length 9 will be formed. If this process is
  continued, what is the side length of the square spiral for which
  the ratio of primes along both diagonals first falls below 10%?"
  (:require [euler.toolset :refer [prime?]]
            [clojure.math.numeric-tower :as math]))

(defn squared [n] (math/expt n 2))

;; Given the bottom right corner number, return the square length.
(defn square-length [n] (math/sqrt n))

(defn corners [n]
  ;; Given the bottom right corner number, return the four corner numbers.
  ;; 49 => (49 43 37 31)
  (let [x (dec (square-length n))]
    (take 4 (iterate #(- % x) n))))

(defn problem58 []
  (loop [x 7         ; The current square-length.
         primes 8    ; The number of primes found thus far.
         total 13]   ; The total number of corner numbers checked.
    (if (< (/ primes total) 0.10)
      x
      (recur (+ x 2)
             (+ primes (count (filter prime? (corners (squared (+ x 2))))))
             (+ total 4)))))
