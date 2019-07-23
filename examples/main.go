/**
 * Created by Wangwei on 2019-07-23 09:00.
 */

package main

import (
	"fmt"
)

func main() {
	role1 := new(Role)
	role2 := Role{}
	role3 := &Role{}

	var num int = 5
	fmt.Printf("num: %d\n", num)

	fmt.Printf("role: %v,%v,%v", role1,role2,role3)
}

type Role struct {
	Code        string
}
