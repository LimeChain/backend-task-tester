package main

import (
	"errors"
	"log"
	"os"
)

var expectedMyTxns = []Transaction{
	{
		TransactionHash:   "0x04fcf6e9f6451163a231cc7b5f9ae4b2835dfeb7ff2cb5d29609343f09a30f79",
		TransactionStatus: 1,
		BlockHash:         "0xefdf0a8a757ae93591bae7b308d4b9f6747d6a2a789ad2e6d76e9e21ff96f604",
		BlockNumber:       7976374,
		From:              "0xf29a6c0f8ee500dc87d0d4eb8b26a6fac7a76767",
		To:                "0xe97aff2b33383db8235a9b4616dc134f740243ad",
		Input:             "0x",
		Value:             50000000000000000,
	},
	{
		TransactionHash:   "0x444b853306b3a4f934c90cb996956966d804edc27fbb67402ebb8294a2cc8fd2",
		TransactionStatus: 1,
		BlockHash:         "0xefdf0a8a757ae93591bae7b308d4b9f6747d6a2a789ad2e6d76e9e21ff96f604",
		BlockNumber:       7976374,
		From:              "0xf29a6c0f8ee500dc87d0d4eb8b26a6fac7a76767",
		To:                "0xf020c6df9839b4e919539260b56b55ab40a72451",
		Input:             "0x",
		Value:             50000000000000000,
	},
}

func testAuthenticate(rpcClient *LimeClient) testable {
	return func() bool {
		res, err := rpcClient.PostAuthenticate("carol", "carol")

		if err != nil || res == nil || len(res.Token) == 0 {
			log.Println("[testAuthenticate] FAIL: No response for carol")
			return false
		}

		carolToken := res.Token

		res, err = rpcClient.PostAuthenticate("dave", "dave")

		if err != nil || res == nil || len(res.Token) == 0 {
			log.Println("[testAuthenticate] FAIL: No response for dave")
			return false
		}

		daveToken := res.Token

		if carolToken == daveToken {
			log.Println("[testAuthenticate] FAIL: Carol and Dave got the same token")
			return false
		}

		_, err = rpcClient.PostAuthenticate("george", "george")

		if err == nil {
			log.Println("[testAuthenticate] FAIL: george got a token, but should not have done so")
			return false
		}

		log.Println("[testAuthenticate] SUCCESS")
		return true
	}
}

func testGetMyTransactions(rpcClient *LimeClient) testable {
	return func() bool {

		authRes, err := rpcClient.PostAuthenticate("carol", "carol")

		if err != nil || authRes == nil || len(authRes.Token) == 0 {
			log.Println("[testAuthenticate] FAIL: No response for carol")
			return false
		}

		carolToken := authRes.Token

		rlpString := "/f888b842307834343462383533333036623361346639333463393063623939363935363936366438303465646332376662623637343032656262383239346132636338666432b842307830346663663665396636343531313633613233316363376235663961653462323833356466656237666632636235643239363039333433663039613330663739"
		res, err := rpcClient.GetEth(rlpString, carolToken)

		if err != nil || res == nil {
			log.Println("[testGetMyTransactions] FAIL: No response")
			return false
		}

		if len(res.Transactions) != len(expectedMyTxns) {
			log.Printf("[testGetMyTransactions] FAIL: Wrong count of transactions in the db; expected %d, got %d\n", len(expectedMyTxns), len(res.Transactions))
			return false
		}

		if ok := compare("testGetMyTransactions", expectedMyTxns, res.Transactions); !ok {
			return false
		}

		res, err = rpcClient.GetMy(carolToken)
		if err != nil || res == nil {
			log.Println("[testGetMyTransactions] FAIL: No response")
			return false
		}

		if len(res.Transactions) != len(expectedMyTxns) {
			log.Printf("[testGetMyTransactions] FAIL: Wrong count of transactions in the db for carol; expected %d, got %d\n", len(expectedMyTxns), len(res.Transactions))
			return false
		}

		log.Println("[testGetMyTransactions] SUCCESS")
		return true
	}
}

func testDockerfileExists() testable {
	return func() bool {

		if _, err := os.Stat("./Dockerfile"); errors.Is(err, os.ErrNotExist) {
			log.Println("[testDockerfileExists] FAIL: Dockerfile does not exist")
			return false
		}

		log.Println("[testDockerfileExists] SUCCESS")
		return true
	}
}
