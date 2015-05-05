package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type wrapper struct {
	cmd  string
	args []string
}

func (w *wrapper) setArgs(args []string) {
	w.args = make([]string, 0, len(args))

	skipNext := false
	for a := range args {

		if skipNext {
			skipNext = false
			continue
		}

		switch {
		case strings.HasPrefix(args[a], "-o"):
			b := a + 1
			if args[a] != "-o" {
				b = a
				args[a] = args[a][2:]
			}
			ext := filepath.Ext(args[b])
			switch ext {
			case ".exe":
				args[a] = "/Fe:" + strings.Trim(args[b], " ")
			case ".obj":
				args[a] = "/Fo:" + strings.Trim(args[b], " ")
			}
			skipNext = b != a
		}

		w.args = append(w.args, args[a])
	}
}

func (w wrapper) argString() string {
	return strings.Join(w.args, " ")
}

func (w wrapper) run() error {
	c := exec.Command(w.cmd, w.argString())

	err := c.Run()
	return err
}

func main() {

	numArgs := len(os.Args)

	if numArgs < 2 {
		fmt.Errorf("Wrong number of arguments")
	}

	w := wrapper{cmd: os.Args[1]}
	if numArgs > 2 {
		w.setArgs(os.Args[2:])
	}

	err := w.run()
	if err != nil {
		fmt.Println(err)
	}

}
