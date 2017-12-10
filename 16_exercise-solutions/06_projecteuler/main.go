package main

import "fmt"

func main() {
	// fmt.Println(maxprimefactor(600851475143)) // 6857

	maxprod, fac1, fac2 := largestpalindromeproduct(100, 1000) // 913 × 993 = 906609
	fmt.Printf("%d × %d = %d\n", fac1, fac2, maxprod)
}

/*
  Project Euler - problem 3 - Largest prime factor

  The prime factors of 13195 are 5, 7, 13 and 29.
  What is the largest prime factor of the number 600851475143 ?
*/

// This works but slow, needs to be optimized

func maxprimefactor(n int) int {
	var mprimfac int
	if n == 1 {
		return 1
	}
	for i := 1; i < n/2; i++ {
		if (n%i == 0) && (maxprimefactor(i) == 1) && (mprimfac < i) {
			mprimfac = i
		}
	}
	return mprimfac
}

/*
  Project Euler - problem 4 - Largest palindrome product

  A palindromic number reads the same both ways. The largest palindrome
  made from the product of two 2-digit numbers is 9009 = 91 × 99.
  Find the largest palindrome made from the product of two 3-digit numbers.
*/

func largestpalindromeproduct(lbound int, ubound int) (int, int, int) {
	var maxprod int
	var n1 int
	var n2 int

	for i := lbound; i < ubound; i++ {
		for j := i; j < ubound; j++ {
			curprod := i * j
			if ispalindromicnum(curprod) && (curprod > maxprod) {
				maxprod = curprod
				n1, n2 = i, j
			}
		}
	}

	return maxprod, n1, n2
}

func ispalindromicnum(n int) bool {
	return n == reverseint(n)
}

func reverseint(n int) int {
	var revn int
	for n > 0 {
		revn *= 10
		revn += n % 10
		n /= 10
	}
	return revn
}
