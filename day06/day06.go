package day06

import (
	"encoding/csv"
	"io"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// Population represent a population of laternfish
type Population struct {
	adult    []int
	maturing []int
}

func Solve() error {

	population, err := readAndParseInputFile("day06/input")
	if err != nil {
		return err
	}

	solutionPartOne := solvePartOne(population, 80)

	zap.L().Info(
		"Solution",
		zap.Int("Day", 6),
		zap.Int("Part", 1),
		zap.Int("Solution (increments)", solutionPartOne),
	)

	solutionPartTwo := solvePartTwo(population, 256)
	if err != nil {
		return err
	}

	zap.L().Info(
		"Solution",
		zap.Int("Day", 6),
		zap.Int("Part", 2),
		zap.Int("Solution (increments)", solutionPartTwo),
	)

	return nil
}

func solvePartTwo(population Population, days int) int {
	for i := 0; i < days; i++ {
		population.Tick()
	}
	return population.census()
}

func solvePartOne(population Population, days int) int {
	for i := 0; i < days; i++ {
		population.Tick()
	}
	return population.census()
}

func (p Population) census() int {
	populationCount := 0
	for _, n := range p.adult {
		populationCount += n
	}
	for _, n := range p.maturing {
		populationCount += n
	}
	return populationCount
}

func NewPopulation(spawnTime int, maturationTime int) *Population {
	return &Population{
		adult:    make([]int, spawnTime, spawnTime),
		maturing: make([]int, maturationTime, maturationTime),
	}
}

func (p *Population) initFromState(states []int) error {
	// 	Given a list of integers representing the current countdown states of the population, return a Population
	for _, v := range states {
		if v > len(p.adult) {
			// We're out of bounds
			return errors.New("Mass production error! Boop, beep!")
		}
		p.adult[v]++
	}
	return nil
}

func (p *Population) Tick() {
	// Take the front value and add its value startingPopulation children to the end of
	spawnCount := p.adult[0]
	adolescents := p.maturing[0]
	p.adult = append(p.adult[1:], p.adult[0]+adolescents)
	p.maturing = append(p.maturing[1:], spawnCount)
}

func readAndParseInputFile(fname string) (Population, error) {
	rc, err := readerFromFileContents(fname)
	if err != nil {
		return Population{}, err
	}
	return parseInput(rc)
}

func parseInput(input io.Reader) (Population, error) {
	p := NewPopulation(7, 2)
	r := csv.NewReader(input)
	r.ReuseRecord = true
	record, err := r.Read()
	if err != nil {
		return *p, err
	}
	states := make([]int, len(record))
	// We know it's just one line
	for i, f := range record {
		val, err := strconv.Atoi(f)
		if err != nil {
			return *p, errors.Wrap(err, "Beep! Boop! Your string was very fishy")
		}
		states[i] = val
	}
	err = p.initFromState(states)
	if err != nil {
		return *p, errors.Wrap(err, "Bloop! Blup! Something fishy going on!")
	}
	return *p, nil
}

func readerFromFileContents(fname string) (io.Reader, error) {
	contents, err := ioutil.ReadFile(fname)
	if err != nil {

	}
	return strings.NewReader(string(contents)), nil
}
