package main

import (
    "fmt"
    "log"
    "net/http"
    "net/http/cookiejar"
    "net/url"
    "github.com/PuerkitoBio/goquery"
)

func loginAndAccessForm(formURL string,  name string, password string) {
	jar, _ := cookiejar.New(nil)
	client := &http.Client{Jar: jar}

    // Step 1: Log in to the site
    loginPayload := url.Values{
        "username": {name},
		"password": {password},
		// "csrf_token": {csrfToken}, // Add CSRF token if needed
    }

    // Send POST request to the login URL
    loginResp, err := client.PostForm(formURL, loginPayload)
    if err != nil {
        log.Fatalf("Login request failed: %v", err)
    }
    defer loginResp.Body.Close()

    // Check login success
    if loginResp.StatusCode != http.StatusOK {
        log.Fatalf("Login failed with status: %d", loginResp.StatusCode)
    }
    fmt.Println("Login successful!")

    // Step 2: Access the form page after logging in
    formResp, err := client.Get(formURL)
    if err != nil {
        log.Fatalf("Failed to load form page: %v", err)
    }
    defer formResp.Body.Close()

    // Step 3: Use goquery to parse the form page
    doc, err := goquery.NewDocumentFromReader(formResp.Body)
    if err != nil {
        log.Fatalf("Failed to parse form page: %v", err)
    }

    // Step 4: Find and print form fields
    doc.Find("form").Each(func(i int, s *goquery.Selection) {
        action, _ := s.Attr("action")
        method, _ := s.Attr("method")
        fmt.Printf("Form %d:\n", i+1)
        fmt.Printf("Action: %s, Method: %s\n", action, method)

        s.Find("input").Each(func(j int, input *goquery.Selection) {
            name, _ := input.Attr("name")
            inputType, _ := input.Attr("type")
            placeholder, _ := input.Attr("placeholder")
            fmt.Printf("Field %d - Name: %s, Type: %s, Placeholder: %s\n", j+1, name, inputType, placeholder)
        })
        fmt.Println("========")
    })
}
