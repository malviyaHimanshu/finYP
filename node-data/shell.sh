#!/bin/sh

#OLD
#geth1="geth console --datadir node1 --networkid 1729 --nodiscover --http --http.port 8000 --http.addr 0.0.0.0 --http.corsdomain * --http.api eth,net,web3,miner,debug,personal,rpc --rpc.allow-unprotected-txs"
#geth2="geth console --datadir node2 --networkid 1729 --port 30304 --nodiscover --http --authrpc.port 8051 --http.addr 0.0.0.0 --http.corsdomain * --http.api eth,net,web3,miner,debug,personal,rpc --unlock 0x9d2d449fa0d6423ca09dbbada4064282881952dd --password password.txt --mine --miner.threads=1 --allow-insecure-unlock  --rpc.allow-unprotected-txs"


geth1="/home/ajay/go-ethereum/build/bin/geth console --datadir node1 --networkid 1729 --rpc.enabledeprecatedpersonal"
geth2="/home/ajay/go-ethereum/build/bin/geth console --datadir node2 --networkid 1729 --port 30304 --authrpc.port 8552 --rpc.enabledeprecatedpersonal --mine --miner.threads=1 --miner.etherbase 0xe56b88900cdee8d109540d1d9fc4e195422d9e28"


# if you want to use xterm
#xterm -title "Node1" -hold -e ${geth1}| xterm -title "Node2" -hold -e ${geth2}

# if you want to use terminal
gnome-terminal --title="Node1" -p -x ${geth1}| gnome-terminal --title="Node2" -p -x ${geth2}
#gnome-terminal -- /bin/bash - '${geth1};read'

# if you want to use xfce-terminal
#xfce4-terminal -H --title="Node1" -x ${geth1}| xfce4-terminal -H --title="Node2" -x ${geth2}
