package christmas

import (
	"fmt"
	"strings"

	"github.com/goodleby/playground/writeout"
)

func SongLines() []string {
	lines := []string{}

	for day := 1; day <= len(christmasGifts); day++ {
		lines = append(lines, fmt.Sprintf(mainLyric, writeout.NumberAsOrdinalText(day)))

		todaysGifts := christmasGifts[0:day]
		for i := len(todaysGifts) - 1; i >= 0; i-- {
			var prefix string
			if i == 0 {
				if day == 1 {
					prefix = "A"
				} else {
					prefix = "And a"
				}
			} else {
				prefix = ToTitle(writeout.NumberAsText(i + 1))
			}
			gift := todaysGifts[i]

			lines = append(lines, fmt.Sprintf("%s %s", prefix, gift))
		}
	}

	return lines
}

const mainLyric = "On the %s day of Christmas my true love sent to me"

var christmasGifts = []string{
	"partridge in a pear tree",
	"turtle doves",
	"French hens",
	"calling birds",
	"gold rings",
	"geese a-laying",
	"swans a-swimming",
	"maids a-milking",
	"ladies dancing",
	"lords a-leaping",
	"pipers piping",
	"drummers drumming",
}

func ToTitle(str string) string {
	letters := strings.Split(str, "")
	return strings.ToUpper(letters[0]) + strings.Join(letters[1:], "")
}
