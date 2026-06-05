package wcs201

import (
	"fmt"

	"github.com/flywave/ogc-osgeo/pkg/wsc200"
)

// NoSuchCoverage exception
func NoSuchCoverage(coverageID string) wsc200.Exception {
	return exception{
		ExceptionText: fmt.Sprintf("Coverage not found: %s", coverageID),
		ExceptionCode: "NoSuchCoverage",
		LocatorCode:   coverageID,
	}
}

// InvalidSubsetting exception
func InvalidSubsetting(axis string) wsc200.Exception {
	return exception{
		ExceptionText: fmt.Sprintf("Invalid subset for axis: %s", axis),
		ExceptionCode: "InvalidSubsetting",
		LocatorCode:   axis,
	}
}

// SubsettingNotSupported exception
func SubsettingNotSupported(axis string) wsc200.Exception {
	return exception{
		ExceptionText: fmt.Sprintf("Subsetting not supported for axis: %s", axis),
		ExceptionCode: "SubsettingNotSupported",
		LocatorCode:   axis,
	}
}

// CoverageNotDefined exception
func CoverageNotDefined(coverageID string) wsc200.Exception {
	return exception{
		ExceptionText: fmt.Sprintf("Coverage not defined: %s", coverageID),
		ExceptionCode: "CoverageNotDefined",
		LocatorCode:   coverageID,
	}
}
