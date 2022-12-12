#!/bin/sh

wget https://github.com/LimeChain/backend-task-tester/raw/main/limebackendtester-osx
chmod u+x ./limebackendtester-osx
ETH_NODE_URL="https://eth-goerli.g.alchemy.com/v2/tbOepyPHC7Bqub--gHJzkZr97UT16cSz" DB_CONNECTION_URL="postgresql://postgres:myzkevmpassword@localhost:5432/postgres" ./limebackendtester-osx
rm ./limebackendtester-osx
