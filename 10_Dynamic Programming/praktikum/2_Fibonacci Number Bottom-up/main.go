package main

import "fmt"

func fibo(n int) int {
  // your code here
  if (n == 1 || n == 0){
	  return n
  }
  n1 := 0
  n2 := 1
  var n3 int
  for i := 1; i<n; i++ {
	   n3 = n1+n2
	   n1=n2
	   n2=n3
  }

  return n3
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