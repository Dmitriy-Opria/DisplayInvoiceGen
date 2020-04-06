login:
	@docker login -u $(DOCKER_USER) -p $(DOCKER_PASS) repo.rmgops.com

order_matcher_stop:
	@docker stop order_matcher && \
    docker rm order_matcher
order_matcher_start:
	@docker pull repo.rmgops.com/docker/enp_update/order_matcher:develop && \
    docker run -it -d --name order_matcher  --log-opt max-size=2048m --log-opt max-file=10 repo.rmgops.com/docker/enp_update/order_matcher:develop
