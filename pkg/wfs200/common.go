package wfs200

import (
	"github.com/flywave/ogc-osgeo/pkg/utils"
	"github.com/flywave/ogc-osgeo/pkg/wsc110"
)

//
const (
	getcapabilities     = `GetCapabilities`
	getfeature          = `GetFeature`
	describefeaturetype = `DescribeFeatureType`
	getpropertyvalue    = `GetPropertyValue`

	/* Unused/support transactional requesttypes
	lockfeature = `LockFeature`
	getfeaturewithlock = `GetFeatureWithLock`
	*/

	// TODO StoredQuery
	// liststoredqueries     = `ListStoredQueries`
	// describestoredqueries = `DescribeStoredQueries`
	/* Unused/support tranactional storedquery requesttypes
	createstoredquery = `CreateStoredQuery`
	dropstoredquery = `DropStoredQuery`
	*/

	Service = `WFS`
	Version = `2.0.0`
)

// WFS 2.0.0 Tokens
const (
	SERVICE = `SERVICE`
	REQUEST = `REQUEST`
	VERSION = `VERSION`

	OUTPUTFORMAT = `OUTPUTFORMAT`
)

const (
	gml32 string = `text/xml' subtype=gml/3.2`
)

// baseParameterValueRequest struct
type baseParameterValueRequest struct {
	version string `yaml:"version,omitempty"`
	request string `yaml:"request,omitempty"`
}

// BaseRequest based on Table 5 WFS2.0.0 spec
// Note: not usable for GetCapabilities request regarding deviation of Optional/Mandatory parameters SERVICE and VERSION
type BaseRequest struct {
	Service string             `xml:"service,attr" yaml:"service,omitempty"`
	Version string             `xml:"version,attr" yaml:"version"`
	Attr    utils.XMLAttribute `xml:",attr"`
}

// parseBaseParameterValueRequest builds a BaseRequest Struct based on the given parameters
func (b *BaseRequest) parseBaseParameterValueRequest(bpv baseParameterValueRequest) []wsc110.Exception {
	// Service is optional, because it's implicit for a GetFeature request
	b.Service = Service

	if bpv.version != `` {
		b.Version = bpv.version
	} else {
		// Version is mandatory
		return wsc110.MissingParameterValue(VERSION).ToExceptions()
	}
	return nil
}
