package json_example

import (
	"encoding/json"
	"fmt"
)

type Response1 struct {
	Page   int
	Fruits []string
}

func PrintValue() {

	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	stringB, _ := json.Marshal([]string{"apple", "punch"})
	fmt.Println(string(stringB))

	response1 := Response1{
		Page:   1,
		Fruits: []string{"apple", "punch"}}

	resB, _ := json.Marshal(response1)
	fmt.Println(string(resB))

	var resC Response1
	err := json.Unmarshal(resB, &resC)
	if err != nil {
		panic("error")
	}
	fmt.Println(resC.Page)
	fmt.Println(resC.Fruits)

	var data map[string]interface{}

	s := `{"nums": 123, "strs": ["apple", "punch"]}`
	err = json.Unmarshal([]byte(s), &data)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
	fmt.Println(data["nums"].(float64))
	strs := data["strs"].([]interface{})
	fmt.Println(strs[1].(string))

	s1 := `{"Page": 1, "Fruits": ["a", "b"]}`

	var response2 Response1
	err1 := json.Unmarshal([]byte(s1), &response2)
	if err1 != nil {
		panic(err1)
	}
	fmt.Println(response2.Page)
	fmt.Println(response2.Fruits[0])
}
