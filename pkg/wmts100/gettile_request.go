package wmts100

import (
	"encoding/xml"
	"net/url"
	"regexp"
	"strings"

	"github.com/flywave/ogc-osgeo/pkg/utils"
	"github.com/flywave/ogc-osgeo/pkg/wsc110"
)

// GetTileRequest struct with the needed parameters/attributes needed for making a GetTile request
type GetTileRequest struct {
	XMLName xml.Name           `xml:"GetTile" yaml:"gettile"`
	Service string             `xml:"service,attr" yaml:"service"`
	Version string             `xml:"version,attr" yaml:"version"`
	Attr    utils.XMLAttribute `xml:",attr"`
}

// ParseXML builds a GetTile object based on a XML document
func (gc *GetTileRequest) ParseXML(body []byte) wsc110.Exceptions {
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

// ParseQueryParameters builds a GetTile object based on the available query parameters
func (gc *GetTileRequest) ParseQueryParameters(query url.Values) wsc110.Exceptions {
	for k, v := range query {
		switch strings.ToUpper(k) {
		case REQUEST:
			if strings.EqualFold(v[0], gettile) {
				gc.XMLName.Local = gettile
			}
		case SERVICE:
			gc.Service = strings.ToUpper(v[0])
		case VERSION:
			gc.Version = strings.ToUpper(v[0])
		}
	}
	return nil
}

// ToQueryParameters builds a new query string that will be proxied
func (gc GetTileRequest) ToQueryParameters() url.Values {
	querystring := make(map[string][]string)
	querystring[REQUEST] = []string{gc.XMLName.Local}
	querystring[SERVICE] = []string{gc.Service}
	querystring[VERSION] = []string{gc.Version}

	return querystring
}

// ToXML builds a 'new' XML document 'based' on the 'original' XML document
func (gc GetTileRequest) ToXML() []byte {
	si, _ := xml.MarshalIndent(gc, "", "")
	re := regexp.MustCompile(`><.*>`)
	return []byte(xml.Header + re.ReplaceAllString(string(si), "/>"))
}
