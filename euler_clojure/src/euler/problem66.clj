(ns euler.problem66
  "https://projecteuler.net/problem=66

  Consider quadratic Diophantine equations of the form: x**2 – Dy**2 = 1. For
  example, when D=13, the minimal solution in x is 6492 – 13×1802 = 1. It can
  be assumed that there are no solutions in positive integers when D is square.
  By finding minimal solutions in x for D = {2, 3, 5, 6, 7}, we obtain the
  following:

    3**2 – 2×2**2 = 1
    2**2 – 3×1**2 = 1
    9**2 – 5×4**2 = 1
    5**2 – 6×2**2 = 1
    8**2 – 7×3**2 = 1

  Hence, by considering minimal solutions in x for D ≤ 7, the largest x is
  obtained when D=5.  Find the value of D ≤ 1000 in minimal solutions of x for
  which the largest value of x is obtained."
  (:require [clojure.math.numeric-tower :refer [expt floor sqrt]]))

;; Each iteration through the convergents of the continued fraction of sqrt(D),
;; we want to check whether the numerator and denominator provide a solution to
;; the Diophantine equation: https://en.wikipedia.org/wiki/Pell%27s_equation
;; See the section entitled 'Fundamental solution via continued fractions'

(defn cont-frac-values
  "Yield the values in the continued fraction representation of sqrt(D),
  e.g. sqrt(23) = [4;(1,3,1,8)], so yield (4, 1, 3, 1, 8, 1, 3, 1, 8, ...)"
  ;; See problem64.clj for a link explaining this algorithm. Here we use 'D'
  ;; in place of 'S' to be consistent with the wording of the question.
  [D]
  (map (comp int :a)
       (iterate
         (fn [{m :m d :d a :a}]
           (let [next-m (- (* d a) m)
                 next-d (/ (- D (expt next-m 2)) d)
                 next-a (floor (/ (+ (floor (sqrt D)) next-m) next-d))]
             {:m next-m
              :d next-d
              :a next-a}))
         {:m 0
          :d 1
          :a (floor (sqrt D))})))

(defn cont-frac
  "Return the continued fraction of the given values."
  ;; For example, given (4, 1, 3, 1, 8) representing the fifth
  ;; convergent of sqrt(23), return 4 + (1/(1 + (1/(3 + (1/8)))))
  [values]
  (+ (first values)
     (if (empty? (rest values)) 0 (/ 1 (cont-frac (rest values))))))

(defn convergent
  "Return the nth convergent of the continued fraction for sqrt(D),
  where D is a non-square positive integer."
  [D n]
  (clojure.lang.Numbers/toRatio
   (cont-frac (take n (cont-frac-values D)))))

(defn solve-pells-equation [D]
  (let [solution? (fn [conv]
                    ;; Check whether the convergent satisfies the
                    ;; Diophantine equation.
                    (let [x (numerator conv)
                          y (denominator conv)]
                      (= 1 (- (expt x 2)
                              (* D (expt y 2))))))
        ;; Find the solution with the minimal value of x
        ;; which satisfies the equation.
        answer (first
                (filter solution?
                 (map (partial convergent D) (iterate inc 1))))]
    ;; For this problem, we only need the value of x.
    (numerator answer)))

(defn square? [n]
  ;; Is n a perfect square?
  (= (sqrt n) (floor (sqrt n))))

(defn problem66 []
  (let [solutions (for [i (range 1 1001) :when ((complement square?) i)]
                    [i, (solve-pells-equation i)])
        ;; Find the solution with the largest value of x.
        answer (apply max-key second solutions)]
    ;; Return the value of D for which that value was obtained.
    (first answer)))
