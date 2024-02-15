#!/bin/sh

wget https://github.com/LimeChain/backend-task-tester/raw/main/limebackendtester-linux
chmod u+x ./limebackendtester-linux
./limebackendtester-linux
rm ./limebackendtester-linux
