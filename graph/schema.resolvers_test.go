package graph

import (
	"context"
	"errors"
	"testing"

	"okp4/cosmos-faucet/graph/generated"
	"okp4/cosmos-faucet/graph/model"
	"okp4/cosmos-faucet/pkg"
	"okp4/cosmos-faucet/pkg/client"

	gql "github.com/99designs/gqlgen/client"
	"github.com/cosmos/cosmos-sdk/types"

	"github.com/99designs/gqlgen/graphql/handler"
	. "github.com/smartystreets/goconvey/convey"
)

var config = pkg.Config{
	GrpcAddress: "127.0.0.1:9090",
	ChainID:     "my-chain",
	Denom:       "denom",
	AmountSend:  58,
	FeeAmount:   74,
	Memo:        "my memo",
	Prefix:      "okp4",
	GasLimit:    42,
	Mnemonic:    "nasty random alter chronic become keen stadium test chaos fashion during claim rug thing trade swap bleak shuffle bronze gun tobacco length aim hazard",
}

type mockFaucet struct {
	config    pkg.Config
	withError bool
}

func (f mockFaucet) GetConfig() pkg.Config {
	return config
}

func (f mockFaucet) GetFromAddr() types.AccAddress {
	bech32, _ := types.AccAddressFromBech32("okp41jse8senm9hcvydhl8v9x47kfe5z82zmwtw8jvj")
	return bech32
}

func (f mockFaucet) Close() error {
	panic("implement me")
}

func (f mockFaucet) SendTxMsg(_ context.Context, _ string) (*types.TxResponse, error) {
	var code uint32
	if f.withError {
		code = 12
	} else {
		code = 0
	}

	return &types.TxResponse{
		Height:    0,
		TxHash:    "HASH",
		Codespace: "",
		Code:      code,
		Data:      "",
		RawLog:    "",
		Logs:      nil,
		Info:      "",
		GasWanted: 10,
		GasUsed:   20,
		Tx:        nil,
		Timestamp: "",
		Events:    nil,
	}, nil
}

func TestMutationResolver_Send(t *testing.T) {
	Convey("Given send mutation with a wrong address", t, func() {
		faucet, err := client.NewFaucet(config)
		if err != nil {
			return
		}

		srv := gql.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{Faucet: faucet}})))

		m := `
                mutation {
                    send(input: {
                        toAddress: "wrong formated address"
                    }) {
                        hash
                    }
                }
                `

		Convey("When send mutation", func() {
			err := srv.Post(m, nil)

			Convey("Mutation should return error", func() {
				So(err, ShouldNotBeNil)

				var jsonError gql.RawJsonError
				So(errors.As(err, &jsonError), ShouldBeTrue)
				So(jsonError.Error(), ShouldContainSubstring, "decoding bech32 failed: invalid character in string: ' '")
			})
		})
	})

	Convey("Given good configured faucet ", t, func() {
		faucet := mockFaucet{
			config:    config,
			withError: false,
		}

		srv := gql.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{Faucet: faucet}})))

		m := `
                mutation {
                    send(input: {
                        toAddress: "okp41jse8senm9hcvydhl8v9x47kfe5z82zmwtw8jvj"
                    }) {
                        hash
                        code
                        rawLog
                        gasWanted
                        gasUsed
                    }
                }
                `
		var result struct {
			Send model.TxResponse
		}

		Convey("When post mutation", func() {
			err := srv.Post(m, &result)

			Convey("Mutation should be successful", func() {
				So(err, ShouldBeNil)
				So(result.Send.Code, ShouldEqual, 0)
				So(result.Send.Hash, ShouldEqual, "HASH")
				So(result.Send.GasUsed, ShouldEqual, 20)
				So(result.Send.GasWanted, ShouldEqual, 10)
			})
		})
	})

	Convey("Given faucet with error", t, func() {
		faucet := mockFaucet{
			config:    config,
			withError: true,
		}

		srv := gql.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{Faucet: faucet}})))

		m := `
                mutation {
                    send(input: {
                        toAddress: "okp41jse8senm9hcvydhl8v9x47kfe5z82zmwtw8jvj"
                    }) {
                        hash
                        code
                        rawLog
                        gasWanted
                        gasUsed
                    }
                }
                `
		var result struct {
			Send model.TxResponse
		}

		Convey("When post mutation", func() {
			err := srv.Post(m, &result)

			Convey("Mutation should be successful but with error code returned with hash of failed transaction", func() {
				So(err, ShouldNotBeNil)
				So(result.Send.Code, ShouldEqual, 12)
				So(result.Send.Hash, ShouldEqual, "HASH")
				So(result.Send.GasUsed, ShouldEqual, 20)
				So(result.Send.GasWanted, ShouldEqual, 10)
			})
		})
	})
}

func TestQueryResolver_Configuration(t *testing.T) {
	Convey("Given a faucet configuration context to the resolver", t, func() {
		faucet, err := client.NewFaucet(config)
		if err != nil {
			return
		}

		Convey("When create query context with faucet and configuration", func() {
			srv := gql.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{Faucet: faucet}})))

			var result struct {
				Configuration model.Configuration
			}
			q := `
                query {
                    configuration {
                        chainId
                        denom
                        prefix
                        amountSend
                        feeAmount
                        memo
                        gasLimit
                    }
                }
                `
			srv.MustPost(q, &result)

			Convey("Configuration should be the same as the given server initialisation", func() {
				So(err, ShouldBeNil)
				So(result, ShouldNotBeNil)
				So(result.Configuration.ChainID, ShouldEqual, config.ChainID)
				So(result.Configuration.Denom, ShouldEqual, config.Denom)
				So(result.Configuration.AmountSend, ShouldEqual, config.AmountSend)
				So(result.Configuration.FeeAmount, ShouldEqual, config.FeeAmount)
				So(result.Configuration.Memo, ShouldEqual, config.Memo)
				So(result.Configuration.Prefix, ShouldEqual, config.Prefix)
				So(result.Configuration.GasLimit, ShouldEqual, config.GasLimit)
			})
		})
	})
}
