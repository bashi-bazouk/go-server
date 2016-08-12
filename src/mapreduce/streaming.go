package mapreduce

import (
	"bufio"
	"strings"
	"io"
)

type KVPair struct {
	Key string
	Value string
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

func ReadStream(r io.Reader, c Consumer) {
	br := bufio.NewReader(r)
	line, err := br.ReadString("\n")
	for !err {
		split := strings.SplitN(line, "\t", 2)
		if len(split) < 2 {
			append(split, "")
		}
		c <- KVPair { Key: split[0], Value: split[1] }
		line, err = br.ReadString("\n")
	}
	close(c)
}

func WriteStream(w io.Writer, p Producer) {
	bw := bufio.NewWriter(w)
	for kvPair := range p {
		bw.WriteString(kvPair.Key + "\t")
		bw.Flush()
		bw.WriteString(kvPair.Value + "\n")
		bw.Flush()
	}
}