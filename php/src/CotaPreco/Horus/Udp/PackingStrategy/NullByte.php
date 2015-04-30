<?php

namespace CotaPreco\Horus\Udp\PackingStrategy;

use CotaPreco\Horus\Message\Message;
use CotaPreco\Horus\Message\MessageInterface;
use CotaPreco\Horus\Message\TaggedMessage;
use CotaPreco\Horus\Message\TagSequencedMessage;
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
        $escape = function ($tagOrMessage) {
            return str_replace(chr(0), null, $tagOrMessage);
        };

        /* @var TaggedMessage $message */
        if ($message instanceof TaggedMessage) {
            return $escape($message->getTag()) . chr(0) . $escape($message->getMessage());
        }

        /* @var TagSequencedMessage $message */
        if ($message instanceof TagSequencedMessage) {
            return (
                implode(chr(0) . chr(0), array_map($escape, $message->getTags())) .
                chr(0) .
                $escape($message->getMessage())
            );
        }

        /* @var Message $message */
        return $escape($message->getMessage());
    }
}
