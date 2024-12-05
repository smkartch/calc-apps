package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"reflect"
	"testing"
)

func assertEqual(t *testing.T, expected, actual any) {
	t.Helper()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got: [%v] want: [%v]", actual, expected)
	}
}

func TestHTTPServer_Add(t *testing.T) {
	assertHTTP(t, http.MethodGet, "/bogus?a=1&b=2", http.StatusNotFound, "text/plain; charset=utf-8", "404 page not found\n")
	assertHTTP(t, http.MethodPost, "/add?a=1&b=2", http.StatusMethodNotAllowed, "text/plain; charset=utf-8", "Method Not Allowed\n")
	assertHTTP(t, http.MethodGet, "/add?a=Na&b=2", http.StatusUnprocessableEntity, "text/plain; charset=utf-8", "a was invalid\n")
	assertHTTP(t, http.MethodGet, "/add?a=1&b=Na", http.StatusUnprocessableEntity, "text/plain; charset=utf-8", "b was invalid\n")
	assertHTTP(t, http.MethodGet, "/add?a=1&b=2", http.StatusOK, "text/plain; charset=utf-8", "3")
}

func assertHTTP(t *testing.T, method string, target string, expectedStatus int, expectedContentType string, expectedResponse string) {
	t.Run(fmt.Sprintf("%s %s", method, target), func(t *testing.T) {
		request := httptest.NewRequest(method, target, nil)
		response := httptest.NewRecorder()

		dumpRequest, _ := httputil.DumpRequest(request, true)
		t.Log("\n" + string(dumpRequest))

		NewRouter(nil).ServeHTTP(response, request)

		dumpResponse, _ := httputil.DumpResponse(response.Result(), true)
		t.Log("\n" + string(dumpResponse))

		assertEqual(t, expectedStatus, response.Code)
		assertEqual(t, expectedContentType, response.Header().Get("Content-Type"))
		assertEqual(t, expectedResponse, response.Body.String())
	})
}
