package main

import (
	"fmt"
	"strings"
	"time"
)

/*
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

const (
	suffixAM = "AM"
	suffixPM = "PM"
)

var tests []string = []string{"12:01:00PM", "12:01:00AM"}

func militaryTime(s string) (string, error) {
	t, err := time.Parse("3:04:05PM", s)
	if err != nil {
		return "", err
	}

	if t.Hour() == 12 {
		if strings.HasSuffix(s, suffixAM) {
			t = t.Add(-12 * time.Hour)
		}
	}

	mt := t.Format("15:04:05")

	return mt, nil
}

func main() {
	for _, ts := range tests {
		t, err := militaryTime(ts)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Military time: ", t)
	}
}
