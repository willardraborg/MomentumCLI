package main

import (
	"fmt"
	"strconv"
	"time"
)

func cmdLog(args []string) error{
	if len(args) != 2{
		return fmt.Errorf("usage: momentum log \"<note>\" <hours>")
	}
	note := args[0]

	hours, err := strconv.ParseFloat(args[1], 64)
	if err != nil{
		return fmt.Errorf("hours must be a number, got %q", args[1])
	}
	if hours <= 0{
		return fmt.Errorf("hours must be greater than zero")
	}

	entries, err := loadEntries()
	if err != nil{
		return err
	}

	today := time.Now().Format("2006-01-02")

	for i := range entries{
		if entries[i].Date == today{
			entries[i].Hours += hours
			entries[i].Note = note
			if err := saveEntries(entries); err != nil{
				return err
			}
			fmt.Printf("updated %s: %.1f hours total\n", today, entries[i].Hours)
			return nil
		}
	}

	entries = append(entries, Entry{Date: today, Note: note, Hours: hours})
	if err := saveEntries(entries); err != nil{
		return err
	}
	fmt.Printf("logged %s: %q (%.1f hours)\n", today, note, hours)
	return nil
}

func cmdList(args []string) error{
	n := 10
	if len(args) == 1{
		parsed, err := strconv.Atoi(args[0])
		if err != nil || parsed < 1{
			return fmt.Errorf("list: number must be a positive integer, got %q", args[0])
		}
		n = parsed
	}

	entries, err := loadEntries()
	if err != nil{
		return err
	}
	if len(entries) == 0{
		fmt.Println("no entries yet, log your first day with: momentum log \"<note>\" <hours>")
		return nil
	}

	start := len(entries) - n
	if start < 0 {
		start = 0
	}
	for _, e := range entries[start:] {
		fmt.Printf("%s %4.1f h %s\n", e.Date, e.Hours, e.Note)
	}
	return nil
}


