package wmts100

import (
	"encoding/xml"
)

// Type returns GetFeatureInfo
func (g *GetFeatureInfoResponse) Type() string {
	return getfeatureinfo
}

// Service returns the service type
func (g *GetFeatureInfoResponse) Service() string {
	return Service
}

// Version returns the service version
func (g *GetFeatureInfoResponse) Version() string {
	return Version
}

// GetFeatureInfoResponse struct
type GetFeatureInfoResponse struct {
	XMLName      xml.Name      `xml:"GetFeatureInfoResponse" yaml:"getfeatureinforesponse"`
	TextPayload  TextPayload   `xml:"TextPayload,omitempty" yaml:"textpayload"`
	BinaryPayload BinaryPayload `xml:"BinaryPayload,omitempty" yaml:"binarypayload"`
}
