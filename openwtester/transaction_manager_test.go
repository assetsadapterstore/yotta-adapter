/*
 * Copyright 2018 The openwallet Authors
 * This file is part of the openwallet library.
 *
 * The openwallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The openwallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package openwtester

import (
	"github.com/astaxie/beego/config"
	"github.com/blocktree/openwallet/v2/openw"
	"path/filepath"
	"testing"

	"github.com/blocktree/openwallet/v2/log"
	"github.com/blocktree/openwallet/v2/openwallet"
)

func TestWalletManager_GetTransactions(t *testing.T) {
	tm := testInitWalletManager()
	list, err := tm.GetTransactions(testApp, 0, -1, "", false)
	if err != nil {
		log.Error("GetTransactions failed, unexpected error:", err)
		return
	}
	for i, tx := range list {
		log.Info("trx[", i, "] :", tx)
	}
	log.Info("trx count:", len(list))
}

func TestWalletManager_GetTransactionByWxID(t *testing.T) {
	tm := testInitWalletManager()
	wxID := openwallet.GenTransactionWxID(&openwallet.Transaction{
		TxID: "aa162ca1c48323c4dca7386e9846b291fd56f0ea7cb9c74e0f004430f9f0aebb",
		Coin: openwallet.Coin{
			Symbol:     "YTA",
			IsContract: false,
			ContractID: "",
		},
	})
	log.Info("wxID:", wxID)
	tx, err := tm.GetTransactionByWxID(testApp, wxID)
	if err != nil {
		log.Error("GetTransactionByTxID failed, unexpected error:", err)
		return
	}
	log.Info("tx:", tx)
}

func TestWalletManager_GetAssetsAccountBalance(t *testing.T) {
	tm := testInitWalletManager()
	walletID := "W5aLsVkgMAWdsuytgQFhViU7pYo9Tn6zu4"
	accountID := "CVoJCngmG4EMALtjeSDE861wSmYwqNeyiLk9yLeDBsTk"

	balance, err := tm.GetAssetsAccountBalance(testApp, walletID, accountID)
	if err != nil {
		log.Error("GetAssetsAccountBalance failed, unexpected error:", err)
		return
	}
	log.Info("balance:", balance)
}

func TestWalletManager_GetAssetsAccountTokenBalance(t *testing.T) {
	tm := testInitWalletManager()
	walletID := "W3c9iRRyyjGkGoKLM6CX5Jkc6mMdDPmK6y"
	accountID := "3utXCeZyz75agfNDsMM9vdFEhfEy8yK5PXX6AfGpDCfm"
	//accountID := "ra8PKbjnB9tjyFoEo5qADcibtA21u1BcyMttxmtQWf1"

	contract := openwallet.SmartContract{
		Address:  "ysr.ystar:YSR",
		Symbol:   "YTA",
		Name:     "YSR",
		Decimals: 4,
	}

	balance, err := tm.GetAssetsAccountTokenBalance(testApp, walletID, accountID, contract)
	if err != nil {
		log.Error("GetAssetsAccountTokenBalance failed, unexpected error:", err)
		return
	}
	log.Info("balance:", balance.Balance)
}

func TestWalletManager_GetEstimateFeeRate(t *testing.T) {
	tm := testInitWalletManager()
	coin := openwallet.Coin{
		Symbol: "YTA",
	}
	feeRate, unit, err := tm.GetEstimateFeeRate(coin)
	if err != nil {
		log.Error("GetEstimateFeeRate failed, unexpected error:", err)
		return
	}
	log.Std.Info("feeRate: %s %s/%s", feeRate, coin.Symbol, unit)
}

func TestGetAddressVerify(t *testing.T) {
	symbol := "YTA"
	assetsMgr, err := openw.GetAssetsAdapter(symbol)
	if err != nil {
		log.Error(symbol, "is not support")
		return
	}
	//读取配置
	absFile := filepath.Join(configFilePath, symbol+".ini")

	c, err := config.NewConfig("ini", absFile)
	if err != nil {
		return
	}
	assetsMgr.LoadAssetsConfig(c)
	addrDec := assetsMgr.GetAddressDecoderV2()

	flag := addrDec.AddressVerify("")
	log.Infof("flag: %v", flag)

}