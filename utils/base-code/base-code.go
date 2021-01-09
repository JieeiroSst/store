package base_code

import (
	"encoding/base64"
	"log"
	"strconv"
)

func Encode(input int) string{
	return base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(input)))
}

func DeCode(input string) int{
	raw,err:=base64.StdEncoding.DecodeString(input)
	if err != nil {
		log.Fatal(err)
	}

	num,_ :=strconv.Atoi(string(raw))

	return num
}