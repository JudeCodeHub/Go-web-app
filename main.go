package main

import (
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	// Render the home html page from static folder
	http.ServeFile(w, r, "static/home.html")
}

func coursePage(w http.ResponseWriter, r *http.Request) {
	// Render the course html page
	http.ServeFile(w, r, "static/courses.html")
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	// Render the about html page
	http.ServeFile(w, r, "static/about.html")
}

func contactPage(w http.ResponseWriter, r *http.Request) {
	// Render the contact html page
	http.ServeFile(w, r, "static/contact.html")
}

func main() {
	// Handle the root path (loading http://localhost:8080/)
	http.HandleFunc("/", homePage)

	// Handle specific page requests matching the HTML href links
	http.HandleFunc("/home.html", homePage)
	http.HandleFunc("/courses.html", coursePage)
	http.HandleFunc("/about.html", aboutPage)
	http.HandleFunc("/contact.html", contactPage)

	log.Println("Server starting on port 8080...")
	
	// Listen on port 8080
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}