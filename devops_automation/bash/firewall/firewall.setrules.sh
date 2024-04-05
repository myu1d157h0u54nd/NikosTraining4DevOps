#!/bin/bash

# Politicas por defecto
        # Don’t set the default policy to DROP
        # iptables -P INPUT DROP

        # pro
                iptables -P INPUT DROP
                iptables -P OUTPUT ACCEPT
                iptables -P FORWARD ACCEPT

# infra, trafico local permitido
        # dhcp
        iptables -I INPUT -i lxcbr0 -p udp --dport 11:11 --sport 11:11 -j ACCEPT

# estas reglas son necesarias para permitir las comunicaciones internas
        #  Allow all loopback (lo0) traffic
                iptables -A INPUT -i lo -j ACCEPT
        # Drop all traffic to 127/8 that doesn't use lo0
                iptables -A INPUT -d 127.0.0.0/8 -j REJECT

        #  Accept all established inbound connections
                iptables -A INPUT -m state --state ESTABLISHED,RELATED -j ACCEPT
        # If you have a network, that you want to attach to the outside
                iptables -t nat -A POSTROUTING -s 1.1.1.1/24 -o eth0 -j MASQUERADE


# Allow port
        #  HTTP and HTTPS and SSH
        iptables -A INPUT -p tcp --dport 11 -j ACCEPT
        iptables -A INPUT -p tcp --dport 111 -j ACCEPT
        iptables -A INPUT -p tcp --dport 1111 -j ACCEPT
        iptables -A INPUT -p tcp --dport 1111 -j ACCEPT
        iptables -A INPUT -p tcp --dport 11 -j ACCEPT

# Deny port
        #  Ping
        iptables -A INPUT -p icmp --icmp-type echo-request -j DROP

# Allow Linode Longview
        #  Allow incoming Longview connections
        iptables -A INPUT -s asd.asd.com -j ACCEPT

        # Allow metrics to be provided Longview
        iptables -A OUTPUT -d asd.asd.com -j ACCEPT

# Forwarding port # host ip service/port
# =CONCATENATE("iptables -t nat -A PREROUTING -i eth0 -p ";F3;" -m ";F3;" --dport ";D3;" -j DNAT --to-destination ";G3;":";E3)

#liintu Proxy
iptables -t nat -A PREROUTING -i eth0 -p tcp -m tcp --dport 1111 -j DNAT --to-destination 1.1.1.1:1111
