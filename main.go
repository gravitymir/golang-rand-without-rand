package main

import (
	"errors"
	"fmt"
)

//uIntRand random uint without packege rand
func uIntRand(min, max uint) (uint, error) {
	if min > max {
		return 0, fmt.Errorf("get min > max")
	}
	m := make(map[uint]uint, max-min)
	for i := min; i <= max; i++ {
		m[i] = i
	}
	for _, v := range m {
		return v, nil
	}
	return 0, errors.New("error! uIntRand")
}

func main() {
	var min uint = 0
	var max uint = 10000

	for min < max {
		if rand, err := uIntRand(min, max); err != nil {
			fmt.Println(err)
		} else {
			if rand <= 100 {
				fmt.Println(rand)
				break
			}
			fmt.Println(rand)
		}
		min++
	}
}

//go build  -ldflags "-s -w" -o  rand.go  cmd/main.go
