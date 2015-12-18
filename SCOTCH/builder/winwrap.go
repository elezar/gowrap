// winwrap is a tool that acts as a compiler driver for icl and lib on Windows platforms.
// This maps the Linux-like command line arguments for output file specification to
// the windows equivalents.
//   For example:
//     icl -o outfilename.exe => /Fe:outfilename
//     icl -o outfilename.obj => /Fo:outfilename
//     lib outfilename.lib => lib /out:outfilename.lib

package main

import (
	"bytes"
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

func findLibDir(indirs []string, fullname string) string {
	for l := len(indirs) - 1; l >= 0; l-- {
		f := fmt.Sprintf("%s/%s", indirs[l], fullname)
		if _, err := os.Stat(f); err == nil {
			return indirs[l]
		}
	}

	return "."
}

func (w *wrapper) setArgs(args []string) {
	w.args = make([]string, 0, len(args))

	skipNext := false
	var libdirs []string

	libdirs = make([]string, 0)

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

		case strings.HasPrefix(args[a], "-L"):
			fmt.Println("Adding", args[a])
			libdirs = append(libdirs, args[a][2:])
			continue

		case strings.HasPrefix(args[a], "-l"):
			fullname := fmt.Sprintf("lib%s%s", args[a][2:], ".lib")
			indir := findLibDir(libdirs, fullname)
			args[a] = fmt.Sprintf("%s/%s", indir, fullname)

		case w.cmd == "lib":
			if strings.HasSuffix(args[a], ".lib") {
				args[a] = "/out:" + strings.Trim(args[a], " ")
			}

		}

		w.args = append(w.args, fmt.Sprint(args[a]))
	}
}

func (w wrapper) argString() string {

	return strings.Join(w.args, " ")
}

func (w wrapper) run() error {
	fmt.Println("Executing command: ", w.cmd, " ", w.argString())

	c := exec.Command(w.cmd, w.args...)
	var out bytes.Buffer
	var errOut bytes.Buffer
	c.Stdout = &out
	c.Stderr = &errOut

	err := c.Run()
	if err != nil {
		fmt.Println(errOut.String())
	}
	fmt.Println(out.String())

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
