using System;

class MainClass {
  public static void Main (string[] args) {
    int limit = 100;
    for (int start = 0; start <= limit; start++)
    {
       if (start % 3 == 0 & start % 5 == 0)
            {
                Console.WriteLine("FizzBuzz");
            }
            else if (start % 3 == 0)
            {
                Console.WriteLine("Fizz");
            }
            else if (start % 5 == 0)
            {
                Console.WriteLine("Buzz");
            }
            else
            {
                Console.WriteLine(start);
            }
    }
  }
}
