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

	// Split each of the strings into slices by the dot character into a String array
	splitVer1 := strings.Split(version1, ".")
	splitVer2 := strings.Split(version2, ".")

	// Get length of each split version to use for later
	ver1Length := len(splitVer1)
	ver2Length := len(splitVer2)

	// Start to iterate through the first split version string in order to compare the versions
	for index, strValue := range splitVer1 {

		if (index < ver1Length) && (index < ver2Length) {
			// values are still in String form so we want to convert them
			currentVersion1, err1 := strconv.Atoi(strValue)
			currentVersion2, err2 := strconv.Atoi(splitVer2[index])

			if err1 != nil {
				var b strings.Builder
				b.WriteString("Version 1 has malformed data")
				return 0, errors.New(b.String())
			} else if err2 != nil {
				var b strings.Builder
				b.WriteString("Version 2 has malformed data")
				return 0, errors.New(b.String())

			}

			// These are where both currentVersion1 and currentVersion2 exist
			if currentVersion1 > currentVersion2 {
				return 1, nil
			} else if currentVersion1 < currentVersion2 {
				return -1, nil
			}
		}

		if (index+1 == ver1Length) && (index+1 < ver2Length) {
			// This is when when a version exists for THE NEXT currentVersion2 but doesn't exist for THE NEXT currentVersion1.
			//	Therefore, version2 is newer than version1
			return -1, nil
		} else if (index+1 == ver1Length) && (index+1 > ver2Length) {
			// This is when when a version exists for currentVersion1, but doesn't exist for currentVersion2
			return 1, nil
		}
	}

	// If no returns before, they are equal.
	return 0, nil
}
