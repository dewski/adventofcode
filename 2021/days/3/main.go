package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dewski/adventofcode/2021/inputs"
)

type BitCounter struct {
	// TODO: pluralize fields (ones zeros)
	One  int
	Zero int
}

func (bc *BitCounter) Process(bit string) {
	switch bit {
	case "0":
		bc.Zero += 1
	case "1":
		bc.One += 1
	default:
		panic(fmt.Sprintf("Invalid input: %s", bit))
	}
}

func (bc BitCounter) GammaRate() string {
	if bc.One > bc.Zero {
		return "1"
	}

	return "0"
}

func (bc BitCounter) EpsilonRate() string {
	if bc.One > bc.Zero {
		return "0"
	}

	return "1"
}

func NewDiagnosticReport(elements []string) DiagnosticReport {
	size := 0
	// Determine number of bit counting structs we'll need
	if len(elements) > 0 {
		size = len(elements[0])
	}

	// Ensure all elements are the same size
	for _, element := range elements {
		if len(element) != size {
			panic("Invalid input")
		}
	}

	dr := DiagnosticReport{
		BitCounters: make([]BitCounter, size),
	}

	dr.ProcessDiagnostics(elements)

	return dr
}

type DiagnosticReport struct {
	BitCounters []BitCounter
}

func (dr *DiagnosticReport) ProcessDiagnostic(input string) {
	bits := strings.Split(input, "")

	for i, bit := range bits {
		dr.BitCounters[i].Process(bit)
	}
}

func (dr *DiagnosticReport) ProcessDiagnostics(diagnostics []string) {
	for _, diagnostic := range diagnostics {
		dr.ProcessDiagnostic(diagnostic)
	}
}

// GammaRate is determined by finding the most common bit in the corresponding
// position of all numbers in the diagnostic report.
//
// For example, given the following diagnostic report:
//
// 00100
// 11110
// 10110
// 10111
// 10101
// 01111
// 00111
// 11100
// 10000
// 11001
// 00010
// 01010
//
// Considering only the first bit of each number, there are five 0 bits and
// seven 1 bits. Since the most common bit is 1, the first bit of the gamma rate
// is 1.
//
// The most common second bit of the numbers in the diagnostic report is 0, so
// the second bit of the gamma rate is 0.
//
// The most common value of the third, fourth, and fifth bits are 1, 1, and 0,
// respectively, and so the final three bits of the gamma rate are 110.
//
// So, the gamma rate is the binary number 10110, or 22 in decimal.
func (dr DiagnosticReport) GammaRate() int64 {
	parts := make([]string, len(dr.BitCounters))
	for _, bc := range dr.BitCounters {
		parts = append(parts, bc.GammaRate())
	}

	return binaryPartsToInt(parts)
}

// EpsilonRate is calculated in a similar way to GammaRate; rather than use the
// most common bit, the least common bit from each position is used. So, the
// epsilon rate is 01001, or 9 in decimal. Multiplying the gamma rate (22) by
// the epsilon rate (9) produces the power consumption, 198.
func (dr DiagnosticReport) EpsilonRate() int64 {
	parts := make([]string, len(dr.BitCounters))
	for _, bc := range dr.BitCounters {
		parts = append(parts, bc.EpsilonRate())
	}

	return binaryPartsToInt(parts)
}

func binaryPartsToInt(parts []string) int64 {
	binary := strings.Join(parts, "")

	i, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		panic(err)
	}

	return i
}

func (dr DiagnosticReport) PowerConsumption() int64 {
	return dr.GammaRate() * dr.EpsilonRate()
}

type LifeSupportReport struct {
	BitCounters []BitCounter

	size  int
	table [][]string
}

func NewLifeSupportReport(elements []string) LifeSupportReport {
	size := 0
	// Determine number of bit counting structs we'll need
	if len(elements) > 0 {
		size = len(elements[0])
	}

	// Ensure all elements are the same size
	for _, element := range elements {
		if len(element) != size {
			panic("Invalid input")
		}
	}

	table := [][]string{}
	for _, element := range elements {
		table = append(table, strings.Split(element, ""))
	}

	lsr := LifeSupportReport{
		size:  size,
		table: table,
	}

	return lsr
}

func (lsr LifeSupportReport) reduce(reduceFunc func(BitCounter) string) []string {
	elements := lsr.table

	for i := 0; i < lsr.size; i++ {
		bc := BitCounter{}
		cache := map[string][][]string{}
		for _, bits := range elements {
			b := bits[i]
			bc.Process(b)
			cache[b] = append(cache[b], bits)
		}

		keep := reduceFunc(bc)

		elements = cache[keep]

		if len(elements) == 1 {
			return elements[0]
		}
	}

	return elements[0]
}

func (lsr LifeSupportReport) OxygenGeneratorRating() int64 {
	rating := lsr.reduce(func(bc BitCounter) string {
		if bc.One >= bc.Zero {
			return "1"
		}

		return "0"
	})

	return binaryPartsToInt(rating)
}

func (lsr LifeSupportReport) C02ScrubberRating() int64 {
	rating := lsr.reduce(func(bc BitCounter) string {
		if bc.Zero <= bc.One {
			return "0"
		}

		return "1"
	})

	return binaryPartsToInt(rating)
}

func (lsr LifeSupportReport) Rating() int64 {
	return lsr.OxygenGeneratorRating() * lsr.C02ScrubberRating()
}

func main() {
	dr := NewDiagnosticReport(inputs.DayThreeDiagnosticReport)

	fmt.Printf("Gamma Rate: %d\n", dr.GammaRate())
	fmt.Printf("Epsilon Rate: %d\n", dr.EpsilonRate())
	fmt.Printf("Power Consumption: %d\n", dr.PowerConsumption())

	ls := NewLifeSupportReport(inputs.DayThreeDiagnosticReport)
	fmt.Printf("Oxygen Generator Rating: %d\n", ls.OxygenGeneratorRating())
	fmt.Printf("C02 Scrubber Rating: %d\n", ls.C02ScrubberRating())
	fmt.Printf("Rating: %d\n", ls.Rating())
}
