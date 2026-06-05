package wmts100

import (
	"encoding/xml"
	"net/url"
	"strings"

	"github.com/flywave/ogc-osgeo/pkg/utils"
	"github.com/flywave/ogc-osgeo/pkg/wsc110"
)

// GetFeatureInfoRequest struct with the needed parameters/attributes needed for making a GetFeatureInfo request
type GetFeatureInfoRequest struct {
	XMLName            xml.Name           `xml:"GetFeatureInfo" yaml:"getfeatureinfo"`
	Service            string             `xml:"service,attr" yaml:"service"`
	Version            string             `xml:"version,attr" yaml:"version"`
	Layer              string             `xml:"Layer" yaml:"layer"`
	Style              string             `xml:"Style" yaml:"style"`
	Format             string             `xml:"Format" yaml:"format"`
	TileMatrixSet      string             `xml:"TileMatrixSet" yaml:"tilematrixset"`
	TileMatrix         string             `xml:"TileMatrix" yaml:"tilematrix"`
	TileRow            string             `xml:"TileRow" yaml:"tilerow"`
	TileCol            string             `xml:"TileCol" yaml:"tilecol"`
	I                  string             `xml:"I" yaml:"i"`
	J                  string             `xml:"J" yaml:"j"`
	InfoFormat         string             `xml:"InfoFormat" yaml:"infoformat"`
	DimensionNameValue []DimensionNameValue `xml:"DimensionNameValue,omitempty" yaml:"dimensionnamevalue"`
	Attr               utils.XMLAttribute `xml:",attr"`
}

// Type returns GetFeatureInfo
func (gc GetFeatureInfoRequest) Type() string {
	return getfeatureinfo
}

// Validate validates the GetFeatureInfo request
func (gc GetFeatureInfoRequest) Validate(c wsc110.Capabilities) wsc110.Exceptions {
	var exceptions wsc110.Exceptions
	if gc.Layer == "" {
		exceptions = append(exceptions, wsc110.MissingParameterValue("LAYER"))
	}
	if gc.TileMatrixSet == "" {
		exceptions = append(exceptions, wsc110.MissingParameterValue("TILEMATRIXSET"))
	}
	if gc.I == "" {
		exceptions = append(exceptions, wsc110.MissingParameterValue("I"))
	}
	if gc.J == "" {
		exceptions = append(exceptions, wsc110.MissingParameterValue("J"))
	}
	return exceptions
}

// ParseXML builds a GetFeatureInfo object based on a XML document
func (gc *GetFeatureInfoRequest) ParseXML(body []byte) wsc110.Exceptions {
	var xmlattributes utils.XMLAttribute
	if err := xml.Unmarshal(body, &xmlattributes); err != nil {
		return wsc110.Exceptions{wsc110.MissingParameterValue()}
	}
	if err := xml.Unmarshal(body, &gc); err != nil {
		return wsc110.Exceptions{wsc110.MissingParameterValue("REQUEST")}
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

	gc.Attr = utils.StripDuplicateAttr(n)
	return nil
}

// ParseQueryParameters builds a GetFeatureInfo object based on the available query parameters
func (gc *GetFeatureInfoRequest) ParseQueryParameters(query url.Values) wsc110.Exceptions {
	fpv := getFeatureInfoRequestParameterValue{}

	if exceptions := fpv.parseQueryParameters(query); exceptions != nil {
		return exceptions
	}

	if exceptions := gc.parseGetFeatureInfoRequestParameterValue(fpv); exceptions != nil {
		return exceptions
	}
	return nil
}

func (gc *GetFeatureInfoRequest) parseGetFeatureInfoRequestParameterValue(fpv getFeatureInfoRequestParameterValue) wsc110.Exceptions {
	gc.XMLName.Local = getfeatureinfo
	gc.Service = Service
	gc.Version = Version
	gc.Layer = fpv.Layer
	gc.Style = fpv.Style
	gc.Format = fpv.Format
	gc.TileMatrixSet = fpv.TileMatrixSet
	gc.TileMatrix = fpv.TileMatrix
	gc.TileRow = fpv.TileRow
	gc.TileCol = fpv.TileCol
	gc.I = fpv.I
	gc.J = fpv.J
	gc.InfoFormat = fpv.InfoFormat
	gc.DimensionNameValue = fpv.DimensionNameValue
	return nil
}

// ToQueryParameters builds a new query string that will be proxied
func (gc GetFeatureInfoRequest) ToQueryParameters() url.Values {
	fpv := getFeatureInfoRequestParameterValue{}
	fpv.parseGetFeatureInfoRequest(gc)

	q := fpv.toQueryParameters()
	return q
}

// ToXML builds a 'new' XML document 'based' on the 'original' XML document
func (gc GetFeatureInfoRequest) ToXML() []byte {
	si, _ := xml.Marshal(gc)
	return append([]byte(xml.Header), si...)
}
