(ns euler.problem71
  "https://projecteuler.net/problem=71

  Consider the fraction, n/d, where n and d are positive integers. If n<d and
  HCF(n,d)=1, it is called a reduced proper fraction.

  If we list the set of reduced proper fractions for d ≤ 8 in ascending order
  of size, we get:

  1/8, 1/7, 1/6, 1/5, 1/4, 2/7, 1/3, 3/8, 2/5, 3/7, 1/2, 4/7, 3/5, 5/8, 2/3, 5/7, 3/4, 4/5, 5/6, 6/7, 7/8

  It can be seen that 2/5 is the fraction immediately to the left of 3/7.

  By listing the set of reduced proper fractions for d ≤ 1,000,000 in ascending
  order of size, find the numerator of the fraction immediately to the left of
  3/7.")

;; See https://en.wikipedia.org/wiki/Farey_sequence
;; The value of the new term in between neighbours 2/5 and 3/7 is found by
;; computing the mediant of those neighbours. We can take result to be the next
;; left-hand neighbour of 3/7 iteratively until the denominator reaches 1e6.

(defn mediant [a b]
  (/ (+ (numerator a) (numerator b))
     (+ (denominator a) (denominator b))))

(defn problem71 []
  (numerator
   (last
    (take-while #(< (denominator %) 1e6)
                (iterate (partial mediant (/ 3 7)) (/ 2 5))))))
