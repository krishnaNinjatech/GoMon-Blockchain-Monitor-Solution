package main

import (
    "context"
    "encoding/json"
    "io/ioutil"
    "log"
    "math/big"
    "time"

    // "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/sirupsen/logrus"
)

type SecurityAlert struct {
    AlertType   string `json:"alert_type"`
    Description string `json:"description"`
    Timestamp   string `json:"timestamp"`
}

func main() {
    client, err := ethclient.Dial("http://localhost:8546")
    if err != nil {
        log.Fatalf("Failed to connect to the Ethereum client: %v", err)
    }

    logrus.Info("Monitoring new blocks...") //Monitor 

    var alerts []SecurityAlert
    var lastBlockNumber uint64

    for {
        // Poll for the latest block
        header, err := client.HeaderByNumber(context.Background(), nil)
        if err != nil {
            logrus.Errorf("Failed to retrieve latest block header: %v", err)
            time.Sleep(10 * time.Second)
            continue
        }

        if header.Number.Uint64() > lastBlockNumber {
            lastBlockNumber = header.Number.Uint64()
            logrus.Infof("New Block: %s", header.Hash().Hex())
            block, err := client.BlockByHash(context.Background(), header.Hash())
            if err != nil {
                logrus.Errorf("Failed to retrieve block: %v", err)
            } else {
                logrus.Infof("Block Transactions: %d", len(block.Transactions()))
                for index, tx := range block.Transactions() {
                    from, err := client.TransactionSender(context.Background(), tx, header.Hash(), uint(index))
                    if err != nil {
                        logrus.Errorf("Failed to get transaction sender: %v", err)
                        continue
                    }
                    logrus.Infof("Transaction from: %s", from.Hex()) // logging address details
                    logrus.Infof("Transaction to: %s", tx.To().Hex())
                    logrus.Infof("Transaction value: %s", tx.Value().String())

                    // Security alert test case: High-value transaction, we implemented a test case of crossing 2 ETH
                    threshold := big.NewInt(1000000000000000000) // 1 ETH in Wei
                    if tx.Value().Cmp(threshold) > 0 {
                        alert := SecurityAlert{
                            AlertType:   "High-value Transaction",
                            Description: "A transaction with a high value was detected.",
                            Timestamp:   time.Now().Format(time.RFC3339), //setting an alert timestamp
                        }
                        alerts = append(alerts, alert)
                        logrus.Warn("High-value transaction detected!")
                    }
                }
            }
        }

        time.Sleep(10 * time.Second) // Adjust the polling interval as needed
    }

    // Write alerts to file
    jsonData, err := json.MarshalIndent(alerts, "", "  ")
    if err != nil {
        log.Fatalf("Failed to marshal alerts: %v", err)
    }

    err = ioutil.WriteFile("security_alerts.json", jsonData, 0644)
    if err != nil {
        log.Fatalf("Failed to write alerts to file: %v", err)
    }
}
