package handlers

import "github.com/blockfreight/go-bftx/lib/app/blockchain"

func GetInfo() (interface{}, error) {
	info, err := blockchain.GetInfo()
	if err != nil {
		return nil, err
	}

	return info, nil
}
