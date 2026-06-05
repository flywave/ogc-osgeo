package wmts100

//
const (
	getcapabilities = `GetCapabilities`
	gettile         = `GetTile`
	getfeatureinfo  = `GetFeatureInfo`

	Service = `WMTS`
	Version = `1.0.0`
)

// KVP parameter tokens
const (
	SERVICE      = `SERVICE`
	REQUEST      = `REQUEST`
	VERSION      = `VERSION`
	LAYER        = `LAYER`
	STYLE        = `STYLE`
	FORMAT       = `FORMAT`
	TILEMATRIXSET = `TILEMATRIXSET`
	TILEMATRIX   = `TILEMATRIX`
	TILEROW      = `TILEROW`
	TILECOL      = `TILECOL`
	I            = `I`
	J            = `J`
	INFOFORMAT   = `INFOFORMAT`
)
