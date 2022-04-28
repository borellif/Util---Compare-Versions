package main

import (
	"testing"
)

func TestCompareVersionsV1Shorter(t *testing.T) {
	version1 := "1.2.3.4.5"
	version2 := "1.2.3.4.5.6"

	test1, err := compareVersions(version1, version2)

	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	if test1 != -1 {
		t.Errorf("Version 1: %s, Version 2: %s, want -1, got %d", version1, version2, test1)
	}
}

func TestCompareVersionsV1Longer(t *testing.T) {
	version1 := "1.2.3.4.5.6"
	version2 := "1.2.3.4.5"

	test1, err := compareVersions(version1, version2)

	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	if test1 != 1 {
		t.Errorf("Version 1: %s, Version 2: %s, want 1, got %d", version1, version2, test1)
	}
}

func TestCompareVersionsV1(t *testing.T) {
	version1 := "0.1.2.3.4"
	version2 := "1.2.3.4"

	test1, err := compareVersions(version1, version2)

	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	if test1 != -1 {
		t.Errorf("Version 1: %s, Version 2: %s, want -1, got %d", version1, version2, test1)
	}
}

func TestCompareVersionsV2(t *testing.T) {
	version1 := "1.2.99.4"
	version2 := "1.2.3.4"

	test1, err := compareVersions(version1, version2)

	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	if test1 != 1 {
		t.Errorf("Version 1: %s, Version 2: %s, want 1, got %d", version1, version2, test1)
	}
}

func TestCompareVersionsEqual(t *testing.T) {
	version1 := "0.1.2.3.4"
	version2 := "0.1.2.3.4"

	test1, err := compareVersions(version1, version2)

	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	if test1 != 0 {
		t.Errorf("Version 1: %s, Version 2: %s, want 0, got %d", version1, version2, test1)
	}
}

func TestCompareVersionsBadData(t *testing.T) {
	version1 := "a.b.c.d.e"
	version2 := "1.3.5c"

	test1, err := compareVersions(version1, version2)

	if err != nil && test1 != 0 {
		t.Errorf("Error: %s", err.Error())
		t.Errorf("Version 1: %s, Version 2: %s, want 0, got %d", version1, version2, test1)
	}
}

func TestCompareVersionsEdgeCase(t *testing.T) {
	version1 := "1.1.1.1.1.1.1.1.1.1.1.1.1.1.1.1.1.1.1.1.1.1.1.1.1.1.1.1.1.1.1.1.1.1.1.1."
	version2 := "1.1.1.1.1.1.1.1.1.1.1.1.1.1.1.1.1.1.1.1.1.1.1.1.1.1.1.1.1.1.1.1.1.1.1.1.1"

	test1, err := compareVersions(version1, version2)

	if err != nil && test1 != 0 {
		t.Errorf("Error: %s", err.Error())
		t.Errorf("Version 1: %s, Version 2: %s, want 0, got %d", version1, version2, test1)
	}
}

func TestCompareVersionsEdgeCase2(t *testing.T) {
	version1 := ".1.1.1.1.1."
	version2 := "1"

	test1, err := compareVersions(version1, version2)

	if err != nil && test1 != 0 {
		t.Errorf("Error: %s", err.Error())
		t.Errorf("Version 1: %s, Version 2: %s, want 0, got %d", version1, version2, test1)
	}
}
