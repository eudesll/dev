package main

import r "../../runner"

func main() {
	questions := []func(){q1, q4, q5}
	r.RunSolutions(questions)
}
