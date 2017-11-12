# https://projecteuler.net/problem=53
#
# How many, not necessarily distinct, values of nCr, for 1 ≤ n ≤ 100, are greater
# than one-million?

require_relative "utils"

def num_combinations(n, r)
  n.factorial / (r.factorial * (n-r).factorial)
end

def problem53
  (1..100).flat_map { |n|
    (1..n).flat_map { |r|
      num_combinations(n, r)
    }.select { |x| x > 1e6 }
  }.count
end

puts problem53
