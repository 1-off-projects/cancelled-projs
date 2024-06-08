package main

import (
	"fmt"
)

func main() {
	var n_terms int
	n_terms = 0
	n1, n2 := 0, 1
	count := 1
	c1 := (n_terms < 1)
	c2 := (n_terms == 1)
	c3 := (n_terms >= 4300)
	fmt.Println("# of terms: ")
	fmt.Scanf("%b\n", &n_terms)
	switch n_terms {
	case c1 == true:
		fmt.Println("enter a valid number")
	case c2 == true:
		fmt.Println("sequence is: 0")
	case c3 == true:
		fmt.Println("too large")
	default:
		fmt.Println("Sequence is:")
		for count < n_terms {
			fmt.Print(n1)
			math := n1 + n2
			n1 = n2
			n2 = math
			count += 1
		}
	}
	/*if n_terms < 1 {
		fmt.Println("enter a valid number")
	} else if n_terms == 1 {
		fmt.Println("sequence is: 0")
	} else if n_terms >= 4300 {
		fmt.Println("too large")
	} else {
		fmt.Println("Sequence is:")
		for count < n_terms {
			fmt.Print(n1)
			math := n1 + n2
			n1 = n2
			n2 = math
			count += 1
		}
	}*/
}
