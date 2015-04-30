`Horus` [![Build Status](https://travis-ci.org/CotaPreco/Horus.svg?branch=develop)](https://travis-ci.org/CotaPreco/Horus)
=====
An simple and minimalist event-hub that we've been using for pipelining events to the client from any direction.

Horus is pretty much something like the following art:
```
                                         +----------+
                                         |   PHP    |
                                         +----------+
                                           |
                                           |
                                           v
+--------------------+     +-------+     +----------+     +-----+
| Browser (ws://...) | <-- | Horus | <-- | Receiver | <-- | ... |
+--------------------+     +-------+     +----------+     +-----+
  |                          ^             ^
  +--------------------------+             |
                                           |
                                         +----------+
                                         |    JS    |
                                         +----------+
```

As you can see, you have one `Receiver` and any client (...such as PHP) which sends stuff to the so-called `Receiver`, then Horus deliver it to your front-end. Very much simple.

### `Receiver`?
Someone who receives something, lol. But seriously, c'mon, the `Receiver` is just someone waiting for messages, which will be delivered to somewhere, of course, your front-end.

### Why the currently native receiver is UDP?
Mainly because we don't care about losing packets. If you're transmitting transactional stuff, I mean, stuff that needs some reliability, don't do it first of all. You can't ensure that the client will receive it in the end; just think in the old *US Postal Service*, drop your mail in the mailbox and hopes the *postal service* will deliver it to the proper location.

Aware of this, UDP is fast because, unlike TCP, there's no handshaking, connection setup, congestion control, packet sequencing, etc, you know, such things. Fast as hell, :japanese_goblin:.

#### `ReceiveStrategy`
As `Receiver` doesn't know what you send over the protocol and in which format *(..and shouldn't)*, and as this needs to be abstract, you need to implement it.

For example, **we**'re using a null-byte approach, but you can have your own format, etc. So, in few words: `ZmqReceiver`, fine, but: the message arrives in which format? **XML**? **JSON**? You get the idea.

Anyway, just implement the interface and attempt receive using the specified strategy.

### *...Work in progress*
