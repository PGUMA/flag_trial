package main

import (
	"flag"
	"fmt"
)

func main() {
	run()
}

func run() {
	var (
		s = flag.Int("p", 10000, "principal")
		r = flag.Float64("r", 0.01, "rate of interest")
		n = flag.Int("n", 5, "running term")
	)

	flag.Parse()

	fmt.Printf("param -p : %d\n", *s)
	fmt.Printf("param -r : %g\n", *r)
	fmt.Printf("param -b : %d\n", *n)

	fmt.Printf("元金：%dを年利%gで%d年運用した場合の予定実績\n", *s, *r, *n)
	var i = 0
	var c = float64(*s)
	for i < *n {
		c = c * (float64(1) + *r)
		i++
		fmt.Printf("%d年後:%d円\n", i, int(c))
	}
}
