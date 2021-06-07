# Author: nghiatc
# Since: Jun 07, 2021

.PHONY: web
web:
	@cd ./web; ran -b 127.0.0.1 -p 5000; cd ..;

.PHONY: signal
signal:
	# @cd ./signal; gin -p 5001 --certFile=localhost.pem --keyFile=localhost.key; cd ..;
	@cd ./signal; gin -p 5001; cd ..;
