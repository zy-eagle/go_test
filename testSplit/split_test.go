package testSplit

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestSplit(t *testing.T) {
	//str1 := "schema.canal.*"
	//fmt.Println(strings.SplitN(str1, ".", 2))
	//
	//str2 := "schema"
	//fmt.Println(strings.SplitN(str2, ".", 2))
	//
	//str3 := `13d80f1d-566c-11f0-baae-fa163e7e863f:1-2,\\n13dc7c32-566c-11f0-bb08-fa163eed320f:1-8,\\\n13dc7c32-566c-11f0-bb08-fa163eed320f:1-8\n`
	//
	//parts := strings.Split(str3, ",")
	//re := regexp.MustCompile(`^(\\\\+n|\\n)|(\\\\+n|\\n)$`)
	//res := make([]string, 0, len(parts))
	//for _, part := range parts {
	//	temp := re.ReplaceAllString(part, "")
	//	res = append(res, temp)
	//}
	//
	//fmt.Println(strings.Join(res, ","))

	str4 := "ddd"
	fmt.Println(str4[:2])

	//tempGTID := strings.Replace(strings.TrimSpace(str3), "\\\\n", "", -1)
	//fmt.Println(tempGTID)
	//tempGTID1 := strings.Replace(tempGTID, "\\n", "", -1)
	//fmt.Println(tempGTID1)

	m := map[string]interface{}{
		"name": "eagle",
		"age":  18,
	}

	res, _ := json.Marshal(m)

	fmt.Println(string(res))

}
