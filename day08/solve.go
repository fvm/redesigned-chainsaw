package day08

import (
	"bufio"
	"io"
	"strings"
	"unicode"

	"github.com/pkg/errors"
)

// Today's puzzle is figuring out a permutation of characters which maps to one of ten output permutations

type DisplayLine struct {
	inputFields   []string
	displayFields []string
}

type PuzzleInput struct {
	displayLines []DisplayLine
}

type Number struct {
	value        int
	permutations []string
}

/*

We need to find the Horn clause which represents the possible wirings of the displayFields
The displayFields is rigged that it should be able to displayFields the numbers based on its wiring
abcefg		= 0
cf				= 1
acdeg		= 2 ->
acdfg		= 3
bcdf			= 4
abdfg		= 5
abdefg		= 6
acf			= 7
abcdefg	= 8
abcdfg		= 9


a Wiring is a set of 10 wires which map an input symbol to an output symbol
A valid Wiring will be able to produce the regular numbers
So what does this imply?

Let's take an output vallue "ca". This has to be some permutation of "cf", meaning that:
c -> c and f -> f -> a or
c -> a and f -> c

If c -> a, in order to be able to displayFields 1, f -> d
If c -> a, in order to be able to displayFields 7, a -> f, f -> d
If c -> a, in order to be able to displayFields 4, b -> g, d -> b, f -> d

Basically:
If c shifts -2, everything else should shift -2

If
*/

type Mapping struct {
	input  []Number
	output Number
}

// Q: How many ones, fours, sevens and eights are there in the output value?
func solvePartOne(input PuzzleInput) (int, error) {
	count := 0
	for _, displayLine := range input.displayLines {
		// Count all the strings of length two (one), three (seven), four (four) and eight in the displayFields
		for _, digitString := range displayLine.displayFields {
			switch len(digitString) {
			case 2, 3, 4, 7:
				count++
			default:
				continue
			}
		}
	}
	return count, nil
}

func solvePartTwo(input PuzzleInput) (int, error) {
	/*
			Get the one, four, seven and eight from a line
			Given `ab` is one, this means 1) a -> c and b -> f, or 2) b -> c and a -> f (ab -> cf)
			Given seven is `dab` -> `abd`, this means that d -> a (abd -> acf)
			Given four is eafb -> `abef`, this means that 3) e -> d and f -> b or 4) that e -> b and f -> d (abdef -> abcdf)
			Given eight is `acedgfb` -> `abcdefg` this means that: 5) c -> e and g -> g or g -> e and c -> g

			So, using the four inputs we get to 8 possible sets of wirings

		... TBD (or not)
	*/

	return 0, nil
}

func parseInput(rc io.Reader) (PuzzleInput, error) {
	p := PuzzleInput{displayLines: make([]DisplayLine, 200)}

	s := bufio.NewScanner(rc)
	s.Split(bufio.ScanLines)
	lineNumber := 0
	for s.Scan() {
		l := s.Text()
		fields := strings.FieldsFunc(
			l, func(r rune) bool {
				// 	Returns true if whitespace or pipe
				return unicode.IsSpace(r) || r == rune('|')
			},
		)
		if len(fields) > 14 {
			return PuzzleInput{}, errors.Errorf("Whups, your displayFields had a weird amount of values: %d", len(fields))
		}
		displayLine := DisplayLine{
			inputFields:   fields[0:10],
			displayFields: fields[10:],
		}
		p.displayLines[lineNumber] = displayLine
		lineNumber++
	}
	return p, nil
}

// A mapping is valid when it solves the
