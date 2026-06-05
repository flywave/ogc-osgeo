package wcs201

import (
	"net/url"
	"strings"

	"github.com/flywave/ogc-osgeo/pkg/wsc200"
)

type describeCoverageRequestParameterValue struct {
	request    string   `yaml:"request"`
	service    string   `yaml:"service"`
	version    string   `yaml:"version"`
	coverageID []string `yaml:"coverageid"`
}

func (dpv *describeCoverageRequestParameterValue) parseQueryParameters(query url.Values) []wsc200.Exception {
	var exceptions []wsc200.Exception
	for k, v := range query {
		if len(v) != 1 {
			exceptions = append(exceptions, wsc200.InvalidParameterValue(k, strings.Join(v, ",")))
		} else {
			switch strings.ToUpper(k) {
			case SERVICE:
				dpv.service = strings.ToUpper(v[0])
			case VERSION:
				dpv.version = v[0]
			case REQUEST:
				dpv.request = v[0]
			case COVERAGEID:
				dpv.coverageID = strings.Split(v[0], ",")
			}
		}
	}
	if len(exceptions) > 0 {
		return exceptions
	}
	return nil
}

func (dpv *describeCoverageRequestParameterValue) parsedescribeCoverageRequest(d DescribeCoverageRequest) {
	dpv.request = describecoverage
	dpv.version = d.Version
	dpv.service = d.Service
	dpv.coverageID = d.CoverageID
}

func (dpv describeCoverageRequestParameterValue) toQueryParameters() url.Values {
	querystring := make(map[string][]string)
	querystring[REQUEST] = []string{dpv.request}
	querystring[SERVICE] = []string{dpv.service}
	querystring[VERSION] = []string{dpv.version}
	if len(dpv.coverageID) > 0 {
		querystring[COVERAGEID] = []string{strings.Join(dpv.coverageID, ",")}
	}
	return querystring
}
