package main

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io"
)

func main() {
	fmt.Println("Please enter a word to hash: ")
	var input string
	fmt.Scanln(&input)

	md5 := md5.New()
	sha_256 := sha256.New()
	sha_512 := sha512.New()

	io.WriteString(md5, input)
	sha_256.Write([]byte(input))
	sha_512.Write([]byte(input))
	

	fmt.Printf("md5:\t\t%x\n", md5.Sum(nil))

	//eb7a03c377c28da97ae97884582e6bd07fa44724af99798b42593355e39f82cb
	fmt.Printf("sha256:\t\t%x\n", sha_256.Sum(nil))

	//5cdaf0d2f162f55ccc04a8639ee490c94f2faeab3ba57d3c50d41930a67b5fa6915a73d6c78048729772390136efed25b11858e7fc0eed1aa7a464163bd44b1c
	fmt.Printf("sha512:\t\t%x\n", sha_512.Sum(nil))

}