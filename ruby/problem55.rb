# https://projecteuler.net/problem=55
#
# If we take 47, reverse and add, 47 + 74 = 121, which is palindromic. A number
# that never forms a palindrome through the reverse and add process is called a
# Lychrel number. How many Lychrel numbers are there below ten-thousand? (Only
# consider fifty iterations)

require_relative "utils"

# Return the reverse of n's digits.
def rev(n)
  n.to_s.chars.reverse.join('').to_i
end

def is_palindromic?(n)
  n == rev(n)
end

def is_lychrel?(n)
  start = n + rev(n)
  iterations = Utils.iterate(start) { |x| x + rev(x) }
  iterations.take(50).none? { |y| is_palindromic?(y) }
end

def problem55
  (1..9999).count { |x| is_lychrel?(x) }
end

puts problem55
