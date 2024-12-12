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
	router.Handle("GET /sub", NewHTTPHandler(logger, &calc.Subtraction{}))
	router.Handle("GET /mult", NewHTTPHandler(logger, &calc.Multiplication{}))
	router.Handle("GET /div", NewHTTPHandler(logger, &calc.Division{}))
	return router
}

type HTTPHandler struct {
	logger     *log.Logger
	calculator Calculator
}

func NewHTTPHandler(logger *log.Logger, calculator Calculator) http.Handler {
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
	_, err = fmt.Fprint(response, c)
	if err != nil {
		this.logger.Println(err)
	}
}
