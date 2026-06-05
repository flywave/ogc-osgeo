package wcs201

import (
	"encoding/xml"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/flywave/ogc-osgeo/pkg/utils"
	"github.com/flywave/ogc-osgeo/pkg/wsc200"
)

// Type returns GetCoverage
func (g *GetCoverageRequest) Type() string {
	return getcoverage
}

// Validate validates the GetCoverage request
func (g *GetCoverageRequest) Validate(c Capabilities) []wsc200.Exception {
	var exceptions []wsc200.Exception
	if g.CoverageID == "" {
		exceptions = append(exceptions, wsc200.MissingParameterValue(COVERAGEID))
	}
	return exceptions
}

// ParseXML builds a GetCoverage object based on a XML document
func (g *GetCoverageRequest) ParseXML(body []byte) []wsc200.Exception {
	var xmlattributes utils.XMLAttribute
	if err := xml.Unmarshal(body, &xmlattributes); err != nil {
		return wsc200.MissingParameterValue().ToExceptions()
	}
	if err := xml.Unmarshal(body, &g); err != nil {
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
	g.Attr = utils.StripDuplicateAttr(n)
	return nil
}

// QueryParameters builds a GetCoverage object based on the available query parameters
func (g *GetCoverageRequest) QueryParameters(query url.Values) []wsc200.Exception {
	gpv := getCoverageRequestParameterValue{}

	if exceptions := gpv.parseQueryParameters(query); exceptions != nil {
		return exceptions
	}

	if exceptions := g.parsegetCoverageRequestParameterValue(gpv); exceptions != nil {
		return exceptions
	}
	return nil
}

func (g *GetCoverageRequest) parsegetCoverageRequestParameterValue(gpv getCoverageRequestParameterValue) []wsc200.Exception {
	g.XMLName.Local = getcoverage
	g.Service = Service
	g.Version = Version
	g.CoverageID = gpv.coverageID
	g.Format = gpv.format
	g.DimensionSubset = gpv.dimensionSubset
	g.SubsettingCRS = gpv.subsettingCRS
	g.OutputCRS = gpv.outputCRS
	g.MediaType = gpv.mediaType
	return nil
}

// ToQueryParameters builds a new query string that will be proxied
func (g GetCoverageRequest) ToQueryParameters() url.Values {
	gpv := getCoverageRequestParameterValue{}
	gpv.parseGetCoverageRequest(g)

	q := gpv.toQueryParameters()
	return q
}

// ToXML builds a 'new' XML document 'based' on the 'original' XML document
func (g GetCoverageRequest) ToXML() []byte {
	si, _ := xml.Marshal(&g)
	return append([]byte(xml.Header), si...)
}

// DimensionSubset models a WCS 2.0.1 DimensionSubset element
type DimensionSubset struct {
	Dimension string `xml:"Dimension" yaml:"dimension"`
	Trim      *Trim  `xml:"Trim,omitempty" yaml:"trim"`
	Slice     *Slice `xml:"Slice,omitempty" yaml:"slice"`
}

// Trim models a WCS 2.0.1 trim subset
type Trim struct {
	Low  float64 `xml:"Low" yaml:"low"`
	High float64 `xml:"High" yaml:"high"`
}

// Slice models a WCS 2.0.1 slice subset
type Slice struct {
	Value float64 `xml:"Value" yaml:"value"`
}

// parseSubsetString parses a KVP SUBSET parameter into a DimensionSubset.
// Format: "axis(low,high)" for trim or "axis(value)" for slice.
func parseSubsetString(s string) (*DimensionSubset, error) {
	idx := strings.Index(s, "(")
	if idx == -1 {
		return nil, fmt.Errorf("invalid subset format: %s", s)
	}
	axis := s[:idx]
	rest := s[idx+1:]
	if !strings.HasSuffix(rest, ")") {
		return nil, fmt.Errorf("invalid subset format: %s", s)
	}
	rest = strings.TrimSuffix(rest, ")")

	parts := strings.Split(rest, ",")
	ds := &DimensionSubset{Dimension: axis}

	if len(parts) == 1 {
		v, err := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64)
		if err != nil {
			return nil, fmt.Errorf("invalid slice value: %s", parts[0])
		}
		ds.Slice = &Slice{Value: v}
	} else if len(parts) == 2 {
		low, err := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64)
		if err != nil {
			return nil, fmt.Errorf("invalid trim low value: %s", parts[0])
		}
		high, err := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)
		if err != nil {
			return nil, fmt.Errorf("invalid trim high value: %s", parts[1])
		}
		ds.Trim = &Trim{Low: low, High: high}
	} else {
		return nil, fmt.Errorf("invalid subset format: %s", s)
	}
	return ds, nil
}

// subsetToKVP converts a DimensionSubset back to KVP string format.
func subsetToKVP(ds DimensionSubset) string {
	if ds.Slice != nil {
		return fmt.Sprintf("%s(%v)", ds.Dimension, ds.Slice.Value)
	}
	if ds.Trim != nil {
		return fmt.Sprintf("%s(%v,%v)", ds.Dimension, ds.Trim.Low, ds.Trim.High)
	}
	return ds.Dimension
}

// GetCoverageRequest struct
type GetCoverageRequest struct {
	XMLName          xml.Name           `xml:"GetCoverage" yaml:"getcoverage"`
	Service          string             `xml:"service,attr" yaml:"service"`
	Version          string             `xml:"version,attr" yaml:"version"`
	CoverageID       string             `xml:"CoverageId" yaml:"coverageid"`
	Format           *string            `xml:"Format,omitempty" yaml:"format"`
	DimensionSubset  []DimensionSubset  `xml:"DimensionSubset,omitempty" yaml:"dimensionsubset"`
	SubsettingCRS    *string            `xml:"subsettingCrs,attr,omitempty" yaml:"subsettingcrs"`
	OutputCRS        *string            `xml:"outputCrs,attr,omitempty" yaml:"outputcrs"`
	MediaType        *string            `xml:"mediaType,attr,omitempty" yaml:"mediatype"`
	Attr             utils.XMLAttribute `xml:",attr"`
}
