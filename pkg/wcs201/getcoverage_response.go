package wcs201

import (
	"encoding/xml"
)

// Type returns GetCoverage
func (g *GetCoverageResponse) Type() string {
	return getcoverage
}

// Service returns the service type
func (g *GetCoverageResponse) Service() string {
	return Service
}

// Version returns the service version
func (g *GetCoverageResponse) Version() string {
	return Version
}

// GetCoverageResponse struct
type GetCoverageResponse struct {
	XMLName      xml.Name      `xml:"Coverage" yaml:"coverage"`
	CoverageID   string        `xml:"CoverageId" yaml:"coverageid"`
	BinaryPayload BinaryPayload `xml:"BinaryPayload,omitempty" yaml:"binarypayload"`
}

// BinaryPayload struct
type BinaryPayload struct {
	Format   string `xml:"Format" yaml:"format"`
	Contents string `xml:"BinaryContent" yaml:"binarycontents"`
}
