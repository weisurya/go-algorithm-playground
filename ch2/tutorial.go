package main

import "fmt"

func main() {
	a := 23
	n := 9
	d := 12

	fmt.Println(arithmeticProgression(a, n, d))
	fmt.Println(sumOfAP(a, n, d))
}

/*
Arithmetic Progression (AP)
the different between any 2 terms is constant

a = first term of sequence
d = constant difference
n = nth term of AP
T = output og AP
S = sum of AP
*/
func arithmeticProgression(a, n, d int) int {
	return a + (n-1)*d
}

func sumOfAP(a, n, d int) int {
	return n * (2 * arithmeticProgression(a, n, d)) / 2
}
