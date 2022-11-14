package main

import (
	"fmt"
	"net/http"
	"time"
)


func main()  {

	links := []string {
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://amazon.com",
		"https://golang.org",
		"http://invalid_link.com", 
	}

	channel := make(chan string) // creating a new channel


	
	for _, link := range links {
		go checkLink(link, channel) 
		// run this function insde a new routine - create a new thread
		// create a child routine

		// since main routine ends before the for loop, ignores the for loop. (without channel)
		// main doesnt care about his children 
	}

//	fmt.Println(<- channel) // awaits message - blocking channel

	for link := range channel { // run though for loop every single time a channel receives a message
		// fmt.Println(<-channel)

		// go checkLink(link, channel)

		// function literal -> Lambda
		go func(l string) {
			time.Sleep(time.Second * 1) // pauses the current go routine (5 seconds)
			checkLink(l, channel)
		}(link) // extra () to invoke it - passing link to literal function
	}
}

func checkLink(link string, channel chan string) {
	_, err := http.Get(link) // blocking function

	if err != nil {
		fmt.Println(link, "might be down!")
		// channel <- "Might be down I think"
		channel <- link
	} else {
		fmt.Println(link, "is up!")
		// channel <- "Yep, it's up"
		channel <- link
	}

}