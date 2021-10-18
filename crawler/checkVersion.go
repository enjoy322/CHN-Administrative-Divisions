package crawler

import (
	"CHN-Administrative-Divisions/model"
	"CHN-Administrative-Divisions/service"
	"fmt"
)

func CheckVersion() {
	var v model.Version
	service.Read("./file/version.json", &v)
	fmt.Println(v)

}
