package day00

import (
	"log"

	"gitlab.com/frankvanmeurs/redesigned-chainsaw/ninetynine"
)

type problemSolver struct{}

func (ps problemSolver) Problem() ninetynine.Problem {
	return problem{}
}

type problem struct{}

type solution struct{}

func (p problem) Read(b []byte) (n int, err error) {
	return n, err
}

func (p problem) Write(b []byte) (n int, err error) {
	return n, err
}

func (s solution) Write(b []byte) (n int, err error) {
	return n, err
}

func (s solution) Read(p []byte) (n int, err error) {
	return n, err
}

func Solver() ninetynine.Solver {
	return problemSolver{}
}

func (ps problemSolver) Solve(p ninetynine.Problem) (ninetynine.Solution, error) {
	defer log.Printf("%s", "plop")
	log.Print("piep")

	return ps.solve(p)
}

func (ps problemSolver) solve(p ninetynine.Problem) (ninetynine.Solution, error) {
	return solution{}, nil
}
