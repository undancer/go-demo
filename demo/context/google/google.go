// Package google provides a function to do Google searches using the Google Web
// Search API. See https://developers.google.com/web-search/docs/
//
// This package is an example to accompany https://blog.golang.org/context.
// It is not intended for use by others.
//
// Google has since disabled its search API,
// and so this package is no longer useful.
package google

import (
	"context"
	"encoding/json"
	"github.com/undancer/go-demo/demo/context/query"
	"log"
	"net/http"
	"time"
)

// Results is an ordered list of search results.
type Results []Result

// A Result contains the title and URL of a search result.
type Result struct {
	Title, SubTitle string
}

// Search sends query to Google search and returns the results.
func Search(ctx *query.QueryCtx) (Results, error) {
	// Prepare the Google Search API request.
	req, err := http.NewRequest("GET", "http://localhost:9000/context_demo", nil)
	if err != nil {
		return nil, err
	}

	ctx.SetReq(req)
	// Issue the HTTP request and handle the response. The httpDo function
	// cancels the request if ctx.Done is closed.
	var results Results
	err = httpDo(ctx, req, func(resp *http.Response, err error) error {
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		// Parse the JSON search result.
		// https://developers.google.com/web-search/docs/#fonje
		var data struct {
			ResponseData struct {
				Results []struct {
					Title, SubTitle string
				}
			}
		}
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			return err
		}
		for _, res := range data.ResponseData.Results {
			results = append(results, Result{Title: res.Title, SubTitle: res.SubTitle})
		}
		return nil
	})
	// httpDo waits for the closure we provided to return, so it's safe to
	// read results here.
	return results, err
}

// httpDo issues the HTTP request and calls f with the response. If ctx.Done is
// closed while the request or f is running, httpDo cancels the request, waits
// for f to exit, and returns ctx.Err. Otherwise, httpDo returns f's error.
func httpDo(ctx *query.QueryCtx, req *http.Request, f func(*http.Response, error) error) error {
	// Run the HTTP request in a goroutine and pass the response to f.
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	// WithCancel会在ctx的children中增加cancelDb，这样当
	// ctx 结束的时候，cancelDb也会受到消息
	cancelDb, cancel := context.WithCancel(ctx.Context)
	defer cancel()
	c := make(chan error, 1)
	go func() { c <- f(client.Do(req)) }()
	go func(ctx context.Context) {
		t := time.NewTimer(2 * time.Second)

		select {
		case <-t.C:
			log.Println("db access finished!")
		case <-ctx.Done():
			log.Println("canceld by parent, release resource")
		}
	}(cancelDb)
	select {
	case <-ctx.Done():
		tr.CancelRequest(req)
		<-c // Wait for f to return.
		return ctx.Err()
	case err := <-c:
		return err
	}
}
