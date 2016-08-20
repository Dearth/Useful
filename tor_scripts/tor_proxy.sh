#!/bin/sh

command=$1

if [ "$command" = "clear" ]
then
	iptables -F
	iptables -t nat -F
	iptables-restore < /etc/iptables.rules
	echo "nameserver 127.0.1.1" > /etc/resolv.conf
else
	# destinations you don't want routed through Tor
	NON_TOR="172.20.20.0/24"

	# the UID Tor runs as
	TOR_UID="120"

	# Tor's TransPort
	TRANS_PORT="9040"

	iptables -F
	iptables -t nat -F

	iptables -P INPUT ACCEPT

	iptables -t nat -A OUTPUT -m owner --uid-owner $TOR_UID -j RETURN
	iptables -t nat -A OUTPUT -p udp --dport 53 -j REDIRECT --to-ports 53
	for NET in $NON_TOR 127.0.0.0/8 127.128.0.0/10; do
		iptables -t nat -A OUTPUT -d $NET -j RETURN
	done
	iptables -t nat -A OUTPUT -p tcp --syn -j REDIRECT --to-ports $TRANS_PORT

	iptables -A OUTPUT -m state --state ESTABLISHED,RELATED -j ACCEPT
	for NET in $NON_TOR 127.0.0.0/8; do
		iptables -A OUTPUT -d $NET -j ACCEPT
	done
	iptables -A OUTPUT -m owner --uid-owner $TOR_UID -j ACCEPT
	iptables -A OUTPUT -j REJECT

	echo "nameserver 127.0.0.1" > /etc/resolv.conf
fi
