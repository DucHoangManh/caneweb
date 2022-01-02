package cane

import (
	"fmt"
	"testing"
)

type s struct {
	Name string
}

func Test(t *testing.T) {
	sl := make([]*s, 0)
	sl = append(sl, &s{Name: "er"})
	sl = append(sl, &s{Name: "ad"})
	for _, v := range sl {
		fmt.Println(v.Name)
	}
}
