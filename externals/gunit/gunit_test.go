package gunit_test

import (
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

func TestMySuperCoolFixture(t *testing.T) {
	gunit.Run(new(MySuperCoolFixture), t)
}

type MySuperCoolFixture struct {
	*gunit.Fixture
}

func (this *MySuperCoolFixture) Test1() {
	this.So(1, should.Equal, 1)
}

func (this *MySuperCoolFixture) SkipTest1() {
	this.So(1, should.Equal, 2)
}
