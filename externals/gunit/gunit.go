package gunit

import (
	"reflect"
	"strings"
	"testing"
)

type Setup interface {
	Setup()
}

func Run(t *testing.T, fixture any) {
	// scan fixture type, looking for Setup(), Test...()
	// for each test:
	//   start new subtest
	//     create new instance of fixture
	//     call .Setup()
	//     call .Test...()

	fixtureType := reflect.TypeOf(fixture)
	for i := 0; i < fixtureType.NumMethod(); i++ {
		name := fixtureType.Method(i).Name
		if strings.HasPrefix(name, "SkipTest") {
			t.Run(name, func(t *testing.T) {
				t.Skip()
			})
		} else if strings.HasPrefix(name, "Test") {
			t.Run(name, func(t *testing.T) {
				fixtureValue := reflect.New(fixtureType.Elem())
				fixtureValue.Elem().FieldByName("Fixture").Set(
					reflect.ValueOf(&Fixture{T: t}),
				)
				fixtureWithSetup, ok := fixtureValue.Interface().(Setup)
				if ok {
					fixtureWithSetup.Setup()
				}
				fixtureValue.MethodByName(name).Call(nil)
			})
		}
	}
}

type Fixture struct{ *testing.T }

func (this *Fixture) So(actual any, assert assertion, expected ...any) {
	err := assert(actual, expected...)
	if err != nil {
		this.Helper()
		this.Error(err)
	}
}

type assertion func(actual any, expected ...any) error
