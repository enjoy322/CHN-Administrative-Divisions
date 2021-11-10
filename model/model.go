package model

type DivisionYear struct {
	Year       int64
	YearStr    string
	UpdatedStr string
	UpdatedAt  int64
}

type Division struct {
	Url          string
	SimpleCode   string
	Code         string
	Name         string
	VillageType  string
	TownCode     string
	CountyCode   string
	CityCode     string
	ProvinceCode string

	FullName string
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

type Fail struct {
	City    []Division
	County  []Division
	Town    []Division
	Village []Division
}

const CodeProvince = 1
const CodeCity = 2
const CodeCounty = 3
const CodeTown = 4
const CodeVillage = 5
