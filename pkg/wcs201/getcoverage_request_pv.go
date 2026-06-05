package wcs201

import (
	"net/url"
	"strings"

	"github.com/flywave/ogc-osgeo/pkg/wsc200"
)

type getCoverageRequestParameterValue struct {
	request         string            `yaml:"request"`
	service         string            `yaml:"service"`
	version         string            `yaml:"version"`
	coverageID      string            `yaml:"coverageid"`
	format          *string           `yaml:"format"`
	dimensionSubset []DimensionSubset `yaml:"dimensionsubset"`
	subsettingCRS   *string           `yaml:"subsettingcrs"`
	outputCRS       *string           `yaml:"outputcrs"`
	mediaType       *string           `yaml:"mediatype"`
}

func (gpv *getCoverageRequestParameterValue) parseQueryParameters(query url.Values) []wsc200.Exception {
	var exceptions []wsc200.Exception
	for k, v := range query {
		switch strings.ToUpper(k) {
		case SERVICE:
			gpv.service = strings.ToUpper(v[0])
		case VERSION:
			gpv.version = v[0]
		case REQUEST:
			gpv.request = v[0]
		case COVERAGEID:
			gpv.coverageID = v[0]
		case FORMAT:
			vp := v[0]
			gpv.format = &vp
		case SUBSET:
			for _, sv := range v {
				ds, err := parseSubsetString(sv)
				if err != nil {
					exceptions = append(exceptions, wsc200.InvalidParameterValue(sv, SUBSET))
				} else if ds != nil {
					gpv.dimensionSubset = append(gpv.dimensionSubset, *ds)
				}
			}
		case SUBSETTINGCRS:
			vp := v[0]
			gpv.subsettingCRS = &vp
		case OUTPUTCRS:
			vp := v[0]
			gpv.outputCRS = &vp
		case MEDIATYPE:
			vp := v[0]
			gpv.mediaType = &vp
		}
	}
	if len(exceptions) > 0 {
		return exceptions
	}
	return nil
}

func (gpv *getCoverageRequestParameterValue) parseGetCoverageRequest(g GetCoverageRequest) {
	gpv.request = getcoverage
	gpv.version = g.Version
	gpv.service = g.Service
	gpv.coverageID = g.CoverageID
	gpv.format = g.Format
	gpv.dimensionSubset = g.DimensionSubset
	gpv.subsettingCRS = g.SubsettingCRS
	gpv.outputCRS = g.OutputCRS
	gpv.mediaType = g.MediaType
}

func (gpv getCoverageRequestParameterValue) toQueryParameters() url.Values {
	querystring := make(map[string][]string)
	querystring[REQUEST] = []string{gpv.request}
	querystring[SERVICE] = []string{gpv.service}
	querystring[VERSION] = []string{gpv.version}
	querystring[COVERAGEID] = []string{gpv.coverageID}
	if gpv.format != nil {
		querystring[FORMAT] = []string{*gpv.format}
	}
	for _, ds := range gpv.dimensionSubset {
		querystring[SUBSET] = append(querystring[SUBSET], subsetToKVP(ds))
	}
	if gpv.subsettingCRS != nil {
		querystring[SUBSETTINGCRS] = []string{*gpv.subsettingCRS}
	}
	if gpv.outputCRS != nil {
		querystring[OUTPUTCRS] = []string{*gpv.outputCRS}
	}
	if gpv.mediaType != nil {
		querystring[MEDIATYPE] = []string{*gpv.mediaType}
	}
	return querystring
}
