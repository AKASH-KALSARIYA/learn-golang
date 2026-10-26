package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"CourseName"`
	Price    int
	PlatForm string   `json:"WebSite"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("JSON KA JADU!!!")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	MyCourses := []course{
		{
			"ReactJs BootCamp",
			299,
			"React.dev",
			"AK071",
			[]string{"Frontend-dev", "JS"},
		},
		{
			"NodeJs BootCamp",
			399,
			"Node.dev",
			"AK072",
			[]string{"BAckend-dev", "JS"},
		},
		{
			"Mern BootCamp",
			399,
			"Mern.dev",
			"AK073",
			[]string{"FullStack-dev", "JS"},
		},
	}

	//package this data as JSON data
	// MarshalIndent(data, prefix, indent)
	finaljson, err := json.MarshalIndent(MyCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finaljson)
}

func DecodeJson() {
	jsonDataFromweb := []byte(`
	        {
                "CourseName": "ReactJs BootCamp",
                "Price": 299,
                "WebSite": "React.dev",
                "tags": [
                        "Frontend-dev",
                        "JS"
               			 ]
        	}
	`)

	var MyCourses course

	checkValid := json.Valid(jsonDataFromweb)
	if checkValid {
		fmt.Println("Json Was valid!!")
		json.Unmarshal(jsonDataFromweb, &MyCourses)
		fmt.Printf("%#v\n", MyCourses)
	} else {
		fmt.Println("Json Was Not Valid!!!")
	}

	//sometimes cases have where youjust want to add data to key value
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromweb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and Vlaue is %v", k, v)
	}

}
