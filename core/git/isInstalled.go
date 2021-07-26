package git

import "os/exec"

func IsInstalled(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}
