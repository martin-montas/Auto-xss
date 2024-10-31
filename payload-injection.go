package main

import (
    "fmt"
    "net/http"
    "net/url"
    "sync"

    "github.com/PuerkitoBio/goquery"
)

func threadedQuery(actionURL *url.URL, form *goquery.Selection, wg *sync.WaitGroup) {
    /*
    * 	This function takes a form and sends a POST request
    * 	with the form data. It also uses go routines to make
    * 	it faster.
    */
    defer wg.Done()

    action, exists := form.Attr("action")
    if !exists {
        fmt.Println(Red + "[!] " + Reset + "Form action not found")
        return
    }

    actionURL, err := actionURL.Parse(action)
    fmt.Println(Blue + "[*] " + Reset + "Parsing the action URL")
    if err != nil {
        fmt.Println(Red+"[!] "+Reset+"Error parsing action URL:", err)
        return
    }

    // Create the payload based on form input fields
    formData := url.Values{}
    fmt.Println(Blue + "[*] " + Reset + "Parsing the action URL")
    form.Find("input").Each(func(index int, item *goquery.Selection) {
        name, exists := item.Attr("name")
        if exists {
            formData.Set(name, payloads[0])
        }
    })

    // Send the POST request
    resp, err := http.PostForm(actionURL.String(), formData)
    fmt.Println(Blue + "[*] " + Reset + "Sending the POST request")
    if err != nil {
        fmt.Println(Red+"[!] "+"Error sending request:", err)
        return
    }
    defer resp.Body.Close()

    // Print the response status
    fmt.Println(Green+"[+] "+Reset+"Response Status: ", resp.Status)
}
