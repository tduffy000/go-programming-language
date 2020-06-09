package main

import (
	"ch7/eval"
	"fmt"
	"math"
)

func main() {

	var exprs = make(map[string]eval.Env)

	exprs["sqrt(A / pi)"] = eval.Env{"A": 87616, "pi": math.Pi}
	exprs["pow(x,3)+pow(y,3)"] = eval.Env{"x": 9, "y": 3}
	exprs["5/9*(F-32)"] = eval.Env{"F": 212}
	exprs["-x"] = eval.Env{"x": 1}

	for f, _ := range exprs {
		expr, _ := eval.Parse(f)
		fmt.Printf("Expr: %v\n", expr)
	}

}
