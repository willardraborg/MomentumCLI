package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	switch os.Args[1] {
	case "log":
		if err := cmdLog(os.Args[2:]); err != nil{
			fmt.Fprintln(os.Stderr, "momentum:", err)
			os.Exit(1)
		}
	case "streak":
		fmt.Println("streak: not implemented yet")
	case "stats":
		fmt.Println("stats: not implemented yet")
	case "list":
		if err := cmdList(os.Args[2:]); err != nil{
			fmt.Fprintln(os.Stderr, "momentum:", err)
			os.Exit(1)
		}
	case "graph":
		fmt.Println("graph: not implemented yet")
	default:
		fmt.Printf("momentum: unknown command %q\n", os.Args[1])
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("usage: momentum <command>")
	fmt.Println()
	fmt.Println(" log: \"<note>\" <hours>   log today as shipped")
	fmt.Println(" streak: show current streak")
	fmt.Println(" stats: show total days and hours")
	fmt.Println(" list <number>: show latest <number> of logged days")
	fmt.Println(" graph: show graph of days shipped")
}
