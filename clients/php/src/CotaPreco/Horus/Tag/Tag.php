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

    /**
     * @param string $tag
     */
    public function __construct($tag)
    {
        $this->tag = (string) $tag;
    }

    /**
     * @param  string $tag
     * @return self
     */
    public static function fromString($tag)
    {
        return new self($tag);
    }

    /**
     * {@inheritDoc}
     */
    public function __toString()
    {
        return $this->tag;
    }
}
