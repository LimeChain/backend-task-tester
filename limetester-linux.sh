#!/bin/sh

wget https://github.com/LimeChain/backend-task-tester/raw/main/limebackendtester-linux
chmod u+x ./limebackendtester-linux
API_PORT=3061 ETH_NODE_URL="https://eth-goerli.g.alchemy.com/v2/tbOepyPHC7Bqub--gHJzkZr97UT16cSz" DB_CONNECTION_URL="postgresql://postgres:myzkevmpassword@localhost:5432/postgres" ./limebackendtester-linux
rm ./limebackendtester-linux
