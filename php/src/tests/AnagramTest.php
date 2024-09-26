<?php

use PHPUnit\Framework\TestCase;
use App\Algorithm\Anagram;


class AnagramTest extends TestCase
{
   
    public function testAnagramMapApproach()
    {
        $anagramChecker =  new Anagram();

        $this->assertTrue($anagramChecker->anagramCharMapApproach("Hello","lleho"));
    }

}