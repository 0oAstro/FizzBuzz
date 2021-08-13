#include <stdio.h>
int fizz_buzz(int limit){
  for (int start = 0; start <= limit; start++){
    if (start % 3 == 0 && start % 5 == 0){
      printf("FizzBuzz\n");
    } else if (start % 3 == 0){
      printf("Fizz\n");
    } else if (start % 5 == 0){
      printf("Buzz\n");
    } else {
      printf("%d \n", start);
    }
  }
  return 0;
}
int main(void) {
  fizz_buzz(100);
  return 0;
}
