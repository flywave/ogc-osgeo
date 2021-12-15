package tms100

import (
	"encoding/xml"
	"fmt"
	"testing"
)

func TestException(t *testing.T) {
	err := Exception{Message: "error"}
	output, _ := xml.MarshalIndent(err, "", "  ")
	xmlstr := string(output)
	fmt.Print(xmlstr)
}

func TestCapabilities(t *testing.T) {
	ser := Services{TileMapService: []ServiceMapinfo{{Title: "Example Tile Map Service", Version: "1.0.0", Href: "http://tms.osgeo.org/1.0.0/"}}, FancyFeatureService: []FancyFeatureService{{Title: "Features!", Version: "0.9", Href: "http://ffs.osgeo.org/0.9/"}}}
	output, _ := xml.MarshalIndent(ser, "", "  ")
	xmlstr := string(output)
	t.Log(xmlstr)
	ser1 := &Services{}
	ser1.ParseXML(output)

	if len(ser1.TileMapService) != 1 {
		t.FailNow()
	}
}

func TestTileMapService(t *testing.T) {
	xml := `
	<?xml version="1.0" encoding="UTF-8" ?>
 <TileMapService version="1.0.0" services="http://www.osgeo.org/services/root.xml">
   <Title>Example Static Tile Map Service</Title>
   <Abstract>This is a longer description of the static tiling map service.</Abstract>
 | <KeywordList>example tile service static</KeywordList>
 | <ContactInformation>
 |   <ContactPersonPrimary>
 |     <ContactPerson>Paul Ramsey</ContactPerson>
 |     <ContactOrganization>Refractions Research</ContactOrganization>
 |   </ContactPersonPrimary>
 |   <ContactPosition>Manager</ContactPosition>
 |   <ContactAddress>
 |     <AddressType>postal</AddressType>
 |     <Address>300 - 1207 Douglas Street</Address>
 |     <City>Victoria</City>
 |     <StateOrProvince>British Columbia</StateOrProvince>
 |     <PostCode>V8W2E7</PostCode>
 |     <Country>Canada</Country>
 |   </ContactAddress>
 |   <ContactVoiceTelephone>12503833022</ContactVoiceTelephone>
 |   <ContactFacsimileTelephone>12503832140</ContactFacsimileTelephone>
 |   <ContactElectronicMailAddress>pramsey@refractions.net</ContactElectronicMailAddress>
 | </ContactInformation>
   <TileMaps>
     <TileMap 
       title="Vancouver Island Base Map" 
       srs="EPSG:26910" 
       profile="none" 
       href="http://www.osgeo.org/services/basemap.xml" />
   </TileMaps>
 </TileMapService>`

	ser1 := &TileMapService{}
	ser1.ParseXML([]byte(xml))

	if len(ser1.TileMaps.TileMap) != 1 {
		t.FailNow()
	}
}

func TestTileMap(t *testing.T) {
	xml := `
	<?xml version="1.0" encoding="UTF-8" ?>
	<TileMap version="1.0.0" tilemapservice="http://tms.osgeo.org/1.0.0">
	 <Title>British Columbia Landsat Imagery (2000)</Title>
	 <Abstract>Landsat data collected in the year 2000 over British Columbia.</Abstract>
   | <KeywordList></KeywordList>
   | <Metadata type="TC211" mime-type="text/xml" href="http://www.org" />
   | <Attribution>
   |   <Title>Government of British Columbia</Title>
   |   <Logo width="10" height="10" href="http://gov.bc.ca/logo.png" mime-type="image/png" />
   | </Attribution>
   | <WebMapContext href="http://wms.gov.bc.ca" />
   | <Face>0</Face>
	 <SRS>EPSG:3005</SRS>
	 <BoundingBox minx="100000" miny="100000" maxx="1800000" maxy="1800000" />
	 <Origin x="100000" y="100000" />
	 <TileFormat width="256" height="256" mime-type="image/png" extension="png" />
	 <TileSets profile="local">
	   <TileSet href="http://tms.osgeo.org/1.0.0/landsat2000/2048" units-per-pixel="2048" order="0" />
	   <TileSet href="http://tms.osgeo.org/1.0.0/landsat2000/1024" units-per-pixel="1024" order="1" />
	   <TileSet href="http://tms.osgeo.org/1.0.0/landsat2000/512" units-per-pixel="512" order="2" />
	   <TileSet href="http://tms.osgeo.org/1.0.0/landsat2000/256" units-per-pixel="256" order="3" />
	   <TileSet href="http://tms.osgeo.org/1.0.0/landsat2000/128" units-per-pixel="128" order="4" />
	   <TileSet href="http://tms.osgeo.org/1.0.0/landsat2000/64" units-per-pixel="64" order="5" />
	 </TileSets>
   </TileMap>
  `

	ser1 := &TileMap{}
	ser1.ParseXML([]byte(xml))

	if len(ser1.TileSets.TileSet) != 6 {
		t.FailNow()
	}
}
