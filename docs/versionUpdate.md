# Version Updates

## 2024/07/06
1. **Executable Files**: We have added **precompiled executable files** for **blockEmulator**. With these executables, users can run the precompiled files *without having Golang installed or downloading the required libraries*. The **scripts** for compiling Golang code into executable files are located in the ``./zPrecompileScripts`` folder. 

2. **Run Executable Files**: The executable files for the three operating systems are as follows:
    - **Windows**: ``./blockEmulator_Windows_Precompile.exe``
    - **Linux**: ``./blockEmulator_Linux_Precompile``
    - **MacOS**: ``./blockEmulator_MacOS_Precompile``

- Users can use these executables to generate a **batch script for the executable** (for example, on Windows):
```./blockEmulator_Windows_Precompile.exe -S 2 -N 4 -m 0 -g --shellForExe```

- Here, `--shellForExe` indicates generating a batch script for the **executable**.
    
- Then, the `WinExe_bat_shardNum=2_NodeNum=4_mod=CLPA_Broker.bat` file will generated, users can run this `.bat` file on the Windows OS. 

3. **Compatibility**: This update will **not affect** existing code.


## 2024/07/01
1. **Modified Dataset**：We have substituted the provided dataset ``./TestTx_1M.csv`` with ``./selectedTxs_300K.csv``. Transactions in the new dataset are curated from real-world historical Ethereum transactions. Algorithms reliant on historical transaction data are anticipated to demonstrate improved performance with this updated dataset.

## 2024/06/12
1. **Fine-tune Lock**: The source codes of **Fine-tune Lock**(INFOCOM'24) has been added to the branch **Fine-tune Lock**. 

## 2024/06/11 
1. **Bug Fixed**: Fixed some bugs in the PBFT consensus, which could cause PBFT consensus nodes to discard early arriving messages in cases of unstable network latency.

## 2024/05/20
1. **New Dataset**：Users can take experiments on a test dataset ``./TestTx_1M.csv``.

## 2024/03/01
1. **FAQ Added**: We have added the **[*FAQ.md*](https://github.com/HuangLab-SYSU/block-emulator/blob/main/docs/FAQ.md)**

## 2023/12/12
1. **New Features**: We have added the *shell generator function* in *Linux*. 
- When users use the command 
`go run main.go -g ` to
 generate the *.bat* file for *Windows*, a *.shell* file will be generated, too.

## 2023/07/21
**Problem**: The account transferring led to the slight disorganization of txpool. 
- **Reason**: The algorithm saved the cost, but mixed up the txpool. 
- **Solution**: We rewrite the algorithm so that it can keep the order of txpool with a small cost. 

## 2023/07/18
**Problem**: The TCP connection will report err if the number of shards or nodes is too large.   
- **Reason**: The TCP connection is implemented as a *short connection*, which results in too many simultaneous connections due to TCP's wait_close time delaying the closing of the TCP connection. 
- **Solution1**: We solved the problem by changing the *short connection* to *long connection*, which will not build or close tcp connection frequently.
- **Solution2**: Users can split the commands and run blockEmulator on different PCs, and the relevant IP settings can be modified in *params/global_config.go* and *build/build.go*. 

## 2023/06/22
Debugs:
1. **Problem**: The PBFT consensus will interrupt in a large-scale experiment (large-size data & a number of shards).   
- **Reason**: The implementation of the func *HandleforSequentialRequest* and *HandleforSequentialRequest* has bugs, which can result in laggards not being able to get old PBFT messages. 
- **Solution**: We solved the logic errors in the func, and now the function can operate correctly. 

## 2023/05/25
1. **New Features**: We have added query function. 
- After the experiment is completed, users can query the account balance now. 
- After the experiment is completed, users can query the information of blocks and transactions on the blockchain now. 

## 2023/05/24
1. Future work of [**2023/05/21 Debugs - 1**](#20230521)
- 2023/05/21: **Future**: We will solve this problem without "Virtual transactions", because this implementation cost more.
- Now this function *blockChain.AddAccount* can operate correctly without invoking *GetUpdateStatusTrie* (the resource cost is reduced). 

## 2023/05/21
Debugs:
1. **Problem**: The function *blockChain.AddAccount* has no effect on the storage.  
- **Reason**: The implementation of this func cannot change the Merkle root...
- **Solution**: We use "Virtual transactions" to replace the *AddAccount* operation. For a just added account, we consider it as a virtual transaction (whose *Sender* or *Recipient* is "00000000000000"), so that we can use *GetUpdateStatusTrie* to do this operation. 
- **Future**: We will solve this problem without "Virtual transactions", because this implementation costs more.

2. **Problem**: The transactions migrate to the incorrect shard (in CLPA + Broker mechanism)
- **Reason**: The judge function is incorrect, leading to the wrong behaviors. 
- **Solution**: We have added a new attribute *SenderIsBroker* in broker tx, to identify whether the sender is a broker account (if *HasBroker && !SenderIsBroker* is true, then the *recipient* is a broker account), and modified the *sendAccounts_and_Txs* function in *accountTransfermod_Broker.go* file with this attribute. 

## 2023/05/19
Debugs: 
1. **Problem**: The init_balances of some accounts were wrong.  
- **Reason**: The init_balances were set by shallow copy ("="), but *Balance* is a pointer (*big.Int)
- **Solution**: We have replaced the shallow copy with a deep copy (with an "adding-zero" form)

