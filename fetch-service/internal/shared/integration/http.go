package integration

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Request struct {
	HttpClient *http.Client
	Method     string
	URL        string
}

func PerformRequest(ctx context.Context, req Request, result interface{}) error {
	// for now we only Get method
	resp, err := http.Get(req.URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, result)
	if err != nil {
		return err
	}

	return nil
}
