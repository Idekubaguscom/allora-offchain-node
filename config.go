package main

import (
	"allora_offchain_node/lib"
	reputerCoinGecko "allora_offchain_node/pkg/reputer_coingecko_l1_norm"
	worker10min "allora_offchain_node/pkg/worker_coin_predictor_10min_eth"
	worker20min "allora_offchain_node/pkg/worker_coin_predictor_20min"
)

var UserConfig = lib.UserConfig{
	Wallet: lib.WalletConfig{
		AddressKeyName:           "secret",               // load a address by key from the keystore
		AddressRestoreMnemonic:   "secret",               // mnemonic for the allora account
		AddressAccountPassphrase: "secret",               // passphrase for the allora account
		AlloraHomeDir:            "/home/allora/.allora", // home directory for the allora keystore
		Gas:                      "1000000",              // gas to use for the allora client in uallo
		GasAdjustment:            1.0,                    // gas adjustment to use for the allora client
		SubmitTx:                 true,                   // set to false to run in dry-run processes without committing to the chain. useful for development and testing
		LoopWithinWindowSeconds:  5,
		NodeRpc:                  "http://rpc.allora.network",
		MaxRetries:               3,
		MinDelay:                 1,
		MaxDelay:                 6,
		EarlyArrivalPercent:      0.6,
		LateArrivalPercent:       0.1,
	},
	Worker: []lib.WorkerConfig{
		{
			TopicId:             1,
			InferenceEntrypoint: worker10min.NewAlloraEntrypoint(),
			ForecastEntrypoint:  nil,
			ExtraData: map[string]string{
				"inferenceEndpoint": "http://localhost:8000/inference",
				"token":             "ETH",
				"forecastEndpoint":  "http://localhost:8000/forecast",
			},
		},
		{
			TopicId:             2,
			InferenceEntrypoint: worker20min.NewAlloraEntrypoint(),
			ForecastEntrypoint:  worker20min.NewAlloraEntrypoint(),
		},
	},
	Reputer: []lib.ReputerConfig{
		{
			TopicId:           1,
			ReputerEntrypoint: reputerCoinGecko.NewAlloraEntrypoint(),
			MinStake:          1000000,
		},
	},
}
