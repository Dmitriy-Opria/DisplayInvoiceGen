login:
	@docker login -u $(DOCKER_USER) -p $(DOCKER_PASS) repo.rmgops.com

invoice_gen_stop:
	@docker stop invoice_gen && \
    docker rm invoice_gen
invoice_gen_start:
	@docker pull repo.rmgops.com/docker/display_invoice/invoice_gen:develop && \
   docker run -it -d -p 9065:9065 --log-opt max-size=2048m --log-opt max-file=10 --name invoice_gen repo.rmgops.com/docker/display_invoice/invoice_gen:develop

invoice_gen_prod_stop:
	@docker stop invoice_gen && \
    docker rm invoice_gen
invoice_gen_prod_start:
	@docker pull repo.rmgops.com/docker/display_invoice/invoice_gen:master && \
	docker run -it -d -p 9066:9065 --log-opt max-size=2048m --log-opt max-file=10 --name invoice_gen repo.rmgops.com/docker/display_invoice/invoice_gen:master

sales_force_prod_stop:
	@docker stop salesforce_uploader && \
    docker rm salesforce_uploader
sales_force_prod_start:
	@docker pull repo.rmgops.com/docker/upload_pdf/salesforce_uploader:master && \
	docker run -it -d -p 9065:9065 --name salesforce_uploader -v $(pwd)/current:/src/current -v $(pwd)/archived:/src/archived repo.rmgops.com/docker/upload_pdf/salesforce_uploader:master

