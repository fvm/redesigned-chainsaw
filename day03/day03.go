package day03

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func transpose(input []string) ([]string, error) {
	// Going over each string in the array
	// Count the row
	columns := len(input[0])
	for _, word := range input {
		if len(word) != columns {
			return nil, errors.New("String 'array' with varying length, go f* yourself")
		}
	}

	// Make the same amount of rows as the input has for the output`
	output := make([]string, columns)
	for _, word := range input {
		s := bufio.NewScanner(strings.NewReader(word))
		s.Split(bufio.ScanRunes)
		c := 0
		for s.Scan() {
			r := s.Text()
			output[c] += r
			c++
		}
	}
	return output, nil
}

// type flippedFloppedStrings struct {
// 	transposed []string
// 	original   []string
// }

func Solve() error {

	data, err := readInput("day03/input")
	if err != nil {
		return err
	}

	solutionPartOne, err := solvePartOne(data)
	if err != nil {
		return err
	}

	zap.L().Info(
		"Solution",
		zap.Int("Day", 3),
		zap.Int("Part", 1),
		zap.Int("Solution (increments)", solutionPartOne),
	)

	solutionPartTwo, err := solvePartTwo(data)
	if err != nil {
		return err
	}

	zap.L().Info(
		"Solution",
		zap.Int("Day", 3),
		zap.Int("Part", 2),
		zap.Int("Solution (increments)", solutionPartTwo),
	)

	return nil
}

func getCommonBitsies(d []string) (string, string, error) {
	var mostCommon, leastCommon string
	for _, s := range d {
		if c := strings.Count(s, "1"); float32(c) >= float32(len(s))/2.0 {
			mostCommon += "1"
			leastCommon += "0"
		} else {
			mostCommon += "0"
			leastCommon += "1"
		}
	}
	return mostCommon, leastCommon, nil
}

func solvePartOne(d []string) (int, error) {
	// Count the ones in each rowx
	// If it's less than 500, zeros win
	d, err := transpose(d)
	if err != nil {
		return 0, err
	}
	aGamma, aEpsilon, err := getCommonBitsies(d)
	if err != nil {
		return 0, err
	}
	iGamma, err := strconv.ParseUint(aGamma, 2, 12)
	if err != nil {
		return 0, err
	}
	iEpsilon, err := strconv.ParseUint(aEpsilon, 2, 12)
	if err != nil {
		return 0, err
	}
	return int(iGamma * iEpsilon), nil
}

func readInput(fname string) ([]string, error) {
	var out = make([]string, 1000)
	fptr, err := os.Open(fname)
	if err != nil {
		return out, err
	}
	defer func(fptr *os.File) {
		err := fptr.Close()
		if err != nil {
			// 	Whups
			zap.L().Error("Error closing file", zap.Error(err))
		}
	}(fptr)

	s := bufio.NewScanner(fptr)
	s.Split(bufio.ScanLines)
	count := 0
	for s.Scan() {
		out[count] = s.Text()
		count++
	}
	return out, nil
}

func okayOnceMoreWithFeeling(data []string) (string, string, error) {
	// 	Go through each column and figure out the most common and least common bitsies
	o2 := make([]string, len(data))
	copy(o2, data)
	var err error
	var target string
	for column := 0; column < len(o2[0]); column++ {
		// Transposed
		o2, err = transpose(o2)
		if err != nil {
			return "", "", err
		}
		if strings.Count(o2[column], "0") == strings.Count(o2[column], "1") {
			target = "1"
		} else {
			target, _, err = getCommonBitsies([]string{o2[column]})
		}
		// Normal
		o2, err = transpose(o2)
		// Filter mode
		// Get rid of all the entries which are not equal to the target at the current position
		o2, err = reduceDataByTarget(o2, target, column)
		if len(o2) == 1 {
			break
		}
	}

	co2 := make([]string, len(data))
	copy(co2, data)
	for column := 0; column < len(co2[0]); column++ {
		// Transposed
		co2, err = transpose(co2)
		if err != nil {
			return "", "", err
		}
		if strings.Count(co2[column], "0") == strings.Count(co2[column], "1") {
			target = "0"
		} else {
			_, target, err = getCommonBitsies([]string{co2[column]})
		}
		// Normal
		co2, err = transpose(co2)
		// Filter mode
		// Get rid of all the entries which are not equal to the target at the current position
		co2, err = reduceDataByTarget(co2, target, column)
		if len(co2) == 1 {
			break
		}
	}
	return o2[0], co2[0], err
}

func reduceDataByTarget(input []string, target string, position int) ([]string, error) {
	// Get rid of all entries in which don't have `target` at `position`, or you just have one left
	currentLength := len(input)
	for i := 0; i < currentLength; i++ {
		if string(input[i][position]) != target {
			currentLength--
			// If this would leave us with an empty list, return
			if currentLength == 0 {
				return input, nil
			}
			input = append(input[:i], input[i+1:]...)
			// Set the counter back
			i--
		}
	}
	return input, nil
}

// To ne shot behind the barn
// func doThatWeirdFilterShit(f flippedFloppedStrings) (string, string, error) {
// 	mostCommon, leastCommon, err := getCommonBitsies(f.transposed)
// 	if err != nil {
// 		return "", "", nil
// 	}
// 	o2 := make([]string, len(f.original))
// 	copy(o2, f.original)
// 	// Use the strings to filter out the values
// 	for i, m := range mostCommon {
// 		// The character at position i is the most common one in column i
// 		ms := string(m)
// 		// // Go over each string in the list
// 		st := len(o2)
// 		if st <= 1 {
// 			break
// 		}
// 		for j := 0; j < st; j++ {
// 			if string(o2[j][i]) != ms {
// 				st--
// 				// If this would make the length 0, stop
// 				if st == 0 {
// 					break
// 				}
// 				o2 = append(o2[:j], o2[j+1:]...)
// 				j--
// 			}
// 		}
// 		if st == 1 {
// 			break
// 		}
// 	}
// 	co2 := make([]string, len(f.original))
// 	copy(co2, f.original)
// 	// Use the strings to filter out the values
// 	for i, l := range leastCommon {
// 		// The character at position i is the most common one in column i
// 		ls := string(l)
// 		// // Go over each string in the list
// 		st := len(co2)
// 		if st <= 1 {
// 			break
// 		}
// 		for j := 0; j < st; j++ {
// 			if string(co2[j][i]) != ls {
// 				st--
// 				// If this would make the length 0, stop
// 				if st == 0 {
// 					break
// 				}
// 				// Drop it like it's hot
// 				co2 = append(co2[:j], co2[j+1:]...)
// 				j--
//
// 			}
// 		}
// 	}
// 	return o2[0], co2[0], err
// }

func solvePartTwo(f []string) (int, error) {
	co2, o2, err := okayOnceMoreWithFeeling(f)
	co2rating, err := strconv.ParseUint(co2, 2, 16)
	if err != nil {
		return 0, err
	}
	o2rating, err := strconv.ParseUint(o2, 2, 16)
	if err != nil {
		return 0, err
	}
	return int(co2rating * o2rating), nil
}
