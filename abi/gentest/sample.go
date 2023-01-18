package gentest

import (
	"math/big"

	"github.com/indexsupply/x/abi"
)

var EEvent = abi.Event{
	Name: "e",
	Inputs: []abi.Input{
		abi.Input{
			Name: "address",
			Type: "address",
		},
		abi.Input{
			Name: "address_list",
			Type: "address[]",
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
			Inputs: []abi.Input{
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
					Inputs: []abi.Input{
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
	it := x.it.At(1)
	res := make([][20]byte, it.Len())
	for i, v := range it.List() {
		res[i] = v.Address()
	}
	return res
}

func (x *E) Bool() bool {
	return x.it.At(2).Bool()
}

func (x *E) BoolList() []bool {
	it := x.it.At(3)
	res := make([]bool, it.Len())
	for i, v := range it.List() {
		res[i] = v.Bool()
	}
	return res
}

func (x *E) Bytes() []byte {
	return x.it.At(4).Bytes()
}

func (x *E) BytesList() [][]byte {
	it := x.it.At(5)
	res := make([][]byte, it.Len())
	for i, v := range it.List() {
		res[i] = v.Bytes()
	}
	return res
}

func (x *E) String() string {
	return x.it.At(6).String()
}

func (x *E) StringList() []string {
	it := x.it.At(7)
	res := make([]string, it.Len())
	for i, v := range it.List() {
		res[i] = v.String()
	}
	return res
}

func (x *E) Uint8() uint8 {
	return x.it.At(8).Uint8()
}

func (x *E) Uint8List() []uint8 {
	it := x.it.At(9)
	res := make([]uint8, it.Len())
	for i, v := range it.List() {
		res[i] = v.Uint8()
	}
	return res
}

func (x *E) Uint64() uint64 {
	return x.it.At(10).Uint64()
}

func (x *E) Uint64List() []uint64 {
	it := x.it.At(11)
	res := make([]uint64, it.Len())
	for i, v := range it.List() {
		res[i] = v.Uint64()
	}
	return res
}

func (x *E) Uint256() *big.Int {
	return x.it.At(12).BigInt()
}

func (x *E) Uint256List() []*big.Int {
	it := x.it.At(13)
	res := make([]*big.Int, it.Len())
	for i, v := range it.List() {
		res[i] = v.BigInt()
	}
	return res
}

func (x *E) I2() *I2 {
	i := x.it.At(14)
	return &I2{&i}
}

func (x *I2) F1() [20]byte {
	return x.it.At(0).Address()
}

func (x *I2) F2() [][20]byte {
	it := x.it.At(1)
	res := make([][20]byte, it.Len())
	for i, v := range it.List() {
		res[i] = v.Address()
	}
	return res
}

func (x *I2) F3() *F3 {
	i := x.it.At(2)
	return &F3{&i}
}

func (x *F3) F4() [20]byte {
	return x.it.At(0).Address()
}
