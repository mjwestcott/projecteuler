(ns euler.problem60
  "The primes 3, 7, 109, and 673, are quite remarkable. By taking any two primes
  and concatenating them in any order the result will always be prime. For
  example, taking 7 and 109, both 7109 and 1097 are prime. The sum of these four
  primes, 792, represents the lowest sum for a set of four primes with this
  property.

  Find the lowest sum for a set of five primes for which any two primes
  concatenate to produce another prime."
  (:require [euler.toolset :refer [prime? get-primes]]
            [clojure.string :as str]))

(def concats-to-prime?
  "Tests whether concatenating ints x and y in either order makes a prime."
  (memoize
   (fn [x y]
     (let [join #(Integer/parseInt (str/join "" [(str %1) (str %2)]))]
       (and (prime? (join x y)) (prime? (join y x)))))))

(defn all-concat-to-prime?
  [candidates]
  (every? #(apply concats-to-prime? %) (for [x candidates
                                             y candidates
                                             :when (not= x y)]
                                         [x y])))

;; Use a depth-first search. Nodes are lists. They represent candidate
;; solutions.
(defn goal-state? [node]
  (= 5 (count node)))

(defn successors [node primes]
  (for [p primes
        :when (and (> p (apply max node))
                   (all-concat-to-prime? (conj node p)))]
    (conj node p)))

(defn depth-first-search [primes]
  (loop [frontier (apply list (for [p primes] [p]))]
    (let [node (peek frontier)]
      (if (goal-state? node)
        (reduce + node)
        (let [children (successors node primes)]
          (recur (into (pop frontier) (reverse children))))))))

;; It's not clear how many primes to search through.
;; Experimentation suggests around 9000.
(defn problem60 [] (depth-first-search (get-primes 1 9000)))
