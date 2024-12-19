package handlers

import (
	"bytes"
	"errors"
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smkartch/calc-lib"
)

func assertError(t *testing.T, actual, target error) {
	t.Helper()
	should.So(t, actual, should.Wrap, target)
}

func TestHandler_WrongNumberOfArguments(t *testing.T) {
	handler := NewHandler(nil, nil)
	err := handler.Handle(nil)
	should.So(t, err, should.Wrap, errWrongNumberOfArgs)
}
func TestHandler_InvalidFirstArgument(t *testing.T) {
	handler := NewHandler(nil, nil)
	err := handler.Handle([]string{"INVALID", "47"})
	should.So(t, err, should.Wrap, errInvalidArg)
}
func TestHandler_InvalidSecondArgument(t *testing.T) {
	handler := NewHandler(nil, nil)
	err := handler.Handle([]string{"47", "INVALID"})
	should.So(t, err, should.Wrap, errInvalidArg)
}
func TestHandler_OutputWriterError(t *testing.T) {
	oops := errors.New("oops")
	writer := &ErringWriter{err: oops}
	handler := NewHandler(writer, &calc.Addition{})
	err := handler.Handle([]string{"2", "3"})
	should.So(t, err, should.Wrap, oops)
	should.So(t, err, should.Wrap, errWriterFailure)
}
func TestHandler_HappyPath(t *testing.T) {
	writer := &bytes.Buffer{}
	handler := NewHandler(writer, &calc.Addition{})
	err := handler.Handle([]string{"2", "3"})
	should.So(t, err, should.BeNil)
	should.So(t, writer.String(), should.Equal, "5")
}

type ErringWriter struct {
	err error
}

func (this *ErringWriter) Write(p []byte) (n int, err error) {
	return 0, this.err
}
