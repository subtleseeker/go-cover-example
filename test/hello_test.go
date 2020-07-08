package test

import (
	"pkg1"
	"testing"
)

func TestPkg1_Hello(t *testing.T) {
	ret := pkg1.Hello()
	if ret != "hello" {
		t.Error("error")
	}
	t.Log(ret)
}
