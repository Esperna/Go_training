package main

import (
	"ch7/ex7_16/eval"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/calc", calcurator)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func calcurator(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("expr")
	if query == "" {
		fmt.Fprintf(w, "no expression\n")
		return
	}
	var expr eval.Expr
	var err error
	if expr, err = eval.Parse(query); err != nil {
		log.Fatal(err)
	}

	vars := make(map[eval.Var]bool)
	if err := expr.Check(vars); err != nil {
		fmt.Fprintf(w, "%s\n", err)
		return
	}
	env := make(map[eval.Var]float64)
	for k := range vars {
		query := r.URL.Query().Get(string(k))
		if query == "" {
			fmt.Fprintf(w, "missing query value of %v\n", k)
			continue
		}
		val, err := strconv.ParseFloat(query, 64)
		if err != nil {
			fmt.Fprintf(w, "invalid parameter %s", err)
			continue
		}
		env[k] = val
	}
	result := expr.Eval(env)
	fmt.Fprintf(w, "%s = %v\n", expr, result)
}
