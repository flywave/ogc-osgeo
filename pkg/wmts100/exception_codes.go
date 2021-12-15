package wmts100

import "github.com/flywave/ogc-osgeo/pkg/wsc110"

// TileOutOfRange exception
func TileOutOfRange() wsc110.Exception {
	return exception{
		ExceptionCode: "TileOutOfRange",
		ExceptionText: "TileRow or TileCol out of rangeName",
		LocatorCode:   "", // TODO parse the right parameter TileRow or TileCol
	}
}
