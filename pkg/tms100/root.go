package tms100

import (
	"encoding/xml"
)

type ServiceMapinfo struct {
	XMLName xml.Name `xml:"TileMapService"`
	Title   string   `xml:"title,attr"`
	Version string   `xml:"version,attr"`
	Href    string   `xml:"href,attr"`
}

type FancyFeatureService struct {
	Title   string `xml:"title,attr"`
	Version string `xml:"version,attr"`
	Href    string `xml:"href,attr"`
}

type Services struct {
	XMLName             xml.Name              `xml:"Services"`
	TileMapService      []ServiceMapinfo      `xml:"TileMapService"`
	FancyFeatureService []FancyFeatureService `xml:"FancyFeatureService"`
}

func (s *Services) ParseXML(body []byte) *Exception {
	if err := xml.Unmarshal(body, s); err != nil {
		return &Exception{Message: "Xml parse error!"}
	}

	return nil
}

func (s *Services) ToXML() []byte {
	si, _ := xml.MarshalIndent(s, "", "")
	return si
}
