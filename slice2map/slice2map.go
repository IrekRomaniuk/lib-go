package slice2map

import (
	"strings"
	"strconv"
)
/*
input: ["ae1:1","ae2:0","ae3:0"]
output: {"ae1":1,"ae2":0,"ae3":0}
 */

const testVersion = 1
//convert slice to map
func Slice2Map(inputSlice []string) map[string]int {
	outputMap :=  make(map[string]int,len(inputSlice))
	for _, element:= range inputSlice {
		s:= strings.Split(element,":")
		i, _ := strconv.Atoi(s[1])
		outputMap[s[0]]= i
	}
	return outputMap
}
