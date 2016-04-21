(ns euler.problem68
  "https://projecteuler.net/problem=68

  What is the maximum 16-digit string for a 'magic' 5-gon ring?"
  (:require [euler.toolset :refer [digits]]
            [clojure.string :as str]))

(defn five-gon-rings
  "Return a list of solutions to the 'magic' 5-gon ring problem, each line
  summing to n. The empty list will be returned if there are no solutions"
  [n]
  (for [a (range 1 11)
        b (range 1 11) :when (not= b a)
        c (range 1 11) :when (not (contains? #{a b} c))
        :when (= n (+ a b c))
        d (range 1 11) :when (not (contains? #{a b c} d))
        e (range 1 11) :when (not (contains? #{a b c d} e))
        :when (= n (+ d c e))
        f (range 1 11) :when (not (contains? #{a b c d e} f))
        g (range 1 11) :when (not (contains? #{a b c d e f} g))
        :when (= n (+ f e g))
        h (range 1 11) :when (not (contains? #{a b c d e f g} h))
        i (range 1 11) :when (not (contains? #{a b c d e f g h} i))
        :when (= n (+ h g i))
        j (range 1 11) :when (not (contains? #{a b c d e f g h i} j))
        :when (= n (+ j i b))
        :when (< a (min d f h j))]
    ;; Each solution can be described uniquely starting from the group of three
    ;; with the numerically lowest external node and working clockwise. So we
    ;; specified at the end that (< a (min d f h j)
    [[a b c] [d c e] [f e g] [h g i] [j i b]]))

(defn problem68 []
  (let [start 6  ; Each line cannot sum to less than 6 (1+2+3)
        end 28]  ; Or greater than 27 (8+9+10)
    (->> (mapcat five-gon-rings (range start end))
         (filter (complement empty?))
         (map (comp str/join #(map str %) flatten))
         (filter #(= (count %) 16))
         ((comp bigint last sort)))))
