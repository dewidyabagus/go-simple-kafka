SHELL := /bin/bash

.PHONY: run/service
run/service:
	@echo 'Running Order Service'
	go run .

.PHONY: run/consumer-inventory
run/consumer-inventory:
	@echo 'Running Consumer Inventory'
	go run app/consumer_inventory/main.go

.PHONY: run/consumer-shipment
run/consumer-shipment:
	@echo 'Running Consumer Shipment'
	go run app/consumer_shipment/main.go
