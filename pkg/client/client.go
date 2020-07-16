// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package client

import "net/http"

// Stubs to keep test errors from happening when bringing in the client.
// This will be overritten the first time you run the openapi-generator
type Configuration struct {
	HTTPClient *http.Client
}

func NewConfiguration() *Configuration {
	return &Configuration{}
}

type APIClient struct {
	cfg *Configuration
}

func NewAPIClient(cfg *Configuration) *APIClient {
	return &APIClient{
		cfg: NewConfiguration(),
	}
}
