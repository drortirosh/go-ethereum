// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package tests

import (
	"encoding/json"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
)

var _ = (*stAuthorizationMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (s stAuthorization) MarshalJSON() ([]byte, error) {
	type stAuthorization struct {
		ChainID *math.HexOrDecimal256
		Address common.Address        `json:"address" gencodec:"required"`
		Nonce   math.HexOrDecimal64   `json:"nonce" gencodec:"required"`
		V       *math.HexOrDecimal256 `json:"v" gencodec:"required"`
		R       *math.HexOrDecimal256 `json:"r" gencodec:"required"`
		S       *math.HexOrDecimal256 `json:"s" gencodec:"required"`
	}
	var enc stAuthorization
	enc.ChainID = (*math.HexOrDecimal256)(s.ChainID)
	enc.Address = s.Address
	enc.Nonce = math.HexOrDecimal64(s.Nonce)
	enc.V = (*math.HexOrDecimal256)(s.V)
	enc.R = (*math.HexOrDecimal256)(s.R)
	enc.S = (*math.HexOrDecimal256)(s.S)
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (s *stAuthorization) UnmarshalJSON(input []byte) error {
	type stAuthorization struct {
		ChainID *math.HexOrDecimal256
		Address *common.Address       `json:"address" gencodec:"required"`
		Nonce   *math.HexOrDecimal64  `json:"nonce" gencodec:"required"`
		V       *math.HexOrDecimal256 `json:"v" gencodec:"required"`
		R       *math.HexOrDecimal256 `json:"r" gencodec:"required"`
		S       *math.HexOrDecimal256 `json:"s" gencodec:"required"`
	}
	var dec stAuthorization
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.ChainID != nil {
		s.ChainID = (*big.Int)(dec.ChainID)
	}
	if dec.Address == nil {
		return errors.New("missing required field 'address' for stAuthorization")
	}
	s.Address = *dec.Address
	if dec.Nonce == nil {
		return errors.New("missing required field 'nonce' for stAuthorization")
	}
	s.Nonce = uint64(*dec.Nonce)
	if dec.V == nil {
		return errors.New("missing required field 'v' for stAuthorization")
	}
	s.V = (*big.Int)(dec.V)
	if dec.R == nil {
		return errors.New("missing required field 'r' for stAuthorization")
	}
	s.R = (*big.Int)(dec.R)
	if dec.S == nil {
		return errors.New("missing required field 's' for stAuthorization")
	}
	s.S = (*big.Int)(dec.S)
	return nil
}
