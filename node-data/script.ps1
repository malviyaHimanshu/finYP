#$gethsub="geth --help"
#geth1="geth console --datadir node1 --networkid 1729 --nodiscover --http --http.port 8000 --http.addr 0.0.0.0 --http.corsdomain * --http.api eth,net,web3,miner,debug,personal,rpc --rpc.allow-unprotected-txs"
#geth2="geth console --datadir node2 --networkid 1729 --port 30304 --nodiscover --http --authrpc.port 8051 --http.addr 0.0.0.0 --http.corsdomain * --http.api eth,net,web3,miner,debug,personal,rpc --unlock 0x9d2d449fa0d6423ca09dbbada4064282881952dd --password password.txt --mine --miner.threads=1 --allow-insecure-unlock  --rpc.allow-unprotected-txs"


start powershell {geth --help; Read-Host}
#start powershell {echo hi; Read-Host}