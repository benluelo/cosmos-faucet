package send

import (
	"context"
	"errors"
	"fmt"
	client "github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	crypto "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	auth "github.com/cosmos/cosmos-sdk/x/auth/types"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"
	"google.golang.org/grpc"
	"okp4/cosmos-faucet/util"
)

func SendTx(config util.Config, address string) error {
	conf := types.GetConfig()
	conf.SetBech32PrefixForAccount(config.Prefix, config.Prefix)

	fromPrivKey, err := GeneratePrivateKey(config.Mnemonic)
	if err != nil {
		return err
	}

	fromAddr := types.AccAddress(fromPrivKey.PubKey().Address())

	toAddr, err := types.GetFromBech32(address, config.Prefix)
	if err != nil {
		return err
	}

	fmt.Println(fromPrivKey.PubKey().String())
	fmt.Println("Addre from pub key : ", fromPrivKey.PubKey().Address().String())
	fmt.Println("Addre 1 : ", fromAddr.String())

	msg := bank.NewMsgSend(fromAddr, toAddr, types.NewCoins(types.NewInt64Coin(config.Denom, config.AmountSend)))

	encCfg := simapp.MakeTestEncodingConfig() // TODO:
	txBuilder := encCfg.TxConfig.NewTxBuilder()

	err = txBuilder.SetMsgs(msg)
	if err != nil {
		return err
	}
	txBuilder.SetGasLimit(config.GasLimit)
	txBuilder.SetMemo(config.Memo)
	txBuilder.SetFeeAmount(types.NewCoins(types.NewInt64Coin(config.Denom, config.FeeAmount)))

	fmt.Println(msg) // TODO:

	grpcConn, _ := grpc.Dial(
		"127.0.0.1:9090", // TODO:
		grpc.WithInsecure(),
	)
	defer grpcConn.Close()

	account, err := GetAccount(grpcConn, fromAddr.String())
	if err != nil {
		return err
	}

	sigV2 := signing.SignatureV2{
		PubKey: account.GetPubKey(),
		Data: &signing.SingleSignatureData{
			SignMode:  encCfg.TxConfig.SignModeHandler().DefaultMode(),
			Signature: nil,
		},
		Sequence: account.GetSequence(),
	}

	err = txBuilder.SetSignatures(sigV2)
	if err != nil {
		return err
	}

	signerData := xauthsigning.SignerData{
		ChainID:       config.ChainId,
		AccountNumber: account.GetAccountNumber(),
		Sequence:      account.GetSequence(),
	}
	sigV2, err = client.SignWithPrivKey(
		encCfg.TxConfig.SignModeHandler().DefaultMode(),
		signerData,
		txBuilder,
		fromPrivKey,
		encCfg.TxConfig,
		account.GetSequence())
	if err != nil {
		return err
	}

	err = txBuilder.SetSignatures(sigV2)
	if err != nil {
		return err
	}

	txBytes, err := encCfg.TxConfig.TxEncoder()(txBuilder.GetTx())
	if err != nil {
		return err
	}

	txJSONBytes, err := encCfg.TxConfig.TxJSONEncoder()(txBuilder.GetTx())
	if err != nil {
		return err
	}
	txJSON := string(txJSONBytes)
	fmt.Println(txJSON)

	txClient := tx.NewServiceClient(grpcConn)
	grpcRes, err := txClient.BroadcastTx(
		context.Background(),
		&tx.BroadcastTxRequest{
			Mode:    tx.BroadcastMode_BROADCAST_MODE_SYNC,
			TxBytes: txBytes,
		},
	)
	if err != nil {
		return err
	}

	if grpcRes.TxResponse.Code != 0 {
		return errors.New(grpcRes.TxResponse.RawLog)
	}

	return nil
}

func GeneratePrivateKey(mnemonic string) (crypto.PrivKey, error) {
	algo, err := keyring.NewSigningAlgoFromString("secp256k1", keyring.SigningAlgoList{hd.Secp256k1})
	if err != nil {
		return nil, err
	}

	hdPath := hd.CreateHDPath(118, 0, 0).String()

	// create master key and derive first key for keyring
	derivedPriv, err := algo.Derive()(mnemonic, "", hdPath)
	if err != nil {
		return nil, err
	}

	return algo.Generate()(derivedPriv), nil
}

func GetAccount(grpcConn *grpc.ClientConn, address string) (*auth.BaseAccount, error) {
	authClient := auth.NewQueryClient(grpcConn)
	query, err := authClient.Account(context.Background(), &auth.QueryAccountRequest{Address: address})
	if err != nil {
		return nil, err
	}

	var account auth.BaseAccount
	err = account.Unmarshal(query.GetAccount().Value)
	if err != nil {
		return nil, err
	}

	return &account, nil
}
