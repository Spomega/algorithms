<?php

/**
 * Given an array of integers sorted in ascending order and a target value,
 * return the indexes of any pair of numbers in the array that sum to the target.
 * The order of the indexes in the result doesnt matter.
 * if no pair is found return an empty array
 * 
 * Example : Input: nums = [-5, -2 , 3, 4, 6], target = 7
 * Output: [2,3]
 * Explanation: nums[2] + nums[3] = 3 + 4 =7
 */

namespace Src\Algorithm\Patterns\PairSum;

class PairSum {

    //Brute force solution (O(n^2))
    public function pairSumBruteForce(array $nums, int $target): array
    {
        // loop through the array and sum each element with the next one
        // if sum is equal to target return the indexes
        
        $result = [];

        foreach ($nums as $key => $num){
            foreach( $nums as $key2 => $num2) {
                 if ($key != $key2 && $num + $num2 == $target) {
                    $result[] = $key;
                    $result[] = $key2;
                    return $result;
                 }
            }
        }

        return $result;
    }
}





