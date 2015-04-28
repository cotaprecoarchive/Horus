<?php

namespace CotaPreco\Horus\Message;

use CotaPreco\Horus\Tag\TagInterface;

/**
 * @author Andrey K. Vital <andreykvital@gmail.com>
 */
class TaggedMessage extends Message
{
    /**
     * @var TagInterface
     */
    private $tag;

    /**
     * @param TagInterface $tag
     * @param string       $message
     */
    public function __construct(TagInterface $tag, $message)
    {
        parent::__construct($message);

        $this->tag = $tag;
    }

    /**
     * @return TagInterface
     */
    public function getTag()
    {
        return $this->tag;
    }
}
