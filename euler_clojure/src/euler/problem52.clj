(ns euler.problem52
  "https://projecteuler.net/problem=52

  It can be seen that the number, 125874, and its double, 251748, contain exactly
  the same digits, but in a different order.

  Find the smallest positive integer, x, such that 2x, 3x, 4x, 5x, and 6x,
  contain the same digits.")

(defn multiples [x]
  (for [i (range 2 7)] (* i x)))

(defn same-digits? [x y]
  (= (sort (str x)) (sort (str y))))

(defn all-same-digits? [x]
  (every? true? (for [y (multiples x)] (same-digits? x y))))

(defn problem52 []
  (first (for [x (iterate inc 1) :when (all-same-digits? x)] x)))
