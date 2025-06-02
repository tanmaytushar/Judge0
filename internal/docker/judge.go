package docker

import (
	"fmt"
	"os/exec"
	"path/filepath"

)

func RunDockerJudge(submissionID string,language string) error {
	absPath, err := filepath.Abs("submissions/" + submissionID)
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %v", err)
	}

	var imageName, runCommand string

	if language == "cpp" {
		imageName = "cpp-judge-image"
		runCommand = "g++ /app/main.cpp -o /app/a.out && timeout 2s /app/a.out < /app/input.txt > /app/output.txt"
	} else if language == "java" {
		imageName = "java-judge-image"
		runCommand = "javac /app/Main.java && timeout 2s java -cp /app Main < /app/input.txt > /app/output.txt"
	} else if language == "python" {
		imageName = "python-judge-image"
		runCommand = "timeout 2s python3 /app/main.py < /app/input.txt > /app/output.txt"
	}

	cmd := exec.Command(
		"docker", "run", "--rm",
		"-v", fmt.Sprintf("%s:/app", absPath),
		imageName,
		"bash", "-c",
		runCommand,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("docker execution failed: %v\nOutput: %s", err, string(output))
	}
	return nil
}
