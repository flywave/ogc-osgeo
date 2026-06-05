package wmts100

import (
	"net/url"
	"strings"

	"github.com/flywave/ogc-osgeo/pkg/wsc110"
)

type getTileRequestParameterValue struct {
	Layer              string               `yaml:"layer"`
	Style              string               `yaml:"style"`
	Format             string               `yaml:"format"`
	TileMatrixSet      string               `yaml:"tilematrixset"`
	DimensionNameValue []DimensionNameValue `yaml:"dimensionnamevalue"`
	TileMatrix         string               `yaml:"tilematrix"`
	TileRow            string               `yaml:"tilerow"`
	TileCol            string               `yaml:"tilecol"`
}

func (tpv *getTileRequestParameterValue) parseQueryParameters(query url.Values) wsc110.Exceptions {
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
				tpv.Layer = v[0]
			case STYLE:
				tpv.Style = v[0]
			case FORMAT:
				tpv.Format = v[0]
			case TILEMATRIXSET:
				tpv.TileMatrixSet = v[0]
			case TILEMATRIX:
				tpv.TileMatrix = v[0]
			case TILEROW:
				tpv.TileRow = v[0]
			case TILECOL:
				tpv.TileCol = v[0]
			default:
				// Custom dimensions are passed as DIM_NAME=VALUE
			}
		}
	}
	if len(exceptions) > 0 {
		return exceptions
	}
	return nil
}

func (tpv *getTileRequestParameterValue) parseGetTileRequest(gt GetTileRequest) {
	tpv.Layer = gt.Layer
	tpv.Style = gt.Style
	tpv.Format = gt.Format
	tpv.TileMatrixSet = gt.TileMatrixSet
	tpv.TileMatrix = gt.TileMatrix
	tpv.TileRow = gt.TileRow
	tpv.TileCol = gt.TileCol
	tpv.DimensionNameValue = gt.DimensionNameValue
}

func (tpv getTileRequestParameterValue) toQueryParameters() url.Values {
	querystring := make(map[string][]string)
	querystring[LAYER] = []string{tpv.Layer}
	querystring[STYLE] = []string{tpv.Style}
	querystring[FORMAT] = []string{tpv.Format}
	querystring[TILEMATRIXSET] = []string{tpv.TileMatrixSet}
	querystring[TILEMATRIX] = []string{tpv.TileMatrix}
	querystring[TILEROW] = []string{tpv.TileRow}
	querystring[TILECOL] = []string{tpv.TileCol}
	return querystring
}
