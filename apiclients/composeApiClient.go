package apiclients 

import (
  "net/http"
  "io/ioutil"
)

type ApiClient func(subPath string) (string, error)
type BodyExtractor func(response *http.Response) (string, error)

func ComposeApiClient(apiBaseUrl string) ApiClient {
	return func(subPath string) (string, error) {
		response, apiErr := http.Get(apiBaseUrl + subPath)

		if apiErr != nil {
			return "", apiErr
		}

		bodyBytes, ioErr := ioutil.ReadAll(response.Body)

		if ioErr != nil {
			return "", ioErr
		}

		return string(bodyBytes), nil
	}
}

