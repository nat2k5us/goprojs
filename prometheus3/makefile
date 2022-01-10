SHELL := /bin/bash -O extglob -c

GO_FILES = $(shell ls +(cmd|pkg)/*/!(*_test).go)
TEST_FILES = $(shell ls +(cmd|pkg)/*/@(*_test).go)
GOFMT_FILES = $(GO_FILES) $(TEST_FILES)
BUILD_DIR = build

# From https://stackoverflow.com/questions/10858261/how-to-abort-makefile-if-variable-not-set
# Check that given variables are set and all have non-empty values,
# die with an error otherwise.
#
# Params:
#   1. Variable name(s) to test.
#   2. (optional) Error message to print.
check_defined = \
    $(strip $(foreach 1,$1, \
        $(call __check_defined,$1,$(strip $(value 2)))))
__check_defined = \
    $(if $(value $1),, \
      $(error Undefined $1$(if $2, ($2))))


build: 
	go build main.go

clean: 
	go clean

run-docker:
	./run/run-docker.sh

build-docker:
	docker build --tag myapp .

helm-install-myapp: build-docker
	helm upgrade --install myapp \
	--namespace=default \
	charts/myapp \
	-f charts/myapp/values.yaml \
	--set image.tag=v1.0 \
	--set image.repository=myapp

helm-uninstall-myapp:
	helm uninstall myapp

grafana-uninstall:
	helm uninstall my-grafana

grafana-install:
	helm install my-grafana bitnami/grafana

prometheus-install:
	helm install my-prometheus prometheus-community/prometheus

prometheus-uninstall:
	helm uninstall my-prometheus