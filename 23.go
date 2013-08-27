//A perfect number is a number for which the sum of its proper divisors is exactly
//equal to the number. For example, the sum of the proper divisors of 28 would be
//1 + 2 + 4 + 7 + 14 = 28, which means that 28 is a perfect number.

//A number n is called deficient if the sum of its proper divisors is less than n
//and it is called abundant if this sum exceeds n.

//As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16, the smallest
//number that can be written as the sum of two abundant numbers is 24. By
//mathematical analysis, it can be shown that all integers greater than 28123 can
//be written as the sum of two abundant numbers. However, this upper limit cannot
//be reduced any further by analysis even though it is known that the greatest
//number that cannot be expressed as the sum of two abundant numbers is less than
//this limit.

//Find the sum of all the positive integers which cannot be written as the sum of
//two abundant numbers.
package main

import (
	"container/list"
	"fmt"
	"math"
)

type IntSlice []int

func (s *IntSlice) push(i int) {
	*s = append(*s, i)
}

func (s IntSlice) sum() int {
	var sum int
	for _, val := range s {
		sum += val
	}
	return sum
}

func (s IntSlice) end() int {
	return s[len(s)-1]
}

func (s IntSlice) isSumOfTwo(n int) bool {
	left := 0
	right := len(s) - 1

	for left <= right {
		sum := s[left] + s[right]
		if sum < n {
			left++
		} else if sum > n {
			right--
		} else {
			return true
		}
	}

	return false
}

type IntList struct {
	*list.List
}

func newIntList(a, b int) IntList {
	l := list.New()
	for i := a; i < b; i++ {
		l.PushBack(i)
	}
	return IntList{l}
}

func (l IntList) Print() {
	fmt.Print("[ ")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%d ", e.Value.(int))
	}
	fmt.Println("]")
}

func intMod(x, y int) int {
	return int(math.Mod(float64(x), float64(y)))
}

func advance(e *list.Element, n int) *list.Element {
	e2 := e
	for i := 0; i < n; i++ {
		e2 = e2.Next()
		if e2 == nil {
			return nil
		}
	}
	return e2
}

func getProperDivisors(n int) []int {
	divisors := make(IntSlice, 0, 10)
	divisors.push(1)

	sqrt := int(math.Sqrt(float64(n)))
	potentialFactors := newIntList(2, sqrt+1)
	for e := potentialFactors.Front(); e != nil; e = e.Next() {
		val := e.Value.(int)

		if intMod(n, val) == 0 {
			divisors.push(val)
			if val != n/val {
				divisors.push(n / val)
			}
		} else {
			for e2 := advance(e, val); e2 != nil; e2 = advance(e2, val) {
				e2 = e2.Prev()
				potentialFactors.Remove(e2.Next())
			}
		}
	}

	return divisors
}

func isAbundant(n int) bool {
	divisors := getProperDivisors(n)
	sum := IntSlice(divisors).sum()
	return sum > n
}

func makeAbundantNumbersGenerator() chan int {
	out := make(chan int)
	go func() {
		for i := 1; ; i++ {
			if isAbundant(i) {
				out <- i
			}
		}
	}()
	return out
}

func ex23() {
	// Numbers that cannot be expressed as the sum of two abundant numbers
	nums := make(IntSlice, 0, 100)

	abun := make(IntSlice, 0, 1000)
	for i := 1; i < 28123; i++ {
		if isAbundant(i) {
			abun.push(i)
		}

		if !abun.isSumOfTwo(i) {
			nums.push(i)
		}
	}

	fmt.Println(nums.sum())
}
