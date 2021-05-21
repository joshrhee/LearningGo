package main

import (
	"errors"
	"fmt"
	"net/http"
)

type requestResult struct {
	url    string
	status string
}

var errRequestFailed = errors.New("Request failed")

func main() {

	results := make(map[string]string)
	c := make(chan requestResult)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}
	for _, url := range urls {
		go hitURL(url, c)
	}

	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}
}

func hitURL(url string, c chan<- requestResult) {
	resp, err := http.Get(url)
	//If status code >= 400, this means there is some problem
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- requestResult{url: url, status: status}
}

// //Optimal way using go routine channel
// func main() {
// 	//Creating a channel to recieve data from go routine
// 	c := make(chan string)
// 	people := [3]string{"Josh", "Chloe", "Jake"}
// 	for _, person := range people {
// 		//Using go routine to call functions concurrently
// 		go myChannel(person, c)
// 	}

// 	//If we also put go in-front-of this, go would terminate program without looping
// 	//BC go routine is only working when at least one function is working
// 	//If we make this one also go routine, there is no normal function, so it will be terminated.

// 	//printing channel's data
// 	//<-c means getting a message from channel
// 	//<-c is a blocking operation, which stops program until it receives a message and continue
// 	for i := 0; i < len(people); i++ {
// 		fmt.Println("Waiting for ", i)
// 		fmt.Println(<-c)
// 	}
// }

// func myChannel(person string, c chan string) {
// 	//Sending a message to channel
// 	c <- person + " is sexy"
// }
