package main

import (
	"bufio"
	"ch7/eval"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// parse call & find required env vars
		expr, err := eval.Parse(scanner.Text())
		if err != nil {
			// parse error
		}
		vars := make(map[eval.Var]bool)
		expr.Check(vars) // pass empty so we know what we need

		// populate empty Env with required vars
		env := eval.Env{}
		reader := bufio.NewReader(os.Stdin)
		for v, _ := range vars {
			fmt.Fprintf(os.Stdout, "Give me a value for %q: ", v)
			val, _ := reader.ReadString('\n')
			x, _ := strconv.ParseFloat(strings.TrimSpace(val), 64)
			env[v] = x
		}
		fmt.Printf("%v=%v\n", expr, expr.Eval(env))
	}
}
