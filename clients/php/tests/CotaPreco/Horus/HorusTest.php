<?php

namespace CotaPreco\Horus;

use CotaPreco\Horus\Message\MessageInterface;
use PHPUnit_Framework_TestCase as TestCase;

/**
 * @author Andrey K. Vital <andreykvital@gmail.com>
 */
class HorusTest extends TestCase
{
    /**
     * @test
     */
    public function send()
    {
        /* @var MessageInterface $message */
        $message = $this->getMock(MessageInterface::class);

        /* @var MessageTransportInterface|\PHPUnit_Framework_MockObject_MockObject $transport */
        $transport = $this->getMock(MessageTransportInterface::class);

        $transport->expects($this->once())
            ->method('send')
            ->with($this->equalTo($message));

        $horus = new Horus($transport);
        $horus->send($message);
    }

    /**
     * @test
     */
    public function sendAll()
    {
        $that = $this;

        /**
         * @return \PHPUnit_Framework_MockObject_MockObject
         */
        $createMessage = function () use ($that) {
            return $this->getMock(MessageInterface::class);
        };

        /* @var MessageTransportInterface|\PHPUnit_Framework_MockObject_MockObject $transport */
        $transport = $this->getMock(MessageTransportInterface::class);

        $transport->expects($this->exactly(3))
            ->method('send')
            ->with($this->isInstanceOf(MessageInterface::class));

        $horus = new Horus($transport);

        $horus->sendAll([
            $createMessage(),
            $createMessage(),
            $createMessage()
        ]);
    }
}
