"""
problem54.py

https://projecteuler.net/problem=54

The file, poker.txt, contains one-thousand random hands dealt to two players.
How many hands does Player 1 win?
"""
from collections import Counter

def suits(hand):
    """Return the suits of a given poker hand"""
    # hands are given as strings of rank, suit e.g. 'AC' or '8D'
    # suits(['AC', '8D', '8H', '3C', '2S']) --> ['C', 'D', 'H', 'C', 'S']
    return [h[1] for h in hand]

def ranks(hand):
    """Return the ranks of a given poker hand."""
    # Note: the ranks are first sorted in descending order then sorted by count.
    # e.g. hand1 = ranks(['AC', '8D', '8H', '3S', '2S']) --> [8, 8, 14, 3, 2]
    #      hand2 = ranks(['KC', '9D', '9H', '3C', '2C']) --> [9, 9, 13, 3, 2]
    # This allows us to correctly judge that hand2 > hand1.
    # Also, an ace is played 'low' i.e. with rank 1 if it makes a straight.
    trans = {'A': 14, 'K': 13, 'Q': 12, 'J': 11, 'T': 10}
    convert = lambda lst: [trans[x] if x in trans else int(x) for x in lst]
    revsorted = lambda lst: sorted(lst, reverse=True)
    modify_ace = lambda lst: lst if lst != [14, 5, 4, 3, 2] else [5, 4, 3, 2, 1]
    sort_by_count = lambda lst: sorted(lst, key=lambda x: lst.count(x), reverse=True)
    return sort_by_count(modify_ace(revsorted(convert([h[0] for h in hand]))))

def group(ranks):
    """Return a sorted list of counts of the card ranks"""
    # group([8, 8, 14, 3, 2]) --> [2, 1, 1, 1]
    # group([9, 9, 9, 5, 5]) --> [3, 2]
    return sorted(Counter(ranks).values(), reverse=True)

def onepair(hand): return group(ranks(hand)) == [2, 1, 1, 1]
def twopair(hand): return group(ranks(hand)) == [2, 2, 1]
def threeofakind(hand): return group(ranks(hand)) == [3, 1, 1]
def fourofakind(hand): return group(ranks(hand)) == [4, 1]
def fullhouse(hand): return group(ranks(hand)) == [3, 2]
def straightflush(hand): return (flush(hand) and straight(hand))
def flush(hand): return len(set(suits(hand))) == 1
def straight(hand): return ((max(ranks(hand)) - min(ranks(hand)) == 4)
                            and len(set(ranks(hand))) == 5)

def value(hand):
    """Return a tuple of (numerical value, ranks) for the given hand. The ranks
    are provided in order to break ties."""
    return ((8 if straightflush(hand) else
             7 if fourofakind(hand) else
             6 if fullhouse(hand) else
             5 if flush(hand) else
             4 if straight(hand) else
             3 if threeofakind(hand) else
             2 if twopair(hand) else
             1 if onepair(hand) else
             0), ranks(hand))

def compare(hand1, hand2):
    """Return 1 if hand1 wins, else return 0."""
    return (1 if max((hand1, hand2), key=value) == hand1 else 0)

def players(row):
    """Split a row of 'poker.txt' into the two players' hands."""
    cards = row.split()
    return (cards[:5], cards[5:])

def problem54():
    with open("data/poker.txt", "r") as f:
        rows = f.readlines()
    return sum(compare(*players(row)) for row in rows)
