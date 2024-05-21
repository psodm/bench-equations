package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// columnString converts a column number to it's equivalent Excel string representation
// eg. column 1 = 'A', column 26 = 'Z', column 27 = 'AA'
//
// Modified from the approach outlined here:
// https://www.geeksforgeeks.org/find-excel-column-name-given-number/
func columnString(n int) string {
	// Only valid for positive integers
	if n < 1 {
		return ""
	}

	char := []string{}
	rem := 0

	for n > 0 {
		rem = n % 26
		if rem == 0 {
			char = append(char, "Z")
			n = (n / 26) - 1
		} else {
			char = append(char, string(toChar(rem)))
			n = n / 26
		}
	}

	str := strings.Join(char[:], "")
	return reverse(str)
}

// generateFormula creates a bench tracker formula for the Resources Tracker spreadsheet
func generateFormula(startColumn string, row, formulaWeeks, benchWeek int) string {
	sByte := []byte(strings.ToUpper(startColumn))
	sInt := int(sByte[0])
	str := "=("
	for i := 0; i < formulaWeeks; i++ {
		column := columnString(sInt - 64 + i*6)
		f := fmt.Sprintf(
			"SUMIFS('Resource Schedule'!$%s%d,'Resource Schedule'!$%s$2\">\"&TODAY()-7,'Resource Schedule'!$%s2\"<=\"&TODAY()+%d)+",
			column, row, column, column, benchWeek*7)
		str = fmt.Sprintf("%s%s", str, f)
	}
	str = strings.ReplaceAll(str, "+0", "")     // remove all '+0' in first week formula
	return fmt.Sprintf("%s)", str[:len(str)-1]) // remove last '+' and add closing bracket
}

// reverse reverses a string
func reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

// toChar converts an integer to it's equivalent ASCII char
func toChar(i int) rune {
	return rune('A' + i - 1)
}

func main() {
	flags := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	start := flags.String("column", "T", "first column to calculate from")
	weeks := flags.Int("weeks", 52, "number of weeks to include")
	row := flags.Int("row", 3, "row to calculate for")

	flags.Parse(os.Args[1:])

	for i := 0; i < 4; i++ {
		formula := generateFormula(*start, *row, *weeks, i)
		fmt.Printf("Week %d:\n%s\n", i+1, formula)
	}
}
