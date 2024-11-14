package handlers

import (
	"fmt"
	"io"
	"strconv"
)

type Calculator interface {
	Calculate(a, b int) int
}

type Handler struct {
	stdout     io.Writer
	calculator Calculator
}

func NewHandler(stdout io.Writer, calculator Calculator) *Handler {
	return &Handler{
		stdout:     stdout,
		calculator: calculator,
	}
}
func (this *Handler) Handle(args []string) error {
	if len(args) != 2 {
		return errWrongNumberOfArgs
	}
	a, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("%w: [%s] %w", errInvalidArg, args[0], err)
	}
	b, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("%w: [%s] %w", errInvalidArg, args[1], err)
	}
	result := this.calculator.Calculate(a, b)
	_, err = fmt.Fprint(this.stdout, result)
	if err != nil {
		return fmt.Errorf("%w: %w", errWriterFailure, err)
	}
	return nil
}

var (
	errWrongNumberOfArgs = fmt.Errorf("usage: calc [a] [b]")
	errInvalidArg        = fmt.Errorf("invalid argument")
	errWriterFailure     = fmt.Errorf("error writing")
)
