class Integer
  def factorial
    self.downto(1).reduce(1, :*)
  end
end

class Utils
  # https://ruby-doc.org/core-2.4.2/Enumerable.html#method-i-lazy-label-Example
  def self.pythagorean_triples
    (1..Float::INFINITY).lazy.flat_map {|z|
      (1..z).flat_map {|x|
        (x..z).select {|y|
          x**2 + y**2 == z**2
        }.map {|y|
          [x, y, z]
        }
      }
    }
  end
end
