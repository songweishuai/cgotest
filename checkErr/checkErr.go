package checkErr

import (
	"fmt"
	"os"
)

func CheckErr(err error, exit bool) {
	if err == nil {
		return
	}
	fmt.Println(err)

	if err != nil && exit == true {
		os.Exit(-1)
	}
}
