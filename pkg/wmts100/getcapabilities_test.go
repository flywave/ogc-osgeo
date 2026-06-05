package wmts100

import (
	"encoding/xml"
	"net/url"
	"testing"
)

func TestGetCapabilitiesType(t *testing.T) {
	gc := GetCapabilitiesRequest{}
	if gc.Type() != getcapabilities {
		t.Errorf("expected: %s\n got: %s", getcapabilities, gc.Type())
	}
}

func TestGetCapabilitiesValidate(t *testing.T) {
	var tests = []struct {
		request    GetCapabilitiesRequest
		exceptions int
	}{
		0: {request: GetCapabilitiesRequest{}, exceptions: 0},
	}
	for k, test := range tests {
		exceptions := test.request.Validate(nil)
		if len(exceptions) != test.exceptions {
			t.Errorf("test: %d, expected: %d exceptions\n got: %d", k, test.exceptions, len(exceptions))
		}
	}
}

func TestGetCapabilitiesParseXML(t *testing.T) {
	var tests = []struct {
		body       []byte
		result     GetCapabilitiesRequest
		exceptions int
	}{
		0: {body: []byte(`<GetCapabilities service="WMTS" version="1.0.0"></GetCapabilities>`),
			result:     GetCapabilitiesRequest{Service: "WMTS", Version: "1.0.0"},
			exceptions: 0,
		},
		1: {body: []byte(`not xml`),
			exceptions: 1,
		},
	}
	for k, test := range tests {
		var gc GetCapabilitiesRequest
		exceptions := gc.ParseXML(test.body)
		if len(exceptions) != test.exceptions {
			t.Errorf("test: %d, expected: %d exceptions\n got: %d", k, test.exceptions, len(exceptions))
		}
		if exceptions == nil {
			if gc.Service != test.result.Service {
				t.Errorf("test: %d, expected: %s\n got: %s", k, test.result.Service, gc.Service)
			}
		}
	}
}

func TestGetCapabilitiesParseQueryParameters(t *testing.T) {
	var tests = []struct {
		query      url.Values
		result     GetCapabilitiesRequest
		exceptions int
	}{
		0: {query: url.Values{
			REQUEST: {getcapabilities}, SERVICE: {Service}, VERSION: {Version},
		},
			result:     GetCapabilitiesRequest{Service: Service, Version: Version},
			exceptions: 0,
		},
	}
	for k, test := range tests {
		var gc GetCapabilitiesRequest
		exceptions := gc.ParseQueryParameters(test.query)
		if len(exceptions) != test.exceptions {
			t.Errorf("test: %d, expected: %d exceptions\n got: %d", k, test.exceptions, len(exceptions))
		}
		if exceptions == nil {
			if gc.Version != test.result.Version {
				t.Errorf("test: %d, expected: %s\n got: %s", k, test.result.Version, gc.Version)
			}
		}
	}
}

func TestGetCapabilitiesToQueryParameters(t *testing.T) {
	gc := GetCapabilitiesRequest{
		Service: Service,
		Version: Version,
	}
	gc.XMLName.Local = getcapabilities

	q := gc.ToQueryParameters()
	if q.Get(REQUEST) != getcapabilities {
		t.Errorf("expected: %s\n got: %s", getcapabilities, q.Get(REQUEST))
	}
	if q.Get(SERVICE) != Service {
		t.Errorf("expected: %s\n got: %s", Service, q.Get(SERVICE))
	}
	if q.Get(VERSION) != Version {
		t.Errorf("expected: %s\n got: %s", Version, q.Get(VERSION))
	}
}

func TestGetCapabilitiesToXML(t *testing.T) {
	gc := GetCapabilitiesRequest{
		Service: Service,
		Version: Version,
	}
	gc.XMLName.Local = getcapabilities

	body := gc.ToXML()
	var parsed GetCapabilitiesRequest
	if err := xml.Unmarshal(body, &parsed); err != nil {
		t.Fatalf("failed to unmarshal: %s", err)
	}
	if parsed.Service != Service {
		t.Errorf("expected Service: %s\n got: %s", Service, parsed.Service)
	}
}

func TestGetCapabilitiesResponse(t *testing.T) {
	r := GetCapabilitiesResponse{}
	if r.Type() != getcapabilities {
		t.Errorf("expected: %s\n got: %s", getcapabilities, r.Type())
	}
	if r.Service() != Service {
		t.Errorf("expected Service: %s\n got: %s", Service, r.Service())
	}
	if r.Version() != Version {
		t.Errorf("expected Version: %s\n got: %s", Version, r.Version())
	}
}

func TestGetCapabilitiesResponseToXML(t *testing.T) {
	r := GetCapabilitiesResponse{
		Namespaces: Namespaces{
			Xmlns: "http://www.opengis.net/wmts/1.0",
		},
	}
	body := r.ToXML()
	if len(body) == 0 {
		t.Fatal("expected non-empty XML")
	}
}
