package main

import (
	"ch7/eval"
	"fmt"
	"math"
)

func main() {

	var exprs = make(map[string]eval.Env)

	exprs["min(a,b,c)"] = eval.Env{"a": 87616, "b": math.Pi, "c": -1}
	exprs["max(a,b,c)"] = eval.Env{"a": 87616, "b": math.Pi, "c": -1}

	for f, env := range exprs {
		expr, _ := eval.Parse(f)
		fmt.Printf("Expr: %v\n", expr)
		fmt.Printf("Env: %+v\n", env)
		x := expr.Eval(env)
		fmt.Printf("result = %v\n", x)
	}

}
