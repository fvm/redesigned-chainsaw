package day04

import (
	"bufio"
	"encoding/csv"
	"io"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Entry struct {
	Value   int  `yaml:"value" json:"value"`
	Crossed bool `yaml:"crossed" json:"crossed"`
}
type BoardNumber struct {
	row   int
	col   int
	entry Entry
}

type Board struct {
	Numbers  []BoardNumber `json:"boardnumbers,omitempty" yaml:"boardnumbers"`
	ByRow    [][]*Entry    `json:"byRow,omitempty" yaml:"byRow"`
	ByColumn [][]*Entry    `json:"byColumn,omitempty" yaml:"byColumn"`
}

func Solve() error {

	boards, calledNumbers, err := readAndParseInputFile("day04/input")
	if err != nil {
		return err
	}

	solutionPartOne, err := solvePartOne(boards, calledNumbers)
	if err != nil {
		return err
	}

	zap.L().Info(
		"Solution",
		zap.Int("Day", 4),
		zap.Int("Part", 1),
		zap.Int("Solution (increments)", solutionPartOne),
	)

	solutionPartTwo, err := solvePartTwo(boards, calledNumbers)
	if err != nil {
		return err
	}

	zap.L().Info(
		"Solution",
		zap.Int("Day", 4),
		zap.Int("Part", 2),
		zap.Int("Solution (increments)", solutionPartTwo),
	)

	return nil
}

func solvePartOne(boards []Board, calledNumbers []int) (int, error) {
	winningBoard, _, winningNumber, _ := goBingoOrGoHome(boards, calledNumbers)
	sum := 0
	for _, boardNumber := range winningBoard.Numbers {
		if boardNumber.entry.Crossed {
			continue
		}
		sum += boardNumber.entry.Value
	}
	return sum * winningNumber, nil
}

func solvePartTwo(boards []Board, calledNumbers []int) (int, error) {
	var winningBoard Board
	var winningNumber int
	var remainingBoards = boards
	var remainingNumbers = calledNumbers
	for {
		if len(remainingNumbers) == 0 || len(remainingBoards) == 0 {
			// We're done
			break
		}
		winningBoard, remainingBoards, winningNumber, remainingNumbers = goBingoOrGoHome(remainingBoards, remainingNumbers)
	}
	sum := 0
	for _, boardNumber := range winningBoard.Numbers {
		if boardNumber.entry.Crossed {
			continue
		}
		sum += boardNumber.entry.Value
	}
	return sum * winningNumber, nil
}

func goBingoOrGoHome(boards []Board, calledNumbers []int) (Board, []Board, int, []int) {
	for idx, calledNumber := range calledNumbers {
		for jdx, board := range boards {
			board.cross(calledNumber)
			if board.check() {
				return board, append(boards[:jdx], boards[jdx+1:]...), calledNumber, calledNumbers[idx:]
			}
		}
	}
	return Board{}, []Board{}, 0, []int{}
}

func bingo(es ...*Entry) bool {
	b := true
	for _, e := range es {
		b = b && e.Crossed
	}
	return b
}

func (board Board) check() bool {
	for _, row := range board.ByRow {
		if bingo(row...) {
			return true
		}
	}
	for _, column := range board.ByColumn {
		if bingo(column...) {
			return true
		}
	}
	return false
}

func readAndParseInputFile(fname string) ([]Board, []int, error) {
	rc, err := readerFromFileContents(fname)
	if err != nil {
		return nil, nil, err
	}
	return parseInput(rc)
}
func readerFromFileContents(fname string) (io.Reader, error) {
	contents, err := ioutil.ReadFile(fname)
	if err != nil {

	}
	return strings.NewReader(string(contents)), nil
}

func parseInput(in io.Reader) ([]Board, []int, error) {
	// Read the first line into called numbers
	// The rest are boards
	var err error
	var boards []Board
	var calledNumbers []int

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		line := s.Text()
		if line == "" {
			break
		}
		records, err := csv.NewReader(strings.NewReader(line)).Read()
		if err != nil {
			return boards, calledNumbers, errors.Wrap(err, "Unable to get a CSV reader for the first line")
		}

		calledNumbers = make([]int, len(records))
		for i := 0; i < len(records); i++ {
			number, err := strconv.Atoi(records[i])
			if err != nil {
				return boards, calledNumbers, errors.Wrap(err, "error while parsing called numbers")
			}
			calledNumbers[i] = number
		}
	}

	// Let's start with the boards
	boardString := strings.Builder{}

	for s.Scan() {
		line := s.Text()
		switch line {
		case "":
			// new Board
			b, err := fillBoard(boardString.String(), 5, 5)
			if err != nil {
				return boards, calledNumbers, errors.Wrap(err, "Error while parsing baord strings")
			}
			boardString.Reset()
			boards = append(boards, b)

		default:
			_, err = boardString.WriteString(line + "\n")
			if err != nil {
				return boards, calledNumbers, errors.Wrap(err, "Unable to write line into boardstring while reading file")
			}
		}
	}

	return boards, calledNumbers, nil
}

func (board Board) find(value int) (int, int, bool) {
	for _, n := range board.Numbers {
		if n.entry.Value == value {
			return n.row, n.col, true
		}
	}
	return -1, -1, false
}

func (board *Board) cross(value int) bool {
	r, c, f := board.find(value)
	if f {
		board.ByRow[r][c].Crossed = true
	}
	return f
}

func newBoard(numRows, numCols int) Board {
	b := Board{
		Numbers:  make([]BoardNumber, numRows*numCols),
		ByRow:    nil,
		ByColumn: nil,
	}
	b.ByColumn = make([][]*Entry, numCols)
	rows := make([]*Entry, numCols*numRows)

	for i := 0; i < numCols; i++ {
		b.ByColumn[i] = rows[i*numRows : (i+1)*numRows]
	}
	b.ByRow = make([][]*Entry, numRows)
	columns := make([]*Entry, numRows*numCols)
	for i := 0; i < numRows; i++ {
		b.ByRow[i] = columns[i*numCols : (i+1)*numCols]
	}
	return b
}

func fillBoard(input string, numRows, numCols int) (Board, error) {
	b := newBoard(numRows, numCols)
	r := csv.NewReader(strings.NewReader(input))
	r.Comma = ' '
	r.TrimLeadingSpace = true
	var record []string
	var err error
	idx := 0
	for row := 0; row < 5; row++ {
		record, err = r.Read()
		if err != nil {
			return b, errors.Wrap(err, "Trouble reading the record")
		}
		for col, field := range record {
			value, err := strconv.Atoi(field)
			if err != nil {
				return b, errors.Wrap(err, "Couldnt read a Value into a number")
			}
			b.Numbers[idx] = BoardNumber{
				row: row,
				col: col,
				entry: Entry{
					Value:   value,
					Crossed: false,
				},
			}
			b.ByRow[row][col] = &b.Numbers[idx].entry
			b.ByColumn[col][row] = &b.Numbers[idx].entry
			idx++
		}

	}
	return b, nil
}
