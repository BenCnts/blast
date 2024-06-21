// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package native

import (
	"encoding/json"

	"github.com/BenCnts/blast-geth/common"
	"github.com/BenCnts/blast-geth/common/hexutil"
)

var _ = (*flatCallResultMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (f flatCallResult) MarshalJSON() ([]byte, error) {
	type flatCallResult struct {
		Address *common.Address `json:"address,omitempty"`
		Code    *hexutil.Bytes  `json:"code,omitempty"`
		GasUsed *hexutil.Uint64 `json:"gasUsed,omitempty"`
		Output  *hexutil.Bytes  `json:"output,omitempty"`
	}
	var enc flatCallResult
	enc.Address = f.Address
	enc.Code = (*hexutil.Bytes)(f.Code)
	enc.GasUsed = (*hexutil.Uint64)(f.GasUsed)
	enc.Output = (*hexutil.Bytes)(f.Output)
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (f *flatCallResult) UnmarshalJSON(input []byte) error {
	type flatCallResult struct {
		Address *common.Address `json:"address,omitempty"`
		Code    *hexutil.Bytes  `json:"code,omitempty"`
		GasUsed *hexutil.Uint64 `json:"gasUsed,omitempty"`
		Output  *hexutil.Bytes  `json:"output,omitempty"`
	}
	var dec flatCallResult
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Address != nil {
		f.Address = dec.Address
	}
	if dec.Code != nil {
		f.Code = (*[]byte)(dec.Code)
	}
	if dec.GasUsed != nil {
		f.GasUsed = (*uint64)(dec.GasUsed)
	}
	if dec.Output != nil {
		f.Output = (*[]byte)(dec.Output)
	}
	return nil
}
