package main

import (
    "bytes"
    "os"
    "os/exec"
)

func main() {
    cmd := exec.Command("cat")
    cmd.Stdin = bytes.NewBuffer([]byte("hi"))
    cmd.Stdout = os.Stdout

    cmd.Run()

}
