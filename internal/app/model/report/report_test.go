package report_test

import (
	"github.com/stretchr/testify/assert"
	"mobio/internal/app/model/report"
	"testing"
)

func TestSpecialization_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		s       func() *report.Report
		isValid bool
	}{
		{
			name: "valid",
			s: func() *report.Report {
				return report.TestReport(t)
			},
			isValid: true,
		},
		{
			name: "empty",
			s: func() *report.Report {
				s := report.TestReport(t)
				s.Email = ""
				return s
			},
			isValid: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.s().Validate())
			} else {
				assert.Error(t, tc.s().Validate())
			}
		})
	}
}
