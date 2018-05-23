package model

type GeoInfo struct {
	Geometry Geometry
}

type Geometry struct {
	Location Location
}

type Location struct {
	Lat float64
	Lng float64
}