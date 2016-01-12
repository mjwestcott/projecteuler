;; problem55.clj
;;
;; If we take 47, reverse and add, 47 + 74 = 121, which is
;; palindromic. A number that never forms a palindrome through the
;; reverse and add process is called a Lychrel number. How many
;; Lychrel numbers are there below ten-thousand? (Only consider fifty
;; iterations)

(ns euler.problem55)

(defn to-digits [n]
  ;; 1234 => (1 2 3 4)
  (map #(Character/getNumericValue %) (str n)))

(defn to-num [digits]
  ;; [1 2 3 4] => 1234N
  (bigint (clojure.string/join digits)))

(def palindromic?
  (memoize (fn [n] (= (to-digits n)
                      (reverse (to-digits n))))))

(defn lychrel? [n]
  (let [reverse-num #(to-num (reverse (to-digits %)))
        start (+ n (reverse-num n))
        iterations (iterate #(+' % (reverse-num %)) start)]
    (not-any? palindromic? (take 50 iterations))))

(defn problem55 []
  (count (filter lychrel? (range 1 10000))))
