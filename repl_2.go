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
	flag.BoolVar(&funcFlag, "", false, "bytes words lines")

	flag.Usage = func() {
		fmt.Printf("Correct Usage: ./main [flag] [file-name]\n")
		fmt.Printf("Flags: \n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if len(os.Args) == 2 {
		if (os.Args[0]) != "./main" {
			flag.Usage()
			os.Exit(0)
		}
		defaultFlag(os.Args[1])
	} else if len(os.Args) != 3 {
		flag.Usage()
		os.Exit(0)
	}

	args := flag.Args()
	flagName := string(string(os.Args[1])[1])

	if funcFlag {
		val, err := Flags()[flagName].callback(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		fmt.Printf("	%v %v\n", val, args[0])
	}
}
