package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var verbose bool

var ext string = ".was"

func init() {
	flag.BoolVar(&verbose, "v", false, "verbose output")
}

func main() {

	flag.Usage = func() {

		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, `

	Examples:

was filename1 [filename2 filename3 ...]

Move list of files to files with a .was extension, and move them back if they already have a .was extension.

WIP

Make it return non-zero if there were any errors
Let user choose the extension


`)

		flag.PrintDefaults()
	}

	flag.Parse()

	wasFiles := flag.Args()

	if len(wasFiles) < 1 {

		flag.Usage()
		os.Exit(2)

	}

	flag.Parse()

	if verbose {
		fmt.Println("hello world:%v:%s:", verbose, wasFiles)
	}

	for _, file := range wasFiles {

		if verbose {
			fmt.Fprintf(os.Stderr, "handling file:%s:len(file):%d:\n", file, len(file))
		}

                //chop off slash from directories
                if file[len(file) - 1] == "/"[0] {
                  file = file[0:len(file) - 1]
                }

		if file == ext {
			fmt.Fprintf(os.Stderr, "ignoring .was:%v\n")
			continue
		}

		if _, err := os.Stat(file); err != nil {
			fmt.Fprintf(os.Stderr, "skipping:%v\n", err)
			continue
		}

		targetFile := file + ext
		if strings.HasSuffix(file, ext) {

			if verbose {
				fmt.Fprintf(os.Stderr, "doing unwas on:%s\n", targetFile)
			}

			targetFile = file[0 : len(file)-len(ext)]

		}

		if _, err := os.Stat(targetFile); err == nil {

			if verbose {
				fmt.Fprintf(os.Stderr, "target is blocked:%s\n", targetFile)
			}

			if err := os.Remove(targetFile); err != nil {
				fmt.Fprintf(os.Stderr, "could not clear the way for new was file:skipping:%v\n", err)
				continue
			}
		}

		if verbose {
			fmt.Fprintf(os.Stderr, "target is clear:%s\n", file)
		}

		if err := os.Rename(file, targetFile); err != nil {
			fmt.Fprintf(os.Stderr, "failed to was:%v\n", err)
			continue
		}

		if verbose {
			fmt.Fprintf(os.Stderr, "was'd:%s\n", file)
		}

	}

}
