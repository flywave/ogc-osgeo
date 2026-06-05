package wmts100

import (
	"strconv"
	"strings"
)

// ParseRESTfulPath parses a RESTful WMTS tile URL path into a GetTileRequest.
// Template: /{TileMatrixSet}/{TileMatrix}/{TileRow}/{TileCol}.{Format}
func (gc *GetTileRequest) ParseRESTfulPath(path string) error {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) < 4 {
		return nil
	}

	gc.TileMatrixSet = parts[0]
	gc.TileMatrix = parts[1]

	if _, err := strconv.Atoi(parts[2]); err == nil {
		gc.TileRow = parts[2]
	}

	// TileCol may carry optional format extension: "0.png"
	colParts := strings.SplitN(parts[3], ".", 2)
	if _, err := strconv.Atoi(colParts[0]); err == nil {
		gc.TileCol = colParts[0]
	}
	if len(colParts) > 1 {
		gc.Format = colParts[1]
	}

	return nil
}

// ToRESTfulPath builds a RESTful WMTS tile URL path from a GetTileRequest.
func (gc GetTileRequest) ToRESTfulPath() string {
	colPart := gc.TileCol
	if gc.Format != "" {
		colPart = gc.TileCol + "." + gc.Format
	}
	return strings.Join([]string{gc.TileMatrixSet, gc.TileMatrix, gc.TileRow, colPart}, "/")
}

// ParseRESTfulPath parses a RESTful WMTS feature info URL path into a GetFeatureInfoRequest.
// Template: /{TileMatrixSet}/{TileMatrix}/{TileRow}/{TileCol}/{J}/{I}
func (gc *GetFeatureInfoRequest) ParseRESTfulPath(path string) error {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) < 6 {
		return nil
	}

	gc.TileMatrixSet = parts[0]
	gc.TileMatrix = parts[1]

	if _, err := strconv.Atoi(parts[2]); err == nil {
		gc.TileRow = parts[2]
	}
	if _, err := strconv.Atoi(parts[3]); err == nil {
		gc.TileCol = parts[3]
	}

	gc.J = parts[4]
	gc.I = parts[5]

	return nil
}

// ToRESTfulPath builds a RESTful WMTS feature info URL path from a GetFeatureInfoRequest.
func (gc GetFeatureInfoRequest) ToRESTfulPath() string {
	return strings.Join([]string{gc.TileMatrixSet, gc.TileMatrix, gc.TileRow, gc.TileCol, gc.J, gc.I}, "/")
}
