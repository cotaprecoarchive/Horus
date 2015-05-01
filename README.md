`Horus` [![Build Status](https://travis-ci.org/CotaPreco/Horus.svg?branch=develop)](https://travis-ci.org/CotaPreco/Horus)
=====
Horus is an simple and minimalist event-hub for pipelining events from any direction to the client.

![Horus](https://raw.githubusercontent.com/CotaPreco/Horus/develop/assets/Horus.png "Horus")

- Dependency free: just drop and run; that's why is written in pure **Go**;
- It works very well with what you have today;
- Fast as hell, it does nothing rather than serving a WebSocket and deliver messages;
- Small and tiny, atomized.

# How it works
![How it works](https://raw.githubusercontent.com/CotaPreco/Horus/develop/assets/1.png "How it works")

In few words: your front-end will connect to **Horus** through a [WebSocket](http://en.wikipedia.org/wiki/WebSocket), and start waiting for new messages. And then you have what we call *Receiver*, *Receiver* is basically someone listening for incoming messages to send them to the clients (*...your front-end*).

### Can I use Horus today?
Yes. And if you're familiar with [Docker](http://www.docker.com/), you can be getting started with:

```
docker run -d -p 8000:8000 -p 7600:7600/udp cotapreco/horus:0.1.0
```

### It will scale to thousands of connections?
Maybe yes, maybe not. Well, there isn't much to change, you just need to consider increasing the number servers running **Horus**, very simple. And also [HAProxy](http://www.haproxy.org/) can help you deal with that.

# License
[MIT License](https://github.com/CotaPreco/Horus/blob/develop/LICENSE) &copy; Cota Preço.
