package wmts100

import (
	"encoding/xml"
	"net/url"
	"testing"
)

func TestGetFeatureInfoType(t *testing.T) {
	gc := GetFeatureInfoRequest{}
	if gc.Type() != getfeatureinfo {
		t.Errorf("expected: %s\n got: %s", getfeatureinfo, gc.Type())
	}
}

func TestGetFeatureInfoValidate(t *testing.T) {
	var tests = []struct {
		request    GetFeatureInfoRequest
		exceptions int
	}{
		0: {request: GetFeatureInfoRequest{Layer: "layer", TileMatrixSet: "EPSG:4326", I: "100", J: "200"}, exceptions: 0},
		1: {request: GetFeatureInfoRequest{}, exceptions: 4},
		2: {request: GetFeatureInfoRequest{Layer: "layer"}, exceptions: 3},
	}
	for k, test := range tests {
		exceptions := test.request.Validate(nil)
		if len(exceptions) != test.exceptions {
			t.Errorf("test: %d, expected: %d exceptions\n got: %d", k, test.exceptions, len(exceptions))
		}
	}
}

func TestGetFeatureInfoParseXML(t *testing.T) {
	body := []byte(`<GetFeatureInfo service="WMTS" version="1.0.0">
	<Layer>some_layer</Layer>
	<Style>some_style</Style>
	<Format>image/png</Format>
	<TileMatrixSet>EPSG:4326</TileMatrixSet>
	<TileMatrix>1</TileMatrix>
	<TileRow>0</TileRow>
	<TileCol>0</TileCol>
	<I>100</I>
	<J>200</J>
	<InfoFormat>text/plain</InfoFormat>
</GetFeatureInfo>`)

	var gc GetFeatureInfoRequest
	exceptions := gc.ParseXML(body)
	if exceptions != nil {
		t.Fatalf("unexpected exceptions: %v", exceptions)
	}
	if gc.Layer != "some_layer" {
		t.Errorf("expected Layer: some_layer\n got: %s", gc.Layer)
	}
	if gc.I != "100" {
		t.Errorf("expected I: 100\n got: %s", gc.I)
	}
	if gc.J != "200" {
		t.Errorf("expected J: 200\n got: %s", gc.J)
	}
	if gc.InfoFormat != "text/plain" {
		t.Errorf("expected InfoFormat: text/plain\n got: %s", gc.InfoFormat)
	}
}

func TestGetFeatureInfoParseXMLWrongName(t *testing.T) {
	// Verify the XMLName is GetFeatureInfo, not GetTile
	var gc GetFeatureInfoRequest
	gc.XMLName.Local = getfeatureinfo
	if gc.XMLName.Local != "GetFeatureInfo" {
		t.Errorf("expected XMLName: GetFeatureInfo\n got: %s", gc.XMLName.Local)
	}
}

func TestGetFeatureInfoParseQueryParameters(t *testing.T) {
	var tests = []struct {
		query      url.Values
		result     GetFeatureInfoRequest
		exceptions int
	}{
		0: {query: url.Values{
			REQUEST: {getfeatureinfo}, SERVICE: {Service}, VERSION: {Version},
			LAYER: {"layer"}, STYLE: {"style"}, FORMAT: {"image/png"},
			TILEMATRIXSET: {"EPSG:4326"}, TILEMATRIX: {"1"}, TILEROW: {"0"}, TILECOL: {"0"},
			I: {"100"}, J: {"200"}, INFOFORMAT: {"text/plain"},
		},
			result: GetFeatureInfoRequest{
				Layer: "layer", Style: "style", Format: "image/png",
				TileMatrixSet: "EPSG:4326", TileMatrix: "1", TileRow: "0", TileCol: "0",
				I: "100", J: "200", InfoFormat: "text/plain",
			},
			exceptions: 0,
		},
	}
	for k, test := range tests {
		var gc GetFeatureInfoRequest
		exceptions := gc.ParseQueryParameters(test.query)
		if len(exceptions) != test.exceptions {
			t.Errorf("test: %d, expected: %d exceptions\n got: %d", k, test.exceptions, len(exceptions))
		}
		if exceptions == nil {
			if gc.Layer != test.result.Layer {
				t.Errorf("test: %d, expected Layer: %s\n got: %s", k, test.result.Layer, gc.Layer)
			}
			if gc.I != test.result.I {
				t.Errorf("test: %d, expected I: %s\n got: %s", k, test.result.I, gc.I)
			}
			if gc.J != test.result.J {
				t.Errorf("test: %d, expected J: %s\n got: %s", k, test.result.J, gc.J)
			}
			if gc.InfoFormat != test.result.InfoFormat {
				t.Errorf("test: %d, expected InfoFormat: %s\n got: %s", k, test.result.InfoFormat, gc.InfoFormat)
			}
		}
	}
}

func TestGetFeatureInfoToQueryParameters(t *testing.T) {
	gc := GetFeatureInfoRequest{
		Layer: "layer", Style: "style", Format: "image/png",
		TileMatrixSet: "EPSG:4326", TileMatrix: "1", TileRow: "0", TileCol: "0",
		I: "100", J: "200", InfoFormat: "text/plain",
	}
	gc.XMLName.Local = getfeatureinfo

	q := gc.ToQueryParameters()
	if q.Get(LAYER) != "layer" {
		t.Errorf("expected: layer\n got: %s", q.Get(LAYER))
	}
	if q.Get(I) != "100" {
		t.Errorf("expected: 100\n got: %s", q.Get(I))
	}
	if q.Get(J) != "200" {
		t.Errorf("expected: 200\n got: %s", q.Get(J))
	}
	if q.Get(INFOFORMAT) != "text/plain" {
		t.Errorf("expected: text/plain\n got: %s", q.Get(INFOFORMAT))
	}
}

func TestGetFeatureInfoToXML(t *testing.T) {
	gc := GetFeatureInfoRequest{
		Service: Service, Version: Version,
		Layer: "layer", Style: "style", Format: "image/png",
		TileMatrixSet: "EPSG:4326", TileMatrix: "1", TileRow: "0", TileCol: "0",
		I: "100", J: "200", InfoFormat: "text/plain",
	}
	gc.XMLName.Local = getfeatureinfo

	body := gc.ToXML()
	var parsed GetFeatureInfoRequest
	if err := xml.Unmarshal(body, &parsed); err != nil {
		t.Fatalf("failed to unmarshal: %s", err)
	}
	// Verify the XML element name is GetFeatureInfo
	if parsed.XMLName.Local != "GetFeatureInfo" {
		t.Errorf("expected XMLName: GetFeatureInfo\n got: %s", parsed.XMLName.Local)
	}
	if parsed.Layer != "layer" {
		t.Errorf("expected Layer: layer\n got: %s", parsed.Layer)
	}
	if parsed.I != "100" {
		t.Errorf("expected I: 100\n got: %s", parsed.I)
	}
}

func TestGetFeatureInfoRESTfulParse(t *testing.T) {
	gc := GetFeatureInfoRequest{}
	err := gc.ParseRESTfulPath("EPSG:4326/1/0/0/200/100")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if gc.TileMatrixSet != "EPSG:4326" {
		t.Errorf("expected TileMatrixSet: EPSG:4326\n got: %s", gc.TileMatrixSet)
	}
	if gc.TileCol != "0" {
		t.Errorf("expected TileCol: 0\n got: %s", gc.TileCol)
	}
	if gc.J != "200" {
		t.Errorf("expected J: 200\n got: %s", gc.J)
	}
	if gc.I != "100" {
		t.Errorf("expected I: 100\n got: %s", gc.I)
	}
}

func TestGetFeatureInfoRESTfulToPath(t *testing.T) {
	gc := GetFeatureInfoRequest{
		TileMatrixSet: "EPSG:4326", TileMatrix: "1", TileRow: "0", TileCol: "0",
		J: "200", I: "100",
	}
	path := gc.ToRESTfulPath()
	expected := "EPSG:4326/1/0/0/200/100"
	if path != expected {
		t.Errorf("expected: %s\n got: %s", expected, path)
	}
}

func TestGetFeatureInfoResponse(t *testing.T) {
	r := GetFeatureInfoResponse{}
	if r.Type() != getfeatureinfo {
		t.Errorf("expected Type: %s\n got: %s", getfeatureinfo, r.Type())
	}
	if r.Service() != Service {
		t.Errorf("expected Service: %s\n got: %s", Service, r.Service())
	}
	if r.Version() != Version {
		t.Errorf("expected Version: %s\n got: %s", Version, r.Version())
	}
}
