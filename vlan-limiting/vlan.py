#ip link add link eth0 name eth0.8 type vlan id 8

for i in range(100, 580):
    inf = "eth1."+str(i);
    print "/sbin/ip link add link eth1 name", inf, "type vlan id", i;

print "\n\n"
for i in range(8, 80):
    inf = "eth1."+str(i);
    print "/sbin/ip -d link show", inf
