package ninetynine

import (
	"io"
)

type Problem io.Writer

type Solution io.Reader

type Solver interface {
	Solve(p Problem) (s Solution, err error)
	Problem() Problem
}

type SolverFunc func(p Problem) (s Solution, err error)

type SolverController interface {
	SolveAll() ([]Solution, error)
}

type solverController struct {
	solvers []Solver
}

func (sc solverController) SolveAll() ([]Solution, error) {
	var solutions []Solution

	for _, solver := range sc.solvers {
		s, err := solver.Solve(solver.Problem())
		if err != nil {
			return solutions, err
		}

		solutions = append(solutions, s)
	}

	return solutions, nil
}

func BasicSolverController() SolverController {
	return solverController{
		solvers: []Solver{},
	}
}

func (sc solverController) AddSolver(additionalSolver Solver) solverController {
	return solverController{
		solvers: append(sc.solvers, additionalSolver),
	}
}
