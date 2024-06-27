# Task 1 : Deducing the Assessment for a targeted solution submission: 

# Objective: Build a external agent in Go to monitor a Blockchain Network. [Done for 1 test case]
Evaluation Step: Extract relevant operational metrics that might be used to identify potential vulnerabilities, threats and attacks in real-time. [Done for 1 test case]

Approach: I am going to choose setting up a local block chain network to simulate specific scenarios for monitoring purposes. [Ethereum]

# Requirements checklist:
1. GitHub Repository Submitted solution in GitHub. (Done).
2. Docker Integration: Docker and docker-compose for setting up and executing the solution. (Done).
3. Documentation:
      Setup Instructions: Provide detailed instructions for setting up and executing your solution with either open nodes or a local blockchain setup. (Done, Check below).
      Metric Justification: Justify the choice of metrics monitored and explain how they can be used to detect vulnerabilities, threats, and attacks in real-time. (Done, Check below).
      Discuss methods for securing a private blockchain (i.e. Hyperledger Fabric/Besu). (Done, Check below).

# Bonus Points:
  1. Test case and Implement logic to detect potential vulnerabilities, threats, and attacks based on the metrics extracted.
	Logics Implemented:
	Deduce high volume transaction and display in the monitor console. [Done].

# 1.  Choise of Metric: Detection of High Volume Transaction: High volume trasactions are a way to basically try to breach the set limit 
In a given point in time. Some of the use cases where the high volume transactions are really important from security stand point are:
 > Money Laundering or Terrorist Financing to be able to monitor suspicious activities which could be financial crimes. 
 > To be able to meet with the compliance and regulatory requirements.  
 > Protection against exploits and attacks from hackers.  
 > Fraud detection to prevent.  
 > Validating the transactions happening on a network to exercise control over network security.  
 > Detecting insider threats is another useful scenario to think about as well.  
 > Detecting any ransomware payments from the victims.  

 # Ideal scenario as follows:
 > Transaction Detected: A transaction of more than 1 ETH (1000000000000000000 in wei) is detected on the blockchain.  
 > Immediate Alert: An alert is triggered due to the unusually high value of the transaction.  
 > Investigation: The transaction details are reviewed. The sender and recipient addresses are checked against known addresses for 
   any suspicious activity.  
 > Regulatory Reporting: If the transaction is deemed suspicious, it is reported to the relevant financial authorities for 
   further investigation.  
 > Action Taken: Based on the investigation, appropriate actions are taken. This could involve freezing the assets, reversing the 
   transaction (if possible), or enhancing security measures to prevent further unauthorized transactions.  

# The POC execution output is given below.

Please consider below scenario as well, didn't have time to implement it along with other metrics that I could have tried, thanks in advance.
2. Test Case (Beta- not implemented using code) : Monitoring Agent - `GenesisSecuirtyTest` - `genesis.json` Analysis 
Security Analysis on Block Chain based on `genesis.json` configuration setup using Golang by running `GenesisSecuirtyTest` function

# Background:
The security implication of genesis.json mis-configuration has a significant impact on the block chain environment that uses it. 
Before we begin, below are some of the Blockchains that are based on Ethereum and Ethereum-Compatible Blockchains that uses `genesis.json` configuration file: Ethereum, Binance Smart Chain, Polygon, Fantom, custom and private Ethereum based networks like Quorum, Hyperledger Besu. 

# Why `genesis'.json:
The `genesis`.json files sets up the initial state of the blockchain with conditions like pre-deployed smart contracts, account balances etc. The Blockchain network depends on network configuration like the chain ID, difficulty and gas limit forming a fundamental basis of how the network operates and how the consensus can be achieved. 

# Blockchain upgrades:
Blockchain upgrades are another important phenomena for Blockchain, the Blockchain upgrades enhance the security among other areas through EIPs - Ethereum Improvement Proposals take below examples,  
"eip150Block": 0: This means that EIP-150 is not activated in the network.
"eip155Block": 0: This means that EIP-155 is not activated in the network.
"eip158Block": 0: This means that EIP-158 is not activated in the network.
When these values are set to 0, it effectively means that the protocol changes proposed in these EIPs are not applied to the blockchain network. Hence, there should be a notification sent to the blockchain administrator to fix this. 

# Atuomation to rescue: We may even create a automated GitHub pull request to activate the EIPs with specific block numbers like below. 

# For example:
"eip150Block": 2463000: EIP-150 is activated at block number 2,463,000.     
"eip155Block": 2675000: EIP-155 is activated at block number 2,675,000.    
"eip158Block": 2675000: EIP-158 is activated at block number 2,675,000.    

# For more details:
EIP-150 (Gas Cost Changes for IO-heavy Operations):
Purpose: Adjusts gas costs for certain operations to mitigate DoS attacks and ensure network stability.
Impact of Activation: Increases gas costs for certain operations, reducing the risk of DoS attacks.

EIP-155 (Simple Replay Attack Protection):
Purpose: Introduces a new transaction field to protect against replay attacks across different chains.
Impact of Activation: Enhances transaction security by preventing transactions from being replayed on different chains.

EIP-158 (State Trie Clearing):
Purpose: Proposes changes to the state trie to reduce its size and improve network efficiency.
Impact of Activation: Results in a more compact and efficient state trie, enhancing overall network performance.  

(go - monitor should highlight the security limitations and be able to suggest remedial measures - print above EIP codes to prevent specific attacks).

# When does these EIPs are accepted to be upgraded in a network:
Once an EIP has been reviewed and accepted by the Ethereum community, it may proceed to the next stage, which could involve implementation by client developers and preparation for network upgrade.

# Points to consider, realtime usecase, what if the network is NEW: 
When the blockchain network is new, then how can we define "eip150Block": 2463000, then we only have to define "eip150Block": 0 ,  This means that EIP-150 is not activated in the network leading to risk of DoS attacks. what do you say about this?

As per my understanding, When a new blockchain network is launched, such as a testnet or a private network, it starts with a genesis block (block 0) that contains the initial state and configuration of the network. At this stage, no EIPs are typically activated because the blockchain is just beginning and has not undergone any upgrades or changes specified by EIPs.

Mitigating DoS Attacks: Before an EIP like EIP-150 is activated, the network might rely on default or initial settings that may not include the specific mitigations proposed by the EIP. This could potentially leave the network vulnerable to certain types of attacks until the necessary improvements are implemented and activated through a network upgrade.

# Some of the strategies and best practices to mitigate the risk of attacks while EIPs are not activated still on the network:
1. Monitor and Alert: Network operators and developers should continuously monitor the network for unusual activity or signs of potential DoS attacks. Implementing alert systems can help detect abnormal behavior early.
2. Adjusting Default Parameters: While waiting for EIPs to be activated, considering to adjusting default parameters related to gas limits, block sizes, and transaction fees. These adjustments can help mitigate known vulnerabilities or attack vectors temporarily, until permanent changes are made.
3. Implementing Temporary Fixes: Developers can deploy temporary patches or updates to the network client software to address specific vulnerabilities or threats. These fixes should be thoroughly tested and documented to ensure they do not introduce unintended issues.
4. Community Awareness and Education: Educate network participants, including miners, developers, and users, about potential vulnerabilities and best practices to mitigate risks. This can include guidelines on transaction validation, gas usage, and network behavior.
5. Fast-Tracking EIP Implementation: If the vulnerability poses a significant risk, expedite the implementation and activation process of critical EIPs through community consensus and developer collaboration. This may involve coordinating a network upgrade or hard fork sooner than originally planned.
6. Backup and Contingency Plans: Have contingency plans in place to quickly revert changes or deploy emergency patches if a serious DoS attack is detected. This includes maintaining backups of critical data and configurations.


# 3. Discuss methods for securing a private blockchain (i.e. Hyperledger Fabric/Besu):
 > IAM is definitely a case for protecting the private blockchain.
    > Hyperledger Fabric uses Fabric CA to manage the identities through digital certifiactes by a Certificate Authority such that only 
      the authorized users can have access to the network in question. 
 > Consenus mechanism security to avoid single point of failure through protocols like Kafka or Raft can be ensured. 
 > Smart Contract Security is another goal for a private blockchain network security. 
 > Data Privacy, Audits and Compliance amongst others. 

let's get into implemented POC, takes 5 mins to setup and run!

# SETUP INSTRUCTIONS
1. docker-compose build
2. docker-compose up
3. ganache-cli -p 8547 -m testseed -e 1000 (Be mindful of the port in main.go)
4. New Terminal Instance: geth attach http://127.0.0.1:8547 
5. eth console> eth.sendTransaction({from: '0x5c0b07b93526cd047c193fac6d7c0f321aa8901f', to: '0x0536f6c3e7577bf74a21fe17957118d393452975', value: web3.toWei(1, 'ether')})
6. New Terminal Instance: go run main.go 
# With this local block chain network should be complete
7. Testing the High value transaction test case breach: run eth console> eth.sendTransaction({from: '0x5c0b07b93526cd047c193fac6d7c0f321aa8901f', to: '0x0536f6c3e7577bf74a21fe17957118d393452975', value: web3.toWei(2, 'ether')})

Following log should be prompted:
      go run main.go
      time="2024-06-28T02:43:53+05:30" level=info msg="Monitoring new blocks..."
      time="2024-06-28T02:43:53+05:30" level=info msg="New Block: 0x775b70fd52eac16e9216ef78e156396f87ce23e76c40c70dbed75763ed6ed1d6"
      time="2024-06-28T02:43:53+05:30" level=info msg="Block Transactions: 1"
      time="2024-06-28T02:43:53+05:30" level=info msg="Transaction from: 0x5C0b07b93526CD047C193fAc6d7C0F321AA8901F"
      time="2024-06-28T02:43:53+05:30" level=info msg="Transaction to: 0x0536f6C3E7577BF74a21fe17957118d393452975"
      time="2024-06-28T02:43:53+05:30" level=info msg="Transaction value: 2000000000000000000"
      time="2024-06-28T02:43:53+05:30" level=warning msg="High-value transaction detected!"

End of POC for 1 test case where golang based monitor is tracking the network for changes. 




