package main

import (
	"context"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

func main() {

	// Create a new request
	req, err := http.NewRequest("GET", "https://www.ardanlabs.com/blog/index.xml", nil)

	if err != nil {
		log.Println("Error while connecting link", err)
		return
	}

	// Create a context with a timeout of 50 milli seconds
	// The below code "times-out' at 50 ms and cancels task
	ctx, cancel := context.WithTimeout(req.Context(), 5000*time.Millisecond)
	defer cancel()

	// Declare a new Transport and Client for the call
	tr := http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	client := http.Client{
		Transport: &tr,
	}

	// Make the http call in a separate go routine so that it can be cancelled
	// We use 'ch' channel of type error as a synchronization mechanism here from main to goroutine
	ch := make(chan error, 1)

	go func() {
		log.Println("Starting the URL invocation")

		resp, err := client.Do(req)
		if err != nil {
			ch <- err
			return
		}

		// Close the response body on return
		defer resp.Body.Close()

		// Write the response to stdout
		io.Copy(os.Stdout, resp.Body)

		// The below 'nil' indicates that no error occured while processing.
		// Also, since the 'select' was waiting on 'ch' channel to indicate completion, we must send this indicator on 'ch' channel.
		ch <- nil
	}()

	// Wait for the request or timeout
	select {
	case <-ctx.Done():
		log.Println("Request timed out. Cancel work...")

		// Cancel the request and wait for it to complete
		tr.CancelRequest(req)
		log.Println(<-ch)

	case err := <-ch:
		if err != nil {
			log.Println(err)
		} else {
			log.Println("URL processed ...!!!!!")
		}
	}
}
