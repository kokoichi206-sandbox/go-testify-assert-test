package main

import (
	"fmt"
	"log"

	"github.com/kokoichi206-sandbox/go-testify-assert-test/gen/go/protobuf"
)

func NewProtoMessage(name string) *protobuf.HelloReply {
	x := &protobuf.HelloReply{
		Name: name,
	}

	// 出力がテストに影響を与える。
	// fmt.Printf("x: %v\n", x)
	log.Print(x)

	return x
}

func main() {
}
