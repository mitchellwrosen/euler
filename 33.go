package main

import (
	"fmt"
)

type Num int

func (n Num) ones() Num { return n % 10 }
func (n Num) tens() Num { return n / 10 }

type Fraction struct {
	num Num
	den Num
}

func (f Fraction) reduce() (Fraction, bool) {
	nt, no := f.num.tens(), f.num.ones()
	dt, do := f.den.tens(), f.den.ones()

	if f.num == f.den {
		return f, false
	}

	if nt == dt && nt != 0 {
		return Fraction{no, do}, true
	} else if nt == do && nt != 0 {
		return Fraction{no, dt}, true
	} else if no == dt && no != 0 {
		return Fraction{nt, do}, true
	} else if no == do && no != 0 {
		return Fraction{nt, dt}, true
	}

	return f, false
}

func (f Fraction) val() float64 {
	if f.den == 0 {
		return -1
	}

	return float64(f.num) / float64(f.den)
}

func (f Fraction) String() string {
	return fmt.Sprintf("%d/%d", f.num, f.den)
}

func ex33() {
	var num, den Num
	for num = 10; num < 99; num++ {
		for den = num + 1; den < 99; den++ {
			frac := Fraction{num, den}
			if reduced, ok := frac.reduce(); ok {
				if frac.val() == reduced.val() {
					fmt.Printf("%s -> %s\n", frac, reduced)
				}
			}
		}
	}
}
