package main

import (
	"fmt"
	"github.com/johannm/holdemeq/pkg/eval"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: app.exe AsKd QhQc JhTh9h")
		os.Exit(0)
	}

	hand1 := eval.ParseStr(os.Args[1])
	hand2 := eval.ParseStr(os.Args[2])
	var board []eval.Card
	if len(os.Args) > 3 {
		board = eval.ParseStr(os.Args[3])
	}
	n := 1000000
	win, lose, draw := eval.CalculateHoldemEquity(hand1, hand2, board, n)

	fmt.Printf("n: %v, win: %v, lose: %v, draw: %v\n", n, float64(win)/float64(n), float64(lose)/float64(n), float64(draw)/float64(n))
	fmt.Printf("equity: %v\n", (float64(win)+0.5*float64(draw))/float64(n))
}
