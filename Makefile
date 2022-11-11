#DEV
build-dev:
	docker build -t videocallingapp -f containers/images/Dockerfile . && docker build -t turn -f containers/images/Dockerfile.turn .

clean-dev:
	docker-compose -f containers/composes/dc.dev.yaml.down

run-dev:
	docker-compose -f containers/composes/dc.dev.yaml up