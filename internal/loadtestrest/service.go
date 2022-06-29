package loadtestrest

import (
	"fmt"
)

//go:generate go run -mod=mod github.com/golang/mock/mockgen --source=./service.go --destination=./mock/httpclient.go github.com/vuvuzela/loadtest/internal/loadtestrest/httpclient HTTPClient
type HTTPClient interface {
	MakeRequest(reqMethod, endpoint, reqBody string) (string, error)
}

type Service struct {
	httpClient HTTPClient
}

func NewService(httpClient HTTPClient) Service {
	return Service{
		httpClient: httpClient,
	}
}

type LoadTestRestAPIInput struct {
	NumOfRequests, ConcurentRequests     int
	RequestMethod, Endpoint, RequestBody string
}

func (in LoadTestRestAPIInput) Validate() error {
	return nil
}

func (s Service) LoadTestRestAPI(in LoadTestRestAPIInput) ([]string, error) {
	jobs := make(chan int, in.NumOfRequests)
	results := make(chan string, in.NumOfRequests)

	for i := 1; i <= in.ConcurentRequests; i++ {
		go s.makeRequests(in.RequestMethod, in.Endpoint, in.RequestBody, jobs, results)
	}

	for i := 1; i <= in.NumOfRequests; i++ {
		jobs <- i
	}
	close(jobs)

	resp := []string{}
	for i := 1; i <= in.NumOfRequests; i++ {
		x, ok := <-results
		if ok {
			if i%100 == 0 {
				fmt.Println(i)
			}
		}
		resp = append(resp, x)
	}
	close(results)

	return resp, nil
}

func (s Service) makeRequests(reqMethod, endpoint, reqBody string, jobs <-chan int, results chan<- string) {
	for range jobs {
		resp, err := s.httpClient.MakeRequest(reqMethod, endpoint, reqBody)
		if err != nil {
			fmt.Println(err)
			return
		}
		results <- string(resp)
	}
}
