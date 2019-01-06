# GO HBTCP

go-hbtcp is a simple TCP server which can split the package and forward to external API.

## Installation

The general way to install the project:

```bash
# Download and install the binary
$ go get -u github.com/amobe/go-hbtcp
```

## Introduce

go-hbtcp will launch a TCP server which can accept and forward the command message.

* The server accepts the message from multiple client.
* The server breaks down the incoming message into several commands with the new line symbol.
* The server manages the command with the queue, the element in the queue will be forwarded to external API with the frequency in 30 requests per second. 
* The server terminates the connection when client sends a 'quit' command or timeout.

## Uasge

```
usage: go-hbtcp [options]
```

By default, go-hbtcp starts a TCP server at local port 20000. And, you can open the web browser to see the statistics data with local port 20001.

### Options

```
--listen value, -l value	Network address to listen (default: 127.0.0.1:2000)
--admin value, -a value		Network address for moniting statistics (default: 127.0.0.1:20001) 
--timeout value, -t value	Timeout second for keep hanging client alive (default: 30)
```

### Statistics Server

You can open the web browser to see the statistics data with local port 20001. The server will return the data with Json format.

```
{
  "clientOnlineConn":0,		// display how many client connection on the server
  "serverInPktAcc":0,		// display the total number of the incoming packages 
  "serverOutPktAcc":0,		// display the total number of the outgoing packages
  "reguestQueueSize":1000,	// the size of queue to handle the outgoing packages
  "reguestQueuePadding":0	// the number of the elements which are waiting to transmit in queue
 }
```

## Test

### General Test
It easy to test the server with net-cat command.

```bash
$ echo -e 'hello world\nquit\n' | nc 127.0.0.1 20000
```

The console output of the server:

```
[I] 2019/01/06 02:06:47.972139 server.go:173: CONN ESTABLISH 127.0.0.1:41960	<< establish the connection from the client
[D] 2019/01/06 02:06:47.972286 server.go:129: CONN NEW TIMER 30s		<< create a timer for checking the timeout
[D] 2019/01/06 02:06:47.982390 server.go:132: CONN RESET TIMER 30s		<< reset the timer at first
[I] 2019/01/06 02:06:47.982459 server.go:93: TCP_I hello world			<< receive the command 'hello world' from the TCP client
[I] 2019/01/06 02:06:47.982491 client.go:19: EXT_O hello world			<< forward the command 'hello world' to the external API
[D] 2019/01/06 02:06:47.992429 server.go:132: CONN RESET TIMER 30s		<< reset the timer after receiving the message from the client
[I] 2019/01/06 02:06:47.992499 server.go:93: TCP_I quit				<< receive the command 'quit' from the TCP client
[I] 2019/01/06 02:06:47.992615 server.go:113: CONN CLOSE 127.0.0.1:41960	<< terminate the client connection
```

If the client doesn't send 'quit' command, the connection will be terminated after 30 seconds.

### Stress Test

go-hbtcp-client is a simple TCP client which can send tons of message to the server.

#### Installation

```bash
# Download and install the test client
$ go get -u github.com/amobe/go-hbtcp/cmd/go-hbtcp-client
```

#### Usage

```bash
usage: go-hbtcp-client [options]
```

By default, go-hbtcp-client starts a TCP client which connection to the destination server at 127.0.0.1:2000.

#### Options

```bash
--amount value, -a value	Number of the TCP client instance (default: 1)
--destAddr value, -d value	Network address of the target server (default: 127.0.0.1:20000) 
```

#### Example

Start two clients for testing the server.

```bash
$ go-hbtcp-client -a 2 -d 127.0.0.1:20000
```

The console output of the server:

```
[I] 2019/01/06 02:38:40.804580 server.go:165: START LISTEN 127.0.0.1:20000	<< start the server
[I] 2019/01/06 02:38:48.866450 server.go:173: CONN ESTABLISH 127.0.0.1:42134	<< establish connection from client 1
[I] 2019/01/06 02:38:48.866574 server.go:173: CONN ESTABLISH 127.0.0.1:42136	<< establish connection from client 2
...
[I] 2019/01/06 02:38:51.372147 client.go:19: EXT_O Hello_44			<< forward the command to external API
[I] 2019/01/06 02:38:51.392378 client.go:19: EXT_O Hello_44
[I] 2019/01/06 02:38:51.412594 reguest.go:80: RequestQueue reach the request limit, block the handling process		<< too many requests to external API, block the forwarding thread
[D] 2019/01/06 02:38:51.804915 reguest.go:63: Queue limit reset by ticker	<< unblock the forwarding process
[I] 2019/01/06 02:38:51.804988 client.go:19: EXT_O Hello_45			<< forward the command to external API
[I] 2019/01/06 02:38:51.825180 client.go:19: EXT_O Hello_45
...

```

