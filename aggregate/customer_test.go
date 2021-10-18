package aggregate_test

import (
	"testing"

	"github.com/condezero/domain-driven-go/aggregate"
)

func TestCustomer_NewCustomer(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Empty name",
			name:        "",
			expectedErr: nil,
		},
		{
			test:        "Valid name",
			name:        "My test name",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := aggregate.NewCustomer(tc.name)
			if err != tc.expectedErr {
				t.Errorf("Expected error: %s, got: %s", tc.expectedErr, err.Error())
			}
		})
	}
}
