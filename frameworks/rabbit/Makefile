.PHONY: rb
rb:
	sudo docker-compose -f ./docker-compose-rabbit.yaml up -d --build


.PHONY: pub
pub:
	go run publisher/pub.go

.PHONY: cons
cons:
	go run consumer/consumer.go
