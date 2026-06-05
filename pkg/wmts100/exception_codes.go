package wmts100

import (
	"fmt"

	"github.com/flywave/ogc-osgeo/pkg/wsc110"
)

// TileOutOfRange exception
func TileOutOfRange(locator string) wsc110.Exception {
	text := "TileRow or TileCol out of range"
	if locator != "" {
		text = fmt.Sprintf("TileRow or TileCol out of range: %s", locator)
	}
	return exception{
		ExceptionCode: "TileOutOfRange",
		ExceptionText: text,
		LocatorCode:   locator,
	}
}

// InvalidTileMatrixSet exception
func InvalidTileMatrixSet(value string) wsc110.Exception {
	return exception{
		ExceptionCode: "InvalidTileMatrixSet",
		ExceptionText: fmt.Sprintf("Invalid TileMatrixSet: %s", value),
		LocatorCode:   value,
	}
}

// LayerNotDefined exception
func LayerNotDefined(layer string) wsc110.Exception {
	return exception{
		ExceptionCode: "LayerNotDefined",
		ExceptionText: fmt.Sprintf("Layer not defined: %s", layer),
		LocatorCode:   layer,
	}
}
