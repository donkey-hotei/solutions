#!/usr/bin/ruby
def left_rotate!(a, k = 1)
  (0...k / 2).each do |i|
    a[i], a[k - 1 - i] = a[k - 1 - i], a[i]
  end
  (0...(a.size - k) / 2).each do|i|
    a[i + k], a[a.size - 1 - i] = a[a.size - 1 - i], a[i + k]
  end
  (0...a.size / 2).each do |i|
    a[i], a[a.size - 1 - i] = a[a.size - 1 - i], a[i]
  end
end

def main
    n,k = gets.strip.split(' ')
    n = n.to_i
    k = k.to_i
    a = gets.strip
    a = a.split(' ').map(&:to_i)
    left_rotate!(a, k)

    a.each { |x| print "#{x} " }
end

main
