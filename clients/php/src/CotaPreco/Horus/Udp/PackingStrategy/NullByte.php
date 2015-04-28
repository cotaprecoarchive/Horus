<?php

namespace CotaPreco\Horus\Udp\PackingStrategy;

use CotaPreco\Horus\Message\Message;
use CotaPreco\Horus\Message\MessageInterface;
use CotaPreco\Horus\Message\TaggedMessage;
use CotaPreco\Horus\MessagePackingStrategyInterface;

/**
 * @author Andrey K. Vital <andreykvital@gmail.com>
 */
final class NullByte implements MessagePackingStrategyInterface
{
    /**
     * {@inheritDoc}
     */
    public function pack(MessageInterface $message)
    {
        $escape = function($tagOrMessage) {
            return str_replace(chr(0), null, $tagOrMessage);
        };

        if (! $message instanceof TaggedMessage) {
            /* @var Message $message */
            return $escape($message->getMessage());
        }

        /* @var TaggedMessage $message */
        return $escape($message->getTag()) . chr(0) . $escape($message->getMessage());
    }
}
