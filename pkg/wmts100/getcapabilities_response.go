package wmts100

import (
	"encoding/xml"

	"github.com/flywave/ogc-osgwo/pkg/wsc110"
)

// Type function needed for the interface
func (gc GetCapabilitiesResponse) Type() string {
	return getcapabilities
}

// Service function needed for the interface
func (gc GetCapabilitiesResponse) Service() string {
	return Service
}

// Version function needed for the interface
func (gc GetCapabilitiesResponse) Version() string {
	return Version
}

// Validate function of the wfs200 spec
func (gc GetCapabilitiesResponse) Validate() wsc110.Exceptions {
	return nil
}

// ToXML builds a GetCapabilities response object
func (gc GetCapabilitiesResponse) ToXML() []byte {
	si, _ := xml.MarshalIndent(gc, "", "")
	return si
}

// GetCapabilitiesResponse base struct
type GetCapabilitiesResponse struct {
	XMLName               xml.Name `xml:"Capabilities"`
	Namespaces            `yaml:"namespaces"`
	ServiceIdentification ServiceIdentification   `xml:"ows:ServiceIdentification" yaml:"serviceidentification"`
	ServiceProvider       *wsc110.ServiceProvider `xml:"ows:ServiceProvider,omitempty" yaml:"serviceprovider"`
	OperationsMetadata    *OperationsMetadata     `xml:"ows:OperationsMetadata,omitempty" yaml:"operationsmetadata"`
	Contents              Contents                `xml:"Contents" yaml:"contents"`
	ServiceMetadataURL    *ServiceMetadataURL     `xml:"ServiceMetadataURL,omitempty" yaml:"servicemetadataurl"`
}

// Namespaces struct containing the namespaces needed for the XML document
type Namespaces struct {
	Xmlns          string `xml:"xmlns,attr" yaml:"xmlns"`       //http://www.opengis.net/wmts/1.0
	XmlnsOws       string `xml:"xmlns:ows,attr" yaml:"common"`  //http://www.opengis.net/ows/1.1
	XmlnsXlink     string `xml:"xmlns:xlink,attr" yaml:"xlink"` //http://www.w3.org/1999/xlink
	XmlnsXSI       string `xml:"xmlns:xsi,attr" yaml:"xsi"`     //http://www.w3.org/2001/XMLSchema-instance
	XmlnsGml       string `xml:"xmlns:gml,attr" yaml:"gml"`     //http://www.opengis.net/gml
	Version        string `xml:"version,attr" yaml:"version"`
	SchemaLocation string `xml:"xsi:schemaLocation,attr" yaml:"schemalocation"`
}

// Method in separated struct so to use it as a Pointer
type Method struct {
	Type       string `xml:"xlink:type,attr" yaml:"type"`
	Href       string `xml:"xlink:href,attr" yaml:"href"`
	Constraint []struct {
		Name          string `xml:"name,attr" yaml:"name"`
		AllowedValues struct {
			Value []string `xml:"ows:Value" yaml:"value"`
		} `xml:"ows:AllowedValues" yaml:"allowedvalues"`
	} `xml:"ows:Constraint" yaml:"constraint"`
}

// ServiceIdentification struct should only be fill by the "template" configuration wmts100.yaml
type ServiceIdentification struct {
	Title              string           `xml:"ows:Title" yaml:"title"`
	Abstract           string           `xml:"ows:Abstract" yaml:"abstract"`
	Keywords           *wsc110.Keywords `xml:"ows:Keywords,omitempty" yaml:"keywords"`
	ServiceType        string           `xml:"ows:ServiceType" yaml:"servicetype"`
	ServiceTypeVersion string           `xml:"ows:ServiceTypeVersion" yaml:"servicetypeversion"`
	Fees               string           `xml:"ows:Fees" yaml:"fees"`
	AccessConstraints  string           `xml:"ows:AccessConstraints" yaml:"accessconstraints"`
}

// ServiceMetadataURL in struct for repeatability
type ServiceMetadataURL struct {
	Href string `xml:"xlink:href,attr" yaml:"href"`
}

// ServiceProvider struct containing the provider/organization information should only be fill by the "template" configuration wmts100.yaml
type ServiceProvider struct {
	XMLName      xml.Name `xml:"ows:ServiceProvider"`
	ProviderName string   `xml:"ows:ProviderName" yaml:"providername"`
	ProviderSite struct {
		Type string `xml:"xlink:type,attr" yaml:"type"`
		Href string `xml:"xlink:href,attr" yaml:"href"`
	} `xml:"ows:ProviderSite" yaml:"providersite"`
	ServiceContact struct {
		IndividualName string `xml:"ows:IndividualName" yaml:"individualname"`
		PositionName   string `xml:"ows:PositionName" yaml:"positionname"`
		ContactInfo    struct {
			Text  string `xml:",chardata"`
			Phone struct {
				Voice     string `xml:"ows:Voice" yaml:"voice"`
				Facsimile string `xml:"ows:Facsimile" yaml:"facsmile"`
			} `xml:"ows:Phone" yaml:"phone"`
			Address struct {
				DeliveryPoint         string `xml:"ows:DeliveryPoint" yaml:"deliverypoint"`
				City                  string `xml:"ows:City" yaml:"city"`
				AdministrativeArea    string `xml:"ows:AdministrativeArea" yaml:"administrativearea"`
				PostalCode            string `xml:"ows:PostalCode" yaml:"postalcode"`
				Country               string `xml:"ows:Country" yaml:"country"`
				ElectronicMailAddress string `xml:"ows:ElectronicMailAddress" yaml:"electronicmailaddress"`
			} `xml:"ows:Address" yaml:"address"`
			OnlineResource struct {
				Type string `xml:"xlink:type,attr" yaml:"type"`
				Href string `xml:"xlink:href,attr" yaml:"href"`
			} `xml:"ows:OnlineResource" yaml:"onlineresource"`
			HoursOfService      string `xml:"ows:HoursOfService" yaml:"hoursofservice"`
			ContactInstructions string `xml:"ows:ContactInstructions" yaml:"contactinstructions"`
		} `xml:"ows:ContactInfo" yaml:"contactinfo"`
		Role string `xml:"ows:Role" yaml:"role"`
	} `xml:"ows:ServiceContact" yaml:"servicecontact"`
}
