package main

import (
	"os/exec"
	"time"
)

func main() {
	c1 := exec.Command("ls")
	//pipe, _ := c1.StdoutPipe()
	c1.Start()
	time.Sleep(10 * time.Second)
	c1.Wait()
}
