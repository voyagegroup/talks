package main

import "fmt"
import "net/http"

func main() {
	// START OMIT
	resp, err := http.Get("https://example.com")
	// ...
	resp, err := http.Post("https://example.com")
	// ...
	req, err := http.NewRequest("GET", "https://example.com", nil)
	// ...
	req.Header.Add("X-Treasure", "🍺")
	resp, err := http.DefaultClient.Do(req)
	// END OMIT
}
