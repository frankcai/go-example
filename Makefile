.PHONY: buildall

buildall:
	$(MAKE) -C monolith build
	$(MAKE) -C client build
	$(MAKE) -C microservices/documentsvc build		
	$(MAKE) -C microservices/ingestsvc build		

proto:
	$(MAKE) -C monolith proto
	$(MAKE) -C microservices/documentsvc proto		
	$(MAKE) -C microservices/ingestsvc proto		
	$(MAKE) -C microservices/proto proto		