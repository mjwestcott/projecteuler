(ns euler.problem55
  "If we take 47, reverse and add, 47 + 74 = 121, which is
  palindromic. A number that never forms a palindrome through the
  reverse and add process is called a Lychrel number. How many
  Lychrel numbers are there below ten-thousand? (Only consider fifty
  iterations)")

(ns euler.problem55)

(defn reverse-num [n]
  ;; 123456 => 654321N
  (bigint (clojure.string/join (reverse (str n)))))

(def palindrome?
  (memoize (fn [n] (= n (reverse-num n)))))

(defn lychrel? [n]
  (let [start (+ n (reverse-num n))
        iterations (iterate #(+' % (reverse-num %)) start)]
    (not-any? palindrome? (take 50 iterations))))

(defn problem55 []
  (count (filter lychrel? (range 1 10000))))
