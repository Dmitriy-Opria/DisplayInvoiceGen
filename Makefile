login:
	@docker login -u $(DOCKER_USER) -p $(DOCKER_PASS) repo.rmgops.com

invoice_gen_stop:
	@docker stop invoice_gen && \
    docker rm invoice_gen
invoice_gen_start:
	@docker pull repo.rmgops.com/docker/display_invoice/invoice_gen:develop && \
   docker run -it -d -p 9065:9065 --log-opt max-size=2048m --log-opt max-file=10 --name invoice_gen repo.rmgops.com/docker/display_invoice/invoice_gen:develop