package util

import (
	"CHN-Administrative-Divisions/base"
)

// FindNeed 查询仍需要爬取的
func FindNeed(level int, upLevelList []base.Division, doneList []base.Division) (needList []base.Division) {
	tempMap := map[string]int{} // 存放不重复主键
	switch level {
	case 1:
		for _, division := range doneList {
			if _, ok := tempMap[division.ProvinceCode]; !ok {
				tempMap[division.ProvinceCode] = 1
			}
		}
		for _, division := range upLevelList {
			if _, ok := tempMap[division.ProvinceCode]; !ok {
				needList = append(needList, division)
			}
		}
		return
	case 2:
		for _, division := range doneList {
			if _, ok := tempMap[division.CityCode]; !ok {
				tempMap[division.CityCode] = 1
			}
		}
		for _, division := range upLevelList {
			if _, ok := tempMap[division.CityCode]; !ok {
				needList = append(needList, division)
			}
		}
		return
	case 3:
		for _, division := range doneList {
			if _, ok := tempMap[division.CountyCode]; !ok {
				tempMap[division.CountyCode] = 1
			}
		}
		for _, division := range upLevelList {
			if _, ok := tempMap[division.CountyCode]; !ok {
				if !division.Branch {
					continue
				}
				needList = append(needList, division)
			}
		}
		return
	case 4:
		for _, division := range doneList {
			if _, ok := tempMap[division.TownCode]; !ok {
				tempMap[division.TownCode] = 1
			}
		}
		for _, division := range upLevelList {
			if _, ok := tempMap[division.TownCode]; !ok {
				needList = append(needList, division)
			}
		}
		return
	default:
		return
	}
}
