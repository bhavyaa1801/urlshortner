package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received")
	fmt.Println("Method :", r.Method)
	fmt.Println("Path   :", r.URL.Path)
	fmt.Fprint(w, "URL Shortener Running 🚀")
}

func shortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }

    var request ShortenRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "invalid JSON" , http.StatusBadRequest)
	    return
	}
    fmt.Println("Original URL:", request.URL)

	shortcode := getUniqueShortCode()

    urlStore[shortcode] = request.URL
 
    fmt.Println(urlStore)


	response := ShortenResponse{
		ShortURL: shortcode,
	}
// thunderc ko json bhej re 
	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "failed to encode json" , http.StatusInternalServerError)
		return
	}
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {

    if r.URL.Path == "/" {
	fmt.Fprint(w, "URL Shortener Running 🚀")
	return
    }

	shortCode := r.URL.Path[1:]

	originalURL, exists := urlStore[shortCode]
	if !exists {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusFound)
}


func main() {
	http.HandleFunc("/", redirectHandler)
	http.HandleFunc("/shorten", shortenHandler) // map the routing

	fmt.Println("Server running on :8080")

	err := http.ListenAndServe(":8080", nil)  // start the server
	if err != nil {
		fmt.Println(err)
	}
}