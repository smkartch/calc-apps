package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/smkartch/calc-lib"
)

func NewRouter(logger *log.Logger) http.Handler {
	router := http.NewServeMux()
	router.Handle("GET /add", NewHTTPHandler(logger, &calc.Addition{}))
	return router
}

type HTTPHandler struct {
	logger     *log.Logger
	calculator *calc.Addition
}

func NewHTTPHandler(logger *log.Logger, calculator *calc.Addition) http.Handler {
	return &HTTPHandler{logger: logger, calculator: calculator}
}

func (this *HTTPHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	a, err := strconv.Atoi(query.Get("a"))
	if err != nil {
		http.Error(response, "a was invalid", http.StatusUnprocessableEntity)
		return
	}
	b, err := strconv.Atoi(query.Get("b"))
	if err != nil {
		http.Error(response, "b was invalid", http.StatusUnprocessableEntity)
		return
	}
	c := this.calculator.Calculate(a, b)
	_, _ = fmt.Fprint(response, c)
}
