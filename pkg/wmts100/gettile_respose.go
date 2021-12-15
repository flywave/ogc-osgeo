package wmts100

import (
	"encoding/xml"
	"regexp"

	"github.com/flywave/ogc-osgwo/pkg/wsc110"
)

// Type function needed for the interface
func (gc GetTileResponse) Type() string {
	return gettile
}

// Service function needed for the interface
func (gc GetTileResponse) Service() string {
	return Service
}

// Version function needed for the interface
func (gc GetTileResponse) Version() string {
	return Version
}

// Validate function of the wfs200 spec
func (gc GetTileResponse) Validate() wsc110.Exceptions {
	return nil
}

// ToXML builds a GetCapabilities response object
func (gc GetTileResponse) ToXML() []byte {
	si, _ := xml.MarshalIndent(gc, "", "")
	re := regexp.MustCompile(`><.*>`)
	return []byte(xml.Header + re.ReplaceAllString(string(si), "/>"))
}

type GetTileResponse struct {
	BinaryPayload BinaryPayload `xml:"BinaryPayload,omitempty"`
	TextPayload   TextPayload   `xml:"TextPayload,omitempty"`
}

type BinaryPayload struct {
	Format   string `xml:"Format" yaml:"format"`
	Contents string `xml:"BinaryContent" yaml:"binarycontents"`
}

type TextPayload struct {
	Format   string `xml:"Format" yaml:"format"`
	Contents string `xml:"TextContent" yaml:"textcontent"`
}
