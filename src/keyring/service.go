package keyring

/*
	Start a local blockchain service

*/

import (
    //"encoding/hex"
    //"errors"
    //"fmt"
    "github.com/skycoin/skycoin/src/coin"
    //"github.com/skycoin/skycoin/src/keyring"

    //"log"
    //"math/rand"
    //"encoding/hex"
)


type BlockChainService struct {
	PendingBlock coin.Block
	BC *coin.BlockChain
	PendingTransactions []coin.Transaction
}


func (self *BlockChainService) Run() {
	//TODO, set genesis address

	seckey_hex := "5a42c0643bdb465d90bf673b99c14f5fa02db71513249d904573d2b8b63d353d"
   	seckey := coin.SecKeyFromHex(seckey_hex)
    pubkey := coin.PubKeyFromSecKey(seckey)
    address := coin.AddressFromPubKey(pubkey) //genesis address

	self.BC = coin.NewBlockChain(address)


}

func (self *BlockChainService) InsertTransaction(transaction coin.Transaction) {
	self.PendingTransactions = append(self.PendingTransactions, transaction)
}


