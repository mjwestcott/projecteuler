(ns euler.problem51
  "By replacing the 1st digit of the 2-digit number *3, it turns out that
  six of the nine possible values: 13, 23, 43, 53, 73, and 83, are all
  prime.

  By replacing the 3rd and 4th digits of 56**3 with the same digit, this
  5-digit number is the first example having seven primes among the ten
  generated numbers, yielding the family: 56003, 56113, 56333, 56443,
  56663, 56773, and 56993. Consequently 56003, being the first member of
  this family, is the smallest prime with this property.

  Find the smallest prime which, by replacing part of the number (not
  necessarily adjacent digits) with the same digit, is part of an eight
  prime value family."
  (:require [euler.toolset :refer [prime? get-primes]]))

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
  (let [digits (str num)]
    (for [target "012"]
      (set (for [[i x] (map-indexed vector digits) :when (= x target)] i)))))

(defn replace-digits
  "Replace elements of template with value x at given indices."
  ;; (replace-digits "56003" #{2 3} "9") => "56993"
  [template indices x]
  (clojure.string/join (map #(if (contains? indices %1) x %2)
                            (iterate inc 0)
                            template)))

(defn family
  "Return the family of numbers resulting from replacing digits at the
  specified indices with the digits 0 to 9. If the number would lead
  with a zero, it is replaced with a sentinel value (-1)."
  ;; (family 56003 #{2 3}) => 56003, 56113, 56223, 56333, 56443, ...
  [num indices]
  (let [template (str num)]
    (for [x "0123456789"]
      (if (and (= x \0) (contains? indices 0))
        -1
        (bigint (replace-digits template indices x))))))

(defn smallest-member?
  "Does the number satisfy the problem specification?"
  [num]
  (some true? (for [indices (find-indices num)]
                (= 8 (count (filter prime? (family num indices)))))))

(defn problem51 []
  (first (filter smallest-member? (get-primes 56995))))
