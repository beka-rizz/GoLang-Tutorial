package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type user struct {
	Name string `json:"name"`
	Age  string `json:"age"`
	City string `json:"city"`
}

func main() {
	http.HandleFunc("/teachers", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println(r.Method)
		switch r.Method {
		case http.MethodGet:
			w.Write([]byte("Hello GET Method on Teachers Route"))
			fmt.Println("Hello GET Method on Teachers Route")
			return
		case http.MethodPost:

			// Parse form data (necessary for x-www-form-urlencoded)
			// err := r.ParseForm()
			// if err != nil {
			// 	http.Error(w, "Error parsing form", http.StatusBadRequest)
			// 	return
			// }
			// fmt.Println("Form:", r.Form)

			// Prepare response data
			// response := make(map[string]any)
			// for key, value := range r.Form {
			// 	response[key] = value[0]
			// }
			// fmt.Println("Processed response map:", response)

			// RAW Body
			body, err := io.ReadAll(r.Body)
			if err != nil {
				return
			}
			defer r.Body.Close()

			// fmt.Println("Raw body:", body)
			fmt.Println("Raw body:", string(body))

			// If you expect json data, then unmarshal it
			var user1 user
			err = json.Unmarshal(body, &user1)
			if err != nil {
				return
			}
			fmt.Println("Unmarshalled JSON into a user struct:", user1)

			response := make(map[string]any)
			err = json.Unmarshal(body, &response)
			if err != nil {
				return
			}
			fmt.Println("Unmarshalled JSON into a map:", response)

			// Access the request details
			fmt.Println("Body:", r.Body)
			fmt.Println("Form:", r.Form)
			fmt.Println("Header:", r.Header)
			fmt.Println("Context:", r.Context())
			fmt.Println("ContentLength:", r.ContentLength)
			fmt.Println("Host:", r.Host)
			fmt.Println("Method:", r.Method)
			fmt.Println("Protocol:", r.Proto)
			fmt.Println("Remote Addr:", r.RemoteAddr)
			fmt.Println("Request URI:", r.RequestURI)
			fmt.Println("TLS:", r.TLS)
			fmt.Println("Trailer:", r.Trailer)
			fmt.Println("Transfer Encoding:", r.TransferEncoding)
			fmt.Println("URL:", r.URL)
			fmt.Println("User Agent:", r.UserAgent())
			fmt.Println("Port:", r.URL.Port())
			fmt.Println("URL Scheme:", r.URL.Scheme)

			w.Write([]byte("Hello POST Method on Teachers Route"))
			fmt.Println("Hello POST Method on Teachers Route")
			return
		}
	})
}
