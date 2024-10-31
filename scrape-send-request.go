package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net"
	"net/http"
	"net/url"
	"sync"
)

func scrapeAndSendRequest(pageURL string, wg *sync.WaitGroup, localAddr string) {
	/*
	 * 	This function get the vulnerable website and parses the
	 * 	html to find form tags. It also uses go routines to make
	 * 	it faster.
	 */

    dialer := &net.Dialer{
        LocalAddr: &net.TCPAddr{
            IP: net.ParseIP(localAddr),
        },
    }

    transport := &http.Transport{
        DialContext: dialer.DialContext,
    }

    client := &http.Client{
        Transport: transport,
    }

	// Fetch the page
	resp, err := client.Get(pageURL)
	fmt.Println(Blue + "[*] " + Reset + "Fetching the page")
	if err != nil {
		fmt.Println(Red+"[-] "+Reset+"Error fetching page:", err)
		wg.Done()
		return
	}
	defer resp.Body.Close()

	// Parse the HTML
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	fmt.Println(Blue + "[*] " +Reset+"Parsing the html")
	if err != nil {
		fmt.Println(Red+"[-] "+Reset+"Error parsing page:", err)
		wg.Done()
		return
	}

	// Construct the full action URL
	actionURL, err := url.Parse(pageURL)
	fmt.Println(Blue + "[*] " + Reset + "Constructing the action URL")
	if err != nil {
		fmt.Println(Red+"[-] "+Reset+"Error parsing page:", err)
		wg.Done()
		return
	}
	// Find all allowed form elements and process them
	doc.Find("form").Each(func(index int, item *goquery.Selection) {
		wg.Add(1)
		go threadedQuery(actionURL, item, wg)
	})
}
