<?php

namespace CotaPreco\Horus\Message;

use CotaPreco\Horus\Tag\Tag;
use PHPUnit_Framework_TestCase as TestCase;

/**
 * @author Andrey K. Vital <andreykvital@gmail.com>
 */
class TaggedMessageTest extends TestCase
{
    /**
     * @var TaggedMessage
     */
    private $message;

    /**
     * {@inheritDoc}
     */
    protected function setUp()
    {
        $this->message = new TaggedMessage(
            Tag::fromString('tag'),
            'message'
        );
    }

    /**
     * @test
     */
    public function getTag()
    {
        $this->assertInstanceOf(Tag::class, $this->message->getTag());
    }

    /**
     * @test
     */
    public function getMessage()
    {
        $this->assertSame('message', $this->message->getMessage());
    }
}
