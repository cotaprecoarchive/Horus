<?php

namespace CotaPreco\Horus\Message;

use PHPUnit_Framework_TestCase as TestCase;

/**
 * @author Andrey K. Vital <andreykvital@gmail.com>
 */
class MessageTest extends TestCase
{
    /**
     * @test
     */
    public function getMessage()
    {
        $message = new Message('message');

        $this->assertSame('message', $message->getMessage());
    }
}
