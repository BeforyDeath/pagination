package main

import (
	"fmt"
	"github.com/BeforyDeath/pagination"
)

func main() {

	Paginator := pagination.Create(50, 3, 7)

	//	Paginator.SetTotal(102)

	pages := Paginator.Get(5)
	fmt.Printf("First:%v\nPrev:%v\nActive:%v\nPage:%v\nNext:%v\nLast:%v\n", pages.First, pages.Prev, pages.Active, pages.Page, pages.Next, pages.Last)
}
