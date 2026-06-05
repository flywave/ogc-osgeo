package wcs201

import (
	"encoding/xml"
	"net/url"
	"testing"
)

func TestDescribeCoverageType(t *testing.T) {
	d := DescribeCoverageRequest{}
	if d.Type() != describecoverage {
		t.Errorf("expected: %s\n got: %s", describecoverage, d.Type())
	}
}

func TestDescribeCoverageValidate(t *testing.T) {
	var tests = []struct {
		request    DescribeCoverageRequest
		exceptions int
	}{
		0: {request: DescribeCoverageRequest{CoverageID: []string{"id1"}}, exceptions: 0},
		1: {request: DescribeCoverageRequest{}, exceptions: 1},
	}
	for k, test := range tests {
		exceptions := test.request.Validate(Capabilities{})
		if len(exceptions) != test.exceptions {
			t.Errorf("test: %d, expected: %d exceptions\n got: %d", k, test.exceptions, len(exceptions))
		}
	}
}

func TestDescribeCoverageParseXML(t *testing.T) {
	var tests = []struct {
		body       []byte
		result     DescribeCoverageRequest
		exceptions int
	}{
		0: {body: []byte(`<DescribeCoverage service="WCS" version="2.0.1"><CoverageId>id1</CoverageId><CoverageId>id2</CoverageId></DescribeCoverage>`),
			result: DescribeCoverageRequest{
				CoverageID: []string{"id1", "id2"},
			},
			exceptions: 0,
		},
		1: {body: []byte(`not xml`),
			exceptions: 1,
		},
	}
	for k, test := range tests {
		var d DescribeCoverageRequest
		exceptions := d.ParseXML(test.body)
		if len(exceptions) != test.exceptions {
			t.Errorf("test: %d, expected: %d exceptions\n got: %d", k, test.exceptions, len(exceptions))
		}
		if exceptions == nil {
			if len(d.CoverageID) != len(test.result.CoverageID) {
				t.Errorf("test: %d, expected %d coverage IDs\n got: %d", k, len(test.result.CoverageID), len(d.CoverageID))
			}
			for i, id := range test.result.CoverageID {
				if d.CoverageID[i] != id {
					t.Errorf("test: %d, expected CoverageID[%d]: %s\n got: %s", k, i, id, d.CoverageID[i])
				}
			}
		}
	}
}

func TestDescribeCoverageQueryParameters(t *testing.T) {
	var tests = []struct {
		query      url.Values
		result     DescribeCoverageRequest
		exceptions int
	}{
		0: {query: url.Values{
			REQUEST: {describecoverage}, SERVICE: {Service}, VERSION: {Version}, COVERAGEID: {"id1,id2"},
		},
			result: DescribeCoverageRequest{
				CoverageID: []string{"id1", "id2"},
			},
			exceptions: 0,
		},
	}
	for k, test := range tests {
		var d DescribeCoverageRequest
		exceptions := d.QueryParameters(test.query)
		if len(exceptions) != test.exceptions {
			t.Errorf("test: %d, expected: %d exceptions\n got: %d", k, test.exceptions, len(exceptions))
		}
		if exceptions == nil {
			if len(d.CoverageID) != len(test.result.CoverageID) {
				t.Errorf("test: %d, expected %d coverage IDs\n got: %d", k, len(test.result.CoverageID), len(d.CoverageID))
			}
			for i, id := range test.result.CoverageID {
				if d.CoverageID[i] != id {
					t.Errorf("test: %d, expected CoverageID[%d]: %s\n got: %s", k, i, id, d.CoverageID[i])
				}
			}
		}
	}
}

func TestDescribeCoverageToQueryParameters(t *testing.T) {
	d := DescribeCoverageRequest{
		CoverageID: []string{"id1", "id2"},
	}
	d.XMLName.Local = describecoverage

	q := d.ToQueryParameters()
	if q.Get(REQUEST) != describecoverage {
		t.Errorf("expected: %s\n got: %s", describecoverage, q.Get(REQUEST))
	}
	if q.Get(COVERAGEID) != "id1,id2" {
		t.Errorf("expected: id1,id2\n got: %s", q.Get(COVERAGEID))
	}
}

func TestDescribeCoverageToXML(t *testing.T) {
	d := DescribeCoverageRequest{
		Service:    Service,
		Version:    Version,
		CoverageID: []string{"id1"},
	}
	d.XMLName.Local = describecoverage

	body := d.ToXML()
	var parsed DescribeCoverageRequest
	if err := xml.Unmarshal(body, &parsed); err != nil {
		t.Fatalf("failed to unmarshal: %s", err)
	}
	if len(parsed.CoverageID) != 1 || parsed.CoverageID[0] != "id1" {
		t.Errorf("expected CoverageID: [id1]\n got: %v", parsed.CoverageID)
	}
}

func TestCoverageDescriptionsType(t *testing.T) {
	d := CoverageDescriptions{}
	if d.Type() != describecoverage {
		t.Errorf("expected: %s\n got: %s", describecoverage, d.Type())
	}
}

func TestCoverageDescriptionsToXML(t *testing.T) {
	d := CoverageDescriptions{
		CoverageDescription: []CoverageDescription{
			{CoverageID: "id1", Title: "Test Coverage", SupportedCRS: []string{"EPSG:4326"}},
		},
	}
	body := d.ToXML()
	var parsed CoverageDescriptions
	if err := xml.Unmarshal(body, &parsed); err != nil {
		t.Fatalf("failed to unmarshal: %s", err)
	}
	if len(parsed.CoverageDescription) != 1 {
		t.Fatalf("expected 1 description\n got: %d", len(parsed.CoverageDescription))
	}
	if parsed.CoverageDescription[0].CoverageID != "id1" {
		t.Errorf("expected CoverageID: id1\n got: %s", parsed.CoverageDescription[0].CoverageID)
	}
}
