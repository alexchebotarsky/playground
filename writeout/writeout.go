package writeout

import (
	"fmt"
	"math"
	"strings"
)

func NumberAsText(num int) string {
	if num == 0 {
		return zero
	}

	absNum := int(math.Abs(float64(num)))
	onesPart := absNum % 10
	tensPart := int(absNum/10) % 10
	hundredsPart := int(absNum/100) % 10
	thousandsPart := int(absNum/1000) % 1000
	millionsPart := int(absNum/1000000) % 1000
	billionsPart := int(absNum/1000000000) % 1000

	maxLen := len(fmt.Sprintf("%d%d%d%d%d%d", onesPart, tensPart, hundredsPart, thousandsPart, millionsPart, billionsPart))
	result := make([]string, 0, maxLen)

	if num < 0 {
		result = append(result, negative)
	}

	if billionsPart != 0 {
		result = append(result, NumberAsText(billionsPart), billion)
	}

	if millionsPart != 0 {
		result = append(result, NumberAsText(millionsPart), million)
	}

	if thousandsPart != 0 {
		result = append(result, NumberAsText(thousandsPart), thousand)
	}

	if hundredsPart != 0 {
		result = append(result, NumberAsText(hundredsPart), hundred)
	}

	if tensPart == 1 {
		result = append(result, teens[onesPart])
	} else {
		if tensPart != 0 {
			result = append(result, tens[tensPart-2])
		}

		if onesPart != 0 {
			result = append(result, ones[onesPart])
		}
	}

	return strings.Join(result, " ")
}

var ones = []string{zero, one, two, three, four, five, six, seven, eight, nine}
var teens = []string{ten, eleven, twelve, thirteen, fourteen, fifteen, sixteen, seventeen, eighteen, nineteen}
var tens = []string{twenty, thirty, forty, fifty, sixty, seventy, eighty, ninety}

const negative = "negative"

const zero = "zero"
const one = "one"
const two = "two"
const three = "three"
const four = "four"
const five = "five"
const six = "six"
const seven = "seven"
const eight = "eight"
const nine = "nine"
const ten = "ten"

const eleven = "eleven"
const twelve = "twelve"
const thirteen = "thirteen"
const fourteen = "fourteen"
const fifteen = "fifteen"
const sixteen = "sixteen"
const seventeen = "seventeen"
const eighteen = "eighteen"
const nineteen = "nineteen"

const twenty = "twenty"
const thirty = "thirty"
const forty = "forty"
const fifty = "fifty"
const sixty = "sixty"
const seventy = "seventy"
const eighty = "eighty"
const ninety = "ninety"

const hundred = "hundred"
const thousand = "thousand"
const million = "million"
const billion = "billion"

func NumberAsOrdinalText(num int) string {
	text := NumberAsText(num)

	for suffix, replacement := range ordinalSuffixReplacementMap {
		cut, ok := strings.CutSuffix(text, suffix)
		if ok {
			return fmt.Sprintf("%s%s", cut, replacement)
		}
	}

	return fmt.Sprintf("%sth", text)
}

const first = "first"
const second = "second"
const third = "third"
const eighth = "eighth"
const ninth = "ninth"

var ordinalSuffixReplacementMap = map[string]string{
	one:   first,
	two:   second,
	three: third,
	eight: eighth,
	nine:  ninth,
	"ve":  "fth",
	"y":   "ieth",
}
