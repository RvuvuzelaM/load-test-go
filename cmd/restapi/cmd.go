package restapi

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

var (
	endpoint  string
	reqMethod string
	reqBody   string

	numOfRequests     int
	concurentRequests int
)

var LoadTestCmd = &cobra.Command{
	Use:   "rest",
	Short: "Allows one to test REST API capabilites by bombarding with requests.",
	Long: `
			Allows one to test REST API capabilites by bombarding with requests.
			What is recomended is to first make test request to your API, and then load test it.`,
	Run: func(cmd *cobra.Command, args []string) {
		jobs := make(chan int, numOfRequests)
		results := make(chan string, numOfRequests)

		for i := 1; i <= concurentRequests; i++ {
			go worker(i, jobs, results)
		}

		for i := 1; i <= numOfRequests; i++ {
			jobs <- i
		}
		close(jobs)

		for i := 1; i <= numOfRequests; i++ {
			_, ok := <-results
			if ok {
				if i%100 == 0 || i == 1 {
					fmt.Println(i)
				}
			}
		}
		close(results)
	},
}

func init() {
	LoadTestCmd.Flags().StringVarP(&endpoint, "endpoint", "e", "https://www.example.com", "endpoint to hit")
	LoadTestCmd.Flags().StringVarP(&reqMethod, "method", "m", "GET", "HTTP request method")
	LoadTestCmd.Flags().StringVarP(&reqBody, "body", "b", "", "request body (if applies)")

	LoadTestCmd.Flags().IntVarP(&concurentRequests, "number-of-workers", "w", 1, "number of concurrent workers for requests")
	LoadTestCmd.Flags().IntVarP(&numOfRequests, "num-of-requests", "n", 1, "overall number of requests")
}

func worker(id int, jobs <-chan int, results chan<- string) {
	c := &http.Client{}

	for range jobs {
		var body io.Reader
		if reqBody != "" {
			body = strings.NewReader(reqBody)
		}

		req, err := http.NewRequest(reqMethod, endpoint, body)
		if err != nil {
			fmt.Println(err)
			return
		}

		resp, err := c.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}

		respByte, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		results <- string(respByte)
	}
}
