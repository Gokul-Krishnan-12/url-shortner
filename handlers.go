package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/Gokul-Krishnan-12/go-url-shortner/utils"
)

func shortenUrl(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var body struct {
		URL        string `json:"url"`
		CustomCode string `json:"custom_code,omitempty"`
		ExpiryMins int    `json:"expiry_mins,omitempty"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.URL == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if !utils.IsValidUrl(body.URL) {
		http.Error(w, "Invalid URL. Must be http or https", http.StatusBadRequest)
		return
	}

	code := body.CustomCode
	if code == "" {
		code = utils.GenerateRandomString(7)
	} else {
		if _, exists := urlStore[code]; exists {
			http.Error(w, "Custom code already in use", http.StatusConflict)
			return
		}
	}

	expiry := time.Now().Add(24 * time.Hour)
	if body.ExpiryMins > 0 {
		expiry = time.Now().Add(time.Duration(body.ExpiryMins) * time.Minute)
	}

	urlStore[code] = URLData{
		OriginalURL: body.URL,
		Expiry:      expiry,
		HitCount:    0,
	}

	if err := saveToFile(urlStore); err != nil {
		log.Println("⚠️ Failed to save:", err)
	}

	response := map[string]string{
		"short_url": fmt.Sprintf("http://localhost:9000/%s", code),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func redirectUrl(w http.ResponseWriter, r *http.Request) {
	code := strings.TrimPrefix(r.URL.Path, "/")
	data, exists := urlStore[code]

	if !exists {
		http.NotFound(w, r)
		return
	}

	if time.Now().After(data.Expiry) {
		delete(urlStore, code)
		http.Error(w, "Link has expired", http.StatusGone)
		return
	}

	data.HitCount++
	urlStore[code] = data

	http.Redirect(w, r, data.OriginalURL, http.StatusFound)
}

func statsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	code := strings.TrimPrefix(r.URL.Path, "/stats/")
	data, exists := urlStore[code]

	if !exists {
		http.NotFound(w, r)
		return
	}

	response := map[string]interface{}{
		"original_url": data.OriginalURL,
		"hit_count":    data.HitCount,
		"expires_at":   data.Expiry.Format(time.RFC3339),
		"expired":      time.Now().After(data.Expiry),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
