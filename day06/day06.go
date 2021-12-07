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
	adult    []uint64
	maturing []uint64
}

func Solve() error {

	population, err := readAndParseInputFile("day06/input")
	if err != nil {
		return err
	}

	solutionPartOne := solvePartOne(population, 80)

	zap.L().Info(
		"Solution",
		zap.Uint64("Day", 6),
		zap.Uint64("Part", 1),
		zap.Uint64("Solution (increments)", solutionPartOne),
	)

	solutionPartTwo := solvePartTwo(population, 256)
	if err != nil {
		return err
	}

	zap.L().Info(
		"Solution",
		zap.Uint64("Day", 6),
		zap.Uint64("Part", 2),
		zap.Uint64("Solution (increments)", solutionPartTwo),
	)

	return nil
}

func solvePartTwo(population Population, days uint64) uint64 {
	for i := uint64(0); i < days; i++ {
		population.Tick()
	}
	return population.census()
}

func solvePartOne(population Population, days uint64) uint64 {
	for i := uint64(0); i < days; i++ {
		population.Tick()
	}
	return population.census()
}

func (p Population) census() uint64 {
	populationCount := uint64(0)
	for _, n := range p.adult {
		populationCount += n
	}
	for _, n := range p.maturing {
		populationCount += n
	}
	return populationCount
}

func NewPopulation(spawnTime uint64, maturationTime uint64) *Population {
	return &Population{
		adult:    make([]uint64, spawnTime, spawnTime),
		maturing: make([]uint64, maturationTime, maturationTime),
	}
}

func (p *Population) initFromState(states []uint64) error {
	// 	Given a list of uint64egers representing the current countdown states of the population, return a Population
	for _, v := range states {
		if v > uint64(len(p.adult)) {
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
	states := make([]uint64, len(record))
	// We know it's just one line
	for i, f := range record {
		val, err := strconv.Atoi(f)
		if err != nil {
			return *p, errors.Wrap(err, "Beep! Boop! Your string was very fishy")
		}
		states[i] = uint64(val)
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
