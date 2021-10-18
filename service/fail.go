package service

import "CHN-Administrative-Divisions/model"

func CheckFail(t int) []model.Division {
	var fails []model.Fail
	Read("./file/fail.json", &fails)
	for _, fail := range fails {
		if fail.Type == t {
			return fail.Divisions
		}
	}
	return nil
}
