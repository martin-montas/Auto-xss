/* @author: https://github.com/martin-montas
 * @date: 2024-07-27
 *
 * This is a simple XSS scanner written in Go.
 * It uses go routines to make it faster.
 *
 */
package main

import (
	"flag"
	"sync"
)


const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	White  = "\033[37m"
)

// TODO: test it on this website  http://192.168.56.104/dvwa/vulnerabilities/xss_s/

/* TODO: this is a way to check for xss  validation
// Check if the payload is reflected in the response body
if strings.Contains(string(body), payload) {
    fmt.Printf("Potential XSS found in %s with payload %s\n", url, payload)
}
*/

// Common XSS payloads
var payloads = []string{
	"<script>alert('XSS')</script>",
	"\"><script>alert('XSS')</script>",
	"<img src=x onerror=alert('XSS')>",
	"<svg onload=alert('XSS')>",
	"<a href=\"javascript:alert('XSS')\">Click me</a>",
	"<body onload=alert('XSS')>",
	"<iframe src=\"javascript:alert('XSS')\"></iframe>",
}

func main() {

    // interface IP
    localAddr := "192.168.56.1"
	var (
		url = flag.String("url", "http://127.0.0.1", "The URL for your target")
	)

	// Parse the flags
	flag.Parse()

	var wg sync.WaitGroup
	wg.Add(1)
	scrapeAndSendRequest(*url, &wg, localAddr)
	wg.Wait()
}
