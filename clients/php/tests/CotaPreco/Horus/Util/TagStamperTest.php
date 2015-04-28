<?php

namespace CotaPreco\Horus\Util;

use CotaPreco\Horus\Message\Message;
use CotaPreco\Horus\Message\TaggedMessage;
use CotaPreco\Horus\Tag\Tag;
use PHPUnit_Framework_TestCase as TestCase;

/**
 * @author Andrey K. Vital <andreykvital@gmail.com>
 */
class TagStamperTest extends TestCase
{
    /**
     * @test
     */
    public function tagAll()
    {
        $message = new Message('message');

        $tags = [
            new Tag('a'),
            new Tag('b'),
            new Tag('c')
        ];

        $this->assertCount(3, TagStamper::tagAll($tags, $message));

        $this->assertContainsOnlyInstancesOf(
            TaggedMessage::class,
            TagStamper::tagAll($tags, $message)
        );
    }
}
