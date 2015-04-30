<?php

namespace CotaPreco\Horus;

use CotaPreco\Horus\Message\MessageInterface;

/**
 * @author Andrey K. Vital <andreykvital@gmail.com>
 */
interface MessageTransportInterface
{
    /**
     * @param  MessageInterface $message
     * @return mixed
     */
    public function send(MessageInterface $message);
}
