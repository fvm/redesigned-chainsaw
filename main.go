package main

import (
	"log"

	"gitlab.com/frankvanmeurs/redesigned-chainsaw/day00"
	"gitlab.com/frankvanmeurs/redesigned-chainsaw/ninetynine"
)

func main() {
	var problemForToday ninetynine.Problem

	var solutionForToday ninetynine.Solution

	solverForToday := day00.Solver()
	problemForToday = solverForToday.Problem()

	_, err := problemForToday.Write([]byte{})
	if err != nil {
		log.Fatal(err)
	}

	solutionForToday, err = solverForToday.Solve(problemForToday)
	if err != nil {
		log.Fatalln(err)
	}

	log.Print(solutionForToday)
}
