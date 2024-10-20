<?php

namespace App\Algorithm;

class BinarySearch {


    public function binarySearch(int $item,array $searchParams):string
    {
         $low = 0;
         $high = count($searchParams)-1;

        while ($low <= $high) {
            $mid = intdiv($low + $high,2);

            $guess = $searchParams[$mid];

            if ($guess == $item) {
                return sprintf("%d is found in position %d",$item,$mid);
            }

            if ($guess > $item) {
                $high =  $mid - 1;
            } else {
                $low = $mid + 1;
            }

        }


        return "Value not found";
    }
}