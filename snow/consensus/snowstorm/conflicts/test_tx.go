// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package conflicts

import (
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/choices"
)

var (
	_ Tx = &TestTx{}
)

// TestTx is a useful test tx
type TestTx struct {
	choices.TestDecidable

	TransitionV   Transition
	EpochV        uint32
	RestrictionsV []ids.ID
	VerifyV       error
	BytesV        []byte
}

// Transition implements the Tx interface
func (t *TestTx) Transition() Transition { return t.TransitionV }

// Epoch implements the Tx interface
func (t *TestTx) Epoch() uint32 { return t.EpochV }

// Restrictions implements the Tx interface
func (t *TestTx) Restrictions() []ids.ID { return t.RestrictionsV }

// Verify implements the Tx interface
func (t *TestTx) Verify() error { return t.VerifyV }

// Bytes implements the Tx interface
func (t *TestTx) Bytes() []byte { return t.BytesV }
