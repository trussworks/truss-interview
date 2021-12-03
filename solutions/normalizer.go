package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

// Use array indexes instead of map keys (ruby solution way)
const (
	Timestamp = iota
	Address
	ZIP
	FullName
	FooDuration
	BarDuration
	TotalDuration
	Notes
)

// For parsing the Duration fields
const (
	HOURS = iota
	MINUTES
	SECONDS
)

func normalize(in []string) []string {
	// The entire CSV is in the UTF-8 character set. Invalid characters should
	// be changed to the Unicode replacement char.
	for ind, val := range in {
		if utf8.Valid([]byte(val)) {
			continue
		}

		in[ind] = strings.ToValidUTF8(in[ind], string(unicode.ReplacementChar))
	}

	// Begin Timestamp stuff.
	//
	// The Timestamp column should be assumed to be in US/Pacific time;
	// please convert it to US/Eastern.
	//
	// Example from Data: 5/12/10 4:48:12 PM
	// Go's required default parse string: Mon Jan 2 15:04:05 MST 2006
	pst, err := time.LoadLocation("PST8PDT")
	if err != nil {
		fmt.Println(err)
	}

	est, err := time.LoadLocation("EST5EDT")
	if err != nil {
		fmt.Println(err)
	}

	inFmt := "1/2/06 3:04:05 PM"
	t, err := time.ParseInLocation(inFmt, in[Timestamp], pst)

	if err != nil {
		fmt.Println(err)
	}

	// Convert to EST
	t = t.In(est)

	// The Timestamp column should formatted in RFC3339 format.
	in[Timestamp] = t.Format(time.RFC3339)
	// End Timestamp stuff

	// All ZIP codes should be formatted as 5 digits. If there are less than 5
	// digits, assume 0 as the prefix.
	in[ZIP] = fmt.Sprintf("%05s", in[ZIP])

	// The FullName column should be converted to uppercase. There will be
	// non-English names.
	in[FullName] = strings.Title(in[FullName])

	// The FooDuration and BarDuration columns are in HH:MM:SS.MS format (where
	// MS is milliseconds); please convert them to the total number of seconds.

	// These two vars will be used to hold the sum of the foo and bar durations
	durationSumSeconds := 0
	durationSumMilliseconds := 0

	// Loop through foo and bar, doing the work on each and tracking
	// cumulative total in the above two vars.
	for durIndex, durField := range in[FooDuration : BarDuration+1] {
		// Offset the index otherwise we'll overwrite TimeStamp
		durIndex += FooDuration

		durationTotalSeconds := 0
		durationMilliseconds := 0

		for i, v := range strings.Split(durField, ":") {
			vAsInt, err := strconv.Atoi(v)
			if err != nil && i != SECONDS {
				fmt.Println(err)
			}

			switch i {
			case HOURS:
				durationTotalSeconds += vAsInt * 60 * 60
			case MINUTES:
				durationTotalSeconds += vAsInt * 60
			case SECONDS:
				wholeSeconds, err := strconv.Atoi(strings.Split(v, ".")[0])
				if err != nil {
					fmt.Println(err)
				}

				durationTotalSeconds += wholeSeconds
				durationMilliseconds, err = strconv.Atoi(strings.Split(v, ".")[1])
				if err != nil {
					fmt.Println(err)
				}
			}
		}

		in[durIndex] = fmt.Sprintf(
			"%d.%d",
			durationTotalSeconds,
			durationMilliseconds,
		)

		durationSumSeconds += durationTotalSeconds
		durationSumMilliseconds += durationMilliseconds
	}

	// The TotalDuration column is filled with garbage data. For each row, please
	// replace the value of TotalDuration with the sum of FooDuration and
	// BarDuration.
	in[TotalDuration] = fmt.Sprintf(
		"%d.%d",
		durationSumSeconds,
		durationSumMilliseconds,
	)

	return in
}

func main() {
	line := 0
	stdinReader := csv.NewReader(os.Stdin)
	stdoutWriter := csv.NewWriter(os.Stdout)

	for {
		rec, err := stdinReader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		// All the normalization including UTF8 stuff happens here on each rec,
		// skip line 1 which is the key.
		line += 1
		if line > 1 {
			rec = normalize(rec)
		}

		// Print out rec after it's processed
		if err := stdoutWriter.Write(rec); err != nil {
			fmt.Println(err)
		}

		// Ensure the underlying io writer writes
		stdoutWriter.Flush()
	}
}
