;; problem51.clj
;;
;; By replacing the 1st digit of the 2-digit number *3, it turns out that
;; six of the nine possible values: 13, 23, 43, 53, 73, and 83, are all
;; prime.
;;
;; By replacing the 3rd and 4th digits of 56**3 with the same digit, this
;; 5-digit number is the first example having seven primes among the ten
;; generated numbers, yielding the family: 56003, 56113, 56333, 56443,
;; 56663, 56773, and 56993. Consequently 56003, being the first member of
;; this family, is the smallest prime with this property.
;;
;; Find the smallest prime which, by replacing part of the number (not
;; necessarily adjacent digits) with the same digit, is part of an eight
;; prime value family.

(ns euler.problem51
  (:require [euler.toolset :refer [prime? get-primes]]
            [clojure.math.combinatorics :as comb]
            [clojure.math.numeric-tower :as math]))

(defn to-digits [n]
  ;; 1234 => (1 2 3 4)
  (map #(Character/getNumericValue %) (str n)))

(defn to-num [digits]
  ;; [1 2 3 4] => 1234
  (Integer/parseInt (clojure.string/join digits)))

;; Our strategy is as follows. Since we are seeking an eight prime family, it
;; must be the case that the pattern of digits which are replaced contains
;; either 0, 1, or 2 in the smallest family member. Therefore, we can search
;; through primes and replace digits in patterns specified by the locations 0,
;; 1, and 2. If the family of numbers that results contains eight primes, we
;; have found the solution.
;;
;; In the example given, 56003 is the smallest member of an eight prime family.
;; We would find the pattern of 0s at indices (2, 3) to produce the
;; corresponding family from 56**3.
(defn find-indices
  "Return three sets, where each contains the indices in the num of
  the digits 0, 1 and 2 respectively."
  ;; (find-indices 18209912) => #{3}, #{0 6}, #{2 7}
  ;; (find-indices 56003) => #{2 3}, #{}, #{}
  [num]
  (let [digits (to-digits num)]
    (for [dgt [0 1 2]]
      (set (for [[i x] (map-indexed vector digits) :when (= x dgt)] i)))))

(defn replace-digits
  "Replace elements of digits with value x at given indices.
  Return sentinel value [-1] if the result would lead with a zero."
  ;; (replace-digits [5 6 0 0 3] #{2 3} 9) => [5 6 9 9 3]
  ;; (replace-digits [5 6 0 0 3] #{0} 0) => [-1]
  [digits indices x]
  (if (and (= x 0) (contains? indices 0))
    [-1]
    (map #(if (contains? indices %1) x %2) (iterate inc 0) digits)))

(defn family
  "Return the family of numbers resulting from replacing digits at the specified
  indices with the digits 0 to 9"
  ;; (family 56003 #{2 3}) => 56003, 56113, 56223, 56333, 56443, ...
  [num indices]
  (let [digits (to-digits num)]
    (for [x (range 10)]
      (to-num (replace-digits digits indices x)))))

(defn smallest-member?
  "Does the number statisfy the problem specification?"
  [num]
  (some true? (for [indices (find-indices num)]
                (= 8 (count (filter prime? (family num indices)))))))

(defn problem51 []
  (first (filter smallest-member? (get-primes 56995))))
