package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func PrepareSubmissionFiles(submissionID, code, problemID, language string) error {
	basePath := filepath.Join("submissions", submissionID)
	err := os.MkdirAll(basePath, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create submission directory: %v", err)
	}

	var filename string
	switch language {
	case "cpp":
		filename = "main.cpp"
	case "java":
		filename = "Main.java"
	case "python":
		filename = "main.py"
	default:
		return fmt.Errorf("unsupported language: %s", language)
	}

	// Write the source code file
	err = os.WriteFile(filepath.Join(basePath, filename), []byte(code), 0644)
	if err != nil {
		return fmt.Errorf("failed to write %s: %v", filename, err)
	}

	// Copy input.txt and expected.txt from testcases/<problem_id>/
	inputPath := filepath.Join("testcases", problemID, "input.txt")
	expectedPath := filepath.Join("testcases", problemID, "expected.txt")

	inputBytes, err := os.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("failed to read input.txt: %v", err)
	}
	expectedBytes, err := os.ReadFile(expectedPath)
	if err != nil {
		return fmt.Errorf("failed to read expected.txt: %v", err)
	}

	err = os.WriteFile(filepath.Join(basePath, "input.txt"), inputBytes, 0644)
	if err != nil {
		return fmt.Errorf("failed to write input.txt: %v", err)
	}

	err = os.WriteFile(filepath.Join(basePath, "expected.txt"), expectedBytes, 0644)
	if err != nil {
		return fmt.Errorf("failed to write expected.txt: %v", err)
	}

	return nil
}
