package wmts100

import (
	"encoding/xml"
	"net/url"
	"strings"

	"github.com/flywave/ogc-osgeo/pkg/utils"
	"github.com/flywave/ogc-osgeo/pkg/wsc110"
)

// GetTileRequest struct with the needed parameters/attributes needed for making a GetTile request
type GetTileRequest struct {
	XMLName            xml.Name           `xml:"GetTile" yaml:"gettile"`
	Service            string             `xml:"service,attr" yaml:"service"`
	Version            string             `xml:"version,attr" yaml:"version"`
	Layer              string             `xml:"Layer" yaml:"layer"`
	Style              string             `xml:"Style" yaml:"style"`
	Format             string             `xml:"Format" yaml:"format"`
	TileMatrixSet      string             `xml:"TileMatrixSet" yaml:"tilematrixset"`
	TileMatrix         string             `xml:"TileMatrix" yaml:"tilematrix"`
	TileRow            string             `xml:"TileRow" yaml:"tilerow"`
	TileCol            string             `xml:"TileCol" yaml:"tilecol"`
	DimensionNameValue []DimensionNameValue `xml:"DimensionNameValue,omitempty" yaml:"dimensionnamevalue"`
	Attr               utils.XMLAttribute `xml:",attr"`
}

// Type returns GetTile
func (gc GetTileRequest) Type() string {
	return gettile
}

// Validate validates the GetTile request
func (gc GetTileRequest) Validate(c wsc110.Capabilities) wsc110.Exceptions {
	var exceptions wsc110.Exceptions
	if gc.Layer == "" {
		exceptions = append(exceptions, wsc110.MissingParameterValue("LAYER"))
	}
	if gc.TileMatrixSet == "" {
		exceptions = append(exceptions, wsc110.MissingParameterValue("TILEMATRIXSET"))
	}
	return exceptions
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
	tpv := getTileRequestParameterValue{}

	if exceptions := tpv.parseQueryParameters(query); exceptions != nil {
		return exceptions
	}

	if exceptions := gc.parseGetTileRequestParameterValue(tpv); exceptions != nil {
		return exceptions
	}
	return nil
}

func (gc *GetTileRequest) parseGetTileRequestParameterValue(tpv getTileRequestParameterValue) wsc110.Exceptions {
	gc.XMLName.Local = gettile
	gc.Service = Service
	gc.Version = Version
	gc.Layer = tpv.Layer
	gc.Style = tpv.Style
	gc.Format = tpv.Format
	gc.TileMatrixSet = tpv.TileMatrixSet
	gc.TileMatrix = tpv.TileMatrix
	gc.TileRow = tpv.TileRow
	gc.TileCol = tpv.TileCol
	gc.DimensionNameValue = tpv.DimensionNameValue
	return nil
}

// ToQueryParameters builds a new query string that will be proxied
func (gc GetTileRequest) ToQueryParameters() url.Values {
	tpv := getTileRequestParameterValue{}
	tpv.parseGetTileRequest(gc)

	q := tpv.toQueryParameters()
	return q
}

// ToXML builds a 'new' XML document 'based' on the 'original' XML document
func (gc GetTileRequest) ToXML() []byte {
	si, _ := xml.Marshal(gc)
	return append([]byte(xml.Header), si...)
}
