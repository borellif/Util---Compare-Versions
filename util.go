package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, world.")

	winner, err := compareVersions("0.1", "1.1")

	if err == nil {
		fmt.Println("Winner: ", winner)
	} else {
		fmt.Println(err)
	}
}

func compareVersions(version1, version2 string) (int, error) {
	if version1 == "" {
		var b strings.Builder
		b.WriteString("Version 1 is null or blank")
		return 0, errors.New(b.String())
	} else if version2 == "" {
		var b strings.Builder
		b.WriteString("Version 2 is null or blank")
		return 0, errors.New(b.String())
	}

	splitVar1 := strings.Split(version1, ".")
	splitVar2 := strings.Split(version2, ".")

	for index, value := range splitVar1 {
		if value > splitVar2[index] {
			return 1, nil
		} else if value < splitVar2[index] {
			return -1, nil
		}
	}

	return 0, nil
}
