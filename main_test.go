package main

import (
	"testing"
)

func TestToChar(t *testing.T) {
	chars := []struct {
		num  int
		char rune
	}{
		{1, 'A'},
		{6, 'F'},
		{26, 'Z'},
		{27, '['},
		{33, 'a'},
	}

	for _, ch := range chars {
		result := toChar(ch.num)
		if result != ch.char {
			t.Errorf("toChar conversion error, expected '%c', got '%c'", ch, result)
		}
	}
}

func TestReverse(t *testing.T) {
	str := []struct {
		forward string
		reverse string
	}{
		{"abcd", "dcba"},
		{"ABCD", "DCBA"},
		{"Encyclopaedia", "aideapolcycnE"},
	}

	for _, s := range str {
		r := reverse(s.forward)
		if r != s.reverse {
			t.Errorf("reversing error, expected '%s', got '%s'", s.reverse, r)
		}
	}
}

func TestColumnString(t *testing.T) {
	cols := []struct {
		colNumber int
		colString string
	}{
		{1, "A"},
		{2, "B"},
		{26, "Z"},
		{27, "AA"},
		{52, "AZ"},
		{53, "BA"},
	}

	for _, c := range cols {
		cStr := columnString(c.colNumber)
		if cStr != c.colString {
			t.Errorf("column string error, expected '%s', got '%s'", c.colString, cStr)
		}
	}
}

func TestGenerateFormula(t *testing.T) {
	tests := []struct {
		column       string
		row          int
		formulaWeeks int
		benchWeek    int
		expected     string
	}{
		{"A", 1, 1, 0, "=(SUMIFS('Resource Schedule'!$A1,'Resource Schedule'!$A$2\">\"&TODAY()-7,'Resource Schedule'!$A2\"<=\"&TODAY()))"},
		{"D", 4, 1, 1, "=(SUMIFS('Resource Schedule'!$D4,'Resource Schedule'!$D$2\">\"&TODAY()-7,'Resource Schedule'!$D2\"<=\"&TODAY()+7))"},
		{"A", 1, 2, 2, "=(SUMIFS('Resource Schedule'!$A1,'Resource Schedule'!$A$2\">\"&TODAY()-7,'Resource Schedule'!$A2\"<=\"&TODAY()+14)+SUMIFS('Resource Schedule'!$G1,'Resource Schedule'!$G$2\">\"&TODAY()-7,'Resource Schedule'!$G2\"<=\"&TODAY()+14))"},
	}

	for _, test := range tests {
		formula := generateFormula(test.column, test.row, test.formulaWeeks, test.benchWeek)
		if formula != test.expected {
			t.Errorf("error generating formula, expected `%s`, got `%s`", test.expected, formula)
		}
	}
}
