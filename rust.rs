fn main() {
    let limit = 100;

    fizzbuzz(limit);
}

fn fizzbuzz (limit: i32){
  for i in 0..limit{
    if i % 3 == 0 && i % 5 == 0{
      println!("FizzBuzz");
    }else if i % 3 == 0{
      println!("Fizz");
    }else if i % 3 == 0{
      println!("Buzz");
    }else{
      println!("{}", i);
    }
  }
}
