# `CotaPreco\Horus\Horus`
A simple and minimalist PHP client for **Horus**.

#### Example
```PHP
<?php

use CotaPreco\Horus\Horus;
use CotaPreco\Horus\Message\Message;
use CotaPreco\Horus\Udp\Udp;

$horus = new Horus(new Udp('0.0.0.0', 7600));
$horus->send(new Message('AHOY!'));
```

The example above will attempt to connect on `0.0.0.0:7600` and send the message. Very simple, and also you can:

```PHP
use CotaPreco\Horus\Message\TaggedMessage;
use CotaPreco\Horus\Tag\Tag;

$horus->send(new TaggedMessage(new Tag('tag'), 'AHOY!'));
```

And you're ready to go. **Horus** will delivery *AHOY!* to everyone tagged with `tag`.

##### ...and what about a message to multiple tags?
As we only allow only one tag per message, we provide an utility called `TagStamper`, then:

```PHP
use CotaPreco\Horus\Util\TagStamper;

/* @var Tag[] $tags */
$tags = [
    new Tag('A'),
    new Tag('B'),
    new Tag('C')
];

$message = new Message('AHOY!');

$horus->sendAll(TagStamper::tagAll($tags, $message));
```

And you win. Under the hood, `TagStamper` just creates `TaggedMessage` for you, with every tag provided in `$tags`, you just need to send in [batch](https://github.com/CotaPreco/Horus/blob/develop/clients/php/src/CotaPreco/Horus/Horus.php#L36).

#### How to do a full match delivery? `A && C`
This means: only delivery if contains `A` **and** `C`, but in few words: you can't do this in a native way. Also, what we can say for you *today* is: design your tags.

You can create `AC` tag exclusively and apply it, and then: `new Tag('AC')` and you're okay.

### Currently supported transport strategies
It will only support oficial receivers, I mean, currently we're supporting only [`Udp`](https://github.com/CotaPreco/Horus/blob/develop/clients/php/src/CotaPreco/Horus/Udp/Udp.php) and using [`NullByte`](https://github.com/CotaPreco/Horus/blob/develop/clients/php/src/CotaPreco/Horus/Udp/PackingStrategy/NullByte.php) as the default packing strategy for [`Udp`](https://github.com/CotaPreco/Horus/blob/develop/clients/php/src/CotaPreco/Horus/Udp/Udp.php). But, of course, you can fork and add more strategies and receivers as you need.
