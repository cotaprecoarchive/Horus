<?php

namespace CotaPreco\Horus\Tag;

use PHPUnit_Framework_TestCase as TestCase;

/**
 * @author Andrey K. Vital <andreykvital@gmail.com>
 */
class TagTest extends TestCase
{
    /**
     * @test
     */
    public function fromString()
    {
        $tag = Tag::fromString('tag');

        $this->assertInstanceOf(TagInterface::class, $tag);
    }

    /**
     * @test
     */
    public function tagCastsToString()
    {
        $tag = Tag::fromString('tag');

        $this->assertSame('tag', (string) $tag);
    }
}
