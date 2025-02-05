package bcd

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

type uint8testdata struct {
	bytes   byte
	integer uint8
}

type uint16testdata struct {
	bytes   []byte
	integer uint16
}

type uint32testdata struct {
	bytes   []byte
	integer uint32
}

type uint64testdata struct {
	bytes   []byte
	integer uint64
}

var uint8data = []uint8testdata{
	{0x00, 0},
	{0x09, 9},
	{0x99, 99},
}

var uint8data_special = []uint8testdata{
	// Hi order digits dropped below
	{0x00, 100},
	{0x55, math.MaxUint8},
}

var uint8data_overflow = []uint8testdata{
	// Single digit overflow returned full zero result
	{0x0F, 0},
	{0xF0, 0},
	{0xFF, 0},
}

var uint16data = []uint16testdata{
	{[]byte{0x00, 0x00}, 0},
	{[]byte{0x00, 0x99}, 99},
	{[]byte{0x09, 0x99}, 999},
	{[]byte{0x99, 0x99}, 9999},
}

var uint16data_special_from = []uint16testdata{
	// Hi order digits dropped below
	{[]byte{0x00, 0x00}, 10000},
	{[]byte{0x55, 0x35}, math.MaxUint16},
}

var uint16data_special_to = []uint16testdata{
	{[]byte{0x99}, 99},
	// Single digit overflow returned full zero result
	{[]byte{0x00, 0x0F}, 0},
	{[]byte{0x00, 0xF0}, 0},
	{[]byte{0x00, 0xFF}, 0},
	{[]byte{0x0F, 0x00}, 0},
	{[]byte{0xF0, 0x00}, 0},
	{[]byte{0xFF, 0x00}, 0},
}

var uint32data = []uint32testdata{
	{[]byte{0x00, 0x00, 0x00, 0x00}, 0},
	{[]byte{0x00, 0x00, 0x00, 0x99}, 99},
	{[]byte{0x00, 0x00, 0x09, 0x99}, 999},
	{[]byte{0x00, 0x00, 0x99, 0x99}, 9999},
	{[]byte{0x00, 0x09, 0x99, 0x99}, 99999},
	{[]byte{0x00, 0x99, 0x99, 0x99}, 999999},
	{[]byte{0x09, 0x99, 0x99, 0x99}, 9999999},
	{[]byte{0x99, 0x99, 0x99, 0x99}, 99999999},
}

var uint32data_special_from = []uint32testdata{
	// Hi order digits dropped below
	{[]byte{0x00, 0x00, 0x00, 0x00}, 100000000},
	{[]byte{0x94, 0x96, 0x72, 0x95}, math.MaxUint32},
}

var uint32data_special_to = []uint32testdata{
	{[]byte{0x99}, 99},
	// Single digit overflow returned full zero result
	{[]byte{0x00, 0x00, 0x00, 0x0F}, 0},
	{[]byte{0x00, 0x00, 0x00, 0xF0}, 0},
	{[]byte{0x00, 0x00, 0x00, 0xFF}, 0},
	{[]byte{0x00, 0x00, 0x0F, 0x00}, 0},
	{[]byte{0x00, 0x00, 0xF0, 0x00}, 0},
	{[]byte{0x00, 0x00, 0xFF, 0x00}, 0},
	{[]byte{0x00, 0x0F, 0x00, 0x00}, 0},
	{[]byte{0x00, 0xF0, 0x00, 0x00}, 0},
	{[]byte{0x00, 0xFF, 0x00, 0x00}, 0},
	{[]byte{0x0F, 0x00, 0x00, 0x00}, 0},
	{[]byte{0xF0, 0x00, 0x00, 0x00}, 0},
	{[]byte{0xFF, 0x00, 0x00, 0x00}, 0},
}

var uint64data = []uint64testdata{
	{[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, 0},
	{[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x99}, 99},
	{[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x09, 0x99}, 999},
	{[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x99, 0x99}, 9999},
	{[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x09, 0x99, 0x99}, 99999},
	{[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x99, 0x99, 0x99}, 999999},
	{[]byte{0x00, 0x00, 0x00, 0x00, 0x09, 0x99, 0x99, 0x99}, 9999999},
	{[]byte{0x00, 0x00, 0x00, 0x00, 0x99, 0x99, 0x99, 0x99}, 99999999},
	{[]byte{0x00, 0x00, 0x00, 0x09, 0x99, 0x99, 0x99, 0x99}, 999999999},
	{[]byte{0x00, 0x00, 0x00, 0x99, 0x99, 0x99, 0x99, 0x99}, 9999999999},
	{[]byte{0x00, 0x00, 0x09, 0x99, 0x99, 0x99, 0x99, 0x99}, 99999999999},
	{[]byte{0x00, 0x00, 0x99, 0x99, 0x99, 0x99, 0x99, 0x99}, 999999999999},
	{[]byte{0x00, 0x09, 0x99, 0x99, 0x99, 0x99, 0x99, 0x99}, 9999999999999},
	{[]byte{0x00, 0x99, 0x99, 0x99, 0x99, 0x99, 0x99, 0x99}, 99999999999999},
	{[]byte{0x09, 0x99, 0x99, 0x99, 0x99, 0x99, 0x99, 0x99}, 999999999999999},
	{[]byte{0x99, 0x99, 0x99, 0x99, 0x99, 0x99, 0x99, 0x99}, 9999999999999999},
}

var uint64data_special_from = []uint64testdata{
	// Hi order digits dropped below
	{[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, 10000000000000000},
	{[]byte{0x67, 0x44, 0x07, 0x37, 0x09, 0x55, 0x16, 0x15}, math.MaxUint64},
}

var uint64data_special_to = []uint64testdata{
	{[]byte{0x99}, 99},
	// Single digit overflow returned full zero result
	{[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0F}, 0},
	{[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xF0}, 0},
	{[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xFF}, 0},
	{[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0F, 0x00}, 0},
	{[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xF0, 0x00}, 0},
	{[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xFF, 0x00}, 0},
	{[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x0F, 0x00, 0x00}, 0},
	{[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0xF0, 0x00, 0x00}, 0},
	{[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0xFF, 0x00, 0x00}, 0},
	{[]byte{0x00, 0x00, 0x00, 0x00, 0x0F, 0x00, 0x00, 0x00}, 0},
	{[]byte{0x00, 0x00, 0x00, 0x00, 0xF0, 0x00, 0x00, 0x00}, 0},
	{[]byte{0x00, 0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00}, 0},
	{[]byte{0x00, 0x00, 0x00, 0xF0, 0x00, 0x00, 0x00, 0x00}, 0},
	{[]byte{0x00, 0x00, 0x00, 0x0F, 0x00, 0x00, 0x00, 0x00}, 0},
	{[]byte{0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0x00}, 0},
	{[]byte{0x00, 0x00, 0x0F, 0x00, 0x00, 0x00, 0x00, 0x00}, 0},
	{[]byte{0x00, 0x00, 0xF0, 0x00, 0x00, 0x00, 0x00, 0x00}, 0},
	{[]byte{0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0x00, 0x00}, 0},
	{[]byte{0x00, 0x0F, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, 0},
	{[]byte{0x00, 0xF0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, 0},
	{[]byte{0x00, 0xFF, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, 0},
	{[]byte{0x0F, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, 0},
	{[]byte{0xF0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, 0},
	{[]byte{0xFF, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, 0},
}

func TestFromUint8(t *testing.T) {
	for _, data := range uint8data {
		assert.Equal(t, data.bytes, FromUint8(data.integer))
	}
	for _, data := range uint8data_special {
		assert.Equal(t, data.bytes, FromUint8(data.integer))
	}
}

func TestFromUint16(t *testing.T) {
	for _, data := range uint16data {
		assert.Equal(t, data.bytes, FromUint16(data.integer))
	}
	for _, data := range uint16data_special_from {
		assert.Equal(t, data.bytes, FromUint16(data.integer))
	}
}

func TestFromUint32(t *testing.T) {
	for _, data := range uint32data {
		assert.Equal(t, data.bytes, FromUint32(data.integer))
	}
	for _, data := range uint32data_special_from {
		assert.Equal(t, data.bytes, FromUint32(data.integer))
	}
}

func TestFromUint64(t *testing.T) {
	for _, data := range uint64data {
		assert.Equal(t, data.bytes, FromUint64(data.integer))
	}
	for _, data := range uint64data_special_from {
		assert.Equal(t, data.bytes, FromUint64(data.integer))
	}
}

func TestToUint8(t *testing.T) {
	for _, data := range uint8data {
		assert.Equal(t, data.integer, ToUint8(data.bytes))
	}
	for _, data := range uint8data_overflow {
		assert.Equal(t, data.integer, ToUint8(data.bytes))
	}
}

func TestToUint16(t *testing.T) {
	for _, data := range uint16data {
		assert.Equal(t, data.integer, ToUint16(data.bytes))
	}
	for _, data := range uint16data_special_to {
		assert.Equal(t, data.integer, ToUint16(data.bytes))
	}
}

func TestToUint32(t *testing.T) {
	for _, data := range uint32data {
		assert.Equal(t, data.integer, ToUint32(data.bytes))
	}
	for _, data := range uint32data_special_to {
		assert.Equal(t, data.integer, ToUint32(data.bytes))
	}
}

func TestToUint64(t *testing.T) {
	for _, data := range uint64data {
		assert.Equal(t, data.integer, ToUint64(data.bytes))
	}
	for _, data := range uint64data_special_to {
		assert.Equal(t, data.integer, ToUint64(data.bytes))
	}
}
