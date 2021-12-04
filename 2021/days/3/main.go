package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dewski/adventofcode/2021/inputs"
)

type DiagnosticReport struct {
	BitCounters []BitCounter
}

func (dr *DiagnosticReport) ProcessDiagnostic(input string) {
	bits := strings.Split(input, "")

	for i, bit := range bits {
		if bit == "0" {
			dr.BitCounters[i].Zero += 1
		} else {
			dr.BitCounters[i].One += 1
		}
	}
}

func (dr *DiagnosticReport) ProcessDiagnostics(diagnostics []string) {
	for _, diagnostic := range diagnostics {
		dr.ProcessDiagnostic(diagnostic)
	}
}

func (dr DiagnosticReport) GammaRate() int64 {
	gr := make([]string, len(dr.BitCounters))
	for _, bc := range dr.BitCounters {
		gr = append(gr, bc.GammaRate())
	}

	binary := strings.Join(gr, "")

	i, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		panic(err)
	}

	return i
}

func (dr DiagnosticReport) EpsilonRate() int64 {
	gr := make([]string, len(dr.BitCounters))
	for _, bc := range dr.BitCounters {
		gr = append(gr, bc.EpsilonRate())
	}

	binary := strings.Join(gr, "")

	i, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		panic(err)
	}

	return i
}

func (dr DiagnosticReport) PowerConsumption() int64 {
	return dr.GammaRate() * dr.EpsilonRate()
}

type BitCounter struct {
	One  int
	Zero int
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

var table = make([]BitCounter, 12)

func main() {
	dr := DiagnosticReport{
		BitCounters: make([]BitCounter, 12),
	}

	dr.ProcessDiagnostics(inputs.DayThreeDiagnosticReport)

	fmt.Println(dr)
	fmt.Println(dr.GammaRate())
	fmt.Println(dr.EpsilonRate())
	fmt.Println(dr.PowerConsumption())
}
