<?php

namespace CotaPreco\Horus\Udp\PackingStrategy;

use CotaPreco\Horus\Message\Message;
use CotaPreco\Horus\Message\TaggedMessage;
use CotaPreco\Horus\Message\TagSequencedMessage;
use CotaPreco\Horus\Tag\Tag;
use PHPUnit_Framework_TestCase as TestCase;

/**
 * @author Andrey K. Vital <andreykvital@gmail.com>
 */
class NullByteTest extends TestCase
{
    /**
     * @var NullByte
     */
    private $strategy;

    /**
     * {@inheritDoc}
     */
    protected function setUp()
    {
        $this->strategy = new NullByte();
    }

    /**
     * @test
     */
    public function packOnlyMessage()
    {
        $message = new Message("message\0with\0null\0bytes");

        $this->assertNotContains("\0", $this->strategy->pack($message));
    }

    /**
     * @test
     */
    public function packTaggedMessage()
    {
        $message = new TaggedMessage(Tag::fromString("\0tag\0"), "message\0");

        $packed = $this->strategy->pack($message);

        $this->assertEquals(1, substr_count($packed, "\0"));
        $this->assertEquals("tag\0message", $packed);
    }

    /**
     * @test
     */
    public function packTagSequencedMessage()
    {
        $message = new TagSequencedMessage(
            [
                new Tag('A'),
                new Tag('B'),
                new Tag('C')
            ],
            'message'
        );

        $packed = $this->strategy->pack($message);

        $this->assertEquals(5, substr_count($packed, "\0"));
        $this->assertEquals("A\0\0B\0\0C\0message", $packed);
    }
}
