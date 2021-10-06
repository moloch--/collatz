package main

import (
	"fmt"
	"math/big"
	"time"
)

func main() {
	largest := ""
	largestSteps := uint64(0)
	count := big.NewInt(1)
	for {
		count.Add(count, big.NewInt(1))
		startTime := time.Now()
		value := big.NewInt(0)
		value.Set(count)
		steps := collatz(value)
		fmt.Printf("\r\u001b[2K%s - %d reached 1 in %d step(s) in %s", largest, count, steps, time.Since(startTime))
		if largestSteps < steps {
			largestSteps = steps
			largest = fmt.Sprintf("largest: %d (%d steps)", count, steps)
		}
	}
}

func collatz(start *big.Int) uint64 {
	one := big.NewInt(1)
	two := big.NewInt(2)
	three := big.NewInt(3)

	value := start.Add(start, big.NewInt(0))
	steps := uint64(0)
	for value.Cmp(one) != 0 {
		if value.Bit(0) == 0 {
			value.Div(value, two)
		} else {
			// 3n+1 is always even
			value.Mul(value, three)
			value.Add(value, one)
			value.Div(value, two)
		}
		steps++
	}
	return steps
}
