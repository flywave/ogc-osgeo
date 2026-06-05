package wmts100

import (
	"encoding/xml"
	"net/url"
	"strings"

	"github.com/flywave/ogc-osgeo/pkg/utils"
	"github.com/flywave/ogc-osgeo/pkg/wsc110"
)

// GetCapabilitiesRequest struct with the needed parameters/attributes needed for making a GetCapabilities request
type GetCapabilitiesRequest struct {
	XMLName xml.Name           `xml:"GetCapabilities" yaml:"getcapabilities"`
	Service string             `xml:"service,attr" yaml:"service"`
	Version string             `xml:"version,attr" yaml:"version"`
	Attr    utils.XMLAttribute `xml:",attr"`
}

// Type returns GetCapabilities
func (gc GetCapabilitiesRequest) Type() string {
	return getcapabilities
}

// Validate validates the GetCapabilities request
func (gc GetCapabilitiesRequest) Validate(c wsc110.Capabilities) wsc110.Exceptions {
	return nil
}

// ParseXML builds a GetCapabilities object based on a XML document
func (gc *GetCapabilitiesRequest) ParseXML(body []byte) wsc110.Exceptions {
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

// ParseQueryParameters builds a GetCapabilities object based on the available query parameters
func (gc *GetCapabilitiesRequest) ParseQueryParameters(query url.Values) wsc110.Exceptions {
	gpv := getCapabilitiesRequestParameterValue{}

	if exceptions := gpv.parseQueryParameters(query); exceptions != nil {
		return exceptions
	}

	if exceptions := gc.parseGetCapabilitiesRequestParameterValue(gpv); exceptions != nil {
		return exceptions
	}
	return nil
}

func (gc *GetCapabilitiesRequest) parseGetCapabilitiesRequestParameterValue(gpv getCapabilitiesRequestParameterValue) wsc110.Exceptions {
	gc.XMLName.Local = getcapabilities
	gc.Service = Service
	gc.Version = gpv.version
	return nil
}

// ToQueryParameters builds a new query string that will be proxied
func (gc GetCapabilitiesRequest) ToQueryParameters() url.Values {
	gpv := getCapabilitiesRequestParameterValue{}
	gpv.parseGetCapabilitiesRequest(gc)

	q := gpv.toQueryParameters()
	return q
}

// ToXML builds a 'new' XML document 'based' on the 'original' XML document
func (gc GetCapabilitiesRequest) ToXML() []byte {
	si, _ := xml.Marshal(gc)
	return append([]byte(xml.Header), si...)
}
