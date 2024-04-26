package main

import (
	"errors"
	"log"
	"os"
)

var expectedMyTxns = []Transaction{
	{
		TransactionHash:   "0x495fe00e6c528257290772d134220f7a46433e1dc26b4101513d3e96b87e392c",
		TransactionStatus: 1,
		BlockHash:         "0x0486cf8d8d557245ce2360488bdfb63c3e941c7dda957a724cca922eadbcd7d2",
		BlockNumber:       5774410,
		From:              "0x1a807a67040b63dd683c6129b5b6d301bad5e483",
		To:                "0x17eeffcba1f3e409cedaad5e3b6b8a2670c577ec",
		Input:             "0x",
		Value:             "2000919073266063000",
	},
	{
		TransactionHash:   "0xb6d62b660c201b7b6d716d0e4bb8d0e9f424bd2196afd799e15c691d78084f82",
		TransactionStatus: 1,
		BlockHash:         "0x0486cf8d8d557245ce2360488bdfb63c3e941c7dda957a724cca922eadbcd7d2",
		BlockNumber:       5774410,
		From:              "0xc1796657d07ba7ebb77ff9c1ed93d4334e6a857a",
		To:                "0x3c436a1c367b5e22c83420f2a0971a1e72221069",
		Input:             "0x",
		Value:             "1998780646933123000",
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
		rlpString := "/f888b842307834393566653030653663353238323537323930373732643133343232306637613436343333653164633236623431303135313364336539366238376533393263b842307862366436326236363063323031623762366437313664306534626238643065396634323462643231393661666437393965313563363931643738303834663832"
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
