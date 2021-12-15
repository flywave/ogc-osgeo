package wmts100

type getFeatureInfoRequestParameterValue struct {
	I          int    `xml:"I" yaml:"i"`
	J          int    `xml:"J" yaml:"j"`
	InfoFormat string `xml:"InfoFormat" yaml:"infoformat" default:"text/plain"` // default text/plain
}
