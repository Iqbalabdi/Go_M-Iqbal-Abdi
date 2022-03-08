package main
import "fmt"

func fibo(n int) int {
  // your code here
  if (n == 1 || n == 0){
	  return n
  } else {
	  return fibo(n-1) + fibo(n-2)
  }
}

func main() {
   fmt.Println(fibo(0))  // 0
   fmt.Println(fibo(1))  // 1
   fmt.Println(fibo(2))  // 1
   fmt.Println(fibo(3))  // 2
   fmt.Println(fibo(5))  // 5
   fmt.Println(fibo(6))  // 8
   fmt.Println(fibo(7))  // 13
   fmt.Println(fibo(9))  // 13
   fmt.Println(fibo(10)) // 55
}
