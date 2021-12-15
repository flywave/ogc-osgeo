package tms100

import (
	"encoding/xml"
)

type Exception struct {
	XMLName xml.Name `xml:"TileMapServerError"`
	Message string   `xml:","`
}

func (e Exception) Error() string {
	return e.Message
}
