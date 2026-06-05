package wcs201

import (
	"encoding/xml"
	"net/url"
	"testing"
)

func sp(s string) *string { return &s }

func TestGetCoverageType(t *testing.T) {
	g := GetCoverageRequest{}
	if g.Type() != getcoverage {
		t.Errorf("expected: %s\n got: %s", getcoverage, g.Type())
	}
}

func TestGetCoverageValidate(t *testing.T) {
	var tests = []struct {
		request    GetCoverageRequest
		exceptions int
	}{
		0: {request: GetCoverageRequest{CoverageID: "test"}, exceptions: 0},
		1: {request: GetCoverageRequest{}, exceptions: 1},
	}
	for k, test := range tests {
		exceptions := test.request.Validate(Capabilities{})
		if len(exceptions) != test.exceptions {
			t.Errorf("test: %d, expected: %d exceptions\n got: %d", k, test.exceptions, len(exceptions))
		}
	}
}

func TestGetCoverageParseXML(t *testing.T) {
	var tests = []struct {
		body       []byte
		result     GetCoverageRequest
		exceptions int
	}{
		0: {body: []byte(`<GetCoverage service="WCS" version="2.0.1"><CoverageId>test</CoverageId></GetCoverage>`),
			result:     GetCoverageRequest{CoverageID: "test"},
			exceptions: 0,
		},
		1: {body: []byte(`not xml`),
			exceptions: 1,
		},
	}
	for k, test := range tests {
		var g GetCoverageRequest
		exceptions := g.ParseXML(test.body)
		if len(exceptions) != test.exceptions {
			t.Errorf("test: %d, expected: %d exceptions\n got: %d", k, test.exceptions, len(exceptions))
		}
		if exceptions == nil {
			if g.CoverageID != test.result.CoverageID {
				t.Errorf("test: %d, expected: %s\n got: %s", k, test.result.CoverageID, g.CoverageID)
			}
		}
	}
}

func TestGetCoverageParseXMLWithSubset(t *testing.T) {
	body := []byte(`<GetCoverage service="WCS" version="2.0.1">
	<CoverageId>test</CoverageId>
	<DimensionSubset>
		<Dimension>X</Dimension>
		<Trim>
			<Low>0</Low>
			<High>100</High>
		</Trim>
	</DimensionSubset>
	<Format>image/tiff</Format>
</GetCoverage>`)

	var g GetCoverageRequest
	exceptions := g.ParseXML(body)
	if exceptions != nil {
		t.Fatalf("unexpected exceptions: %v", exceptions)
	}
	if g.CoverageID != "test" {
		t.Errorf("expected CoverageID: test\n got: %s", g.CoverageID)
	}
	if g.Format == nil || *g.Format != "image/tiff" {
		t.Errorf("expected Format: image/tiff\n got: %v", g.Format)
	}
	if len(g.DimensionSubset) != 1 {
		t.Fatalf("expected 1 DimensionSubset\n got: %d", len(g.DimensionSubset))
	}
	if g.DimensionSubset[0].Dimension != "X" {
		t.Errorf("expected Dimension: X\n got: %s", g.DimensionSubset[0].Dimension)
	}
	if g.DimensionSubset[0].Trim == nil {
		t.Fatal("expected Trim, got nil")
	}
	if g.DimensionSubset[0].Trim.Low != 0 {
		t.Errorf("expected Trim.Low: 0\n got: %v", g.DimensionSubset[0].Trim.Low)
	}
	if g.DimensionSubset[0].Trim.High != 100 {
		t.Errorf("expected Trim.High: 100\n got: %v", g.DimensionSubset[0].Trim.High)
	}
}

func TestGetCoverageQueryParameters(t *testing.T) {
	var tests = []struct {
		query      url.Values
		result     GetCoverageRequest
		exceptions int
	}{
		0: {query: url.Values{
			REQUEST: {getcoverage}, SERVICE: {Service}, VERSION: {Version}, COVERAGEID: {"test"},
		},
			result:     GetCoverageRequest{CoverageID: "test"},
			exceptions: 0,
		},
		1: {query: url.Values{
			REQUEST: {getcoverage}, SERVICE: {Service}, VERSION: {Version}, COVERAGEID: {"test"},
			FORMAT: {"image/tiff"}, SUBSET: {"X(0,100)"}, SUBSETTINGCRS: {"EPSG:4326"},
		},
			result: GetCoverageRequest{
				CoverageID: "test", Format: sp("image/tiff"),
				DimensionSubset: []DimensionSubset{{Dimension: "X", Trim: &Trim{Low: 0, High: 100}}},
				SubsettingCRS:   sp("EPSG:4326"),
			},
			exceptions: 0,
		},
	}
	for k, test := range tests {
		var g GetCoverageRequest
		exceptions := g.QueryParameters(test.query)
		if len(exceptions) != test.exceptions {
			t.Errorf("test: %d, expected: %d exceptions\n got: %d", k, test.exceptions, len(exceptions))
		}
		if exceptions == nil {
			if g.CoverageID != test.result.CoverageID {
				t.Errorf("test: %d, expected CoverageID: %s\n got: %s", k, test.result.CoverageID, g.CoverageID)
			}
			if test.result.Format != nil {
				if g.Format == nil || *g.Format != *test.result.Format {
					t.Errorf("test: %d, expected Format: %v\n got: %v", k, *test.result.Format, g.Format)
				}
			}
			if test.result.SubsettingCRS != nil {
				if g.SubsettingCRS == nil || *g.SubsettingCRS != *test.result.SubsettingCRS {
					t.Errorf("test: %d, expected SubsettingCRS: %v\n got: %v", k, *test.result.SubsettingCRS, g.SubsettingCRS)
				}
			}
			if len(test.result.DimensionSubset) > 0 && len(g.DimensionSubset) > 0 {
				if g.DimensionSubset[0].Dimension != test.result.DimensionSubset[0].Dimension {
					t.Errorf("test: %d, expected Dimension: %s\n got: %s", k, test.result.DimensionSubset[0].Dimension, g.DimensionSubset[0].Dimension)
				}
			}
		}
	}
}

func TestGetCoverageToQueryParameters(t *testing.T) {
	g := GetCoverageRequest{
		CoverageID: "test",
		Format:     sp("image/tiff"),
		DimensionSubset: []DimensionSubset{
			{Dimension: "X", Trim: &Trim{Low: 0, High: 100}},
		},
	}
	g.XMLName.Local = getcoverage

	q := g.ToQueryParameters()
	if q.Get(REQUEST) != getcoverage {
		t.Errorf("expected: %s\n got: %s", getcoverage, q.Get(REQUEST))
	}
	if q.Get(COVERAGEID) != "test" {
		t.Errorf("expected: test\n got: %s", q.Get(COVERAGEID))
	}
	if q.Get(FORMAT) != "image/tiff" {
		t.Errorf("expected: image/tiff\n got: %s", q.Get(FORMAT))
	}
	if q.Get(SUBSET) != "X(0,100)" {
		t.Errorf("expected: X(0,100)\n got: %s", q.Get(SUBSET))
	}
}

func TestGetCoverageToXML(t *testing.T) {
	g := GetCoverageRequest{
		Service:    Service,
		Version:    Version,
		CoverageID: "test",
		Format:     sp("image/tiff"),
	}
	g.XMLName.Local = getcoverage

	body := g.ToXML()
	var parsed GetCoverageRequest
	if err := xml.Unmarshal(body, &parsed); err != nil {
		t.Fatalf("failed to unmarshal: %s", err)
	}
	if parsed.CoverageID != "test" {
		t.Errorf("expected CoverageID: test\n got: %s", parsed.CoverageID)
	}
}

func TestGetCoverageToXMLWithDimensionSubset(t *testing.T) {
	g := GetCoverageRequest{
		Service:    Service,
		Version:    Version,
		CoverageID: "test",
		DimensionSubset: []DimensionSubset{
			{Dimension: "X", Trim: &Trim{Low: 0, High: 100}},
			{Dimension: "Y", Slice: &Slice{Value: 50}},
		},
	}
	g.XMLName.Local = getcoverage

	body := g.ToXML()
	var parsed GetCoverageRequest
	if err := xml.Unmarshal(body, &parsed); err != nil {
		t.Fatalf("failed to unmarshal: %s", err)
	}
	if len(parsed.DimensionSubset) != 2 {
		t.Fatalf("expected 2 DimensionSubset\n got: %d", len(parsed.DimensionSubset))
	}
	if parsed.DimensionSubset[0].Dimension != "X" {
		t.Errorf("expected Dimension[0]: X\n got: %s", parsed.DimensionSubset[0].Dimension)
	}
	if parsed.DimensionSubset[0].Trim == nil {
		t.Fatal("expected Trim on [0], got nil")
	}
	if parsed.DimensionSubset[0].Trim.Low != 0 || parsed.DimensionSubset[0].Trim.High != 100 {
		t.Errorf("expected Trim: (0,100)\n got: (%v,%v)", parsed.DimensionSubset[0].Trim.Low, parsed.DimensionSubset[0].Trim.High)
	}
	if parsed.DimensionSubset[1].Slice == nil {
		t.Fatal("expected Slice on [1], got nil")
	}
	if parsed.DimensionSubset[1].Slice.Value != 50 {
		t.Errorf("expected Slice.Value: 50\n got: %v", parsed.DimensionSubset[1].Slice.Value)
	}
}

func TestGetCoverageResponseType(t *testing.T) {
	g := GetCoverageResponse{}
	if g.Type() != getcoverage {
		t.Errorf("expected: %s\n got: %s", getcoverage, g.Type())
	}
}

func TestGetCoverageResponseService(t *testing.T) {
	g := GetCoverageResponse{}
	if g.Service() != Service {
		t.Errorf("expected: %s\n got: %s", Service, g.Service())
	}
}

func TestGetCoverageResponseVersion(t *testing.T) {
	g := GetCoverageResponse{}
	if g.Version() != Version {
		t.Errorf("expected: %s\n got: %s", Version, g.Version())
	}
}

func TestParseSubsetString(t *testing.T) {
	var tests = []struct {
		input    string
		valid    bool
		dimension string
		hasTrim  bool
		hasSlice bool
		low      float64
		high     float64
		value    float64
	}{
		0: {input: "X(0,100)", valid: true, dimension: "X", hasTrim: true, low: 0, high: 100},
		1: {input: "Y(50)", valid: true, dimension: "Y", hasSlice: true, value: 50},
		2: {input: "Time(2020-01-01,2020-12-31)", valid: false},
		3: {input: "X()", valid: false},
		4: {input: "invalid", valid: false},
		5: {input: "X(1,2,3)", valid: false},
	}
	for k, test := range tests {
		ds, err := parseSubsetString(test.input)
		if test.valid {
			if err != nil {
				t.Errorf("test: %d, unexpected error: %s", k, err)
				continue
			}
			if ds.Dimension != test.dimension {
				t.Errorf("test: %d, expected Dimension: %s\n got: %s", k, test.dimension, ds.Dimension)
			}
			if test.hasTrim {
				if ds.Trim == nil {
					t.Errorf("test: %d, expected Trim, got nil", k)
				} else if ds.Trim.Low != test.low || ds.Trim.High != test.high {
					t.Errorf("test: %d, expected Trim: (%v,%v)\n got: (%v,%v)", k, test.low, test.high, ds.Trim.Low, ds.Trim.High)
				}
			}
			if test.hasSlice {
				if ds.Slice == nil {
					t.Errorf("test: %d, expected Slice, got nil", k)
				} else if ds.Slice.Value != test.value {
					t.Errorf("test: %d, expected Slice.Value: %v\n got: %v", k, test.value, ds.Slice.Value)
				}
			}
		} else {
			if err == nil {
				t.Errorf("test: %d, expected error, got nil", k)
			}
		}
	}
}

func TestSubsetToKVP(t *testing.T) {
	var tests = []struct {
		subset   DimensionSubset
		expected string
	}{
		0: {subset: DimensionSubset{Dimension: "X", Trim: &Trim{Low: 0, High: 100}}, expected: "X(0,100)"},
		1: {subset: DimensionSubset{Dimension: "Y", Slice: &Slice{Value: 50}}, expected: "Y(50)"},
	}
	for k, test := range tests {
		result := subsetToKVP(test.subset)
		if result != test.expected {
			t.Errorf("test: %d, expected: %s\n got: %s", k, test.expected, result)
		}
	}
}
