package play

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func AdminTerminal(port string) {
	fmt.Println("Listening @ " + port)
	exec.Command("open", "http://localhost"+port).Run()

	exit := false
	for !exit {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("$ ")
		scanner.Scan()
		input := scanner.Text()
		if input == "shutdown" {
			// fmt.Println("Shutdown server?")
			os.Exit(0)
		}
		// RECEIVE INPUT HERE
		// WRITE TO AND COMPILE
		fmt.Println(input)
	}
}
