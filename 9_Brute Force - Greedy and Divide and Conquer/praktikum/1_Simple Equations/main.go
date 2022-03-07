package main

import "fmt"

func check(x,y,z,A,B,C int) bool {
	return 2* (x*y + x*z + y*z) == A*A - C && x*y*z == B
}

func SimpleEquations(A,B,C int) {
	x,y,z := 1,1,1

	for ;x*x+y*y+z*z<1000; x++{
		if x+y+z > A {
			break
		}
		for ;  x*x+y*y+z*z<1000; y++ {
			if x+y+z > A {
				break
			}
			for ;  x*x+y*y+z*z<1000; z++ {
				if x+y+z > A {
					break
				}
				if check(x,y,z,A,B,C) {
					fmt.Println(x,y,z)
					break
				}
			}
			if check(x,y,z,A,B,C) {
				break
			}
			z = 1
		}
		if check(x,y,z,A,B,C) {
			break
		}
		y = 1
	}

	if x*x+y*y+z*z > C || x+y+z > A || x*y*z > B {
		fmt.Println("No Solution")
	}

}


func main() {
   SimpleEquations(1, 2, 3)  // no solution
   SimpleEquations(6, 6, 14) // 1 2 3
}