package aoc

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func init() {
	registerSolution("2016:5:1", y2016d5part1)
	registerSolution("2016:5:2", y2016d5part2)
}

func y2016d5part1(input string) string {
	doorID := strings.TrimSpace(input)
	res := make([]string, 8)
	ri := 0
	for i := 0; ri < 8; i++ {
		text := fmt.Sprintf("%s%d", doorID, i)
		hash := md5.Sum([]byte(text))
		hashHex := hex.EncodeToString(hash[:])
		if strings.HasPrefix(hashHex, "00000") {
			res[ri] = hashHex[5:6]
			ri += 1
		}
	}
	return strings.Join(res, "")
}

func y2016d5part2(input string) string {
	doorID := strings.TrimSpace(input)
	res := make([]string, 8)
	for i := range 8 {
		res[i] = "_"
	}
	ct := 0
	for i := 0; ct < 8; i++ {
		text := fmt.Sprintf("%s%d", doorID, i)
		hash := md5.Sum([]byte(text))
		hashHex := hex.EncodeToString(hash[:])
		if strings.HasPrefix(hashHex, "00000") {
			// fmt.Println(hashHex)
			ri, err := strconv.Atoi(hashHex[5:6])
			if err != nil {
				// e.g. a,b,c,d,e,f
				continue
			}
			if ri > 7 {
				// not a valid index
				continue
			}
			if res[ri] != "_" {
				// already populated
				continue
			}
			// before := strings.Join(res, "")
			res[ri] = hashHex[6:7]
			// after := strings.Join(res, "")
			// fmt.Println(before, "->", after, " ", hashHex)
			ct++
		}
	}
	return strings.Join(res, "")
}
