package wmts100

type getTileRequestParameterValue struct {
	Layer              string               `xml:"Layer" yaml:"layer"`
	Style              string               `xml:"Style" yaml:"style"`
	Format             string               `xml:"Format" yaml:"format"`
	TileMatrixSet      string               `xml:"TileMatrixSet" yaml:"tilematrixset"`
	DimensionNameValue []DimensionNameValue `xml:"DimensionNameValue" yaml:"dimensionnamevalue"`
	TileMatrix         string               `xml:"TileMatrix" yaml:"tilematrix"`
	TileRow            string               `xml:"TileRow" yaml:"tilerow"`
	TileCol            string               `xml:"TileCol" yaml:"tilecol"`
}
