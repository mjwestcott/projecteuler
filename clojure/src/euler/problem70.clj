(ns euler.problem70
  "https://projecteuler.net/problem=70

  Euler's Totient function, φ(n) [sometimes called the phi function], is used
  to determine the number of positive numbers less than or equal to n which are
  relatively prime to n. For example, as 1, 2, 4, 5, 7, and 8, are all less
  than nine and relatively prime to nine, φ(9)=6. The number 1 is considered to
  be relatively prime to every positive number, so φ(1)=1.

  Interestingly, φ(87109)=79180, and it can be seen that 87109 is a permutation
  of 79180.

  Find the value of n, 1 < n < 10**7, for which φ(n) is a permutation of n and
  the ratio n/φ(n) produces a minimum."
  (:require [euler.toolset :refer [get-primes phi]]))

(defn problem70 []
  ;; The search space is too large for brute-force. So, note that we are
  ;; seeking roughly the inverse of the previous problem -- to minimize
  ;; n/phi(n). Therefore, we want to maximize phi(n), which is acheived for
  ;; numbers with the fewest and largest unique prime factors. But the number
  ;; cannot simply be prime because in that case phi(n) == n-1 which is not a
  ;; permutation of n. Therefore, the best candidates should have two unique
  ;; prime factors.
  (let [permutation? (fn [x y] (= (sort (str x))
                                  (sort (str y))))
        ;; Since we are seeking large values for both prime factors, we can search
        ;; among numbers close to the value of sqrt(1e7) ~ 3162
        ps (take-while #(< % 4000) (get-primes 2000))
        ns (for [x ps, y ps :when (and (not= x y) (< (* x y) 1e7))] (* x y))
        candidates (filter #(permutation? % (phi %)) ns)]
    (apply min-key #(/ % (phi %)) candidates)))
