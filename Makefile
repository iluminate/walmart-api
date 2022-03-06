compile:
	make test
	. build/docker-build.sh

test:
	go test -v ./tests/...

deploy:
	make compile
	docker-compose -f deployments/docker-compose.yml up -d

undeploy:
	docker-compose -f deployments/docker-compose.yml down
