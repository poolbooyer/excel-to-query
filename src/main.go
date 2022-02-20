package main

import (
	"os"
	con "query_generator/src/connect"
	"query_generator/src/tools"

	"flag"
	"fmt"
	"time"
)

type Product struct {
	UpdateDate time.Time `gorm:"column:kaitei_date"`
	Code       string    `gorm:"column:code"`
	Name       string    `gorm:"column:name"`
	Category   *Category `gorm:"embedded"`
}
type Category struct {
	Code string `gorm:"column:code"`
	Name string `gorm:"column:name"`
}
type Output struct {
	Line  int
	Query string
}

func main() {
	flag.Parse()
	cmd := flag.Args()

	ini, err := tools.ParseIni(cmd[0])
	if err != nil {
		return
	}
	rows, err := tools.ReadExcel(cmd[1], cmd[2])
	if err != nil {
		panic(err)
	}

	var res []Product
	for i, row := range rows {
		if i > 0 {
			if len(row) > 1 {
				updateDate, _ := time.Parse("20060102", row[0])
				res = append(res, Product{
					UpdateDate: updateDate,
					Code:       row[1],
					Name:       row[2],
					Category: &Category{
						Name: row[3],
						Code: row[4],
					},
				})
			}
		}
	}
	var q []*Output
	conn, err := con.ConDB(ini)
	for i, item := range res {
		var v Product
		conn.Table("products").Joins("left join categories on products.category_code = categories.code").Where("products.code = ? OR products.name like ?", item.Code, item.Name).Find(&v)
		if err != nil {
			panic(err)
		}
		if v.Code == "" {
			var v = &Output{
				Line:  i,
				Query: fmt.Sprintf("(%s,%s)", item.Code, item.Name),
			}
			q = append(q, v)
		}
	}
	f, err := os.Create("./files/out.sql")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	for _, v := range q {
		o := []byte(v.Query)
		_, err := f.Write(o)
		if err != nil {
			panic(err)
		}
	}
}
