package zendesk

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/linda-lai/golang-ticket-viewer/internal/api"
	"github.com/linda-lai/golang-ticket-viewer/tests/mocks"
)

func init() {
	api.Client = &mocks.MockClient{}
}

func TestUnmarshalZendeskTickets(t *testing.T) {
	json := `{"id: 1", "description": "foobar"}`
	auth := "2snaj"
	url := "helloworld.com"
	data := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	api.Client = &mocks.MockClient{} //
	mocks.GetDo = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       data,
		}, nil
	}

	resp, _ := api.ListTickets(auth, url)
	assert.Nil(t, resp)
	assert.EqualValues(t, 200, http.Response{})
}
