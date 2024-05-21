package main

import (
	"flag"
	"fmt"
	"os"
)

func StartRepl() {
	var funcFlag bool
	flag.BoolVar(&funcFlag, "c", false, Flags()["c"].description)
	flag.BoolVar(&funcFlag, "l", false, Flags()["l"].description)
	flag.BoolVar(&funcFlag, "w", false, Flags()["w"].description)
	flag.BoolVar(&funcFlag, "m", false, Flags()["m"].description)

	flag.Usage = func() {
		fmt.Printf("Correct Usage: ./main [flag] [file-name]\n")
		fmt.Printf("Flags: \n")
		flag.PrintDefaults()
	}

	if len(os.Args) != 3 {
		flag.Usage()
		os.Exit(0)
	}

	flag.Parse()

	args := flag.Args()
	flagName := string(string(os.Args[1])[1])

	if funcFlag {
		_, err := Flags()[flagName].callback(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
	}
}
