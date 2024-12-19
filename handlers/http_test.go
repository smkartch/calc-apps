package handlers

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"testing"

	"github.com/smkartch/calc-apps/externals/should"
)

func assertEqual(t *testing.T, expected, actual any) {
	t.Helper()
	should.So(t, actual, should.Equal, expected)
}

func TestHTTPServer_NotFound(t *testing.T) {
	assertHTTP(t, http.MethodGet, "/bogus?a=1&b=2", http.StatusNotFound, "text/plain; charset=utf-8", "404 page not found\n")
}

func TestHTTPServer_MethodNotAllowed(t *testing.T) {
	assertHTTP(t, http.MethodPost, "/add?a=1&b=2", http.StatusMethodNotAllowed, "text/plain; charset=utf-8", "Method Not Allowed\n")
	assertHTTP(t, http.MethodPut, "/sub?a=1&b=2", http.StatusMethodNotAllowed, "text/plain; charset=utf-8", "Method Not Allowed\n")
	assertHTTP(t, http.MethodDelete, "/mult?a=1&b=2", http.StatusMethodNotAllowed, "text/plain; charset=utf-8", "Method Not Allowed\n")
	assertHTTP(t, http.MethodPatch, "/div?a=1&b=2", http.StatusMethodNotAllowed, "text/plain; charset=utf-8", "Method Not Allowed\n")
}

func TestHTTPServer_Add(t *testing.T) {
	assertHTTP(t, http.MethodGet, "/add?a=Na&b=2", http.StatusUnprocessableEntity, "text/plain; charset=utf-8", "a was invalid\n")
	assertHTTP(t, http.MethodGet, "/add?a=1&b=Na", http.StatusUnprocessableEntity, "text/plain; charset=utf-8", "b was invalid\n")
	assertHTTP(t, http.MethodGet, "/add?a=1&b=2", http.StatusOK, "text/plain; charset=utf-8", "3")
}

func TestHTTPServer_Subtract(t *testing.T) {
	assertHTTP(t, http.MethodGet, "/sub?a=Na&b=2", http.StatusUnprocessableEntity, "text/plain; charset=utf-8", "a was invalid\n")
	assertHTTP(t, http.MethodGet, "/sub?a=1&b=Na", http.StatusUnprocessableEntity, "text/plain; charset=utf-8", "b was invalid\n")
	assertHTTP(t, http.MethodGet, "/sub?a=2&b=1", http.StatusOK, "text/plain; charset=utf-8", "1")
}

func TestHTTPServer_Multiplication(t *testing.T) {
	assertHTTP(t, http.MethodGet, "/mult?a=Na&b=2", http.StatusUnprocessableEntity, "text/plain; charset=utf-8", "a was invalid\n")
	assertHTTP(t, http.MethodGet, "/mult?a=1&b=Na", http.StatusUnprocessableEntity, "text/plain; charset=utf-8", "b was invalid\n")
	assertHTTP(t, http.MethodGet, "/mult?a=2&b=4", http.StatusOK, "text/plain; charset=utf-8", "8")
}

func TestHTTPServer_Division(t *testing.T) {
	assertHTTP(t, http.MethodGet, "/div?a=Na&b=2", http.StatusUnprocessableEntity, "text/plain; charset=utf-8", "a was invalid\n")
	assertHTTP(t, http.MethodGet, "/div?a=1&b=Na", http.StatusUnprocessableEntity, "text/plain; charset=utf-8", "b was invalid\n")
	assertHTTP(t, http.MethodGet, "/div?a=4&b=2", http.StatusOK, "text/plain; charset=utf-8", "2")
}

func assertHTTP(t *testing.T, method string, target string, expectedStatus int, expectedContentType string, expectedResponse string) {
	t.Run(fmt.Sprintf("%s %s", method, target), func(t *testing.T) {
		request := httptest.NewRequest(method, target, nil)
		response := httptest.NewRecorder()

		dumpRequest, _ := httputil.DumpRequest(request, true)
		t.Log("\n" + string(dumpRequest))

		NewRouter(log.Default()).ServeHTTP(response, request)

		dumpResponse, _ := httputil.DumpResponse(response.Result(), true)
		t.Log("\n" + string(dumpResponse))

		assertEqual(t, expectedStatus, response.Code)
		assertEqual(t, expectedContentType, response.Header().Get("Content-Type"))
		assertEqual(t, expectedResponse, response.Body.String())
	})
}
