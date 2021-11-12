package model

type DivisionYear struct {
	Year       int64
	YearStr    string
	UpdatedStr string
	UpdatedAt  int64
}

type DivisionVillage struct {
	Code         string
	Name         string
	TownCode     string
	CountyCode   string
	CityCode     string
	ProvinceCode string

	Level uint8
}

type Division struct {
	Code         string
	Name         string
	TownCode     string
	CountyCode   string
	CityCode     string
	ProvinceCode string

	Level  uint8
	Branch bool
}

type DivisionTree struct {
	Code         string
	Name         string
	TownCode     string
	CountyCode   string
	CityCode     string
	ProvinceCode string

	Level    uint8
	Children []Division
}

type Version struct {
	CreateAt    int
	CreateAtStr string
	URL         string
	Version     int
	VersionStr  string
	Year        int
	YearStr     string
}

const CodeProvince = 1
const CodeCity = 2
const CodeCounty = 3
const CodeTown = 4
const CodeVillage = 5
