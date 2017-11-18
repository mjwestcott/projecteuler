# https://projecteuler.net/problem=56
#
# Considering natural numbers of the form, a**b, where a, b < 100,
# what is the maximum digital sum?

def digit_sum(n)
  n.to_s.chars.map { |c| c.to_i }.sum
end

def problem56
  (1..99).flat_map { |x|
    (1..99).flat_map { |y|
      digit_sum(x**y)
    }
  }.max
end

puts problem56
