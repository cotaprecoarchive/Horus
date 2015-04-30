<?php

namespace CotaPreco\Horus;

use CotaPreco\Horus\Message\MessageInterface;

/**
 * @author Andrey K. Vital <andreykvital@gmail.com>
 */
class Horus
{
    /**
     * @var MessageTransportInterface
     */
    private $transport;

    /**
     * @param MessageTransportInterface $transport
     */
    public function __construct(MessageTransportInterface $transport)
    {
        $this->transport = $transport;
    }

    /**
     * @param MessageInterface $message
     */
    public function send(MessageInterface $message)
    {
        $this->transport->send($message);
    }

    /**
     * @param []MessageInterface $messages
     */
    public function sendAll(array $messages)
    {
        // ...ensure homogeneity of `$messages` (...only `MessageInterface` allowed)
        $messages = array_map(
            function (MessageInterface $message) {
                return $message;
            },
            $messages
        );

        /* @var MessageInterface $message */
        foreach ($messages as $message) {
            $this->transport->send($message);
        }
    }
}
