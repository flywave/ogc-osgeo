package wmts100

import (
	"encoding/xml"
	"net/url"
	"testing"
)

func TestGetTileType(t *testing.T) {
	gc := GetTileRequest{}
	if gc.Type() != gettile {
		t.Errorf("expected: %s\n got: %s", gettile, gc.Type())
	}
}

func TestGetTileValidate(t *testing.T) {
	var tests = []struct {
		request    GetTileRequest
		exceptions int
	}{
		0: {request: GetTileRequest{Layer: "layer", TileMatrixSet: "EPSG:4326"}, exceptions: 0},
		1: {request: GetTileRequest{}, exceptions: 2},
		2: {request: GetTileRequest{Layer: "layer"}, exceptions: 1},
	}
	for k, test := range tests {
		exceptions := test.request.Validate(nil)
		if len(exceptions) != test.exceptions {
			t.Errorf("test: %d, expected: %d exceptions\n got: %d", k, test.exceptions, len(exceptions))
		}
	}
}

func TestGetTileParseXML(t *testing.T) {
	body := []byte(`<GetTile service="WMTS" version="1.0.0">
	<Layer>some_layer</Layer>
	<Style>some_style</Style>
	<Format>image/png</Format>
	<TileMatrixSet>EPSG:4326</TileMatrixSet>
	<TileMatrix>1</TileMatrix>
	<TileRow>0</TileRow>
	<TileCol>0</TileCol>
</GetTile>`)

	var gc GetTileRequest
	exceptions := gc.ParseXML(body)
	if exceptions != nil {
		t.Fatalf("unexpected exceptions: %v", exceptions)
	}
	if gc.Service != Service {
		t.Errorf("expected Service: %s\n got: %s", Service, gc.Service)
	}
	if gc.Layer != "some_layer" {
		t.Errorf("expected Layer: some_layer\n got: %s", gc.Layer)
	}
	if gc.Style != "some_style" {
		t.Errorf("expected Style: some_style\n got: %s", gc.Style)
	}
	if gc.Format != "image/png" {
		t.Errorf("expected Format: image/png\n got: %s", gc.Format)
	}
	if gc.TileMatrixSet != "EPSG:4326" {
		t.Errorf("expected TileMatrixSet: EPSG:4326\n got: %s", gc.TileMatrixSet)
	}
	if gc.TileMatrix != "1" {
		t.Errorf("expected TileMatrix: 1\n got: %s", gc.TileMatrix)
	}
	if gc.TileRow != "0" {
		t.Errorf("expected TileRow: 0\n got: %s", gc.TileRow)
	}
	if gc.TileCol != "0" {
		t.Errorf("expected TileCol: 0\n got: %s", gc.TileCol)
	}
}

func TestGetTileParseQueryParameters(t *testing.T) {
	var tests = []struct {
		query      url.Values
		result     GetTileRequest
		exceptions int
	}{
		0: {query: url.Values{
			REQUEST: {gettile}, SERVICE: {Service}, VERSION: {Version},
			LAYER: {"some_layer"}, STYLE: {"some_style"}, FORMAT: {"image/png"},
			TILEMATRIXSET: {"EPSG:4326"}, TILEMATRIX: {"1"}, TILEROW: {"0"}, TILECOL: {"0"},
		},
			result: GetTileRequest{
				Layer: "some_layer", Style: "some_style", Format: "image/png",
				TileMatrixSet: "EPSG:4326", TileMatrix: "1", TileRow: "0", TileCol: "0",
			},
			exceptions: 0,
		},
		1: {query: url.Values{
			REQUEST: {gettile}, SERVICE: {Service}, VERSION: {Version},
			LAYER: {"layer"}, STYLE: {"style"}, FORMAT: {"image/png"},
			TILEMATRIXSET: {"EPSG:4326"}, TILEMATRIX: {"1"}, TILEROW: {"0"}, TILECOL: {"0"},
		},
			result: GetTileRequest{
				Layer: "layer", Style: "style", Format: "image/png",
				TileMatrixSet: "EPSG:4326", TileMatrix: "1", TileRow: "0", TileCol: "0",
			},
			exceptions: 0,
		},
	}
	for k, test := range tests {
		var gc GetTileRequest
		exceptions := gc.ParseQueryParameters(test.query)
		if len(exceptions) != test.exceptions {
			t.Errorf("test: %d, expected: %d exceptions\n got: %d", k, test.exceptions, len(exceptions))
		}
		if exceptions == nil {
			if gc.Layer != test.result.Layer {
				t.Errorf("test: %d, expected Layer: %s\n got: %s", k, test.result.Layer, gc.Layer)
			}
			if gc.TileMatrixSet != test.result.TileMatrixSet {
				t.Errorf("test: %d, expected TileMatrixSet: %s\n got: %s", k, test.result.TileMatrixSet, gc.TileMatrixSet)
			}
			if gc.TileRow != test.result.TileRow {
				t.Errorf("test: %d, expected TileRow: %s\n got: %s", k, test.result.TileRow, gc.TileRow)
			}
			if gc.TileCol != test.result.TileCol {
				t.Errorf("test: %d, expected TileCol: %s\n got: %s", k, test.result.TileCol, gc.TileCol)
			}
		}
	}
}

func TestGetTileToQueryParameters(t *testing.T) {
	gc := GetTileRequest{
		Layer: "layer", Style: "style", Format: "image/png",
		TileMatrixSet: "EPSG:4326", TileMatrix: "1", TileRow: "0", TileCol: "0",
	}
	gc.XMLName.Local = gettile

	q := gc.ToQueryParameters()
	if q.Get(LAYER) != "layer" {
		t.Errorf("expected: layer\n got: %s", q.Get(LAYER))
	}
	if q.Get(FORMAT) != "image/png" {
		t.Errorf("expected: image/png\n got: %s", q.Get(FORMAT))
	}
	if q.Get(TILEMATRIXSET) != "EPSG:4326" {
		t.Errorf("expected: EPSG:4326\n got: %s", q.Get(TILEMATRIXSET))
	}
}

func TestGetTileToXML(t *testing.T) {
	gc := GetTileRequest{
		Service: Service, Version: Version,
		Layer: "layer", Style: "style", Format: "image/png",
		TileMatrixSet: "EPSG:4326", TileMatrix: "1", TileRow: "0", TileCol: "0",
	}
	gc.XMLName.Local = gettile

	body := gc.ToXML()
	var parsed GetTileRequest
	if err := xml.Unmarshal(body, &parsed); err != nil {
		t.Fatalf("failed to unmarshal: %s", err)
	}
	if parsed.Layer != "layer" {
		t.Errorf("expected Layer: layer\n got: %s", parsed.Layer)
	}
	if parsed.Format != "image/png" {
		t.Errorf("expected Format: image/png\n got: %s", parsed.Format)
	}
}

func TestGetTileRESTfulParse(t *testing.T) {
	gc := GetTileRequest{}
	err := gc.ParseRESTfulPath("EPSG:4326/1/0/0.png")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if gc.TileMatrixSet != "EPSG:4326" {
		t.Errorf("expected TileMatrixSet: EPSG:4326\n got: %s", gc.TileMatrixSet)
	}
	if gc.TileMatrix != "1" {
		t.Errorf("expected TileMatrix: 1\n got: %s", gc.TileMatrix)
	}
	if gc.TileRow != "0" {
		t.Errorf("expected TileRow: 0\n got: %s", gc.TileRow)
	}
	if gc.TileCol != "0" {
		t.Errorf("expected TileCol: 0\n got: %s", gc.TileCol)
	}
	if gc.Format != "png" {
		t.Errorf("expected Format: png\n got: %s", gc.Format)
	}
}

func TestGetTileRESTfulToPath(t *testing.T) {
	gc := GetTileRequest{
		TileMatrixSet: "EPSG:4326", TileMatrix: "1", TileRow: "0", TileCol: "0", Format: "png",
	}
	path := gc.ToRESTfulPath()
	expected := "EPSG:4326/1/0/0.png"
	if path != expected {
		t.Errorf("expected: %s\n got: %s", expected, path)
	}
}

func TestGetTileRESTfulNoFormat(t *testing.T) {
	gc := GetTileRequest{
		TileMatrixSet: "EPSG:4326", TileMatrix: "1", TileRow: "0", TileCol: "0",
	}
	path := gc.ToRESTfulPath()
	expected := "EPSG:4326/1/0/0"
	if path != expected {
		t.Errorf("expected: %s\n got: %s", expected, path)
	}
}

func TestGetTileResponse(t *testing.T) {
	r := GetTileResponse{}
	if r.Type() != gettile {
		t.Errorf("expected Type: %s\n got: %s", gettile, r.Type())
	}
	if r.Service() != Service {
		t.Errorf("expected Service: %s\n got: %s", Service, r.Service())
	}
	if r.Version() != Version {
		t.Errorf("expected Version: %s\n got: %s", Version, r.Version())
	}
}
