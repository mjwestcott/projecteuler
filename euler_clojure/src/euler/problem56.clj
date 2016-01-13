(ns euler.problem56
  "Considering natural numbers of the form, a**b, where a, b < 100,
  what is the maximum digital sum?"
  (:require [clojure.math.numeric-tower :as math]))

(defn digits [n]
  (->> n                      ; 1234
       (iterate #(quot % 10)) ; (1234, 123, 12, 1, 0, ...)
       (take-while pos?)      ; (1234, 123, 12, 1)
       (mapv #(mod % 10))     ; [4, 3, 2, 1]
       rseq))                 ; (1 2 3 4)

(defn digit-sum [n]
  (reduce + (digits n)))

(defn problem56 []
  (apply max (for [a (range 100)
                   b (range 100)]
               (digit-sum (math/expt a b)))))
