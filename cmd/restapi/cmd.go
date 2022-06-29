package restapi

import (
	"net/http"

	"github.com/spf13/cobra"
	"github.com/vuvuzela/loadtest/internal/loadtestrest"
	"github.com/vuvuzela/loadtest/pkg/httpx"
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
		httpClient := &http.Client{}
		httpxClient := httpx.NewClient(httpClient)
		svc := loadtestrest.NewService(httpxClient)

		in := loadtestrest.LoadTestRestAPIInput{
			NumOfRequests:     numOfRequests,
			ConcurentRequests: concurentRequests,
			RequestMethod:     reqMethod,
			Endpoint:          endpoint,
			RequestBody:       reqBody,
		}
		svc.LoadTestRestAPI(in)
	},
}

func init() {
	LoadTestCmd.Flags().StringVarP(&endpoint, "endpoint", "e", "https://www.example.com", "endpoint to hit")
	LoadTestCmd.Flags().StringVarP(&reqMethod, "method", "m", "GET", "HTTP request method")
	LoadTestCmd.Flags().StringVarP(&reqBody, "body", "b", "", "request body (if applies)")

	LoadTestCmd.Flags().IntVarP(&concurentRequests, "number-of-workers", "w", 1, "number of concurrent workers for requests")
	LoadTestCmd.Flags().IntVarP(&numOfRequests, "num-of-requests", "n", 1, "overall number of requests")
}
