package sflow

import "errors"

var (
	ErrEncodingRecord = errors.New("sflow: failed to encode record")
	ErrDecodingRecord = errors.New("sflow: failed to decode record")
)

type Record interface {
	RecordType() int
	//encode(w io.Writer) error
}
