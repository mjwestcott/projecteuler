(ns euler.problem72
  "https://projecteuler.net/problem=72

  Consider the fraction, n/d, where n and d are positive integers. If n<d and
  HCF(n,d)=1, it is called a reduced proper fraction.

  If we list the set of reduced proper fractions for d ≤ 8 in ascending order
  of size, we get:

  1/8, 1/7, 1/6, 1/5, 1/4, 2/7, 1/3, 3/8, 2/5, 3/7, 1/2, 4/7, 3/5, 5/8, 2/3, 5/7, 3/4, 4/5, 5/6, 6/7, 7/8

  It can be seen that there are 21 elements in this set. How many elements
  would be contained in the set of reduced proper fractions for d ≤
  1,000,000?")

;; Note: this solution is much slower than the equivalent python code which
;; mutates a dict in place. Suggestions very welcome.

(defn update-many [m vals f]
  (reduce #(update-in %1 [%2] f) m vals))

(defn phi-sieve
  "Return a vector of all values of phi(i) for i <= limit.
  By convention phi(0)=0 and phi(1)=1, but Project Euler
  appears to disagree, so those values are ignored."
  ;; Solved in the style of the sieve of Eratosthenes.
  [limit]
  (drop 2
   (loop [i 2
          result (vec (range (inc limit)))
          marked #{}]
     (cond
       (> i limit) result
       (contains? marked i) (recur (inc i) result marked)
       :else (recur (inc i)
                    (update-many result (range i (inc limit) i) #(* % (- 1 (/ 1 i))))
                    (into marked (range i (inc limit) i)))))))

(defn problem72 []
  ;; See https://en.wikipedia.org/wiki/Farey_sequence
  ;; for the connection to Euler's phi function.
  ;; In particular, note that |F n| = |F n-1| + phi(n).
  (reduce +' (phi-sieve 1000000)))
