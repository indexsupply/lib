// Code generated by "genabi"; DO NOT EDIT.

package gentest

import (
	"github.com/indexsupply/x/abi"
	"math/big"
)

var EEvent = abi.Event{
	Name: "e",
	Inputs: []abi.Input{
		abi.Input{
			Indexed: true,
			Name:    "address",
			Type:    "address",
		},
		abi.Input{
			Name: "address_list",
			Type: "address[]",
		},
		abi.Input{
			Name: "address_list_4",
			Type: "address[4]",
		},
		abi.Input{
			Name: "bool",
			Type: "bool",
		},
		abi.Input{
			Name: "bool_list",
			Type: "bool[]",
		},
		abi.Input{
			Name: "bytes",
			Type: "bytes",
		},
		abi.Input{
			Name: "bytes_list",
			Type: "bytes[]",
		},
		abi.Input{
			Name: "string",
			Type: "string",
		},
		abi.Input{
			Name: "string_list",
			Type: "string[]",
		},
		abi.Input{
			Name: "uint8",
			Type: "uint8",
		},
		abi.Input{
			Name: "uint8_list",
			Type: "uint8[]",
		},
		abi.Input{
			Name: "uint64",
			Type: "uint64",
		},
		abi.Input{
			Name: "uint64_list",
			Type: "uint64[]",
		},
		abi.Input{
			Name: "uint256",
			Type: "uint256",
		},
		abi.Input{
			Name: "uint256_list",
			Type: "uint256[]",
		},
		abi.Input{
			Name: "i2",
			Type: "tuple",
			Components: []abi.Input{
				abi.Input{
					Name: "f1",
					Type: "address",
				},
				abi.Input{
					Name: "f2",
					Type: "address[]",
				},
				abi.Input{
					Name: "f3",
					Type: "tuple",
					Components: []abi.Input{
						abi.Input{
							Name: "f4",
							Type: "address",
						},
					},
				},
			},
		},
	},
}

type E struct {
	it *abi.Item
}

func MatchE(l abi.Log) (*E, bool) {
	i, ok := abi.Match(l, EEvent)
	return &E{&i}, ok
}

type I2 struct {
	it *abi.Item
}

type F3 struct {
	it *abi.Item
}

func (x *E) Address() [20]byte {
	return x.it.At(0).Address()
}

func (x *E) AddressList() [][20]byte {
	it0 := x.it.At(1)
	res0 := make([][20]byte, it0.Len())
	for i0 := 0; i0 < it0.Len(); i0++ {
		res0[i0] = it0.At(i0).Address()
	}
	return res0
}

func (x *E) AddressList4() [4][20]byte {
	it0 := x.it.At(2)

	if it0.Len() != 4 {
		panic("genabi: E.AddressList4 array size mismatch")
	}
	var res0 [4][20]byte
	for i0 := 0; i0 < it0.Len(); i0++ {
		res0[i0] = it0.At(i0).Address()
	}
	return res0
}

func (x *E) Bool() bool {
	return x.it.At(3).Bool()
}

func (x *E) BoolList() []bool {
	it0 := x.it.At(4)
	res0 := make([]bool, it0.Len())
	for i0 := 0; i0 < it0.Len(); i0++ {
		res0[i0] = it0.At(i0).Bool()
	}
	return res0
}

func (x *E) Bytes() []byte {
	return x.it.At(5).Bytes()
}

func (x *E) BytesList() [][]byte {
	it0 := x.it.At(6)
	res0 := make([][]byte, it0.Len())
	for i0 := 0; i0 < it0.Len(); i0++ {
		res0[i0] = it0.At(i0).Bytes()
	}
	return res0
}

func (x *E) String() string {
	return x.it.At(7).String()
}

func (x *E) StringList() []string {
	it0 := x.it.At(8)
	res0 := make([]string, it0.Len())
	for i0 := 0; i0 < it0.Len(); i0++ {
		res0[i0] = it0.At(i0).String()
	}
	return res0
}

func (x *E) Uint8() uint8 {
	return x.it.At(9).Uint8()
}

func (x *E) Uint8List() []uint8 {
	it0 := x.it.At(10)
	res0 := make([]uint8, it0.Len())
	for i0 := 0; i0 < it0.Len(); i0++ {
		res0[i0] = it0.At(i0).Uint8()
	}
	return res0
}

func (x *E) Uint64() uint64 {
	return x.it.At(11).Uint64()
}

func (x *E) Uint64List() []uint64 {
	it0 := x.it.At(12)
	res0 := make([]uint64, it0.Len())
	for i0 := 0; i0 < it0.Len(); i0++ {
		res0[i0] = it0.At(i0).Uint64()
	}
	return res0
}

func (x *E) Uint256() *big.Int {
	return x.it.At(13).BigInt()
}

func (x *E) Uint256List() []*big.Int {
	it0 := x.it.At(14)
	res0 := make([]*big.Int, it0.Len())
	for i0 := 0; i0 < it0.Len(); i0++ {
		res0[i0] = it0.At(i0).BigInt()
	}
	return res0
}

func (x *E) I2() *I2 {
	i := x.it.At(15)
	return &I2{&i}
}

func (x *E) F1() [20]byte {
	return x.it.At(0).Address()
}

func (x *E) F2() [][20]byte {
	it0 := x.it.At(1)
	res0 := make([][20]byte, it0.Len())
	for i0 := 0; i0 < it0.Len(); i0++ {
		res0[i0] = it0.At(i0).Address()
	}
	return res0
}

func (x *E) F3() *F3 {
	i := x.it.At(2)
	return &F3{&i}
}

func (x *E) F4() [20]byte {
	return x.it.At(0).Address()
}

var FooEvent = abi.Event{
	Name: "foo",
	Inputs: []abi.Input{
		abi.Input{
			Indexed: true,
			Name:    "bar",
			Type:    "uint64",
		},
		abi.Input{
			Name: "baz",
			Type: "string",
		},
	},
}

type Foo struct {
	it *abi.Item
}

func MatchFoo(l abi.Log) (*Foo, bool) {
	i, ok := abi.Match(l, FooEvent)
	return &Foo{&i}, ok
}

func (x *Foo) Bar() uint64 {
	return x.it.At(0).Uint64()
}

func (x *Foo) Baz() string {
	return x.it.At(1).String()
}

var BarEvent = abi.Event{
	Name: "bar",
	Inputs: []abi.Input{
		abi.Input{
			Indexed: true,
			Name:    "bar",
			Type:    "uint64",
		},
		abi.Input{
			Name: "baz",
			Type: "string[][][]",
		},
	},
}

type Bar struct {
	it *abi.Item
}

func MatchBar(l abi.Log) (*Bar, bool) {
	i, ok := abi.Match(l, BarEvent)
	return &Bar{&i}, ok
}

func (x *Bar) Bar() uint64 {
	return x.it.At(0).Uint64()
}

func (x *Bar) Baz() [][][]string {
	it0 := x.it.At(1)
	res0 := make([][][]string, it0.Len())
	for i0 := 0; i0 < it0.Len(); i0++ {
		it1 := it0.At(i0)
		res1 := make([][]string, it1.Len())
		for i1 := 0; i1 < it1.Len(); i1++ {
			it2 := it1.At(i1)
			res2 := make([]string, it2.Len())
			for i2 := 0; i2 < it2.Len(); i2++ {
				res2[i2] = it2.At(i2).String()
			}
			res1[i1] = res2
		}
		res0[i0] = res1
	}
	return res0
}
