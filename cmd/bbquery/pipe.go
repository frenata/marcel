package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func pipeLess(s string) error {
	r, w := io.Pipe()

	go func() {
		fmt.Fprint(w, s)
		w.Close()
	}()

	less := exec.Command("less")
	less.Stdin = r
	less.Stdout = os.Stdout
	less.Start()
	return less.Wait()
}
