package main

import (
	"ch7/eval"
	"html/template"
	"log"
	"net/http"
)

var calculatorTemplate = template.Must(template.New("").Parse(`
  <h1>Simple Calculator</h1>
  <form method="POST">
    <label>Compute:</label><br />
    <input type="text" name="eq"><br />
  </form>
  {{ if .Answered }}
    <h2>Input: {{ .Eq }}</h2>
    <h2>Answer: {{ .Answer }}</h2>
  {{ end }}
`))

type AnswerPayload struct {
	Answered bool
	Eq       string
	Answer   float64
}

func handler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		calculatorTemplate.Execute(w, nil)
		return
	}
	env := eval.Env{} // right now this only supports eval.literal as operands not eval.Vars
	eq := req.FormValue("eq")
	expr, err := eval.Parse(eq)
	if err != nil {
		w.Write([]byte("Something went wrong"))
		return
	}
	calculatorTemplate.Execute(w, AnswerPayload{true, eq, expr.Eval(env)})
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
