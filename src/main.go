package main


import (
	. "utilities"
	. "server"
	. "configuration"
	"os/exec"
	"bufio"
	"io"
	"os"
)


func main() {
	println("♥")
	go runWebpack(true)
	NewApplication(EnvironmentSettings[DEVELOPMENT], ApplicationRouter).Start()
}


func runWebpack(suppress bool) {
	cmd := exec.Command("webpack", "--watch")
	cmd.Dir = Root + "/src/clients"

	if !suppress {
		stdoutPipe, _ := cmd.StdoutPipe()
		output := bufio.NewReader(stdoutPipe)

		cmd.Start()

		var buf = make([]byte, 2048, 2048)
		count, err := output.Read(buf)

		println("oh shit wadup")
		for count > 0 && err != io.EOF {
			println("got buf...")
			os.Stdout.Write(buf)
			count, err = output.Read(buf)
		}
	} else {
		cmd.Run()
	}

}