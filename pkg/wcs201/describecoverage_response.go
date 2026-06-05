package wcs201

import (
	"encoding/xml"
)

// Type returns DescribeCoverage
func (d *CoverageDescriptions) Type() string {
	return describecoverage
}

// ToXML builds a CoverageDescriptions XML document
func (d CoverageDescriptions) ToXML() []byte {
	si, _ := xml.Marshal(d)
	return append([]byte(xml.Header), si...)
}

// CoverageDescriptions is the response to a DescribeCoverage request
type CoverageDescriptions struct {
	XMLName              xml.Name              `xml:"CoverageDescriptions" yaml:"coveragedescriptions"`
	CoverageDescription []CoverageDescription `xml:"CoverageDescription" yaml:"coveragedescription"`
}

// CoverageDescription describes a single coverage
type CoverageDescription struct {
	CoverageID      string   `xml:"CoverageId" yaml:"coverageid"`
	CoverageSubtype string   `xml:"CoverageSubtype,omitempty" yaml:"coveragesubtype"`
	Title           string   `xml:"ows:Title,omitempty" yaml:"title"`
	Abstract        string   `xml:"ows:Abstract,omitempty" yaml:"abstract"`
	SupportedCRS    []string `xml:"SupportedCRS,omitempty" yaml:"supportedcrs"`
	SupportedFormat []string `xml:"SupportedFormat,omitempty" yaml:"supportedformat"`
}
