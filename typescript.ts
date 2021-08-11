function fizzbuzz(start:number, limit:number){
    if(start <= limit) {
      if(start % 3 == 0 && start % 5 == 0) {
        console.log("FizzBuzz");
      }else if(start % 5 == 0) {
        console.log("Buzz");
      }else if(start % 3 == 0 && start % 5 == 0) {
        console.log("Fizz");
      }else{
        console.log(start);
      }
      fizzbuzz(start + 1, limit)
  }else {
      return
  };
};

fizzbuzz(1, 10);
