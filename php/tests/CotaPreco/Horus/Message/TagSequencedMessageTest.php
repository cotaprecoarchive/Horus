<?php

namespace CotaPreco\Horus\Message;

use CotaPreco\Horus\Tag\Tag;
use CotaPreco\Horus\Tag\TagInterface;
use PHPUnit_Framework_TestCase as TestCase;

/**
 * @author Andrey K. Vital <andreykvital@gmail.com>
 */
class TagSequencedMessageTest extends TestCase
{
    /**
     * @test
     */
    public function getTags()
    {
        $tagSequencedMessage = new TagSequencedMessage(
            [
                new Tag('A'),
                new Tag('B')
            ],
            'message'
        );

        $this->assertContainsOnlyInstancesOf(
            TagInterface::class,
            $tagSequencedMessage->getTags()
        );
    }
}
