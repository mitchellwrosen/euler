package main

import (
	"fmt"
	"math"
)

const NIL = -1

type Int int

func (i Int) digitAt(n int) int {
	if n >= i.length() {
		return NIL
	}

	for j := 0; j < n; j++ {
		i /= 10
	}
	return int(i % 10)
}

func (i Int) length() int {
	return int(math.Log10(float64(i)))
}

func champernoweGen() func() int {
	var curNum Int = 1 // The current integer
	var curLoc int = 0 // 0 = ones, 1 = tens, etc.

	return func() int {
		toRet := curNum.digitAt(curLoc)
		if toRet == NIL {
			curNum += 1
			curLoc = 0
		}
		toRet = curNum.digitAt(curLoc)
		return toRet
	}
}

func ex40() {
	champernowe := champernoweGen()

	for i := 0; i < 10; i++ {
		fmt.Println(champernowe())
	}
}
