(ns euler.problem52
  "https://projecteuler.net/problem=52

  It can be seen that the number, 125874, and its double, 251748, contain exactly
  the same digits, but in a different order.

  Find the smallest positive integer, x, such that 2x, 3x, 4x, 5x, and 6x,
  contain the same digits.")

(defn multiples [x]
  (map #(* % x) (range 2 7)))

(defn same-digits? [x y]
  (= (sort (str x))
     (sort (str y))))

(defn all-same-digits? [x]
  (every? true? (map #(same-digits? % x) (multiples x))))

(defn problem52 []
  (first (filter all-same-digits? (iterate inc 1))))
