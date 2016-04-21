(ns euler.toolset
  (:require [clojure.math.numeric-tower :as math]
            [clojure.math.combinatorics :as comb]))

(defn- check-prime [n]
  (cond
    (< n 3) (= n 2)
    (= 0 (mod n 2)) false
    (some #(= 0 (mod n %)) (range 3 (inc (math/floor (math/sqrt n))) 2)) false
    :else true))

(def prime? (memoize check-prime))

(defn get-primes
  ([]
   (get-primes 2))
  ([start]
   (filter prime? (iterate inc start)))
  ([start end]
   (take-while #(< % end) (get-primes start))))

(defn digits [n]
  (->> n                      ; 1234
       (iterate #(quot % 10)) ; (1234, 123, 12, 1, 0, ...)
       (take-while pos?)      ; (1234, 123, 12, 1)
       (mapv #(mod % 10))     ; [4, 3, 2, 1]
       rseq))                 ; (1 2 3 4)
