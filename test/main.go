// @description 测试代码
// @author zkp15
// @date 2023/8/9 11:42
// @version 1.0

package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"v-tiktok/pkg/strs"
)

func main() {
	s := strs.UUID()
	fmt.Println(s)

	id := uuid.NewV4()
	fmt.Println(id.String())
}
