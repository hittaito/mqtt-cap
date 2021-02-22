# Makefile

cap:
	docker exec -it broker tcpdump -i eth0 -X -s 0 -w /tmp/cap/mqtt.pcap

subscribe:
	docker-compose run subscriber go run ./src/subscriber