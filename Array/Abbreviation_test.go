package Array

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAbbrevName(t *testing.T) {
	cases := []struct {
		In  string
		Out string
	}{
		{In: "Super Man", Out: "S. Man"},
		{In: "Batman", Out: "Batman"},
		{In: "Erih Maria Remark", Out: "E. M. Remark"},
	}
	for _, v := range cases {
		fmt.Println(assert.Equal(t, v.Out, GetAbbrevName(v.In)))
	}
}
