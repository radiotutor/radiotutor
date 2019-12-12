#!/bin/sh

echo " "
for i in {3..1}; do echo -e -n "\rKilling Process in $i.      " && sleep 1; done
echo " "
echo " "
echo -n "Killing Process  ... "
sudo pkill -f "radiotutor"
sleep 0.5
echo "Done"
sleep 0.5
echo " "
echo -n "Starting Process ... "
sleep 0.5
redis-server &
sudo -E ./radiotutor &
echo "Done"
echo " "
sleep 1
echo " "

