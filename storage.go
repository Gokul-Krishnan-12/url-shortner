package main

import (
	"encoding/json"
	"os"
	"time"
)

const storageFile = "urls.json"

type URLData struct {
	OriginalURL string    `json:"original_url"`
	Expiry      time.Time `json:"expiry"`
	HitCount    int       `json:"hit_count"`
}

var urlStore = make(map[string]URLData)

func saveToFile(store map[string]URLData) error {
	data, err := json.MarshalIndent(store, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(storageFile, data, 0644)
}

func loadFromFile() (map[string]URLData, error) {
	store := make(map[string]URLData)

	file, err := os.ReadFile(storageFile)
	if err != nil {
		if os.IsNotExist(err) {
			return store, nil
		}
		return nil, err
	}

	err = json.Unmarshal(file, &store)
	if err != nil {
		return nil, err
	}

	return store, nil
}
