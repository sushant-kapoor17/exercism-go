package collatzconjecture

import "fmt"

func CollatzConjecture(n int) (int, error) {
	steps := 0
	if n <= 0 {
		err := fmt.Errorf("incorrect number provided")
		return n, err
	}
	x := collatzconjecture(n, &steps)
	return *x, nil
}

func collatzconjecture(n int, steps *int) *int {
	if n == 1 {
		return new(int)
	}
	if n%2 == 0 {
		*steps += 1
		collatzconjecture(n/2, steps)
	} else {
		*steps += 1
		collatzconjecture(3*n+1, steps)
	}
	return steps

}
