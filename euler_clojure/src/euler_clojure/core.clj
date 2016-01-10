(ns euler.core
  (:require [clojure.math.numeric-tower :as math]
            [clojure.math.combinatorics :as comb]
            [euler.toolset :refer :all]
            [taoensso.timbre.profiling :as profiling :refer (pspy pspy* profile defnp p p*)])
  (:gen-class))

(defn -main
  [& args]
  (println "Hello, World!"))
