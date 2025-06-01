package utils

import (
	"os"
	"strings"
)

func CompareOutput(submissionID string) (bool, error) {
	base := "submissions/" + submissionID
	output, err := os.ReadFile(base + "/output.txt")
	if err != nil {
		return false, err
	}
	expected, err := os.ReadFile(base + "/expected.txt")
	if err != nil {
		return false, err
	}
	return strings.TrimSpace(string(output)) == strings.TrimSpace(string(expected)), nil
}
