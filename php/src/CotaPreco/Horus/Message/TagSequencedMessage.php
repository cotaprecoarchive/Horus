<?php

namespace CotaPreco\Horus\Message;

use CotaPreco\Horus\Tag\TagInterface;

/**
 * @author Andrey K. Vital <andreykvital@gmail.com>
 */
class TagSequencedMessage extends Message
{
    /**
     * @var TagInterface[]
     */
    private $tags;

    /**
     * @param TagInterface[] $tags
     * @param string         $message
     */
    public function __construct(array $tags, $message)
    {
        $this->tags = array_map(
            function (TagInterface $tag) {
                return $tag;
            },
            $tags
        );

        parent::__construct($message);
    }

    /**
     * @return TagInterface[]
     */
    public function getTags()
    {
        return $this->tags;
    }
}
