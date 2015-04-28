<?php

namespace CotaPreco\Horus\Tag;

/**
 * @author Andrey K. Vital <andreykvital@gmail.com>
 */
final class Tag implements TagInterface
{
    /**
     * @var
     */
    private $tag;

    private function __construct()
    {
    }

    /**
     * @param  string $tagAsString
     * @return self
     */
    public static function fromString($tagAsString)
    {
        $tag = new self();
        $tag->tag = (string) $tagAsString;

        return $tag;
    }

    /**
     * {@inheritDoc}
     */
    public function __toString()
    {
        return $this->tag;
    }
}
