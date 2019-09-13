#!/bin/bash

#ip link add link eth0 name eth0.8 type vlan id 8
# ip link delete eth0.1 and eth0.500

declare -a interfaces=("eth1" "eth2" "eth3" "eth4")

for k in "${interfaces[@]}"
do
    for i in {1..512}
    do
        echo "/sbin/ip link delete link $k.$i"
    done
done

