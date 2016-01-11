;; problem53.clj
;;
;; How many, not necessarily distinct, values of nCr, for 1 ≤ n ≤ 100, are greater
;; than one-million?

(ns euler.problem53)

(defn factorial [n]
  (reduce *' (range 1 (inc n))))

(defn num-combinations [n r]
  (/ (factorial n) (*' (factorial r) (factorial (- n r)))))

(defn problem53 []
  (reduce + (for [n (range 1 (inc 100))
                  r (range 1 (inc n))
                  :when (> (num-combinations n r) 1e6)]
              1)))
