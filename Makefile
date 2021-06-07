# Author:       nghiatc
# Email:        congnghia0609@gmail.com

.PHONY: web
web:
	@cd ./web; ran -b 127.0.0.1 -p 5000; cd ..;

