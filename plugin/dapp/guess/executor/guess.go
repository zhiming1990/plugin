// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package executor

import (
	log "github.com/33cn/chain33/common/log/log15"
	drivers "github.com/33cn/chain33/system/dapp"
	"github.com/33cn/chain33/types"
	pkt "github.com/33cn/plugin/plugin/dapp/guess/types"
)

var logger = log.New("module", "execs.guess")

var driverName = pkt.GuessX

func init() {
	ety := types.LoadExecutorType(driverName)
	ety.InitFuncList(types.ListMethod(&Guess{}))
}

type subConfig struct {
	ParaRemoteGrpcClient string `json:"paraRemoteGrpcClient"`
}

var cfg subConfig


// Init Guess
func Init(name string, sub []byte) {
	driverName := GetName()
	if name != driverName {
		panic("system dapp can't be rename")
	}
	if sub != nil {
		types.MustDecode(sub, &cfg)
	}
	drivers.Register(driverName, newGuessGame, types.GetDappFork(driverName, "Enable"))
}

type Guess struct {
	drivers.DriverBase
}

func newGuessGame() drivers.Driver {
	t := &Guess{}
	t.SetChild(t)
	t.SetExecutorType(types.LoadExecutorType(driverName))
	return t
}

func GetName() string {
	return newGuessGame().GetName()
}

func (g *Guess) GetDriverName() string {
	return pkt.GuessX
}

/*
// GetPayloadValue GuessAction
func (g *Guess) GetPayloadValue() types.Message {
	return &pkt.GuessGameAction{}
}*/

// CheckReceiptExecOk return true to check if receipt ty is ok
func (g *Guess) CheckReceiptExecOk() bool {
	return true
}
