<?php
require_once 'vendor/autoload.php';

use App\Algorithm\Anagram;

$anagramChecker = new Anagram();

echo $anagramChecker->anagramSortApproach("Hello","lleho") ? "The words are Anagrams\n": "Not Anagrams\n" ;

echo $anagramChecker->anagramCharMapApproach("Hello","lleho") ? "The words are Anagrams": "Not Anagrams" ;

