package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

// call multiple endpoints concurrently and report response and error
// example input : ["https://go.dev", "https://google.com", "https://badhost"]

type Result struct {
	URL   string
	Size  int
	Error error
}

func fetchURLResponse(url string, wg *sync.WaitGroup, results chan<- Result) {
	defer wg.Done()

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error", err)
		/*
			err = errors.New("BAD_REQUEST")

			err = fmt.Errorf("Error While calling HTTP Get - %s, Status Code - %d", err.Error(),400)
		*/

		results <- Result{URL: url, Error: err}
		return
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error", err)
		results <- Result{URL: url, Error: err}
		return
	}
	results <- Result{URL: url, Size: len(body)}
}
func fetchAllURLResponse(urls []string) []Result {

	var wg sync.WaitGroup

	results := make(chan Result, len(urls))

	for _, url := range urls {
		wg.Add(1)
		go fetchURLResponse(url, &wg, results)
	}

	wg.Wait()
	close(results)

	var out []Result
	for result := range results {
		out = append(out, result)
	}
	return out

}

func main() {
	urls := []string{"https://go.dev", "https://google.com", "https://badhost"}

	results := fetchAllURLResponse(urls)

	fmt.Printf("%+v", results)

	var u USER
	fmt.Println(u.Password)

	// while return the response
	u.Password = ""
	// Post

	// JSON(User)

}

type USER struct {
	Name     string `json:"Name,required"`
	Age      string `json:"Age"`
	Password string `json:"password,omitempty"`
}
