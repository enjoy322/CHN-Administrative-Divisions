package main

type DivisionYear struct {
	Year      string
	UpdatedAt int64
}

type Division struct {
	Url          string
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
