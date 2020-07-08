package test

import (
	"pkg2"
	"testing"
)

func TestPkg2_Bye(t *testing.T) {
	ret := pkg2.Bye()
	t.Log(ret)
}
