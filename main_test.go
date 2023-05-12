package main

import (
	"testing"

	"github.com/kokoichi206-sandbox/go-testify-assert-test/gen/go/protobuf"
	"github.com/stretchr/testify/assert"
)

func TestAssertion(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		name     string
		expected *protobuf.HelloReply
		errMsg   string
	}{
		"test1": {
			name: "john",
			expected: &protobuf.HelloReply{
				Name: "john",
			},
		},
	}

	for tcName, tc := range tests {
		tcName := tcName
		tc := tc

		t.Run(tcName, func(t *testing.T) {
			t.Parallel()

			actual := NewProtoMessage(tc.name)

			assert.Equal(t, tc.expected, actual, "result didn't match")
		})
	}
}
