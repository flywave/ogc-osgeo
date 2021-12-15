package tms100

type ProfilesType string

const (
	PROFILES_NONE            = "none"
	PROFILES_GLOBAL_GEODETIC = "global-geodetic"
	PROFILES_GLOBAL_MERCATOR = "global-mercator"
	PROFILES_LOCAL           = "local"
)

const (
	SRS_GLOBAL_GEODETIC = "EPSG:4326"
	SRS_GLOBAL_MERCATOR = "OSGEO:41001"
)
