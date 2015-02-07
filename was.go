package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var verbose bool
var force bool
var copy bool

var ext string = ".was"

var errors bool = false

func init() {
	flag.BoolVar(&copy, "c", false, "copy instead of move")
	flag.StringVar(&ext, "e", ext, "file extension")
	flag.BoolVar(&force, "f", false, "clobber any conflicting files")
	flag.BoolVar(&verbose, "v", false, "verbose output")
}

func main() {
	flag.Usage = usage

	flag.Parse()

	wasFiles := flag.Args()

	if len(wasFiles) < 1 {
		wasFiles = filesFromStdin()
	}

	if len(wasFiles) < 1 {
		flag.Usage()
		os.Exit(2)
	}

	if !strings.HasPrefix(ext, ".") {
		ext = "." + ext
	}

	if verbose {
		fmt.Println("hello world:%v:%s:", verbose, wasFiles)
	}

FileLoop:
	for _, file := range wasFiles {
		if file == "" {
			if verbose {
				fmt.Fprintf(os.Stderr, "ignoring empty file\n")
			}
			continue FileLoop
		}

		if verbose {
			fmt.Fprintf(os.Stderr, "handling file:%s:len(file):%d:\n", file, len(file))
		}

		//chop off slash from directories
		file = filepath.Clean(file)

		if file == ext {
			fmt.Fprintf(os.Stderr, "ignoring %s:%v\n", ext)
			continue FileLoop
		}

		if _, err := os.Stat(file); err != nil {
			fmt.Fprintf(os.Stderr, "skipping:%v\n", err)
			continue FileLoop
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
			if !force {
				fmt.Printf("There's a file in the way:%s:\n", targetFile)
				fmt.Printf("Delete %s? Please type yes or no and then press enter:\n", targetFile)
				if askForConfirmation() {
					if err := os.RemoveAll(targetFile); err != nil {
						fmt.Fprintf(os.Stderr, "could not clear the way for new was file:skipping:%v\n", err)
						errors = true
						continue FileLoop
					}
				} else {
					fmt.Fprintf(os.Stderr, "user chose to not delete target:skipping:%s\n", targetFile)
					continue FileLoop
				}
			}
		}

		if verbose {
			fmt.Fprintf(os.Stderr, "target is clear:%s\n", file)
		}

		if copy {
			copyFileHandle, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "skipping:%v\n", err)
				errors = true
				continue FileLoop
			}
			defer copyFileHandle.Close()

			finfo, err := copyFileHandle.Stat()
			if err != nil {
				fmt.Fprintf(os.Stderr, "skipping:%v\n", err)
				errors = true
				continue FileLoop
			}

			if fmode := finfo.Mode(); fmode.IsDir() {
				fmt.Fprintf(os.Stderr, "skipping:copy is not supported for directories\n")
				errors = true
				continue FileLoop
			}

			targetFileHandle, err := os.Create(targetFile)
			if err != nil {
				fmt.Fprintf(os.Stderr, "skipping:%v\n", err)
				errors = true
				continue FileLoop
			}
			defer targetFileHandle.Close()

			_, err = io.Copy(targetFileHandle, copyFileHandle)
			if err != nil {
				fmt.Fprintf(os.Stderr, "skipping:%v\n", err)
				errors = true
				continue FileLoop
			}
		} else {
			if err := os.Rename(file, targetFile); err != nil {
				fmt.Fprintf(os.Stderr, "failed to was:%v\n", err)
				errors = true
				continue FileLoop
			}
		}

		if verbose {
			fmt.Fprintf(os.Stderr, "was'd:%s\n", file)
		}
	}
	if errors {
		os.Exit(1)
	}
}

//swiped this from a gist:
//https://gist.github.com/albrow/5882501
func askForConfirmation() bool {
	consolereader := bufio.NewReader(os.Stdin)

	response, err := consolereader.ReadString('\n')
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(2)
	}
	okayResponses := []string{"y", "Y", "yes", "Yes", "YES"}
	nokayResponses := []string{"n", "N", "no", "No", "NO"}
	if containsString(okayResponses, response[:len(response)-1]) {
		return true
	} else if containsString(nokayResponses, response[:len(response)-1]) {
		return false
	} else {
		fmt.Println("Please type yes or no and then press enter:")
		return askForConfirmation()
	}
}

// posString returns the first index of element in slice.
// If slice does not contain element, returns -1.
func posString(slice []string, element string) int {
	for index, elem := range slice {
		if elem == element {
			return index
		}
	}
	return -1
}

// containsString returns true iff slice contains element
func containsString(slice []string, element string) bool {
	return !(posString(slice, element) == -1)
}

func filesFromStdin() []string {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(2)
	}

	return strings.Split(string(bytes), "\n")
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, `

        Description:

Stupid simple but useful tool to move a file or directory and move it back later.
Was moves a list of files to files with a .was extension, and/or moves them back if they already have a .was extension.

	Examples:

was thisFile -> thisFile.was
was thisFile.was -> thisFile
was thisFile thatFile.was -> thisFile.was thatFile
was -c someFile -> someFile someFile.was
was -e=saw someFile -> someFile.saw
ls -1 | was -> file1.was file2.was file3.was ...

was filename1 [filename2 filename3 ...]

WIP
`)

	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "\n")
}
