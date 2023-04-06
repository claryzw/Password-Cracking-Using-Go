/ Password Cracking with Go by Clarence Itai Msindo
package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"
)

var (
	wordlistPath string
	charset      string
	maxLength    int
	threads      int
	hashesPath   string
)

func init() {
	flag.StringVar(&wordlistPath, "w", "", "path to wordlist file")
	flag.StringVar(&charset, "c", "abcdefghijklmnopqrstuvwxyz0123456789", "custom character set for brute-force attacks")
	flag.IntVar(&maxLength, "l", 8, "maximum length of passwords to generate for brute-force attacks")
	flag.IntVar(&threads, "t", runtime.NumCPU(), "number of threads to use for processing")
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] hashfile\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

	hashesPath = flag.Arg(0)
}
