`Horus` [![Build Status](https://travis-ci.org/CotaPreco/Horus.svg?branch=develop)](https://travis-ci.org/CotaPreco/Horus)
=====
Horus is a simple and minimalist event-hub for pipelining events from any direction to the client.

![Horus](https://raw.githubusercontent.com/CotaPreco/Horus/develop/assets/Horus.png "Horus")

- Dependency free: just drop and run; that's why is written in pure **Go**;
- It works very well with what you have today;
- Fast as hell, it does nothing rather than serving a WebSocket and deliver messages;
- Small and tiny, atomized.

# How it works
![How it works](https://raw.githubusercontent.com/CotaPreco/Horus/develop/assets/1.png "How it works")

In few words: your front-end will connect to **Horus** through a [WebSocket](http://en.wikipedia.org/wiki/WebSocket), and start waiting for new messages. And then you have what we call *Receiver*, *Receiver* is basically someone listening for incoming messages to send them to the clients (*...your front-end*).

## Demonstration
[![Horus asciicast](https://asciinema.org/a/19437.png)](https://asciinema.org/a/19437?autoplay=1)

## Install
```
curl -sL https://raw.githubusercontent.com/CotaPreco/Horus/develop/install.sh |sh
```

Or via **wget**:
```
wget -qO- https://raw.githubusercontent.com/CotaPreco/Horus/develop/install.sh |sh
```

And if you're familiar with [Docker](http://www.docker.com/), you can be getting started with:

```
docker run -d -p 8080:8000 -p 7500:7600/udp cotapreco/horus:0.1.0
```

## Usage
```
horus [...opts]
```

#### Available `...opts`:
| Option | Usage | Example | Default
| :---- | :---: | :--- | :---
| `-ws-host` | *Optional* | `-ws-host 127.0.0.1` | 0.0.0.0
| `-ws-port` | *Optional* | `-ws-port 8888` | 8000
| `-receiver-udp-host` | *Optional* | `-receiver-udp-host 127.0.0.1` | 0.0.0.0
| `-receiver-udp-port` | *Optional* | `-receiver-udp-port 5000` | 7600
| `-udp-max-packet-size` | *Optional* | `-udp-max-packet-size 65507` | 8192

## Getting started
```
$ horus -ws-port 8080 -receiver-udp-port 7500
```

Then, **Horus** will be listening for WebSocket connections on `0.0.0.0:8080`, and also there's an [UDP](http://en.wikipedia.org/wiki/User_Datagram_Protocol) receiver at `0.0.0.0:7500` which will wait for messages to deliver.

At this point, you're able to **listen** and **send** messages, if you're familiar with *npm*, you can install a simple tool called [wscat](https://www.npmjs.com/package/wscat):
```
$ npm install -g wscat
```

Then: `wscat -c ws://localhost:8080` and you're ready, listening!

But, you also can simple point your browser to `localhost`, open *Chrome Inspector* or whatever you can type some javascript code and execute:
```JS
var ws = new WebSocket('ws://localhost:8080');

ws.onmessage = function(e) {
  console.log(e.data);
}
```

#### How can I send messages to `UDP` receiver?
If you're using **bash**:
```
$ echo -n "Hello, world" >/dev/udp/0.0.0.0/5000
```

Otherwise, you can go with [netcat](http://en.wikipedia.org/wiki/Netcat):
```
$ echo -n "Hello, world" |nc -4u -w1 0.0.0.0 5000
```

And also there's a list of **known clients**:

| Author | URL
| :----: | :---:
| [@CotaPreco](https://github.com/CotaPreco) | #php [HorusPHPClient](https://github.com/CotaPreco/HorusPHPClient)
| [@CotaPreco](https://github.com/CotaPreco) | #javascript [horusjs](https://github.com/julianocomg/horusjs)

### Will it scale to thousands of connections?
Maybe yes, maybe not. Well, there isn't much to change, you just need to consider increasing the number servers running **Horus**, very simple. And also [HAProxy](http://www.haproxy.org/) can help you deal with that.

# License
[MIT](https://github.com/CotaPreco/Horus/blob/develop/LICENSE) &copy; Cota Preço.
