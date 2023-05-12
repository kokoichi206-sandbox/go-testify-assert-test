package mystruct_test

import (
	"reflect"
	"testing"

	"github.com/kokoichi206-sandbox/go-testify-assert-test/mystruct"
	"github.com/stretchr/testify/assert"
)

func TestNewMe(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		expected mystruct.Me
		errMsg   string
	}{
		"test1": {
			expected: mystruct.Me{
				Name: "kokoichi206",
			},
		},
	}

	for tcName, tc := range tests {
		tcName := tcName
		tc := tc

		t.Run(tcName, func(t *testing.T) {
			t.Parallel()

			actual := mystruct.NewMe()

			if tc.expected != actual {
				t.Error("result didn't match")
			}

			if !reflect.DeepEqual(tc.expected, actual) {
				t.Error("result didn't match")
			}
			assert.Equal(t, tc.expected, actual, "result didn't match")
		})
	}
}

func TestMarshal(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		expected string
		errMsg   string
	}{
		"test1": {
			expected: `{"Name":"kokoichi206"}`,
		},
	}

	for tcName, tc := range tests {
		tcName := tcName
		tc := tc

		t.Run(tcName, func(t *testing.T) {
			t.Parallel()

			actual := mystruct.Marshal()

			if tc.expected != actual {
				t.Error("result didn't match")
			}

			if !reflect.DeepEqual(tc.expected, actual) {
				t.Error("result didn't match")
			}
		})
	}
}
