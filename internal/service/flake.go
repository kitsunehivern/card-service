package service

import (
	"log"

	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

func InitFlakeService(id int64) {
	n, err := snowflake.NewNode(id)
	if err != nil {
		log.Fatalf("failed to create new snowflake node: %v", err)
	}

	node = n
}

func GenerateFlakeID() int64 {
	return node.Generate().Int64()
}
