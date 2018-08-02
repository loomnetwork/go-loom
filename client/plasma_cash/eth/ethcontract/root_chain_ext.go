// +build evm

package ethcontract

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
)

func (rc *RootChainFilterer) DepositEventData(txReceipt *types.Receipt) (*RootChainDeposit, error) {
	contractABI, err := abi.JSON(strings.NewReader(RootChainABI))
	if err != nil {
		return nil, err
	}
	eventTopic := contractABI.Events["Deposit"].Id()
	eventData := new(RootChainDeposit)
	for _, log := range txReceipt.Logs {
		for _, topic := range log.Topics {
			if topic.Hex() == eventTopic.Hex() {
				if err := rc.contract.UnpackLog(eventData, "Deposit", *log); err != nil {
					return nil, err
				}
				return eventData, nil
			}
		}
	}
	return nil, nil
}

func (rc *RootChainFilterer) ChallengedExitEventData(txReceipt *types.Receipt) (*RootChainChallengedExit, error) {
	contractABI, err := abi.JSON(strings.NewReader(RootChainABI))
	if err != nil {
		return nil, err
	}
	eventTopic := contractABI.Events["ChallengedExit"].Id()
	eventData := new(RootChainChallengedExit)
	for _, log := range txReceipt.Logs {
		for _, topic := range log.Topics {
			if topic.Hex() == eventTopic.Hex() {
				if err := rc.contract.UnpackLog(eventData, "ChallengedExit", *log); err != nil {
					return nil, err
				}
				return eventData, nil
			}
		}
	}
	return nil, nil
}
