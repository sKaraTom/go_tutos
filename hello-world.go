package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("hello world")
	fmt.Println("1+1 =", 1+1)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	test := "initial"
	fmt.Println(test)

	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(math.Sin(n))

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	for {
		fmt.Println("loop")
		break
	}
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	}

	whatAmI := func(i interface{}) string {
		switch t := i.(type) {
		case bool:
			return "I am a bool"
		case int:
			return "I am an int"
		default:
			return fmt.Sprintf("Can't handle type %T", t)
		}
	}
	fmt.Println(whatAmI(1))
	fmt.Println(whatAmI(true))
	fmt.Println(whatAmI("hey"))

	var a [5]int
	fmt.Println("emp:", a)
	fmt.Println("len:", len(a))

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	s := make([]string, 3)
	fmt.Println("emp:", s)
	s = append(s, "d")
	fmt.Println("emp:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)
}
