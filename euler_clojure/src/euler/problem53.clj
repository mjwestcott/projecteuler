(ns euler.problem53
  "How many, not necessarily distinct, values of nCr, for 1 ≤ n ≤ 100, are greater
  than one-million?")

(defn factorial [n]
  (reduce *' (range 1 (inc n))))

(defn num-combinations [n r]
  (/ (factorial n) (*' (factorial r) (factorial (- n r)))))

(defn problem53 []
  (count (filter true? (for [n (range 1 (inc 100))
                             r (range 1 (inc n))]
                         (> (num-combinations n r) 1e6)))))
