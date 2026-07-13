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
