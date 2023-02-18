// Code generated by "genabi"; DO NOT EDIT.
package example

import (
	"github.com/indexsupply/x/abi"
	"github.com/indexsupply/x/abi/schema"
	"math/big"
)

type Transfer struct {
	From    [20]byte
	To      [20]byte
	Id      *big.Int
	Extra   [3][2]uint8
	Details [][]Details
}

func NewTransfer(item abi.Item) Transfer {
	x := Transfer{}
	extraItem0 := item.At(0)
	extra0 := [3][2]uint8{}
	for i0 := 0; i0 < extraItem0.Len(); i0++ {
		extraItem1 := extraItem0.At(i0)
		extra1 := [2]uint8{}
		for i1 := 0; i1 < extraItem1.Len(); i1++ {
			extra1[i1] = extraItem1.Uint8()
		}
		extra0[i0] = extra1
	}
	x.Extra = extra0
	detailsItem0 := item.At(1)
	details0 := make([][]Details, detailsItem0.Len())
	for i0 := 0; i0 < detailsItem0.Len(); i0++ {
		detailsItem1 := detailsItem0.At(i0)
		details1 := make([]Details, detailsItem1.Len())
		for i1 := 0; i1 < detailsItem1.Len(); i1++ {
			details1[i1] = NewDetails(detailsItem1.At(i1))
		}
		details0[i0] = details1
	}
	x.Details = details0
	return x
}

type Details struct {
	Other [20]byte
	Key   [32]byte
	Value []byte
	Geo   Geo
}

func NewDetails(item abi.Item) Details {
	x := Details{}
	x.Other = item.At(0).Address()
	x.Key = item.At(1).Bytes32()
	x.Value = item.At(2).Bytes()
	x.Geo = NewGeo(item.At(3))
	return x
}

type Geo struct {
	X uint8
	Y uint8
}

func NewGeo(item abi.Item) Geo {
	x := Geo{}
	x.X = item.At(0).Uint8()
	x.Y = item.At(1).Uint8()
	return x
}

// transfer(address,address,uint256,uint8[2][3],(address,bytes32,bytes,(uint8,uint8))[][])
// 70711f9efd2d568665592c1d6245e892eab7d9e56c76714682526066ca69d65e
var (
	TransferSignature = [32]byte{0x70, 0x71, 0x1f, 0x9e, 0xfd, 0x2d, 0x56, 0x86, 0x65, 0x59, 0x2c, 0x1d, 0x62, 0x45, 0xe8, 0x92, 0xea, 0xb7, 0xd9, 0xe5, 0x6c, 0x76, 0x71, 0x46, 0x82, 0x52, 0x60, 0x66, 0xca, 0x69, 0xd6, 0x5e}
	TransferSchema    = schema.Parse("(uint8[2][3],(address,bytes32,bytes,(uint8,uint8))[][])")
)

func MatchTransfer(l abi.Log) (Transfer, bool) {
	if TransferSignature != l.Topics[0] {
		return Transfer{}, false
	}
	item := abi.Decode(l.Data, TransferSchema)
	res := NewTransfer(item)
	res.From = abi.Bytes(l.Topics[1][:]).Address()
	res.To = abi.Bytes(l.Topics[2][:]).Address()
	res.Id = abi.Bytes(l.Topics[3][:]).BigInt()
	return res, true
}
