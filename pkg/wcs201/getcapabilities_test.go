package wcs201

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
		request   GetCapabilitiesRequest
		exceptions int
	}{
		0: {request: GetCapabilitiesRequest{}, exceptions: 0},
	}
	for k, test := range tests {
		exceptions := test.request.Validate(Capabilities{})
		if len(exceptions) != test.exceptions {
			t.Errorf("test: %d, expected: %d exceptions\n got: %d", k, test.exceptions, len(exceptions))
		}
	}
}

func TestGetCapabilitiesParseXML(t *testing.T) {
	var tests = []struct {
		body      []byte
		result    GetCapabilitiesRequest
		exceptions int
	}{
		0: {body: []byte(`<GetCapabilities service="WCS" version="2.0.1"></GetCapabilities>`),
			result: GetCapabilitiesRequest{
				XMLName: xml.Name{Local: "GetCapabilities"},
				Service: "WCS",
				Version: "2.0.1",
			},
			exceptions: 0,
		},
		1: {body: []byte(`not xml`),
			result:    GetCapabilitiesRequest{},
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
			if gc.Version != test.result.Version {
				t.Errorf("test: %d, expected: %s\n got: %s", k, test.result.Version, gc.Version)
			}
		}
	}
}

func TestGetCapabilitiesQueryParameters(t *testing.T) {
	var tests = []struct {
		query      url.Values
		result     GetCapabilitiesRequest
		exceptions int
	}{
		0: {query: url.Values{
			REQUEST: {getcapabilities}, SERVICE: {Service}, VERSION: {Version},
		},
			result: GetCapabilitiesRequest{
				Service: Service,
				Version: Version,
			},
			exceptions: 0,
		},
	}
	for k, test := range tests {
		var gc GetCapabilitiesRequest
		exceptions := gc.QueryParameters(test.query)
		if len(exceptions) != test.exceptions {
			t.Errorf("test: %d, expected: %d exceptions\n got: %d", k, test.exceptions, len(exceptions))
		}
		if exceptions == nil {
			if gc.Service != test.result.Service {
				t.Errorf("test: %d, expected: %s\n got: %s", k, test.result.Service, gc.Service)
			}
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

func TestGetCapabilitiesResponseType(t *testing.T) {
	gc := GetCapabilitiesResponse{}
	if gc.Type() != getcapabilities {
		t.Errorf("expected: %s\n got: %s", getcapabilities, gc.Type())
	}
}

func TestGetCapabilitiesResponseService(t *testing.T) {
	gc := GetCapabilitiesResponse{}
	if gc.Service() != Service {
		t.Errorf("expected: %s\n got: %s", Service, gc.Service())
	}
}

func TestGetCapabilitiesResponseVersion(t *testing.T) {
	gc := GetCapabilitiesResponse{}
	if gc.Version() != Version {
		t.Errorf("expected: %s\n got: %s", Version, gc.Version())
	}
}
