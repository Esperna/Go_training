package main

import (
	"bufio"
	"ch7/ex7_15/eval"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Print("Input your expression>")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			fmt.Fprintf(os.Stderr, "%s\n", scanner.Err())
			continue
		}
		in := scanner.Text()
		expr, err := eval.Parse(in)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			continue
		}
		vars := make(map[eval.Var]bool)
		if err := expr.Check(vars); err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			continue
		}
		env := make(map[eval.Var]float64)
		for k := range vars {
			fmt.Printf("%s=", k)
			if !scanner.Scan() {
				fmt.Fprintf(os.Stderr, "%s\n", scanner.Err())
				continue
			}
			in := scanner.Text()
			val, err := strconv.ParseFloat(in, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s\n", err)
				continue
			}
			env[k] = val
		}
		result := expr.Eval(env)
		fmt.Println(result)
		switch in {
		default:
			fmt.Print("Input your expression>")
			continue
		}
	}
}
