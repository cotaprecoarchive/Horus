<?php

namespace CotaPreco\Horus\Util;

use CotaPreco\Horus\Message\Message;
use CotaPreco\Horus\Message\TaggedMessage;
use CotaPreco\Horus\Tag\TagInterface;

/**
 * @author Andrey K. Vital <andreykvital@gmail.com>
 */
final class TagStamper
{
    private function __construct()
    {
    }

    /**
     * @param  TagInterface[] $tags
     * @param  Message        $message
     * @return TaggedMessage[]
     */
    public static function tagAll($tags, Message $message)
    {
        /* @var TagInterface[] $tags */
        $tags = array_map(
            function (TagInterface $tag) {
                return $tag;
            },
            $tags
        );

        /* @var TaggedMessage[] $messages */
        $messages = [];

        /* @var TagInterface $tag */
        foreach ($tags as $tag) {
            $messages[] = new TaggedMessage($tag, $message->getMessage());
        }

        return $messages;
    }
}
