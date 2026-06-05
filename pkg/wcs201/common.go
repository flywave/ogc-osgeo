package wcs201

//
const (
	getcapabilities  = `GetCapabilities`
	describecoverage = `DescribeCoverage`
	getcoverage      = `GetCoverage`
)

// Type and Version as constant
const (
	Service string = `WCS`
	Version string = `2.0.1`
)

// KVP parameter tokens
const (
	SERVICE      = `SERVICE`
	REQUEST      = `REQUEST`
	VERSION      = `VERSION`
	COVERAGEID   = `COVERAGEID`
	FORMAT       = `FORMAT`
	SUBSET       = `SUBSET`
	SUBSETTINGCRS = `SUBSETTINGCRS`
	OUTPUTCRS   = `OUTPUTCRS`
	MEDIATYPE   = `MEDIATYPE`
)
