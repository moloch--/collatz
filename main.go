package main

import (
	"fmt"
	"math/big"
	"time"
)

func main() {
	startValue := new(big.Int)
	var raw string
	fmt.Printf("Start: ")
	fmt.Scanf("%s", &raw)
	startValue.SetString(raw, 10)

	startTime := time.Now()
	steps := collatz(startValue)
	fmt.Printf("Reached 1 in %d step(s) in %s\n", steps, time.Now().Sub(startTime))
}

func collatz(start *big.Int) uint64 {
	one := big.NewInt(1)
	two := big.NewInt(2)
	three := big.NewInt(3)

	value := start
	steps := uint64(0)
	for value.Cmp(one) != 0 {
		if value.Bit(0) == 0 {
			value = value.Div(value, two)
		} else {
			// 3n+1 is always even
			value = value.Mul(value, three)
			value = value.Add(value, one)
			value = value.Div(value, two)
		}
		fmt.Printf("%d\n", value)
		steps++
	}
	return steps
}
