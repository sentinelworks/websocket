#!/bin/bash

#ip link add link eth0 name eth0.8 type vlan id 8
# vconfig add eth1 10


declare -a interfaces=("eth1" "eth2" "eth3" "eth4")

for k in "${interfaces[@]}"
do
    for i in {1..200}
    do
        echo "vconfig add $k $i"
    done
done

