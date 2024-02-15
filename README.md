# Env variables
```
API_PORT={number, ex.3040}

LANG={RUST|GO|NODE|JAVA}
ETH_NODE_URL={infura or alchemy endpoint}
DB_CONNECTION_URL={postgres connection string}
INTEGRATED=
```

# Using the test tool

## Set env variables
```
export API_PORT="8080"
export DB_CONNECTION_URL="postgresql://localhost/postgres?user=postgres&password=password1234"
export ETH_NODE_URL={infura or alchemy endpoint}
```

## Postgres
```
docker run --name postgres --network my-app -p 5432:5432 -e POSTGRES_PASSWORD=password1234 -d postgres
```

## Linux
```
wget https://github.com/LimeChain/backend-task-tester/raw/main/limetester-linux.sh
./limetester-linux.sh
```
## macOS
```
wget https://github.com/LimeChain/backend-task-tester/raw/main/limetester-osx.sh
./limetester-osx.sh
```

# Building the tool

If there are any changes to the tester's codebase make sure to rebuild and commit the binaries.
For building use:
```
./build.sh
```