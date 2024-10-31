package handlers

import (
	"bytes"
	"errors"
	"testing"

	"github.com/smkartch/calc-lib"
)

func assertError(t *testing.T, actual, target error) {
	t.Helper()
	if !errors.Is(actual, target) {
		t.Errorf("expected %v, got %v", target, actual)
	}
}

func TestHandler_WrongNumberOfArguments(t *testing.T) {
	handler := NewHandler(nil, nil)
	err := handler.Handle(nil)
	assertError(t, err, errWrongNumberOfArgs)
}
func TestHandler_InvalidFirstArgument(t *testing.T) {
	handler := NewHandler(nil, nil)
	err := handler.Handle([]string{"INVALID", "47"})
	assertError(t, err, errInvalidArg)
}
func TestHandler_InvalidSecondArgument(t *testing.T) {
	handler := NewHandler(nil, nil)
	err := handler.Handle([]string{"47", "INVALID"})
	assertError(t, err, errInvalidArg)
}
func TestHandler_OutputWriterError(t *testing.T) {
	oops := errors.New("oops")
	writer := &ErringWriter{err: oops}
	handler := NewHandler(writer, nil)
	err := handler.Handle([]string{"2", "3"})
	assertError(t, err, oops)
	assertError(t, err, errWriterFailure)
}
func TestHandler_HappyPath(t *testing.T) {
	writer := &bytes.Buffer{}
	handler := NewHandler(writer, &calc.Addition{})
	err := handler.Handle([]string{"2", "3"})
	assertError(t, err, nil)
	if writer.String() != "5" {
		t.Errorf("expected 5, got %s", writer.String())
	}
}

type ErringWriter struct {
	err error
}

func (this *ErringWriter) Write(p []byte) (n int, err error) {
	return 0, this.err
}
