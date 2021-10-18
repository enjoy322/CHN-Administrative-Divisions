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
	Type      int
	Divisions []Division
}

const CodeFailCrawlCity = 1
const CodeFailCrawlCounty = 2
const CodeFailCrawlTown = 3
const CodeFailCrawlVillage = 4
