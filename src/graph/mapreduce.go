package graph

type KVPair struct {
	Key Node
	Value Object
}

type KVPairs chan KVPair
type Producer KVPairs
type Consumer KVPairs

type Mapper func(Producer, Consumer)
type Reducer func(Producer, Consumer)

type Primitive struct {
	Mapper Mapper
	Reducer Reducer
}