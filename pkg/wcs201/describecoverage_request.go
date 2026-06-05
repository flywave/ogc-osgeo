package wcs201

import (
	"encoding/xml"
	"net/url"
	"strings"

	"github.com/flywave/ogc-osgeo/pkg/utils"
	"github.com/flywave/ogc-osgeo/pkg/wsc200"
)

// Type returns DescribeCoverage
func (d *DescribeCoverageRequest) Type() string {
	return describecoverage
}

// Validate validates the DescribeCoverage request
func (d *DescribeCoverageRequest) Validate(c Capabilities) []wsc200.Exception {
	var exceptions []wsc200.Exception
	if len(d.CoverageID) == 0 {
		exceptions = append(exceptions, wsc200.MissingParameterValue(COVERAGEID))
	}
	return exceptions
}

// ParseXML builds a DescribeCoverage object based on a XML document
func (d *DescribeCoverageRequest) ParseXML(body []byte) []wsc200.Exception {
	var xmlattributes utils.XMLAttribute
	if err := xml.Unmarshal(body, &xmlattributes); err != nil {
		return wsc200.MissingParameterValue().ToExceptions()
	}
	if err := xml.Unmarshal(body, &d); err != nil {
		return wsc200.MissingParameterValue(REQUEST).ToExceptions()
	}
	var n []xml.Attr
	for _, a := range xmlattributes {
		switch strings.ToUpper(a.Name.Local) {
		case VERSION:
		case SERVICE:
		default:
			n = append(n, a)
		}
	}
	d.Attr = utils.StripDuplicateAttr(n)
	return nil
}

// QueryParameters builds a DescribeCoverage object based on the available query parameters
func (d *DescribeCoverageRequest) QueryParameters(query url.Values) []wsc200.Exception {
	dpv := describeCoverageRequestParameterValue{}

	if exceptions := dpv.parseQueryParameters(query); exceptions != nil {
		return exceptions
	}

	if exceptions := d.parsedescribeCoverageRequestParameterValue(dpv); exceptions != nil {
		return exceptions
	}
	return nil
}

func (d *DescribeCoverageRequest) parsedescribeCoverageRequestParameterValue(dpv describeCoverageRequestParameterValue) []wsc200.Exception {
	d.XMLName.Local = describecoverage
	d.Service = Service
	d.Version = Version
	d.CoverageID = dpv.coverageID
	return nil
}

// ToQueryParameters builds a new query string that will be proxied
func (d DescribeCoverageRequest) ToQueryParameters() url.Values {
	dpv := describeCoverageRequestParameterValue{}
	dpv.parsedescribeCoverageRequest(d)

	q := dpv.toQueryParameters()
	return q
}

// ToXML builds a 'new' XML document 'based' on the 'original' XML document
func (d DescribeCoverageRequest) ToXML() []byte {
	si, _ := xml.Marshal(&d)
	return append([]byte(xml.Header), si...)
}

// DescribeCoverageRequest struct with the needed parameters/attributes needed for making a DescribeCoverage request
type DescribeCoverageRequest struct {
	XMLName    xml.Name           `xml:"DescribeCoverage" yaml:"describecoverage"`
	Service    string             `xml:"service,attr" yaml:"service"`
	Version    string             `xml:"version,attr" yaml:"version"`
	CoverageID []string           `xml:"CoverageId" yaml:"coverageid"`
	Attr       utils.XMLAttribute `xml:",attr"`
}
