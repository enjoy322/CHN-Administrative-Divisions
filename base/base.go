package base

const URL = "http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2020"

//level

const CodeProvince = 1
const CodeCity = 2
const CodeCounty = 3
const CodeTown = 4
const CodeVillage = 5

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

const FileDir = "./file"

var (
	ProvinceFile = FileDir + "/province.json"
	CityFile     = FileDir + "/city.json"
	CountyFile   = FileDir + "/county.json"
	TownFile     = FileDir + "/town.json"
	VillageFile  = FileDir + "/village.json"
)
var File = map[string]string{
	"province": ProvinceFile,
	"city":     CityFile,
	"county":   CountyFile,
	"town":     TownFile,
	"village":  VillageFile,
}
