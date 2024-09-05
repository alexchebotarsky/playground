package writeout

import (
	"fmt"
	"testing"
)

func TestNumberAsText(t *testing.T) {
	tests := []struct {
		num  int
		text string
	}{
		{num: 0, text: "zero"},
		{num: 1, text: "one"},
		{num: 2, text: "two"},
		{num: 3, text: "three"},
		{num: 4, text: "four"},
		{num: 5, text: "five"},
		{num: 6, text: "six"},
		{num: 7, text: "seven"},
		{num: 8, text: "eight"},
		{num: 9, text: "nine"},

		{num: 10, text: "ten"},
		{num: 11, text: "eleven"},
		{num: 12, text: "twelve"},
		{num: 13, text: "thirteen"},
		{num: 14, text: "fourteen"},
		{num: 15, text: "fifteen"},
		{num: 16, text: "sixteen"},
		{num: 17, text: "seventeen"},
		{num: 18, text: "eighteen"},
		{num: 19, text: "nineteen"},

		{num: 20, text: "twenty"},
		{num: 21, text: "twenty one"},
		{num: 22, text: "twenty two"},
		{num: 23, text: "twenty three"},
		{num: 24, text: "twenty four"},
		{num: 25, text: "twenty five"},
		{num: 26, text: "twenty six"},
		{num: 27, text: "twenty seven"},
		{num: 28, text: "twenty eight"},
		{num: 29, text: "twenty nine"},

		{num: 30, text: "thirty"},
		{num: 31, text: "thirty one"},
		{num: 32, text: "thirty two"},
		{num: 33, text: "thirty three"},
		{num: 34, text: "thirty four"},
		{num: 35, text: "thirty five"},
		{num: 36, text: "thirty six"},
		{num: 37, text: "thirty seven"},
		{num: 38, text: "thirty eight"},
		{num: 39, text: "thirty nine"},

		{num: 40, text: "forty"},
		{num: 41, text: "forty one"},
		{num: 42, text: "forty two"},
		{num: 43, text: "forty three"},
		{num: 44, text: "forty four"},
		{num: 45, text: "forty five"},
		{num: 46, text: "forty six"},
		{num: 47, text: "forty seven"},
		{num: 48, text: "forty eight"},
		{num: 49, text: "forty nine"},

		{num: 50, text: "fifty"},
		{num: 51, text: "fifty one"},
		{num: 52, text: "fifty two"},
		{num: 53, text: "fifty three"},
		{num: 54, text: "fifty four"},
		{num: 55, text: "fifty five"},
		{num: 56, text: "fifty six"},
		{num: 57, text: "fifty seven"},
		{num: 58, text: "fifty eight"},
		{num: 59, text: "fifty nine"},

		{num: 60, text: "sixty"},
		{num: 61, text: "sixty one"},
		{num: 62, text: "sixty two"},
		{num: 63, text: "sixty three"},
		{num: 64, text: "sixty four"},
		{num: 65, text: "sixty five"},
		{num: 66, text: "sixty six"},
		{num: 67, text: "sixty seven"},
		{num: 68, text: "sixty eight"},
		{num: 69, text: "sixty nine"},

		{num: 70, text: "seventy"},
		{num: 71, text: "seventy one"},
		{num: 72, text: "seventy two"},
		{num: 73, text: "seventy three"},
		{num: 74, text: "seventy four"},
		{num: 75, text: "seventy five"},
		{num: 76, text: "seventy six"},
		{num: 77, text: "seventy seven"},
		{num: 78, text: "seventy eight"},
		{num: 79, text: "seventy nine"},

		{num: 80, text: "eighty"},
		{num: 81, text: "eighty one"},
		{num: 82, text: "eighty two"},
		{num: 83, text: "eighty three"},
		{num: 84, text: "eighty four"},
		{num: 85, text: "eighty five"},
		{num: 86, text: "eighty six"},
		{num: 87, text: "eighty seven"},
		{num: 88, text: "eighty eight"},
		{num: 89, text: "eighty nine"},

		{num: 90, text: "ninety"},
		{num: 91, text: "ninety one"},
		{num: 92, text: "ninety two"},
		{num: 93, text: "ninety three"},
		{num: 94, text: "ninety four"},
		{num: 95, text: "ninety five"},
		{num: 96, text: "ninety six"},
		{num: 97, text: "ninety seven"},
		{num: 98, text: "ninety eight"},
		{num: 99, text: "ninety nine"},

		{num: 100, text: "one hundred"},
		{num: 199, text: "one hundred ninety nine"},
		{num: 900, text: "nine hundred"},

		{num: 1000, text: "one thousand"},
		{num: 1999, text: "one thousand nine hundred ninety nine"},
		{num: 999000, text: "nine hundred ninety nine thousand"},

		{num: 1000000, text: "one million"},
		{num: 1999999, text: "one million nine hundred ninety nine thousand nine hundred ninety nine"},
		{num: 999000000, text: "nine hundred ninety nine million"},

		{num: 1000000000, text: "one billion"},
		{num: 1999999999, text: "one billion nine hundred ninety nine million nine hundred ninety nine thousand nine hundred ninety nine"},
		{num: 2147483647, text: "two billion one hundred forty seven million four hundred eighty three thousand six hundred forty seven"},

		{num: -0, text: "zero"},
		{num: -1, text: "negative one"},
		{num: -2147483647, text: "negative two billion one hundred forty seven million four hundred eighty three thousand six hundred forty seven"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d->%s", tt.num, tt.text), func(t *testing.T) {
			if got := NumberAsText(tt.num); got != tt.text {
				t.Errorf("NumberToText() = %v, want %v", got, tt.text)
			}
		})
	}
}

func TestNumberAsOrdinalText(t *testing.T) {
	tests := []struct {
		num  int
		text string
	}{
		{num: 0, text: "zeroth"},
		{num: 1, text: "first"},
		{num: 2, text: "second"},
		{num: 3, text: "third"},
		{num: 4, text: "fourth"},
		{num: 5, text: "fifth"},
		{num: 6, text: "sixth"},
		{num: 7, text: "seventh"},
		{num: 8, text: "eighth"},
		{num: 9, text: "ninth"},

		{num: 10, text: "tenth"},
		{num: 11, text: "eleventh"},
		{num: 12, text: "twelfth"},
		{num: 13, text: "thirteenth"},
		{num: 14, text: "fourteenth"},
		{num: 15, text: "fifteenth"},
		{num: 16, text: "sixteenth"},
		{num: 17, text: "seventeenth"},
		{num: 18, text: "eighteenth"},
		{num: 19, text: "nineteenth"},

		{num: 20, text: "twentieth"},
		{num: 21, text: "twenty first"},
		{num: 22, text: "twenty second"},
		{num: 23, text: "twenty third"},
		{num: 24, text: "twenty fourth"},
		{num: 25, text: "twenty fifth"},
		{num: 26, text: "twenty sixth"},
		{num: 27, text: "twenty seventh"},
		{num: 28, text: "twenty eighth"},
		{num: 29, text: "twenty ninth"},

		{num: 30, text: "thirtieth"},
		{num: 31, text: "thirty first"},
		{num: 32, text: "thirty second"},
		{num: 33, text: "thirty third"},
		{num: 34, text: "thirty fourth"},
		{num: 35, text: "thirty fifth"},
		{num: 36, text: "thirty sixth"},
		{num: 37, text: "thirty seventh"},
		{num: 38, text: "thirty eighth"},
		{num: 39, text: "thirty ninth"},

		{num: 40, text: "fortieth"},
		{num: 41, text: "forty first"},
		{num: 42, text: "forty second"},
		{num: 43, text: "forty third"},
		{num: 44, text: "forty fourth"},
		{num: 45, text: "forty fifth"},
		{num: 46, text: "forty sixth"},
		{num: 47, text: "forty seventh"},
		{num: 48, text: "forty eighth"},
		{num: 49, text: "forty ninth"},

		{num: 50, text: "fiftieth"},
		{num: 51, text: "fifty first"},
		{num: 52, text: "fifty second"},
		{num: 53, text: "fifty third"},
		{num: 54, text: "fifty fourth"},
		{num: 55, text: "fifty fifth"},
		{num: 56, text: "fifty sixth"},
		{num: 57, text: "fifty seventh"},
		{num: 58, text: "fifty eighth"},
		{num: 59, text: "fifty ninth"},

		{num: 60, text: "sixtieth"},
		{num: 61, text: "sixty first"},
		{num: 62, text: "sixty second"},
		{num: 63, text: "sixty third"},
		{num: 64, text: "sixty fourth"},
		{num: 65, text: "sixty fifth"},
		{num: 66, text: "sixty sixth"},
		{num: 67, text: "sixty seventh"},
		{num: 68, text: "sixty eighth"},
		{num: 69, text: "sixty ninth"},

		{num: 70, text: "seventieth"},
		{num: 71, text: "seventy first"},
		{num: 72, text: "seventy second"},
		{num: 73, text: "seventy third"},
		{num: 74, text: "seventy fourth"},
		{num: 75, text: "seventy fifth"},
		{num: 76, text: "seventy sixth"},
		{num: 77, text: "seventy seventh"},
		{num: 78, text: "seventy eighth"},
		{num: 79, text: "seventy ninth"},

		{num: 80, text: "eightieth"},
		{num: 81, text: "eighty first"},
		{num: 82, text: "eighty second"},
		{num: 83, text: "eighty third"},
		{num: 84, text: "eighty fourth"},
		{num: 85, text: "eighty fifth"},
		{num: 86, text: "eighty sixth"},
		{num: 87, text: "eighty seventh"},
		{num: 88, text: "eighty eighth"},
		{num: 89, text: "eighty ninth"},

		{num: 90, text: "ninetieth"},
		{num: 91, text: "ninety first"},
		{num: 92, text: "ninety second"},
		{num: 93, text: "ninety third"},
		{num: 94, text: "ninety fourth"},
		{num: 95, text: "ninety fifth"},
		{num: 96, text: "ninety sixth"},
		{num: 97, text: "ninety seventh"},
		{num: 98, text: "ninety eighth"},
		{num: 99, text: "ninety ninth"},

		{num: 100, text: "one hundredth"},
		{num: 199, text: "one hundred ninety ninth"},
		{num: 900, text: "nine hundredth"},

		{num: 1000, text: "one thousandth"},
		{num: 1999, text: "one thousand nine hundred ninety ninth"},
		{num: 999000, text: "nine hundred ninety nine thousandth"},

		{num: 1000000, text: "one millionth"},
		{num: 1999999, text: "one million nine hundred ninety nine thousand nine hundred ninety ninth"},
		{num: 999000000, text: "nine hundred ninety nine millionth"},

		{num: 1000000000, text: "one billionth"},
		{num: 1999999999, text: "one billion nine hundred ninety nine million nine hundred ninety nine thousand nine hundred ninety ninth"},
		{num: 2147483647, text: "two billion one hundred forty seven million four hundred eighty three thousand six hundred forty seventh"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d->%s", tt.num, tt.text), func(t *testing.T) {
			if got := NumberAsOrdinalText(tt.num); got != tt.text {
				t.Errorf("NumberToOrdinalText() = %v, want %v", got, tt.text)
			}
		})
	}
}
