(ns euler.core
  (:require [clojure.math.numeric-tower :as math]
            [clojure.math.combinatorics :as comb]
            [clojure.data.json :as json]
            [clojure.java.io :as io]
            [euler.toolset :refer :all])
  (:gen-class))

(defn run-problem [i]
  (let [function (symbol (str "problem" i))]
    (use (vec (list (symbol (str "euler.problem" i)) :only (list function))))
    (eval (list function))))

(defn -main
  [& args]
  (let [answers (json/read-str (slurp "answers.json"))]
    (doseq [i (range 51 73)]
      (let [attempt (run-problem i)
            expected (answers (str i))]
        (if (= attempt expected)
          (println (format "problem%s.clj ✓" i))
          (println (format "** FAIL ** problem%s.clj: attempt=%s, expected=%s" i attempt expected)))))))
