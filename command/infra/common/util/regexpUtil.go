package util

import (
	"regexp"
)

/**
*@Description:
*@Author: BZ
*@date: 2023/11/5 9:26
*@Version: V1.0
 */

func MatchRegexpString(reg string, target string) (bool, error) {
	re, err := regexp.Compile(reg)
	if err != nil {
		return false, err
	}
	res := re.MatchString(target)
	return res, err

}
