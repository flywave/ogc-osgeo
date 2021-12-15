package tms100

import (
	"encoding/xml"
)

type ContactPersonPrimary struct {
	ContactPerson       string `xml:"ContactPerson"`
	ContactOrganization string `xml:"ContactOrganization"`
}

type ContactAddress struct {
	AddressType     string `xml:"AddressType"`
	Address         string `xml:"Address"`
	City            string `xml:"City"`
	StateOrProvince string `xml:"StateOrProvince"`
	PostCode        string `xml:"PostCode"`
	Country         string `xml:"Country"`
}

type ContactInformation struct {
	ContactPersonPrimary         ContactPersonPrimary
	ContactPosition              string `xml:"ContactPosition"`
	ContactAddress               ContactAddress
	ContactVoiceTelephone        string `xml:"ContactVoiceTelephone"`
	ContactFacsimileTelephone    string `xml:"ContactFacsimileTelephone"`
	ContactElectronicMailAddress string `xml:"ContactElectronicMailAddress"`
}

type TileMapInfo struct {
	XMLName xml.Name `xml:"TileMap"`
	Title   string   `xml:"title,attr"`
	Srs     string   `xml:"srs,attr"`
	Profile string   `xml:"profile,attr"`
	Href    string   `xml:"href,attr"`
}

type TileMaps struct {
	XMLName xml.Name      `xml:"TileMaps"`
	TileMap []TileMapInfo `xml:"TileMap"`
}

type TileMapService struct {
	XMLName            xml.Name           `xml:"TileMapService"`
	Version            string             `xml:"version,attr"`
	Services           string             `xml:"services,attr"`
	Title              string             `xml:"Title"`
	Abstract           string             `xml:"Abstract"`
	KeywordList        string             `xml:"KeywordList"`
	ContactInformation ContactInformation `xml:"ContactInformation"`
	TileMaps           TileMaps           `xml:"TileMaps"`
}

func (s *TileMapService) ParseXML(body []byte) *Exception {
	if err := xml.Unmarshal(body, s); err != nil {
		return &Exception{Message: "Xml parse error!"}
	}

	return nil
}

func (s *TileMapService) ToXML() []byte {
	si, _ := xml.MarshalIndent(s, "", "")
	return si
}
