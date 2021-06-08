# ntc-gwebrtc
ntc-gwebrtc is an example WebRTC by golang  

## Dependencies
- Install [golang](https://golang.org/doc/install), [docker](https://docs.docker.com/engine/install/), [docker-compose](https://docs.docker.com/compose/install/).

## Quick start
- Get project
```bash
cd $GOPATH/src
go get github.com/congnghia0609/ntc-gwebrtc
cd $GOPATH/src/github.com/congnghia0609/ntc-gwebrtc
```

- Generate SSL certificates.  
```bash
./gen-ssl-cert.sh
```

**Option 1: Start by DOcker**
- Start docker compose.
```bash
# Start all module
docker-compose up --build

# Show docker
docker ps
```

- Open on browser: [https://localhost:5000/](https://localhost:5000/) (Trust all localhost HTTPS certificates). Open multiple tabs and chat live video.


**Option 2: Start manual**
- Start step-by-step on multi-terminal.
```bash
# Start signal
make signal

# Start turn
make turn

# Start web
make web
```

- Open on browser: [http://localhost:5000/](http://localhost:5000/). Open multiple tabs and chat live video.


## License
This code is under the [Apache License v2](https://www.apache.org/licenses/LICENSE-2.0).  