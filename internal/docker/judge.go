package docker
import (
	"fmt"
	"os/exec"
	"path/filepath"
)

func RunDockerJudge(submissionID string) error {
	absPath, err := filepath.Abs("submissions/" + submissionID)
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %v", err)
	}

	cmd := exec.Command(
		"docker", "run", "--rm",
		"-v", fmt.Sprintf("%s:/app", absPath),
		"cpp-judge-image",
		"bash", "-c",
		"g++ /app/main.cpp -o /app/a.out && timeout 2s /app/a.out < /app/input.txt > /app/output.txt",
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("docker execution failed: %v\nOutput: %s", err, string(output))
	}
	return nil
}
