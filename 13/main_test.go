package main

import (
	"encoding/json"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test13(t *testing.T) {
	tests := []struct {
		left, right string
		expect      bool
	}{
		{
			left:   `[1,1,3,1,1]`,
			right:  `[1,1,5,1,1]`,
			expect: true,
		},
		{
			left:   `[[1],[2,3,4]]`,
			right:  `[[1],4]`,
			expect: true,
		},
		{
			left:   `[9]`,
			right:  `[[8,7,6]]`,
			expect: false,
		},
		{
			left:   `[[4,4],4,4]`,
			right:  `[[4,4],4,4,4]`,
			expect: true,
		},

		{
			left:   `[7,7,7,7]`,
			right:  `[7,7,7]`,
			expect: false,
		},

		{
			left:   `[]`,
			right:  `[3]`,
			expect: true,
		},

		{
			left:   `[[[]]]`,
			right:  `[[]]`,
			expect: false,
		},
		{
			left:   `[1,[2,[3,[4,[5,6,7]]]],8,9]`,
			right:  `[1,[2,[3,[4,[5,6,0]]]],8,9]`,
			expect: false,
		},
		{
			left:   `[]`,
			right:  `[]`,
			expect: true,
		},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var left, right []interface{}
			json.Unmarshal([]byte(tt.left), &left)
			json.Unmarshal([]byte(tt.right), &right)

			res := compare(left, right)
			assert.Equal(t, tt.expect, res == nil || *res)
		})
	}
}
