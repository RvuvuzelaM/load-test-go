package loadtestrest

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	mockLoadTestRest "github.com/vuvuzela/loadtest/internal/loadtestrest/mock"
)

func TestLoadTestRestAPI(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockHTTPClient := mockLoadTestRest.NewMockHTTPClient(mockCtrl)
	svc := NewService(mockHTTPClient)

	mockHTTPClient.EXPECT().
		MakeRequest(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Return(`{"test": "test"}`, nil).
		Times(10)

	in := LoadTestRestAPIInput{
		NumOfRequests:     10,
		ConcurentRequests: 2,
		RequestMethod:     "GET",
		Endpoint:          "https://example.com",
	}
	actual, err := svc.LoadTestRestAPI(in)

	expected := []string{}
	for i := 0; i < 10; i++ {
		expected = append(expected, `{"test": "test"}`)
	}

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}
