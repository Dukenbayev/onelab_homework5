package string_number

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

var rawJson = []byte(`[
  {
    "id": 1,
    "address": {
      "city_id": 5,
      "street": "Satbayev"
    },
    "Age": 20
  },
  {
    "id": 1,
    "address": {
      "city_id": "6",
      "street": "Al-Farabi"
    },
    "Age": "32"
  }
]`)

// XML
var rawXML = []byte(`
	<users>
		<user>
			<id>1</id>
			<address>
				<city_id>5</city_id>
				<street>Satbayev</street>
			</address>
			<age>20</age>
		</user>
		<user>
			<id>1</id>	
			<address>
				<city_id>6</city_id>
				<street>Al-Farabi</street>
			</address>
			<age>32</age>
		</user>
	</users>
`)

func main() {
	var users []User
	if err := json.Unmarshal(rawJson, &users); err != nil {
		panic(err)
	}

	for _, user := range users {
		fmt.Printf("%#v\n", user)
	}
	var usersXml Users
	if err := json.Unmarshal(rawXML, &usersXml); err != nil {
		panic(err)
	}

	for _, user := range usersXml.Users {
		fmt.Printf("%#v\n", user)
	}

}

type User struct {
	ID      int64   `json:"id" xml:"id"`
	Address Address `json:"address" xml:"address"`
	Age     int     `json:"age" xml:"age"`
}

type Address struct {
	CityID int64  `json:"city_id" xml:"city_id"`
	Street string `json:"street" xml:"street"`
}

type Users struct {
	Users []User `xml:"users"`
}

type StringNumber int

func (s *StringNumber) UnmarshalJSON(b  []byte) error{
	result := strings.ReplaceAll(string(b),`"`,``)
	i, err := strconv.Atoi(result)
	if err !=nil{
		return err
	}
	*s = StringNumber(i)
	// Add check by type
	return nil
}

func (s *StringNumber) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error{
	var value string
	d.DecodeElement(&value,&start)
	i, err := strconv.Atoi(strings.ReplaceAll(value,`"`,``))
	if err != nil{
		return err
	}
	*s =(StringNumber)(i)
	return nil
}