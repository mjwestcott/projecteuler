# https://projecteuler.net/problem=51
#
# By replacing the 1st digit of the 2-digit number *3, it turns out that six of
# the nine possible values: 13, 23, 43, 53, 73, and 83, are all prime.
#
# By replacing the 3rd and 4th digits of 56**3 with the same digit, this 5-digit
# number is the first example having seven primes among the ten generated
# numbers, yielding the family: 56003, 56113, 56333, 56443, 56663, 56773, and
# 56993. Consequently 56003, being the first member of this family, is the
# smallest prime with this property.
#
# Find the smallest prime which, by replacing part of the number (not necessarily
# adjacent digits) with the same digit, is part of an eight prime value family.

require "prime"

# Return three arrays, where each array contains the indices
# in the num of the digits 0, 1, and 2 respectively.
# indices(18209912) --> [3], [0, 6], [2, 7]
# indices(56003) --> [2, 3], [], []
def indices(num)
  digits = num.to_s.chars
  "012".chars.map { |c|
    digits.each_index.select { |i| digits[i] == c }
  }
end

# Return the family of numbers resulting from replacing
# digits at the specified indicies with the digits 0 to 9.
# family(56003, [2, 3]) --> 56003, 56113, 56223, 56333, 56443, ...
def family(num, indices)
  template = num.to_s
  "0123456789".chars.map { |c|
    member = template.chars
    indices.each { |i| member[i] = c }
    # Return sentinel value (-1) in case of leading zero.
    member[0] == "0" ? -1 : member.join("").to_i
  }
end

def problem51
  Prime.find { |p|
    indices(p).any? { |indices|
      family(p, indices).count { |x| Prime.prime?(x) } == 8
    }
  }
end

puts problem51
