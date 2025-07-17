package testRegex

import (
	"fmt"
	"regexp"
	"testing"
)

var gCloudVolumeStrRegex = regexp.MustCompile(`^(\d+)G\-([\w|\.]*)$`)

func TestRegexSubMatch(t *testing.T) {
	str := "40G-SAS"

	//re := regexp.MustCompile(`^(\d+)G\-([\w|\.]*)$`)
	subMatch := gCloudVolumeStrRegex.FindStringSubmatch(str)

	if len(subMatch) < 3 {
		t.Errorf("Expected at least 3 matches, got %d", len(subMatch))
		return
	}

	fmt.Println("Matched:", subMatch[0])
	fmt.Println("First group:", subMatch[1])
	fmt.Println("Second group:", subMatch[2])
}

func TestRegex(t *testing.T) {
	reg1 := "canal_.*"
	str1 := "canal_12_3"

	matched, err := regexp.MatchString(reg1, str1)
	if err != nil {
		t.Errorf("Error matching regex: %v", err)
		return
	}
	fmt.Println("Matched:", matched)

	reg2 := ".*canal.*"
	str2 := "canal_12_3"

	matched, err = regexp.MatchString(reg2, str2)
	if err != nil {
		t.Errorf("Error matching regex: %v", err)
		return
	}
	fmt.Println("Matched:", matched)

	reg3 := ".*canal"
	str3 := "3_canal"

	matched, err = regexp.MatchString(reg3, str3)
	if err != nil {
		t.Errorf("Error matching regex: %v", err)
		return
	}
	fmt.Println("Matched:", matched)
}
