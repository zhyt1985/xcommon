/*
 * @Author: ybg
 * @Date: 2022-07-22 16:27:58
 * @LastEditors: ybg
 * @LastEditTime: 2022-08-04 18:00:31
 * @Description: nc
 */
package xexcel

import "testing"

type Person struct {
	Name string `json:"name" xlsx:"name:名称,index:2"`
	Age  int64  `json:"age"  xlsx:"name:年龄,index:1"`
	Sex  int64  `json:"sex"  xlsx:"-"`
}

func TestCreateExcel(t *testing.T) {
	e := NewExeclClient(SetPath("./doc/report/"), SetName("123.xlsx"))
	e.MergeCell("A1", "B1").MergeCell("A2", "B2")
	//d := Person{Name: "ybg", Age: 20, Sex: 1}
	d := []Person{{Name: "y", Age: 10, Sex: 1}, {Name: "b", Age: 20, Sex: 1}}
	//d := Person{Name: "ybg", Age: 20, Sex: 1}
	e.CreateExcel(d)
}
