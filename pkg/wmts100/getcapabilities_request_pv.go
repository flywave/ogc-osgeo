package wmts100

import (
	"net/url"
	"strings"

	"github.com/flywave/ogc-osgeo/pkg/wsc110"
)

type getCapabilitiesRequestParameterValue struct {
	service string `yaml:"service"`
	version string `yaml:"version"`
	request string `yaml:"request"`
}

func (gpv *getCapabilitiesRequestParameterValue) parseQueryParameters(query url.Values) wsc110.Exceptions {
	var exceptions wsc110.Exceptions
	for k, v := range query {
		if len(v) != 1 {
			exceptions = append(exceptions, wsc110.InvalidParameterValue(k, strings.Join(v, ",")))
		} else {
			switch strings.ToUpper(k) {
			case SERVICE:
				gpv.service = strings.ToUpper(v[0])
			case VERSION:
				gpv.version = v[0]
			case REQUEST:
				gpv.request = v[0]
			}
		}
	}
	if len(exceptions) > 0 {
		return exceptions
	}
	return nil
}

func (gpv *getCapabilitiesRequestParameterValue) parseGetCapabilitiesRequest(gc GetCapabilitiesRequest) {
	gpv.request = getcapabilities
	gpv.version = gc.Version
	gpv.service = gc.Service
}

func (gpv getCapabilitiesRequestParameterValue) toQueryParameters() url.Values {
	querystring := make(map[string][]string)
	querystring[SERVICE] = []string{gpv.service}
	querystring[VERSION] = []string{gpv.version}
	querystring[REQUEST] = []string{gpv.request}
	return querystring
}
