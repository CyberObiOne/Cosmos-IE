package exporter

import (
        "fmt"
        "net/http"
        "go.uber.org/zap"

        sdk "github.com/cosmos/cosmos-sdk/types"
	iris "github.com/irisnet/irishub/address"

        terra "github.com/terra-project/core/types"
//	kava "github.com/kava-labs/kava/app"
	emoney "github.com/e-money/em-ledger/types"

        "github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
)

func Go(chain string, port string) {

        log,_ := zap.NewDevelopment()
        defer log.Sync()

	setConfig(chain)

	http.Handle("/metrics", promhttp.Handler())
	go Start(chain, log)

        err := http.ListenAndServe(":" +port, nil)
        // log
        if err != nil {
                // handle error
                log.Fatal("HTTP Handle", zap.Bool("Success", false), zap.String("err", fmt.Sprint(err),))
        } else {
                log.Info("HTTP Handle", zap.Bool("Success", true), zap.String("err", "nil"), zap.String("Listen&Serve", "Prometheus Handler(Port: " +port +")"),)
        }

}

func setConfig(chain string) {

	config := sdk.GetConfig()

	switch chain {
	case "iris":
		iris.ConfigureBech32Prefix()
	case "band":
		bech32MainPrefix := "band"
		var bip44CoinType uint32 = 494

		accountPrefix := bech32MainPrefix
		validatorPrefix := bech32MainPrefix + sdk.PrefixValidator + sdk.PrefixOperator
		consensusPrefix := bech32MainPrefix + sdk.PrefixValidator + sdk.PrefixConsensus

		config.SetBech32PrefixForAccount(accountPrefix, accountPrefix+sdk.PrefixPublic)
		config.SetBech32PrefixForValidator(validatorPrefix, validatorPrefix+sdk.PrefixPublic)
		config.SetBech32PrefixForConsensusNode(consensusPrefix, consensusPrefix+sdk.PrefixPublic)
		config.SetCoinType(bip44CoinType)

//		fmt.Println(accountPrefix, validatorPrefix, consensusPrefix)
		
	case "osmosis":
		Bech32MainPrefix := "osmo"
		

		accountPrefix := Bech32MainPrefix
		validatorPrefix := Bech32MainPrefix + sdk.PrefixValidator + sdk.PrefixOperator
		consensusPrefix := Bech32MainPrefix + sdk.PrefixValidator + sdk.PrefixConsensus

		config.SetBech32PrefixForAccount(accountPrefix, accountPrefix+sdk.PrefixPublic)
		config.SetBech32PrefixForValidator(validatorPrefix, validatorPrefix+sdk.PrefixPublic)
		config.SetBech32PrefixForConsensusNode(consensusPrefix, consensusPrefix+sdk.PrefixPublic)
		
	case "bostrom":
		Bech32MainPrefix := "bostrom"
		accountPrefix := Bech32MainPrefix
		validatorPrefix := Bech32MainPrefix + sdk.PrefixValidator + sdk.PrefixOperator
		consensusPrefix := Bech32MainPrefix + sdk.PrefixValidator + sdk.PrefixConsensus

		config.SetBech32PrefixForAccount(accountPrefix, accountPrefix+sdk.PrefixPublic)
		config.SetBech32PrefixForValidator(validatorPrefix, validatorPrefix+sdk.PrefixPublic)
		config.SetBech32PrefixForConsensusNode(consensusPrefix, consensusPrefix+sdk.PrefixPublic)


	case "terra":
	        config.SetCoinType(terra.CoinType)
	        config.SetFullFundraiserPath(terra.FullFundraiserPath)
	        config.SetBech32PrefixForAccount(terra.Bech32PrefixAccAddr, terra.Bech32PrefixAccPub)
	        config.SetBech32PrefixForValidator(terra.Bech32PrefixValAddr, terra.Bech32PrefixValPub)
	        config.SetBech32PrefixForConsensusNode(terra.Bech32PrefixConsAddr, terra.Bech32PrefixConsPub)
//		case "kava":
//			kava.SetBech32AddressPrefixes(config)
//			kava.SetBip44CoinType(config)
	case "emoney":
		emoney.ConfigureSDK()

	}

	config.Seal()
}
