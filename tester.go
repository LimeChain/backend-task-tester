package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/ybbus/jsonrpc/v3"
)

func chooseServerCommand(lang string) *exec.Cmd {
	switch lang {
	case "RUST":
		return exec.Command("cargo", "run")
	case "NODE":
		return exec.Command("npm", "start")
	case "GO":
		return exec.Command("go", "run", ".")
	default:
		log.Fatalln("Unsupported lang")
	}
	return nil
}

func chooseUnitTestsCommand(lang string) *exec.Cmd {
	switch lang {
	case "RUST":
		return exec.Command("cargo", "test")
	case "NODE":
		return exec.Command("npm", "test")
	case "GO":
		return exec.Command("go", "test")
	default:
		log.Fatalln("Unsupported lang")
	}
	return nil
}

func runServer(lang, apiPort, ethNode, dbConn string) *exec.Cmd {
	cmd := chooseServerCommand(lang)
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "API_PORT="+apiPort+"")
	cmd.Env = append(cmd.Env, "ETH_NODE_URL="+ethNode)
	cmd.Env = append(cmd.Env, "DB_CONNECTION_URL="+dbConn)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err.Error())
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err.Error())
	}

	cmd.Start()

	go func(std io.ReadCloser) {
		scanner := bufio.NewScanner(std)

		for scanner.Scan() {
			fmt.Printf("[ERR] %s\n", scanner.Text())
		}
	}(stderr)

	go func(std io.ReadCloser) {
		scanner := bufio.NewScanner(std)

		for scanner.Scan() {
			fmt.Printf("[OUT] %s\n", scanner.Text())
		}
	}(stdout)

	return cmd

}

func killServer(cmd *exec.Cmd) {
	cmd.Process.Signal(os.Interrupt)
}

func runUnitTests(lang string) {
	cmd := chooseUnitTestsCommand(lang)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err.Error())
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("=============== Running Unit Test Suite ===============")
	log.Println("Evaluate it yourself")

	cmd.Start()

	go func(std io.ReadCloser) {
		scanner := bufio.NewScanner(std)

		for scanner.Scan() {
			fmt.Printf("[ERR] %s\n", scanner.Text())
		}
	}(stderr)

	go func(std io.ReadCloser) {
		scanner := bufio.NewScanner(std)

		for scanner.Scan() {
			fmt.Printf("[OUT] %s\n", scanner.Text())
		}
	}(stdout)

	if err := cmd.Wait(); err != nil {
		log.Print(err)
		log.Println("FAIL: Unit tests could not get started or exited with failure")
		os.Exit(1)
	}

}

func main() {
	lang := os.Getenv("LANG")
	apiPort := os.Getenv("API_PORT")
	ethNode := os.Getenv("ETH_NODE_URL")
	dbConn := os.Getenv("DB_CONNECTION_URL")

	_, isolation := os.LookupEnv("ISOLATION")

	if isolation {
		log.Println("RUNNING IN ISOLATION MODE. Only run this during the test suite development")
	}

	if !isolation {
		cmd := runServer(lang, apiPort, ethNode, dbConn)
		defer killServer(cmd)

		runTestStartServer(cmd)

		var delay time.Duration = 5

		log.Printf("Waiting %d seconds for start to finish", delay)

		time.Sleep(delay * time.Second)
	}

	rpcClient := jsonrpc.NewClient("http://localhost:" + apiPort)

	log.Println("=============== Running Automated Test Suite ===============")
	runTests(rpcClient)
	log.Println("=============== Ended Automated Test Suite ===============")
	log.Printf("TOTAL TESTS: %d\n", total)
	log.Printf("PASSED: %d, FAIL: %d\n", pass, total-pass)

	if !isolation {
		runUnitTests(lang)
	}
}

type testable func() bool

var total, pass uint = 0, 0

func test(testFn testable) {
	total++
	if testFn() {
		pass++
	}
}

// json annotations are only required to transform the structure back to json
type Transaction struct {
	TransactionHash   string `json:"transactionHash"`
	TransactionStatus int    `json:"transactionStatus"`
	BlockHash         string `json:"blockHash"`
	BlockNumber       int    `json:"blockNumber"`
	From              string `json:"from"`
	To                string `json:"to"`
	ContractAddress   string `json:"contractAddress"`
	LogsCount         int    `json:"logsCount"`
	Input             string `json:"input"`
	Value             string `json:"value"`
}

type TransactionResponse struct {
	Transactions []Transaction
}

type AuthenticateResponse struct {
	Token string `json:"token"`
}

func runTestStartServer(cmd *exec.Cmd) {
	go func() {
		if err := cmd.Wait(); err != nil {
			log.Print(err)
			log.Println("FAIL: Server could not get started or exited with failure")
			os.Exit(1)
		}
	}()
}

func runTests(rpcClient jsonrpc.RPCClient) {
	runMidTests(rpcClient)
	runSeniorTests(rpcClient)
}

func runMidTests(rpcClient jsonrpc.RPCClient) {
	log.Println("=============== Mid Test Suite ===============")
	test(testEmptyInitialAllTx(rpcClient))
	test(testExampleTxFetching(rpcClient))
	test(testStoredTxAfterExample(rpcClient))
	test(testMixedTxFetching(rpcClient))
	test(testStoredTxAfterMixed(rpcClient))
}

func runSeniorTests(rpcClient jsonrpc.RPCClient) {
	log.Println("=============== Senior Test Suite ===============")
	test(testAuthenticate(rpcClient))
	test(testGetMyTransactions(rpcClient))
	test(testDockerfileExists())
}
