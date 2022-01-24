package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/mateors/mcb"
	"github.com/xuri/excelize/v2"
)

var db *mcb.DB

type master struct {
	Data []string `json:"data"`
	// Age        int      `json:"age"`
	// Profession string   `json:"profession"`
	// Hobbies    []string `json:"hobbies"`
	// Type       string   `json:"type"`
}

func init() {

	db = mcb.Connect(host, username, password, false)

	res, err := db.Ping()
	if err != nil {

		fmt.Println(res)
		os.Exit(1)
	}
	fmt.Println(res, err)

}

func main() {

	f, err := excelize.OpenFile("file1.xlsx")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
	}
	// rows, err := f.GetCols("Sheet1")
	// if err != nil {
	// 	panic(err)
	// }
	
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, " ")

		}
		fmt.Println()
	}
	
	os.Exit(1)
	//How to insert into couchbase bucket
	var myData master

	form := make(url.Values, 0)
	form.Add("bucket", "royaltypool") //bucket and collection-> namespace:bucket.scope.collection
	form.Add("aid", "d008")           //document ID
	form.Add("data", "rows1")

	// form.Add("age", "36")
	// form.Add("profession", "Developer")
	// form.Add("hobbies", "Programming")
	// form.Add("hobbies", "Problem Solving")
	// form.Add("type", "participant") //what type of data or table name in general (SQL)

	p := db.Insert(form, &myData)    //pass by reference (&myData)
	fmt.Println("Status:", p.Status) //p.Status == Success means data successfully inserted to bucket.

}
