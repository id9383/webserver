# Webserver

The webserver serves web requests.

To build webserver and test utility, run ```make``` command from the root directory of the project
```sh
$ make
```

## Parameters
```sh
$ Usage: webserver <options>

  Options:
    -p <port> : webserver listening port
    -l <path> : webserver log filepath
```

# Running the test
```test-client``` sends GET requests to webserver and prints the response status
```sh
$ ./bin/test-client
Usage: test-client <options>

  Options:
    -p <port> : webserver listening port
```

# Running as a Docker container
The provided ```Dockerfile``` allows to run webserver in a Docker container:
```sh
$ docker build -t webserver .
$ docker run -it webserver /bin/bash
root@d5ff02cd4a76:/# service webserver start
Starting webserver: webserver.
root@d5ff02cd4a76:/# cd /gopath/src/github.com/id9383/webserver
root@d5ff02cd4a76:/gopath/src/github.com/id9383/webserver# ./bin/test-client -p 80
GET /info/v1/data01
Status: 200
GET /info/v1/data02
Status: 200
GET /info/v1/data03
Status: 200
GET /info/v1/data04
Status: 200
GET /info/v1/data05
Status: 200
GET /info/v1/data06
Status: 200
GET /info/v1/data07
Status: 200
GET /info/v1/data08
Status: 200
GET /info/v1/data09
Status: 200
GET /info/v1/data10
Status: 200
Test completed
```
