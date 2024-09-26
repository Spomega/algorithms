<?php

namespace App\Algorithm;
 

class Anagram {

    public function anagramSortApproach($first, $second): bool
    {
        // remove all spaces, set all chars to lowercase and split string into array
        $first = str_split(strtolower(str_replace(" ", "",$first)));
        $second = str_split(strtolower(str_replace(" ", "",$second)));

        // check if number of characters are not the same return false
        if (count($first) != count($second)) {
            return false;
        }
        
        // sort each string array
        sort($first);
        sort($second);


        //convert string array into string and compare both strings
        return implode("",$first) == implode("",$second);
    }


    public function anagramCharMapApproach($first,$second): bool
    {
        $first = strtolower(str_replace(" ", "",$first));
        $second = strtolower(str_replace(" ", "",$second));
        // check if both strings have the same length
        if( strlen($first) != strlen($second)) {
           return false;
        }

        //build character map by looping through first
        
        $charMap1 =  $this->buildCharMap($first);
        $charMap2 =  $this->buildCharMap($second);

        // for($i = 0; $i < strlen($first); $i++) {

        //     $c = $first[$i];

        //     if ($charMap1[$c] != $charMap2[$c]) {
        //         return false;
        //     }
        // }

        foreach($charMap1  as $c => $count) {
            if(!isset($charMap2[$c]) || $charMap2[$c] !== $count) {
                return false;
            }
        }
    
        return true;
    }

    private function buildCharMap($value) : array
    {
        $charMap = [];
        for ($i = 0; $i < strlen($value); $i++) {
        $c = $value[$i];
        // build first character map
        if(isset($charMap[$c])) {
        $charMap[$c]++;
        } else {
        $charMap[$c] = 1;
        }
        }

         return $charMap;
    }

}
