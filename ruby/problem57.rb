# https://projecteuler.net/problem=57
#
# It is possible to show that the square root of two can be expressed as an
# infinite continued fraction.
#
#     √ 2 = 1 + 1/(2 + 1/(2 + 1/(2 + ... ))) = 1.414213...
#
# By expanding this for the first four iterations, we get:
#
#     1 + 1/2 = 3/2 = 1.5
#     1 + 1/(2 + 1/2) = 7/5 = 1.4
#     1 + 1/(2 + 1/(2 + 1/2)) = 17/12 = 1.41666...
#     1 + 1/(2 + 1/(2 + 1/(2 + 1/2))) = 41/29 = 1.41379...
#
# The next three expansions are 99/70, 239/169, and 577/408, but the eighth
# expansion, 1393/985, is the first example where the number of digits in the
# numerator exceeds the number of digits in the denominator. In the first
# one-thousand expansions, how many fractions contain a numerator with more
# digits than denominator?

require_relative "utils"

# How many digits in integer n?
def num_digits(n)
  n.to_s.length
end

# Does the numerator have more digits than the denominator?
def greater_numerator?(frac)
  num_digits(frac.numerator) > num_digits(frac.denominator)
end

def problem57
  Utils.iterate(2) { |x|
    2 + Rational(1, x)
  }.lazy.flat_map { |y|
    1 + Rational(1, y)
  }.take(1000).count { |z|
    greater_numerator?(z)
  }
end

puts problem57
