package melody

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func RunCmd(query string, melodyCmd string) (*string, error) {
	homeDir, err := os.UserHomeDir()
	scriptPath := filepath.Join(homeDir, "Projects", "Melody.CLI", "melody_cli.py")
	cmd := exec.Command("python", scriptPath, melodyCmd, query)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	} else {
		outputStr := string(output)
		res := &outputStr
		return res, nil
	}
}

func StartMelodyDaemon() (*exec.Cmd, error) {
	// Find the path to the melody_daemon.py script
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get user home directory: %v", err)
	}

	scriptPath := filepath.Join(homeDir, "Projects", "Melody.CLI", "melody_daemon.py")
	if _, err := os.Stat(scriptPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("melody daemon script not found at %s", scriptPath)
	}

	// Start the daemon process
	cmd := exec.Command("python", scriptPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("failed to start melody daemon: %v", err)
	}

	log.Printf("Started Melody daemon (PID: %d)", cmd.Process.Pid)
	return cmd, nil
}
