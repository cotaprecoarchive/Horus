<?php

namespace CotaPreco\Horus\Message;

/**
 * @author Andrey K. Vital <andreykvital@gmail.com>
 */
class Message
{
    /**
     * @var string
     */
    private $message;

    /**
     * @param string $message
     */
    public function __construct($message)
    {
        $this->message = (string) $message;
    }

    /**
     * @return string
     */
    public function getMessage()
    {
        return $this->message;
    }
}
