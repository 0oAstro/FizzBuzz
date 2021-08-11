num = 100

for i in 0..num
  if i % 3 == 0 and i % 5 == 0
    puts "FizzBuzz"
  elsif i % 3 == 0
    puts "Fizz"
  elsif i % 5 == 0
    puts "Buzz"
  else
    puts "#{i}"
  end
end
