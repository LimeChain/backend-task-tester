package main

import (
	"log"
	"strings"
)

func testEmptyInitialAllTx(rpcClient *LimeClient) testable {
	return func() bool {
		res, err := rpcClient.GetAll()

		if err != nil || res == nil {
			log.Println("[testEmptyInitialAllTx] FAIL: No response")
			return false
		}

		if len(res.Transactions) > 0 {
			log.Println("[testEmptyInitialAllTx] FAIL: Transactions already found in the database. Hint: Have you dropped the database?")
			return false
		}

		log.Println("[testEmptyInitialAllTx] SUCCESS")
		return true
	}
}

func testExampleTxFetching(rpcClient *LimeClient) testable {
	return func() bool {
		rlpString := "f90110b842307839623266366133633265316165643263636366393262613636366332326430353361643064386135646137616131666435343737646364363537376234353234b842307835613537653330353163623932653264343832353135623037653762336431383531373232613734363534363537626436346131346333396361336639636632b842307837316239653262343464343034393863303861363239383866616337373664306561633062356239363133633337663966366639613462383838613862303537b842307863356639366266316235346433333134343235643233373962643737643765643465363434663763366538343961373438333230323862333238643464373938"
		res, err := rpcClient.GetEth(rlpString, "")

		if err != nil || res == nil || len(res.Transactions) == 0 {
			log.Println("[testExampleTxFetching] FAIL: No response")
			return false
		}

		if res.Transactions[0].TransactionStatus != 1 {
			log.Printf("[testExampleTxFetching] FAIL: Wrong tx status on index %d\n", 0)
			return false
		}

		if res.Transactions[1].TransactionStatus != 1 {
			log.Printf("[testExampleTxFetching] FAIL: Wrong tx status on index %d\n", 3)
			return false
		}

		if res.Transactions[3].TransactionStatus != 0 {
			log.Printf("[testExampleTxFetching] FAIL: Wrong tx status on index %d\n", 3)
			return false
		}

		if strings.ToLower(res.Transactions[0].TransactionHash) != "0x9b2f6a3c2e1aed2cccf92ba666c22d053ad0d8a5da7aa1fd5477dcd6577b4524" {
			log.Printf("[testExampleTxFetching] FAIL: Wrong tx hash on index %d\n", 0)
			return false
		}

		if strings.ToLower(res.Transactions[2].TransactionHash) != "0x71b9e2b44d40498c08a62988fac776d0eac0b5b9613c37f9f6f9a4b888a8b057" {
			log.Printf("[testExampleTxFetching] FAIL: Wrong tx hash on index %d\n", 2)
			return false
		}

		if strings.ToLower(res.Transactions[1].BlockHash) != "0x92557f7e29c39cae6be013ffc817620fcd5233b68405cdfc6e0b5528261e81e5" {
			log.Printf("[testExampleTxFetching] FAIL: Wrong block hash on index %d\n", 1)
			return false
		}

		if strings.ToLower(res.Transactions[3].BlockHash) != "0x3ac55cb392661e0d2239267022dc30f32dc4767cdacfd3e342443122b87101d3" {
			log.Printf("[testExampleTxFetching] FAIL: Wrong block hash on index %d\n", 3)
			return false
		}

		if res.Transactions[0].BlockNumber != 7976382 {
			log.Printf("[testExampleTxFetching] FAIL: Wrong block number on index %d\n", 0)
			return false
		}

		if res.Transactions[2].BlockNumber != 7957369 {
			log.Printf("[testExampleTxFetching] FAIL: Wrong block number on index %d\n", 2)
			return false
		}

		if strings.ToLower(res.Transactions[0].From) != "0xb4d6a98aa8cd5396069c2818adf4ae1a0384b43a" {
			log.Printf("[testExampleTxFetching] FAIL: Wrong from index %d\n", 0)
			return false
		}

		if res.Transactions[0].To != "" {
			log.Printf("[testExampleTxFetching] FAIL: Wrong to index %d\n", 0)
			return false
		}

		if strings.ToLower(res.Transactions[1].From) != "0xf29a6c0f8ee500dc87d0d4eb8b26a6fac7a76767" {
			log.Printf("[testExampleTxFetching] FAIL: Wrong from index %d\n", 1)
			return false
		}

		if strings.ToLower(res.Transactions[1].To) != "0xb0428bf0d49eb5c2239a815b43e59e124b84e303" {
			log.Printf("[testExampleTxFetching] FAIL: Wrong to index %d\n", 1)
			return false
		}

		if strings.ToLower(res.Transactions[3].From) != "0x58fa6ab2931b73a22d85617125b936bd3f74e765" {
			log.Printf("[testExampleTxFetching] FAIL: Wrong from index %d\n", 3)
			return false
		}

		if strings.ToLower(res.Transactions[3].To) != "0x302fd86163cb9ad5533b3952dafa3b633a82bc51" {
			log.Printf("[testExampleTxFetching] FAIL: Wrong to index %d\n", 3)
			return false
		}

		if res.Transactions[1].Input != "0x" {
			log.Printf("[testExampleTxFetching] FAIL: Wrong input index %d\n", 1)
			return false
		}

		if res.Transactions[1].Value != "50000000000000000" {
			log.Printf("[testExampleTxFetching] FAIL: Wrong value index %d\n", 1)
			return false
		}

		if res.Transactions[3].Input != "0x97da873c0000000000000000000000000000000000000000000000056bc75e2d63100000000000000000000000000000000000000000000000000000000000000000000400000000000000000000000058fa6ab2931b73a22d85617125b936bd3f74e76512d1a55b318c0be714e7ce8bc54a96ac48813cfcb73cbaa0a6e933fa9a35b7bb212c8f9a45c4430a6fa3cb8b67a28403c51e494615df4f826280256a8ddabde630818902818100e4dcd34866228be9255cbd322590b92ded49868321f0535734587348c4cb450d2d68367f686faa4688410662e9f38dc62a742f71d8e81b40a3c444381ee1245024467c8f29f04f0f83059dee234f1d4ab13e536eb5958adf91782ed3495b36fd5db6e76626771d998d6e4c75eceb58e1c783b33920dcd7723fbfbc33ba6d5ff902030100010000000000000000000000000000000000000000" {
			log.Printf("[testExampleTxFetching] FAIL: Wrong input index %d\n", 3)
			return false
		}

		if res.Transactions[2].LogsCount != 3 {
			log.Printf("[testExampleTxFetching] FAIL: Wrong logs count on index %d\n", 2)
			return false
		}

		log.Println("[testExampleTxFetching] SUCCESS")
		return true
	}
}

func testStoredTxAfterExample(rpcClient *LimeClient) testable {
	return func() bool {
		res, err := rpcClient.GetAll()

		if err != nil || res == nil {
			log.Println("[testStoredTxAfterExample] FAIL: No response")
			return false
		}

		if len(res.Transactions) != 4 {
			log.Println("[testStoredTxAfterExample] FAIL: Wrong count of transactions in the db")
			return false
		}

		if res.Transactions[0].TransactionStatus != 1 {
			log.Printf("[testStoredTxAfterExample] FAIL: Wrong tx status on index %d\n", 0)
			return false
		}

		if res.Transactions[1].TransactionStatus != 1 {
			log.Printf("[testStoredTxAfterExample] FAIL: Wrong tx status on index %d\n", 3)
			return false
		}

		if res.Transactions[3].TransactionStatus != 0 {
			log.Printf("[testStoredTxAfterExample] FAIL: Wrong tx status on index %d\n", 3)
			return false
		}

		if strings.ToLower(res.Transactions[0].TransactionHash) != "0x9b2f6a3c2e1aed2cccf92ba666c22d053ad0d8a5da7aa1fd5477dcd6577b4524" {
			log.Printf("[testStoredTxAfterExample] FAIL: Wrong tx hash on index %d\n", 0)
			return false
		}

		if strings.ToLower(res.Transactions[2].TransactionHash) != "0x71b9e2b44d40498c08a62988fac776d0eac0b5b9613c37f9f6f9a4b888a8b057" {
			log.Printf("[testStoredTxAfterExample] FAIL: Wrong tx hash on index %d\n", 2)
			return false
		}

		if res.Transactions[1].BlockHash != "0x92557f7e29c39cae6be013ffc817620fcd5233b68405cdfc6e0b5528261e81e5" {
			log.Printf("[testStoredTxAfterExample] FAIL: Wrong block hash on index %d\n", 1)
			return false
		}

		if res.Transactions[3].BlockHash != "0x3ac55cb392661e0d2239267022dc30f32dc4767cdacfd3e342443122b87101d3" {
			log.Printf("[testStoredTxAfterExample] FAIL: Wrong block hash on index %d\n", 3)
			return false
		}

		if res.Transactions[0].BlockNumber != 7976382 {
			log.Printf("[testStoredTxAfterExample] FAIL: Wrong block number on index %d\n", 0)
			return false
		}

		if res.Transactions[2].BlockNumber != 7957369 {
			log.Printf("[testStoredTxAfterExample] FAIL: Wrong block number on index %d\n", 2)
			return false
		}

		if strings.ToLower(res.Transactions[0].From) != "0xb4d6a98aa8cd5396069c2818adf4ae1a0384b43a" {
			log.Printf("[testStoredTxAfterExample] FAIL: Wrong from index %d\n", 0)
			return false
		}

		if res.Transactions[0].To != "" {
			log.Printf("[testStoredTxAfterExample] FAIL: Wrong to index %d\n", 0)
			return false
		}

		if strings.ToLower(res.Transactions[1].From) != "0xf29a6c0f8ee500dc87d0d4eb8b26a6fac7a76767" {
			log.Printf("[testStoredTxAfterExample] FAIL: Wrong from index %d\n", 1)
			return false
		}

		if strings.ToLower(res.Transactions[1].To) != "0xb0428bf0d49eb5c2239a815b43e59e124b84e303" {
			log.Printf("[testStoredTxAfterExample] FAIL: Wrong to index %d\n", 1)
			return false
		}

		if strings.ToLower(res.Transactions[3].From) != "0x58fa6ab2931b73a22d85617125b936bd3f74e765" {
			log.Printf("[testStoredTxAfterExample] FAIL: Wrong from index %d\n", 3)
			return false
		}

		if strings.ToLower(res.Transactions[3].To) != "0x302fd86163cb9ad5533b3952dafa3b633a82bc51" {
			log.Printf("[testStoredTxAfterExample] FAIL: Wrong to index %d\n", 3)
			return false
		}

		if res.Transactions[1].Input != "0x" {
			log.Printf("[testStoredTxAfterExample] FAIL: Wrong input index %d\n", 1)
			return false
		}

		if res.Transactions[1].Value != "50000000000000000" {
			log.Printf("[testStoredTxAfterExample] FAIL: Wrong value index %d\n", 1)
			return false
		}

		if res.Transactions[3].Input != "0x97da873c0000000000000000000000000000000000000000000000056bc75e2d63100000000000000000000000000000000000000000000000000000000000000000000400000000000000000000000058fa6ab2931b73a22d85617125b936bd3f74e76512d1a55b318c0be714e7ce8bc54a96ac48813cfcb73cbaa0a6e933fa9a35b7bb212c8f9a45c4430a6fa3cb8b67a28403c51e494615df4f826280256a8ddabde630818902818100e4dcd34866228be9255cbd322590b92ded49868321f0535734587348c4cb450d2d68367f686faa4688410662e9f38dc62a742f71d8e81b40a3c444381ee1245024467c8f29f04f0f83059dee234f1d4ab13e536eb5958adf91782ed3495b36fd5db6e76626771d998d6e4c75eceb58e1c783b33920dcd7723fbfbc33ba6d5ff902030100010000000000000000000000000000000000000000" {
			log.Printf("[testStoredTxAfterExample] FAIL: Wrong input index %d\n", 3)
			return false
		}

		if res.Transactions[2].LogsCount != 3 {
			log.Printf("[testStoredTxAfterExample] FAIL: Wrong logs count on index %d\n", 2)
			return false
		}

		log.Println("[testStoredTxAfterExample] SUCCESS")
		return true
	}
}

func testMixedTxFetching(rpcClient *LimeClient) testable {
	return func() bool {
		rlpString := "f90198b842307830376166376634393664353532363434613734336630643662313139316234623539616262626633313238346461366263336665653964393662376333366332b842307834353166613264306237333334656266623233346431306564333634356238643034633530303363316631386633373366316630396237626564383536386663b842307839623266366133633265316165643263636366393262613636366332326430353361643064386135646137616131666435343737646364363537376234353234b842307835613537653330353163623932653264343832353135623037653762336431383531373232613734363534363537626436346131346333396361336639636632b842307837316239653262343464343034393863303861363239383866616337373664306561633062356239363133633337663966366639613462383838613862303537b842307863356639366266316235346433333134343235643233373962643737643765643465363434663763366538343961373438333230323862333238643464373938"
		res, err := rpcClient.GetEth(rlpString, "")

		if err != nil || res == nil || len(res.Transactions) == 0 {
			log.Println("[testMixedTxFetching] FAIL: No response")
			return false
		}

		if len(res.Transactions) != 6 {
			log.Println("[testMixedTxFetching] FAIL: Wrong count of transactions in the db")
			return false
		}

		if res.Transactions[0].TransactionStatus != 1 {
			log.Printf("[testMixedTxFetching] FAIL: Wrong tx status on index %d\n", 0)
			return false
		}

		if res.Transactions[1].TransactionStatus != 1 {
			log.Printf("[testMixedTxFetching] FAIL: Wrong tx status on index %d\n", 1)
			return false
		}

		if res.Transactions[0].TransactionHash != "0x07af7f496d552644a743f0d6b1191b4b59abbbf31284da6bc3fee9d96b7c36c2" {
			log.Printf("[testMixedTxFetching] FAIL: Wrong tx status on index %d\n", 0)
			return false
		}

		if res.Transactions[1].TransactionHash != "0x451fa2d0b7334ebfb234d10ed3645b8d04c5003c1f18f373f1f09b7bed8568fc" {
			log.Printf("[testMixedTxFetching] FAIL: Wrong tx hash on index %d\n", 1)
			return false
		}

		// TODO add more cases

		log.Println("[testMixedTxFetching] SUCCESS")
		return true
	}
}

func testStoredTxAfterMixed(rpcClient *LimeClient) testable {
	return func() bool {
		res, err := rpcClient.GetAll()

		if err != nil || res == nil {
			log.Println("[testStoredTxAfterMixed] FAIL: No response")
			return false
		}

		if len(res.Transactions) != 6 {
			log.Println("[testStoredTxAfterMixed] FAIL: Wrong count of transactions in the db")
			return false
		}

		if res.Transactions[0].TransactionStatus != 1 {
			log.Printf("[testStoredTxAfterMixed] FAIL: Wrong tx status on index %d\n", 0)
			return false
		}

		if res.Transactions[1].TransactionStatus != 1 {
			log.Printf("[testStoredTxAfterMixed] FAIL: Wrong tx status on index %d\n", 3)
			return false
		}

		if res.Transactions[3].TransactionStatus != 0 {
			log.Printf("[testStoredTxAfterMixed] FAIL: Wrong tx status on index %d\n", 3)
			return false
		}

		if res.Transactions[0].TransactionHash != "0x9b2f6a3c2e1aed2cccf92ba666c22d053ad0d8a5da7aa1fd5477dcd6577b4524" {
			log.Printf("[testStoredTxAfterMixed] FAIL: Wrong tx hash on index %d\n", 0)
			return false
		}

		if res.Transactions[2].TransactionHash != "0x71b9e2b44d40498c08a62988fac776d0eac0b5b9613c37f9f6f9a4b888a8b057" {
			log.Printf("[testStoredTxAfterMixed] FAIL: Wrong tx hash on index %d\n", 2)
			return false
		}

		if res.Transactions[1].BlockHash != "0x92557f7e29c39cae6be013ffc817620fcd5233b68405cdfc6e0b5528261e81e5" {
			log.Printf("[testStoredTxAfterMixed] FAIL: Wrong block hash on index %d\n", 1)
			return false
		}

		if res.Transactions[3].BlockHash != "0x3ac55cb392661e0d2239267022dc30f32dc4767cdacfd3e342443122b87101d3" {
			log.Printf("[testStoredTxAfterMixed] FAIL: Wrong block hash on index %d\n", 3)
			return false
		}

		if res.Transactions[0].BlockNumber != 7976382 {
			log.Printf("[testStoredTxAfterMixed] FAIL: Wrong block number on index %d\n", 0)
			return false
		}

		if res.Transactions[2].BlockNumber != 7957369 {
			log.Printf("[testStoredTxAfterMixed] FAIL: Wrong block number on index %d\n", 2)
			return false
		}

		if strings.ToLower(res.Transactions[0].From) != "0xb4d6a98aa8cd5396069c2818adf4ae1a0384b43a" {
			log.Printf("[testStoredTxAfterMixed] FAIL: Wrong from index %d\n", 0)
			return false
		}

		if res.Transactions[0].To != "" {
			log.Printf("[testStoredTxAfterMixed] FAIL: Wrong to index %d\n", 0)
			return false
		}

		if strings.ToLower(res.Transactions[1].From) != "0xf29a6c0f8ee500dc87d0d4eb8b26a6fac7a76767" {
			log.Printf("[testStoredTxAfterMixed] FAIL: Wrong from index %d\n", 1)
			return false
		}

		if strings.ToLower(res.Transactions[1].To) != "0xb0428bf0d49eb5c2239a815b43e59e124b84e303" {
			log.Printf("[testStoredTxAfterMixed] FAIL: Wrong to index %d\n", 1)
			return false
		}

		if strings.ToLower(res.Transactions[3].From) != "0x58fa6ab2931b73a22d85617125b936bd3f74e765" {
			log.Printf("[testStoredTxAfterMixed] FAIL: Wrong from index %d\n", 3)
			return false
		}

		if strings.ToLower(res.Transactions[3].To) != "0x302fd86163cb9ad5533b3952dafa3b633a82bc51" {
			log.Printf("[testStoredTxAfterMixed] FAIL: Wrong to index %d\n", 3)
			return false
		}

		if res.Transactions[1].Input != "0x" {
			log.Printf("[testStoredTxAfterMixed] FAIL: Wrong input index %d\n", 1)
			return false
		}

		if res.Transactions[1].Value != "50000000000000000" {
			log.Printf("[testStoredTxAfterMixed] FAIL: Wrong value index %d\n", 1)
			return false
		}

		if res.Transactions[3].Input != "0x97da873c0000000000000000000000000000000000000000000000056bc75e2d63100000000000000000000000000000000000000000000000000000000000000000000400000000000000000000000058fa6ab2931b73a22d85617125b936bd3f74e76512d1a55b318c0be714e7ce8bc54a96ac48813cfcb73cbaa0a6e933fa9a35b7bb212c8f9a45c4430a6fa3cb8b67a28403c51e494615df4f826280256a8ddabde630818902818100e4dcd34866228be9255cbd322590b92ded49868321f0535734587348c4cb450d2d68367f686faa4688410662e9f38dc62a742f71d8e81b40a3c444381ee1245024467c8f29f04f0f83059dee234f1d4ab13e536eb5958adf91782ed3495b36fd5db6e76626771d998d6e4c75eceb58e1c783b33920dcd7723fbfbc33ba6d5ff902030100010000000000000000000000000000000000000000" {
			log.Printf("[testStoredTxAfterMixed] FAIL: Wrong input index %d\n", 3)
			return false
		}

		if res.Transactions[2].LogsCount != 3 {
			log.Printf("[testStoredTxAfterMixed] FAIL: Wrong logs count on index %d\n", 2)
			return false
		}

		log.Println("[testStoredTxAfterMixed] SUCCESS")
		return true
	}
}
