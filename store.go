package main 

import
(
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Entry struct{
	Date string `json:"date"`
	Note string `json:"note"`
	Hours float64 `json:"hours"`
}

func dataPath() (string, error){
	home, err := os.UserHomeDir()
	if err != nil{
		return "", fmt.Errorf("cannot find home directory: %w", err)
	}
	return filepath.Join(home, ".momentum.json"), nil
}

func loadEntries() ([] Entry, error){
	path, err := dataPath()
	if err != nil{
		return nil, err
	}
	data, err := os.ReadFile(path)
	if os.IsNotExist(err){
		return nil, nil
	}
	if err != nil{
		return nil, fmt.Errorf("cannot read %s: %w", path, err)
	}
	var entries []Entry
	if err := json.Unmarshal(data, &entries); err != nil{
		return nil, fmt.Errorf("cannot parse %s: %w", path, err)
	}
	return entries, nil
}

func saveEntries(entries []Entry) error {
	path, err := dataPath()
	if err != nil{
		return err
	}
	data, err := json.MarshalIndent(entries, "", "  ")
	if err != nil{
		return err
	}
	if err := os.WriteFile(path, data, 0o644); err != nil{
		return fmt.Errorf("cannot write %s: %w", path, err)
	}
	return nil
}
