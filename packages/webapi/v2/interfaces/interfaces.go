package interfaces

import (
	"github.com/iotaledger/wasp/packages/chain"
	"github.com/iotaledger/wasp/packages/isc"
	"github.com/iotaledger/wasp/packages/registry"
	"github.com/iotaledger/wasp/packages/webapi/v2/dto"
	"github.com/pangpanglabs/echoswagger/v2"
)

type APIController interface {
	RegisterPublic(publicAPI echoswagger.ApiGroup)
	RegisterAdmin(adminAPI echoswagger.ApiGroup)
}

type Chain interface {
	ActivateChain(chainID *isc.ChainID) error
	DeactivateChain(chainID *isc.ChainID) error
	GetChainByID(chainID *isc.ChainID) chain.Chain
}

type Registry interface {
	GetChainRecordByChainID(chainID *isc.ChainID) (*registry.ChainRecord, error)
}

type Node interface {
	GetNodeInfo(chain chain.Chain) (*dto.ChainNodeInfo, error)
}
