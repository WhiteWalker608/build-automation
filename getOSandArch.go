package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func main() {
	out, err := exec.Command("echo", "\"GOOS="+runtime.GOOS+"\"", ">>", "$GITHUB_ENV").Output()
	if err != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Println("out: ", string(out))
	}
	out, err = exec.Command("echo", "\"GOARCH="+runtime.GOARCH+"\"", ">>", "$GITHUB_ENV").Output()
	if err != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Println("out: ", string(out))
	}
}
