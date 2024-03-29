package client

import (
	pkts1 "github.com/energomonitor/bisquitt/packets1"
	"github.com/energomonitor/bisquitt/transactions"
	"github.com/energomonitor/bisquitt/util"
)

// Transactions states constants type.
type transactionState int

// Transactions states constants.
const (
	transactionDone transactionState = iota
	awaitingPuback
	awaitingPubrec
	awaitingPubrel
	awaitingPubcomp
	awaitingDisconnect
	awaitingPingresp
)

type transaction struct {
	*transactions.RetryTransaction
	client *Client
	log    util.Logger
}

// Transactions involving DISCONNECT packet.
type transactionWithDisconnect interface {
	Disconnect(*pkts1.Disconnect)
}

// Transactions involving PINGRESP packet.
type transactionWithPingresp interface {
	Pingresp(pingresp *pkts1.Pingresp)
}
