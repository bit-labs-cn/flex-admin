package main

import (
	"os"

	"bit-labs.cn/owl/utils"
)

func main() {
	println(utils.BcryptHash(os.Args[1]))
}
