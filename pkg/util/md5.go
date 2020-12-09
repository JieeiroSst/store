package util

import (
	"crypto/md5"
	"encoding/hex"
)

func EncodeMD5(value string)string{
	m:=md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}

func RemoveRepByMap(slc []string) []string{
	result:=[]string{}
	tempMap:=map[string]byte{}
	for _,e:=range slc{
		l:=len(tempMap)
		tempMap[e]=0
		if len(tempMap)!=l{
			result=append(result,e)
		}
	}
	return result
}