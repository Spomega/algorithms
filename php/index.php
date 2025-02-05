<?php
require_once 'vendor/autoload.php';

use App\Algorithm\Anagram;
use App\Algorithm\Fizzbuzz;

$anagramChecker = new Anagram();

echo $anagramChecker->anagramSortApproach("Hello", "lleho") ? "The words are Anagrams\n" : "Not Anagrams\n";

echo $anagramChecker->anagramCharMapApproach("Hello", "lleho") ? "The words are Anagrams" : "Not Anagrams";

$fizzBuzz = new Fizzbuzz();

$fizzBuzz->fizzBuzz(100);
