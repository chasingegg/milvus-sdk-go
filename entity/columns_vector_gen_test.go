// Code generated by go generate; DO NOT EDIT
// This file is generated by go generate 

package entity 

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	schema "github.com/milvus-io/milvus-proto/go-api/schemapb"
	"github.com/stretchr/testify/assert"
)

func TestColumnBinaryVector(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	columnName := fmt.Sprintf("column_BinaryVector_%d", rand.Int())
	columnLen := 12 + rand.Intn(10)
	dim := ([]int{64, 128, 256, 512})[rand.Intn(4)]

	v := make([][]byte,0, columnLen)
	dlen := dim
	dlen /= 8
	
	for i := 0; i < columnLen; i++ {
		entry := make([]byte, dlen)
		v = append(v, entry)
	}
	column := NewColumnBinaryVector(columnName, dim, v)
	
	t.Run("test meta", func(t *testing.T) {
		ft := FieldTypeBinaryVector
		assert.Equal(t, "BinaryVector", ft.Name())
		assert.Equal(t, "[]byte", ft.String())
		pbName, pbType := ft.PbFieldType()
		assert.Equal(t, "[]byte", pbName)
		assert.Equal(t, "", pbType)
	})

	t.Run("test column attribute", func(t *testing.T) {
		assert.Equal(t, columnName, column.Name())
		assert.Equal(t, FieldTypeBinaryVector, column.Type())
		assert.Equal(t, columnLen, column.Len())
		assert.Equal(t, dim, column.Dim())
		assert.Equal(t ,v, column.Data())
		
		var ev []byte
		err := column.AppendValue(ev)
		assert.Equal(t, columnLen+1, column.Len())
		assert.Nil(t, err)
		
		err = column.AppendValue(struct{}{})
		assert.Equal(t, columnLen+1, column.Len())
		assert.NotNil(t, err)
	})

	t.Run("test column field data", func(t *testing.T) {
		fd := column.FieldData()
		assert.NotNil(t, fd)
		assert.Equal(t, fd.GetFieldName(), columnName)

		c, err := FieldDataVector(fd)
		assert.NotNil(t, c)
		assert.NoError(t, err)
	})

	t.Run("test column field data error", func(t *testing.T) {
		fd := &schema.FieldData{
			Type:      schema.DataType_BinaryVector,
			FieldName: columnName,
		}
		_, err := FieldDataVector(fd) 
		assert.Error(t, err)
	})

}

func TestColumnFloatVector(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	columnName := fmt.Sprintf("column_FloatVector_%d", rand.Int())
	columnLen := 12 + rand.Intn(10)
	dim := ([]int{64, 128, 256, 512})[rand.Intn(4)]

	v := make([][]float32,0, columnLen)
	dlen := dim
	
	
	for i := 0; i < columnLen; i++ {
		entry := make([]float32, dlen)
		v = append(v, entry)
	}
	column := NewColumnFloatVector(columnName, dim, v)
	
	t.Run("test meta", func(t *testing.T) {
		ft := FieldTypeFloatVector
		assert.Equal(t, "FloatVector", ft.Name())
		assert.Equal(t, "[]float32", ft.String())
		pbName, pbType := ft.PbFieldType()
		assert.Equal(t, "[]float32", pbName)
		assert.Equal(t, "", pbType)
	})

	t.Run("test column attribute", func(t *testing.T) {
		assert.Equal(t, columnName, column.Name())
		assert.Equal(t, FieldTypeFloatVector, column.Type())
		assert.Equal(t, columnLen, column.Len())
		assert.Equal(t, dim, column.Dim())
		assert.Equal(t ,v, column.Data())
		
		var ev []float32
		err := column.AppendValue(ev)
		assert.Equal(t, columnLen+1, column.Len())
		assert.Nil(t, err)
		
		err = column.AppendValue(struct{}{})
		assert.Equal(t, columnLen+1, column.Len())
		assert.NotNil(t, err)
	})

	t.Run("test column field data", func(t *testing.T) {
		fd := column.FieldData()
		assert.NotNil(t, fd)
		assert.Equal(t, fd.GetFieldName(), columnName)

		c, err := FieldDataVector(fd)
		assert.NotNil(t, c)
		assert.NoError(t, err)
	})

	t.Run("test column field data error", func(t *testing.T) {
		fd := &schema.FieldData{
			Type:      schema.DataType_FloatVector,
			FieldName: columnName,
		}
		_, err := FieldDataVector(fd) 
		assert.Error(t, err)
	})

}

