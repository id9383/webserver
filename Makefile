all: webserver test-client

bin:
	mkdir -p bin

webserver: bin
	cd src/webserver && go build -o ../../bin/webserver

test-client: bin
	cd src/test-client && go build -o ../../bin/test-client

install:
	cp bin/webserver /usr/local/bin/

clean:
	rm -f bin/webserver bin/test-client

.PHONY: bin
