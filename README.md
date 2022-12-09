# Using the test tool

## Linux
```
wget https://github.com/LimeChain/backend-task-tester/raw/main/limetester-linux.sh
API_PORT={PORT} sh ./limetester-linux.sh
```
## macOS
```
wget https://github.com/LimeChain/backend-task-tester/raw/main/limetester-osx.sh
API_PORT={PORT} sh ./limetester-osx.sh
```

# Env variables
```
API_PORT={number, ex.3040}

LANG={RUST|GO|NODE|JAVA}
ETH_NODE_URL={infura or alchemy endpoint}
DB_CONNECTION_URL={postgres connection string}
INTEGRATED=
```