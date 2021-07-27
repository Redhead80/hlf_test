package main

import (
	"github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric-protos-go/peer"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

type LedgerClient struct {
	client *ledger.Client
}

func NewLedgerClient(sdk *fabsdk.FabricSDK, configuration ClientConfiguration) (*LedgerClient, error) {
	clientChannelContext := sdk.ChannelContext(
		configuration.Channel,
		fabsdk.WithUser(configuration.UserName),
		fabsdk.WithOrg(configuration.Organization),
	)

	client, err := ledger.New(clientChannelContext)
	if err != nil {
		return nil, err
	}
	return &LedgerClient{
		client: client,
	}, nil
}

func (lc *LedgerClient) QueryInfo() (*fab.BlockchainInfoResponse, error) {
	return lc.client.QueryInfo()
}

func (lc *LedgerClient) QueryBlockByNumber(blockNumber uint64) (*common.Block, error) {
	return lc.client.QueryBlock(blockNumber)
}

func (lc *LedgerClient) QueryBlockByHash(blockHash []byte) (*common.Block, error) {
	return lc.client.QueryBlockByHash(blockHash)
}

func (lc *LedgerClient) QueryBlockByTxID(txID fab.TransactionID) (*common.Block, error) {
	return lc.client.QueryBlockByTxID(txID)
}

func (lc *LedgerClient) QueryTransactionByID(txID fab.TransactionID) (*peer.ProcessedTransaction, error) {
	return lc.client.QueryTransaction(txID)
}

func (lc *LedgerClient) QueryConfigBlock() (*common.Block, error) {
	return lc.client.QueryConfigBlock()
}

func (lc *LedgerClient) QueryConfig() (fab.ChannelCfg, error) {
	return lc.client.QueryConfig()
}
