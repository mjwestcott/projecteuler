(ns euler.problem56
  "https://projecteuler.net/problem=56

  Considering natural numbers of the form, a**b, where a, b < 100,
  what is the maximum digital sum?"
  (:require [clojure.math.numeric-tower :as math]
            [euler.toolset :refer [digits]]))

(defn digit-sum [n]
  (reduce + (digits n)))

(defn problem56 []
  (apply max (for [a (range 100)
                   b (range 100)]
               (digit-sum (math/expt a b)))))
