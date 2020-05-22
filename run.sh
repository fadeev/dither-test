#!/bin/bash
pkill dithercli ;
pkill ditherd ;
make &&
bash init.sh &&
ditherd start &
sleep 5 &&
dithercli rest-server &
bash ./mock.sh