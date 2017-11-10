(ns euler.problem67
  "https://projecteuler.net/problem=67

  Find the maximum total from top to bottom in triangle.txt a 15K text
  file containing a triangle with one-hundred rows."
  (:require [clojure.string :as str]))

(defn children [triangle]
  ;; Assuming a list of lists such as ((1) (2 3) (4 5 6))
  ;; representing the triangle   1
  ;;                            2 3
  ;;                           4 5 6
  ;; return the two child triangles, ((2) (4 5)) and ((3) (5 6)).
  (list (map drop-last (rest triangle))
        (map rest (rest triangle))))

(def max-route
  "Recursively find the maximum value of the root node plus the
  largest of its children, and so on, all the way to the base."
  (memoize
   (fn [triangle]
     (let [root (first (first triangle))]
       (if (= 1 (count triangle))
         root
         (+ root (apply max (map max-route (children triangle)))))))))

(defn problem67 []
  (let [triangle (for [row (str/split-lines (slurp "resources/triangle.txt"))]
                   (map #(Integer/parseInt %) (str/split row #" ")))]
    (max-route triangle)))
