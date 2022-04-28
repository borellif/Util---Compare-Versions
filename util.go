package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	winner, err := compareVersions("0.1", "1.1")

	if err == nil {
		fmt.Println("Winner: ", winner)
	} else {
		fmt.Println(err)
	}
}

func compareVersions(version1, version2 string) (int, error) {
	// If for some reason version 1 or version 2 are null, throw an error
	//	and provide context as to which version is null or blank
	if version1 == "" {
		var b strings.Builder
		b.WriteString("Version 1 is null or blank")
		return 0, errors.New(b.String())
	} else if version2 == "" {
		var b strings.Builder
		b.WriteString("Version 2 is null or blank")
		return 0, errors.New(b.String())
	}

	// Split each of the strings by the dot character into a String array
	splitVar1 := strings.Split(version1, ".")
	splitVar2 := strings.Split(version2, ".")

	// Start to iterate through the first split version string in order to compare the versions
	for index, strValue := range splitVar1 {

		// values are still in String form so we want to convert them
		currentVersion1, err1 := strconv.Atoi(strValue)
		currentVersion2, err2 := strconv.Atoi(splitVar2[index])

		if (err1 == nil) && (err2 == nil) {
			if currentVersion1 > currentVersion2 {
				return 1, nil
			} else if currentVersion1 < currentVersion2 {
				return -1, nil
			}
		}
	}

	return 0, nil
}
