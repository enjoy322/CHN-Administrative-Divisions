package service

import "CHN-Administrative-Divisions/model"

func CheckFail() model.Fail {
	var fails model.Fail
	Read("./file/fail.json", &fails)
	return fails
}
