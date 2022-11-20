# Using the test tool

## Linux
```
wget https://github.com/LimeChain/backend-task-tester/raw/main/limetester-linux.sh
LANG=RUST sh ./limetester-linux.sh
```
## macOS
```
wget https://github.com/LimeChain/backend-task-tester/raw/main/limetester-osx.sh
LANG=RUST sh ./limetester-osx.sh
```

# Env variables
```
LANG={rust|go|node|java}

```
## During development
```
API_PORT={number, ex.3040}
ETH_NODE_URL={infura or alchemy endpoint}
DB_CONNECTION_URL={postgres connection string}
ISOLATION=
```