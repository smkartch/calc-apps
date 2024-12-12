package should_test

import (
	"errors"
	"testing"
	"time"

	"github.com/smkartch/calc-apps/externals/should"
)

type FakeT struct {
	err    error
	helper bool
}

func (this *FakeT) Helper() {
	this.helper = true
}
func (this *FakeT) Error(values ...any) {
	this.err = values[0].(error)
}

func pass(t *testing.T, actual any, assert should.Assertion, expected ...any) {
	fakeT := &FakeT{}
	should.So(fakeT, actual, assert, expected...)
	if fakeT.err != nil {
		t.Errorf("should not get an error, but did: %v", fakeT.err)
	}
}
func fail(t *testing.T, actual any, assert should.Assertion, expected ...any) {
	fakeT := &FakeT{}
	should.So(fakeT, actual, assert, expected...)
	if !errors.Is(fakeT.err, should.ErrAssertionFailure) {
		t.Errorf("should get an assertion error, but didn't: %v", fakeT.err)
	} else {
		t.Log(fakeT.err)
	}
	if !fakeT.helper {
		t.Errorf("should have caleld t.Helper() but didn't")
	}
}

func TestShouldEqual(t *testing.T) {
	pass(t, 1, should.Equal, 1)
	pass(t, []int{1, 2, 3}, should.Equal, []int{1, 2, 3})
	fail(t, 1, should.Equal, 2)
	fail(t, []int{1, 2, 3}, should.Equal, []int{4, 5, 6})
}
func TestShouldNotEqual(t *testing.T) {
	pass(t, 1, should.NOT.Equal, 2)
	fail(t, 1, should.NOT.Equal, 1)
}
func TestShouldBeTrue(t *testing.T) {
	pass(t, true, should.BeTrue)
	fail(t, false, should.BeTrue)
	fail(t, nil, should.BeTrue)
	fail(t, 1, should.BeTrue)
	fail(t, []int{1}, should.BeTrue)
}
func TestShouldBeFalse(t *testing.T) {
	pass(t, false, should.BeFalse)
	fail(t, true, should.BeFalse)
	fail(t, nil, should.BeFalse)
	fail(t, 1, should.BeFalse)
	fail(t, []int{1}, should.BeFalse)
}
func TestShouldBeNil(t *testing.T) {
	pass(t, nil, should.BeNil)
	fail(t, true, should.BeNil)
	fail(t, false, should.BeNil)
	fail(t, 1, should.BeNil)
	fail(t, []int{1}, should.BeNil)
}
func TestShouldNotBeNil(t *testing.T) {
	pass(t, &time.Time{}, should.NOT.BeNil)
	fail(t, nil, should.NOT.BeNil)
}
