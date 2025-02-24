package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const version = "1.1.0"

func main() {
	dimensionStr := flag.String("d", "", "Dimension in feet and inches, e.g., 8ft6in or 8f6i or 8'6\"")
	numSections := flag.Int("s", 1, "Number of sections to divide the dimension into (default is 1 section)")
	totalInchesFlag := flag.Bool("in", false, "Output the total inches from the input dimension")
	showVersion := flag.Bool("v", false, "Show program version")
	flag.BoolVar(showVersion, "version", false, "Show program version")

	flag.Parse()

	if *showVersion {
		fmt.Printf("IMP version: %v\n", version)
		os.Exit(0)
	}

	if *dimensionStr == "" {
		fmt.Println("Please provide a dimension using the -d flag, e.g., -d 8ft6in")
		return
	}

	if *numSections <= 0 {
		fmt.Println("Number of sections must be a positive integer.")
		return
	}

	totalFeet, err := parseDimension(*dimensionStr)
	if err != nil {
		fmt.Println("Error parsing dimension:", err)
		return
	}

	if *totalInchesFlag {
		totalInches := int(math.Round(totalFeet * 12))
		fmt.Printf("Total inches: %d\n", totalInches)
		return
	}

	sectionSizeFeet := totalFeet / float64(*numSections)
	feet, inches := splitToFeetAndInches(sectionSizeFeet)
	fmt.Printf("Each section size: %d'%d\"\n", feet, inches)
}

func parseDimension(dimStr string) (float64, error) {
	var feet, inches float64
	var err error

	dimStr = strings.ToLower(strings.ReplaceAll(dimStr, " ", ""))

	dimStr = strings.ReplaceAll(dimStr, "'", "f")
	dimStr = strings.ReplaceAll(dimStr, "ft", "f")
	dimStr = strings.ReplaceAll(dimStr, "\"", "i")
	dimStr = strings.ReplaceAll(dimStr, "in", "i")

	// Match format like "8f6i" or "8f6" or "8f"
	re := regexp.MustCompile(`^(\d+)f(\d*)i?$`)
	matches := re.FindStringSubmatch(dimStr)
	if len(matches) > 0 {
		// Parse feet
		feet, err = strconv.ParseFloat(matches[1], 64)
		if err != nil {
			return 0, err
		}

		if len(matches) > 2 && matches[2] != "" {
			inches, err = strconv.ParseFloat(matches[2], 64)
			if err != nil {
				return 0, err
			}
		}
	} else {
		return 0, fmt.Errorf("invalid dimension format")
	}

	totalFeet := feet + inches/12
	return totalFeet, nil
}

func splitToFeetAndInches(totalFeet float64) (int, int) {
	feet := int(math.Floor(totalFeet))
	inches := int(math.Round((totalFeet - float64(feet)) * 12))
	return feet, inches
}
