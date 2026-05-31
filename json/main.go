package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	bollB, _ := json.Marshal(true)
	fmt.Println(string(bollB))

	intB, _ := json.Marshal(12)
	fmt.Println(string(intB))

	floatB, _ := json.Marshal(12.122)
	fmt.Println(string(floatB))

	strB, _ := json.Marshal("farid")
	fmt.Println(strB)

	slcB, _ := json.Marshal([]string{"Farid", "babak", "asma"})
	fmt.Println(slcB)
	mapD := map[string]int{"Apple": 4, "Banana": 3}
	mapB, _ := json.Marshal(mapD)

	fmt.Println(mapB)

	resp1 := &response1{
		Page:   12,
		Fruits: []string{"Appl", "banana", "peach"},
	}

	resp1B, _ := json.Marshal(resp1)

	fmt.Println(resp1B)

	byt := []byte(`{"num":6.12, "strs":["a", "b"]}`)

	var data map[string]interface{}

	if err := json.Unmarshal(byt, &data); err != nil {
		panic(err)
	}

	fmt.Println(data)

	num := data["num"].(float64)
	fmt.Println(num)

	strs := data["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	str := `{"page":12, "fruits":["Apple", "Banana"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)

	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 12, "banana": 3}
	enc.Encode(d)

	dec := json.NewDecoder(strings.NewReader(str))
	res1 := response2{}
	dec.Decode(&res1)
	fmt.Println(res1)
}
