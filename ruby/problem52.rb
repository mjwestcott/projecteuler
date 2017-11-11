# https://projecteuler.net/problem=52
#
# It can be seen that the number, 125874, and its double, 251748, contain exactly
# the same digits, but in a different order.
#
# Find the smallest positive integer, x, such that 2x, 3x, 4x, 5x, and 6x,
# contain the same digits.

def multiples(x)
  (2..6).map { |i| i*x }
end

def same_digits?(x, y)
  x.to_s.chars.sort == y.to_s.chars.sort
end

def all_same_digits?(x)
  multiples(x).all? { |y| same_digits?(x, y) }
end

def problem52
  1.step.find { |x| all_same_digits?(x) }
end

puts problem52
