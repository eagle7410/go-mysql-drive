package lib

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)
var env = Env{}

func init() {
	env.Init()
}

func TestCircle(t *testing.T) {

	Convey("Connect mysql", t, func() {
		_, version, err := Init(&env)
		So(err, ShouldBeNil)
		So(len(version), ShouldBeGreaterThan, 0)
		fmt.Println(version)
	})
}


func BenchmarkConnect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Init(&env)
	}
}
