package main

import (
	"flag"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	dimensionStr := flag.String("d", "", "Dimension in feet and inches, e.g., 8'6\"")
	flag.Parse()

	if *dimensionStr == "" {
		fmt.Println("Please provide a dimension using the -d flag, e.g., -d 8'6\"")
		return
	}

	totalFeet, err := parseDimension(*dimensionStr)
	if err != nil {
		fmt.Println("Error parsing dimension:", err)
		return
	}

	sectionsNeeded := int(math.Ceil(totalFeet / 3))

	fmt.Println(sectionsNeeded)
}

func parseDimension(dimStr string) (float64, error) {
	var feet, inches float64
	var err error

	dimStr = strings.ReplaceAll(dimStr, " ", "")

	re := regexp.MustCompile(`^(\d+)'(\d+)"?$`)
	matches := re.FindStringSubmatch(dimStr)
	if len(matches) == 3 {
		feet, err = strconv.ParseFloat(matches[1], 64)
		if err != nil {
			return 0, err
		}
		inches, err = strconv.ParseFloat(matches[2], 64)
		if err != nil {
			return 0, err
		}
	} else {
		return 0, fmt.Errorf("invalid dimension format")
	}

	totalFeet := feet + inches/12
	return totalFeet, nil
}
