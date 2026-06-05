package wcs201

import (
	"testing"
)

func TestWCSExceptions(t *testing.T) {
	var tests = []struct {
		exception     interface{ Error() string; Code() string; Locator() string }
		exceptionText string
		exceptionCode string
		locatorCode   string
	}{
		0: {exception: NoSuchCoverage("test_coverage"),
			exceptionText: "Coverage not found: test_coverage",
			exceptionCode: "NoSuchCoverage",
			locatorCode:   "test_coverage",
		},
		1: {exception: InvalidSubsetting("X"),
			exceptionText: "Invalid subset for axis: X",
			exceptionCode: "InvalidSubsetting",
			locatorCode:   "X",
		},
		2: {exception: SubsettingNotSupported("Y"),
			exceptionText: "Subsetting not supported for axis: Y",
			exceptionCode: "SubsettingNotSupported",
			locatorCode:   "Y",
		},
		3: {exception: CoverageNotDefined("some_coverage"),
			exceptionText: "Coverage not defined: some_coverage",
			exceptionCode: "CoverageNotDefined",
			locatorCode:   "some_coverage",
		},
	}
	for k, test := range tests {
		if test.exception.Error() != test.exceptionText {
			t.Errorf("test: %d, expected: %s\n got: %s", k, test.exceptionText, test.exception.Error())
		}
		if test.exception.Code() != test.exceptionCode {
			t.Errorf("test: %d, expected: %s\n got: %s", k, test.exceptionCode, test.exception.Code())
		}
		if test.exception.Locator() != test.locatorCode {
			t.Errorf("test: %d, expected: %s\n got: %s", k, test.locatorCode, test.exception.Locator())
		}
	}
}
