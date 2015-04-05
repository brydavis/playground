package play

import "os/exec"

func FormatGo(filepath string) {
	exec.Command("go", "fmt", filepath).Run()
	exec.Command("goimports", "-w", filepath).Run()
}
