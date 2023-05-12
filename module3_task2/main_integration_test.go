package main

import (
  "io"
  "net/http"
  "net/http/httptest"
  "testing"
)

func Test_server(t *testing.T) {
  if testing.Short() {
    t.Skip("Flag `-short` provided: skipping Integration Tests.")
  }

  tests := []struct {
    name         string
    URI          string
    responseCode int
    body         string
  }{
    {
      name:         "Home page",
      URI:          "",
      responseCode: 200,
      body:         "",
    },
    {
	  name:         "Rosalind Franklin",
	  URI: 			"/hello?name=Rosalind",
	  responseCode: 200,
	  body:         "Hello Rosalind!",
	},
	{
	  name:         "Empty name",
	  URI:  		"/hello?name=",
	  responseCode: 400,
	},
	{
	  name:         "No name",
	  URI:  		"/hello",
	  responseCode: 200,
	  body:         "Hello there!",
	},
	{
	  name:         "Healthcheck page",
	  URI:          "/health",
	  responseCode: 200,
	  body:         "ALIVE",
	},
  }

  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      ts := httptest.NewServer(setupRouter())
      defer ts.Close()

      res, err := http.Get(ts.URL + tt.URI)
      if err != nil {
        t.Fatal(err)
      }

      // Check that the status code is what you expect.
      expectedCode := tt.responseCode
      gotCode := res.StatusCode
      if gotCode != expectedCode {
        t.Errorf("handler returned wrong status code: got %q want %q", gotCode, expectedCode)
      }

      // Check that the response body is what you expect.
      expectedBody := tt.body
      bodyBytes, err := io.ReadAll(res.Body)
      res.Body.Close()
      if err != nil {
        t.Fatal(err)
      }
      gotBody := string(bodyBytes)
      if gotBody != expectedBody && len(expectedBody) > 0 {
        t.Errorf("handler returned unexpected body: got %q want %q", gotBody, expectedBody)
      }
    })
  }
}
