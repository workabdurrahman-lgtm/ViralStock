package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type searchRequest struct {
	Keyword string `json:"keyword"`
}

type videoResult struct {
	ThumbnailURL string `json:"thumbnailUrl"`
	VideoURL     string `json:"videoUrl"`
}

type downloadRequest struct {
	VideoURL string `json:"videoUrl"`
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var req downloadRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// In a real scenario, the scraper would provide a direct .mp4 link.
	// For now, this handler assumes req.VideoURL is a direct link.
	// We will simulate a download by fetching a placeholder file.
	// A real video URL would be used once the scraper is fully implemented.
	placeholderVideoURL := "http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4"

	resp, err := http.Get(placeholderVideoURL)
	if err != nil {
		http.Error(w, "Failed to fetch video", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Disposition", "attachment; filename=\"video.mp4\"")
	w.Header().Set("Content-Type", "video/mp4")

	// Stream the video content to the client
	if _, err := io.Copy(w, resp.Body); err != nil {
		http.Error(w, "Failed to stream video", http.StatusInternalServerError)
	}
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var req searchRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// For now, we'll return dummy data.
	// The actual scraping logic will be implemented in the next steps.
	dummyResults := []videoResult{
		{ThumbnailURL: "https://p16-sign-va.tiktokcdn.com/tos-maliva-avt-0068/e19b550186214873946a4a3511013146~c5_720x720.jpeg?lk3s=a5d48078&x-expires=1715983200&x-signature=J%2B%2B5H7%2B2hZ2Xg3w3b7B9Q6M%2F2w%3D", VideoURL: "https://www.tiktok.com/@example/video/123"},
		{ThumbnailURL: "https://p16-sign-va.tiktokcdn.com/tos-maliva-avt-0068/e19b550186214873946a4a3511013146~c5_720x720.jpeg?lk3s=a5d48078&x-expires=1715983200&x-signature=J%2B%2B5H7%2B2hZ2Xg3w3b7B9Q6M%2F2w%3D", VideoURL: "https://www.tiktok.com/@example/video/456"},
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(dummyResults); err != nil {
		http.Error(w, "Failed to encode results", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/api/search", searchHandler)
	http.HandleFunc("/api/download", downloadHandler)

	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
