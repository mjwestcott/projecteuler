(ns euler.problem64
  "https://projecteuler.net/problem=64

  The first ten continued fraction representations of (irrational) square roots are:

      sqrt(2)=[1;(2)]          period=1
      sqrt(3)=[1;(1,2)]        period=2
      sqrt(5)=[2;(4)]          period=1
      sqrt(6)=[2;(2,4)]        period=2
      sqrt(7)=[2;(1,1,1,4)]    period=4
      sqrt(8)=[2;(1,4)]        period=2
      sqrt(10)=[3;(6)]         period=1
      sqrt(11)=[3;(3,6)]       period=2
      sqrt(12)=[3;(2,6)]       period=2
      sqrt(13)=[3;(1,1,1,1,6)] period=5

  Exactly four continued fractions, for N <= 13, have an odd period. How many
  continued fractions for N <= 10000 have an odd period?"
  (:require [clojure.math.numeric-tower :refer [expt floor sqrt]]))

(defn continued-fraction-sqrt [S]
  ;; https://en.wikipedia.org/wiki/Methods_of_computing_square_roots#Continued_fraction_expansion
  ;; Using variables S, m, d, a as in the URL above.
  (let [root-S (sqrt S)
        floor-root-S (floor root-S)]
    (if (= root-S floor-root-S)
      [root-S]
      (loop [seen []
             m 0
             d 1
             a floor-root-S]
        (let [next-seen (conj seen {:m m, :d d, :a a})
              next-m (- (* d a) m)
              next-d (/ (- S (expt next-m 2)) d)
              next-a (floor (/ (+ floor-root-S next-m) next-d))]
          ;; If the pattern repeats, return all the 'a' variables collected in order.
          ;; Otherwise, continue the recursion.
          (if (some #(= % {:m next-m, :d next-d, :a next-a}) next-seen)
            (map :a next-seen)
            (recur next-seen
                   next-m
                   next-d
                   next-a)))))))

(defn problem64 []
  (let [cont-fracs (map continued-fraction-sqrt (range 2 10001))]
    (count (filter #(odd? (count (rest %))) cont-fracs))))
