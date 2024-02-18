package main

import (
	"fmt"

	hoge "github.com/kuropenguin/go-mod-test"
	"github.com/kuropenguin/go-mod-test/fuga"
	util "github.com/kuropenguin/go-mod-test/internal/util"
)

func main() {
	fmt.Println(hoge.HogeFunc())
	fmt.Println(fuga.FugaFunc())
	fmt.Println(util.UtilFunc())

}
