package main

import "math"

func main() {
	for {
		CPU(100)
	}
}

// CPU does stuff
func CPU(times int) []int {
	// Keeps it to a 32 bit int
	num := 40
	var r []int
	for i := 0; i < times; i++ {
		d := bubble(expand(reversePrime(fib(num))))
		r = append(r, d[len(d)-1])
	}
	return r
}

func fib(n int) []int {
	s := []int{1}
	c := 1
	p := 0
	i := n - 1
	for i >= 1 {
		c += p
		p = c - p
		s = append(s, c)
		i--
	}
	return s
}

func prime(n int) bool {
	if n%1 != 0 {
		return false
	} else if n <= 1 {
		return false
	} else if n <= 3 {
		return true
	} else if n%2 == 0 {
		return false
	}
	dl := int(math.Sqrt(float64(n)))
	for d := 3; d <= dl; d += 2 {
		if n%d == 0 {
			return false
		}
	}
	return true
}

func reversePrime(slice []int) []int {
	l := len(slice) - 1
	var r []int
	for l >= 0 {
		if prime(slice[l]) {
			r = append(r, slice[l])
		}
		l--
	}
	return r
}

func expand(slice []int) []int {
	ol := len(slice)
	oc := 0
	l := ol * 10
	var r []int
	for i := 0; i < l; i++ {
		r = append(r, (slice[oc] + (i * 100)))
		if oc < ol-1 {
			oc++
		} else {
			oc = 0
		}
	}
	return r
}

func bubble(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		for y := 0; y < len(slice)-1; y++ {
			if slice[y+1] < slice[y] {
				t := slice[y]
				slice[y] = slice[y+1]
				slice[y+1] = t
			}
		}
	}
	return slice
}
