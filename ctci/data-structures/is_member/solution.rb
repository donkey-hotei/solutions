#!/usr/bin/ruby
require "pry"

class AssertionError < RuntimeError; end

def assert(&block)
  raise AssertionError unless yield
end

#
# Write two functions setup and is_member, one to preprocess a
# set of words so that is_member can efficiently find if a word
# is in that original set. They should be able to handle query
# words with "."'s as wildcards (i.e: "ba." can match "bar",
# "baz", "bat", ...)
#

Tuple = Struct.new(:character, :index) do
  def to_ary
    to_a
  end
end

#
# Preprocess wordset into useful data structure.
#
def setup(wordset)
  # build up a character-index hashmap
  wordset.inject(Hash.new { |h, k| h[k] = [] }) do |hash, word|
    word.split("").each_with_index do |character, index|
      hash[Tuple.new(character, index)] << word
    end

    hash
  end
end

#
# Is word member of words?
#
def member?(word, words)
  # O(n) where n is number of letters in word
  word.split("")
      .each_with_index
      .map { |ch, i| words[Tuple.new(ch, i)] }  # look-up all words w/matchinc letter placements
      .reject { |list| list.empty? }            # reject empty lists
      .inject { |memo, obj| memo & obj }        # find the intersection of all lists found
      .any?                                     # if there are any left, then word is member
end



def main
  wordset = ["he", "foo", "bar", "baz", "ht"]
  hashmap = setup(wordset)
  puts member?("b..", hashmap).inspect

  assert { member?("fo.", hashmap) == true  }
  assert { member?("h..", hashmap) == false }
  assert { member?("...", hashmap) == true  }
end

main