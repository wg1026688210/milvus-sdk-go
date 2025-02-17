package tests

import (
	"context"
	"testing"

	"github.com/milvus-io/milvus-sdk-go/v2/client"
	"github.com/milvus-io/milvus-sdk-go/v2/entity"
	"github.com/stretchr/testify/assert"
)

func TestCreateIndex(t *testing.T) {
	c, err := client.NewGrpcClient(context.Background(), "localhost:19530")
	assert.Nil(t, err)
	assert.NotNil(t, c)
	if c != nil {
		defer c.Close()
	}
	cname := generateCollectionName()
	schema := generateSchema()
	generateCollection(t, c, cname, schema, true)

	idx := entity.NewGenericIndex("", entity.Flat, map[string]string{
		"nlist":       "1024",
		"metric_type": "IP",
	})
	err = c.CreateIndex(context.Background(), cname, "vector", idx, false)
	assert.Nil(t, err)
	indexes, err := c.DescribeIndex(context.Background(), cname, "vector")
	if assert.Nil(t, err) {
		for _, idx := range indexes {
			t.Log(idx.IndexType())
			t.Log(idx.Params())
		}
	}
}
