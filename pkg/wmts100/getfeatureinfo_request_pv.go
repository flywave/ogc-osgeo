package wmts100

import (
	"net/url"
	"strings"

	"github.com/flywave/ogc-osgeo/pkg/wsc110"
)

type getFeatureInfoRequestParameterValue struct {
	Layer              string               `yaml:"layer"`
	Style              string               `yaml:"style"`
	Format             string               `yaml:"format"`
	TileMatrixSet      string               `yaml:"tilematrixset"`
	DimensionNameValue []DimensionNameValue `yaml:"dimensionnamevalue"`
	TileMatrix         string               `yaml:"tilematrix"`
	TileRow            string               `yaml:"tilerow"`
	TileCol            string               `yaml:"tilecol"`
	I                  string               `yaml:"i"`
	J                  string               `yaml:"j"`
	InfoFormat         string               `yaml:"infoformat"`
}

func (fpv *getFeatureInfoRequestParameterValue) parseQueryParameters(query url.Values) wsc110.Exceptions {
	var exceptions wsc110.Exceptions
	for k, v := range query {
		if len(v) != 1 {
			exceptions = append(exceptions, wsc110.InvalidParameterValue(k, strings.Join(v, ",")))
		} else {
			switch strings.ToUpper(k) {
			case SERVICE:
			case VERSION:
			case REQUEST:
			case LAYER:
				fpv.Layer = v[0]
			case STYLE:
				fpv.Style = v[0]
			case FORMAT:
				fpv.Format = v[0]
			case TILEMATRIXSET:
				fpv.TileMatrixSet = v[0]
			case TILEMATRIX:
				fpv.TileMatrix = v[0]
			case TILEROW:
				fpv.TileRow = v[0]
			case TILECOL:
				fpv.TileCol = v[0]
			case I:
				fpv.I = v[0]
			case J:
				fpv.J = v[0]
			case INFOFORMAT:
				fpv.InfoFormat = v[0]
			}
		}
	}
	if len(exceptions) > 0 {
		return exceptions
	}
	return nil
}

func (fpv *getFeatureInfoRequestParameterValue) parseGetFeatureInfoRequest(gf GetFeatureInfoRequest) {
	fpv.Layer = gf.Layer
	fpv.Style = gf.Style
	fpv.Format = gf.Format
	fpv.TileMatrixSet = gf.TileMatrixSet
	fpv.TileMatrix = gf.TileMatrix
	fpv.TileRow = gf.TileRow
	fpv.TileCol = gf.TileCol
	fpv.I = gf.I
	fpv.J = gf.J
	fpv.InfoFormat = gf.InfoFormat
	fpv.DimensionNameValue = gf.DimensionNameValue
}

func (fpv getFeatureInfoRequestParameterValue) toQueryParameters() url.Values {
	querystring := make(map[string][]string)
	querystring[LAYER] = []string{fpv.Layer}
	querystring[STYLE] = []string{fpv.Style}
	querystring[FORMAT] = []string{fpv.Format}
	querystring[TILEMATRIXSET] = []string{fpv.TileMatrixSet}
	querystring[TILEMATRIX] = []string{fpv.TileMatrix}
	querystring[TILEROW] = []string{fpv.TileRow}
	querystring[TILECOL] = []string{fpv.TileCol}
	querystring[I] = []string{fpv.I}
	querystring[J] = []string{fpv.J}
	querystring[INFOFORMAT] = []string{fpv.InfoFormat}
	return querystring
}
