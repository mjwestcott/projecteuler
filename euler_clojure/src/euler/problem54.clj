(ns euler.problem54
  "https://projecteuler.net/problem=54

  The file, poker.txt, contains one-thousand random hands dealt to two players.
  How many hands does Player 1 win?")

(defn suits
  "Return the suits of a given poker hand"
  ;; hands are given as vectors of strings of rank, suit
  ;; e.g. ["AC" "8D" "8H" "3C" "2S"] => [\C \D \H \C \S]
  [hand]
  (for [h hand] (get h 1)))

(defn ranks
  "Return the ranks of a given poker hand."
  ;; Note: the ranks are first sorted in descending order then sorted by count.
  ;; e.g. ["AC" "8D" "8H" "3S" "2S"] => [8 8 14 3 2]
  ;;      ["KC" "9D" "9H" "3C" "2C"] => [9 9 13 3 2]
  ;; This allows us to correctly judge that the second hand > the first.
  ;; Also, an ace is played 'low' i.e. with rank 1 if it makes a straight.
  [hand]
  (let [rs (for [h hand] (get h 0))
        trans {\A 14, \K 13, \Q 12, \J 11, \T 10}
        convert #(for [x %] (get trans x (Character/digit x 10)))
        reverse-sort #(reverse (sort %))
        modify-ace #(if (= % [14 5 4 3 2]) [5 4 3 2 1] %)
        sort-by-freq #(sort-by (frequencies %) > %)]
    (-> rs convert reverse-sort modify-ace sort-by-freq vec)))

(defn group
  "Return a sorted list of counts of the card ranks"
  ;; [8 8 14 3 2] => [2 1 1 1]
  ;; [9 9 9 5 5] => [3 2]
  [ranks]
  (-> ranks frequencies vals sort reverse))

(defn onepair?       [hand] (= (group (ranks hand)) [2 1 1 1]))
(defn twopair?       [hand] (= (group (ranks hand)) [2 2 1]))
(defn threeofakind?  [hand] (= (group (ranks hand)) [3 1 1]))
(defn fourofakind?   [hand] (= (group (ranks hand)) [4 1]))
(defn fullhouse?     [hand] (= (group (ranks hand)) [3 2]))
(defn flush?         [hand] (= 1 (count (set (suits hand)))))
(defn straight?      [hand] (and (= 4 (- (apply max (ranks hand))
                                         (apply min (ranks hand))))
                                 (= 5 (count (set (ranks hand))))))
(defn straightflush? [hand] (and (flush? hand) (straight? hand)))

(defn evaluate
  "Return a dict evaluation of the given hand. If the :value of
  two hands are identical, the :ranks field will break ties."
  ;; ["AD" "KC" "2S" "2D" "2C"] => {:value 3, :ranks [2 2 2 14 13]}
  [hand]
  {:value (cond
            (straightflush? hand) 8
            (fourofakind? hand)   7
            (fullhouse? hand)     6
            (flush? hand)         5
            (straight? hand)      4
            (threeofakind? hand)  3
            (twopair? hand)       2
            (onepair? hand)       1
            :else                 0)
   :ranks (ranks hand)})

(defn player1-wins?
  "Did Player 1 win the hand?"
  ;; Note: the problem specifies no ties are possible. Therefore, we
  ;; sort the hands based on their evaluation and take the first hand
  ;; to be the winner.
  [hands]
  (let [[p1 p2] (map evaluate hands)
        winner (first (reverse (sort-by (juxt :value :ranks) [p1 p2])))]
    (if (= winner p1)
      true
      false)))

(defn deal-two-hands [cards]
    [(subvec cards 0 5) (subvec cards 5 10)])

(defn problem54 []
  (with-open [rdr (clojure.java.io/reader "resources/poker.txt")]
    ;; How many hands did Player 1 win?
    (count (filter true? (for [row (line-seq rdr)]
                           (let [cards (clojure.string/split row #" ")]
                             (player1-wins? (deal-two-hands cards))))))))
