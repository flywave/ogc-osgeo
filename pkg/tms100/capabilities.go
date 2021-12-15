package tms100

import (
	"encoding/xml"
)

type Metadata struct {
	Type     string `xml:"type,attr"`
	MimeType string `xml:"mime-type,attr"`
	Href     string `xml:"href,attr"`
}

type Logo struct {
	Width    string `xml:"width,attr"`
	Height   string `xml:"height,attr"`
	Href     string `xml:"href,attr"`
	MimeType string `xml:"mime-type,attr"`
}

type Attribution struct {
	Title string `xml:"Title"`
	Logo  Logo
}

type WebMapContext struct {
	Href string `xml:"href,attr"`
}

type BoundingBox struct {
	MinX string `xml:"minx,attr"`
	MinY string `xml:"miny,attr"`
	MaxX string `xml:"maxx,attr"`
	MaxY string `xml:"maxy,attr"`
}

type Origin struct {
	X string `xml:"x,attr"`
	Y string `xml:"y,attr"`
}

type TileFormat struct {
	Width     string `xml:"width,attr"`
	Height    string `xml:"height,attr"`
	Extension string `xml:"extension,attr"`
	MimeType  string `xml:"mime-type,attr"`
}

type TileSet struct {
	Href  string `xml:"href,attr"`
	PPM   string `xml:"units-per-pixel,attr"`
	Order string `xml:"order,attr"`
}

type TileSets struct {
	Profile string `xml:"profile,attr"`
	TileSet []TileSet
}

type TileMap struct {
	XMLName       xml.Name      `xml:"TileMap"`
	Version       string        `xml:"version,attr"`
	Services      string        `xml:"tilemapservice,attr"`
	Title         string        `xml:"Title"`
	Abstract      string        `xml:"Abstract"`
	KeywordList   string        `xml:"KeywordList"`
	Metadata      Metadata      `xml:"Metadata"`
	Attribution   Attribution   `xml:"Attribution"`
	WebMapContext WebMapContext `xml:"WebMapContext"`
	Face          int           `xml:"Face"`
	SRS           string        `xml:"SRS"`
	BoundingBox   BoundingBox   `xml:"BoundingBox"`
	Origin        Origin        `xml:"Origin"`
	TileFormat    TileFormat    `xml:"TileFormat"`
	TileSets      TileSets      `xml:"TileSets"`
}

func (s *TileMap) ParseXML(body []byte) *Exception {
	if err := xml.Unmarshal(body, s); err != nil {
		return &Exception{Message: "Xml parse error!"}
	}

	return nil
}

func (s *TileMap) ToXML() []byte {
	si, _ := xml.MarshalIndent(s, "", "")
	return si
}
