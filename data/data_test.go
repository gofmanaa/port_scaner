package data_test

import (
	"testing"

	"github.com/gofmanaa/port_scanner/data"
)

func TestSet(t * testing.T) {
	s := data.NewSet()
	s.Add(12)
	expect :=[]int{12}
	res := s.GetInt() 
	if res[0] != expect[0] {
		t.FailNow()
	}
}