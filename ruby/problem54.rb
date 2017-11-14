# https://projecteuler.net/problem=54
#
# The file, poker.txt, contains one-thousand random hands dealt to two players.
# How many hands does Player 1 win?

class Hand
  attr_reader :cards

  def initialize(cards)
    @cards = cards # e.g. ['AC', '8D' ,'8H' ,'3C' ,'2S']
    @rank_to_i = {"A" => 14, "K" => 13, "Q" => 12, "J" => 11, "T" => 10}
  end

  def suits
    @suits ||= @cards.map { |c| c[1] }
  end

  # Note: the ranks are first sorted in descending order then sorted by count.
  # e.g. Hand.new(['AC', '8D', '8H', '3S', '2S']).ranks --> [8, 8, 14, 3, 2]
  #      Hand.new(['KC', '9D', '9H', '3C', '2C']).ranks --> [9, 9, 13, 3, 2]
  # This allows us to correctly judge that the latter beats the former.
  # Also, an ace is played 'low' i.e. with rank 1 if it makes a straight.
  def ranks
    @ranks ||= begin
      ints = @cards.map { |c| @rank_to_i[c[0]] || c.to_i }.sort.reverse
      return [5, 4, 3, 2, 1] if ints == [14, 5, 4, 3, 2]
      ints.sort_by { |i| -ints.count(i) }
    end
  end

  # A sorted list of counts of the card ranks
  # e.g. Hand.new(['QS', 'QC', '4D', '4H', 'QH']).grouped --> [3, 2]
  def grouped
    @grouped ||= ranks.group_by { |x| x }.map { |k, v| v.length }.sort.reverse
  end

  def onepair?() grouped == [2, 1, 1, 1] end
  def twopair?() grouped == [2, 2, 1] end
  def threeofakind?() grouped == [3, 1, 1] end
  def straight?() (ranks.max - ranks.min == 4) && ranks.uniq.length == 5 end
  def flush?() suits.uniq.length == 1 end
  def fourofakind?() grouped == [4, 1] end
  def fullhouse?() grouped == [3, 2] end
  def straightflush?() flush? && straight? end

  # Return a numerical value for the hand.
  def value
    @value ||= case
      when straightflush? then 8
      when fullhouse? then 7
      when fourofakind? then 6
      when flush? then 5
      when straight? then 4
      when threeofakind? then 3
      when twopair? then 2
      when onepair? then 1
      else 0
    end
  end
end

def compare(hand1, hand2)
  case
  when hand1.value > hand2.value
    1
  when hand1.value == hand2.value
    hand1.ranks <=> hand2.ranks
  else
    -1
  end
end

def hands(row)
  cards = row.split(" ")
  [Hand.new(cards[0..4]), Hand.new(cards[5..10])]
end

def problem54
  rows = File.readlines("data/poker.txt")
  rows.map { |row| compare(*hands(row)) }.count { |x| x == 1 }
end

puts problem54
