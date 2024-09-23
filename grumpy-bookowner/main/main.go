package main

import grumpy_bookowner "algorithms/grumpy-bookowner"

func main() {
	customers := []int{1, 0, 1, 2, 1, 1, 7, 5}
	grumpy := []int{0, 1, 0, 1, 0, 1, 0, 1}
	customers2 := []int{1}
	grumpy2 := []int{0}
	grumpy_bookowner.MaxSatisfied(customers, grumpy, 3)
	grumpy_bookowner.MaxSatisfied(customers2, grumpy2, 1)

}
