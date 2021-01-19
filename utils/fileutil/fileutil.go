package fileutil

import (
	"regexp"
	"strings"

	beego "github.com/beego/beego/v2/adapter"
)

func GetSafeFileName(input string) string {
	reg, err := regexp.Compile("[^A-Za-z0-9_]+")
	if err != nil {
		beego.Error(err)
	}

	safe := reg.ReplaceAllString(input, "-")
	safe = strings.ToLower(strings.Trim(safe, "-"))
	return safe
}
