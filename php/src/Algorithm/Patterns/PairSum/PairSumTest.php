<?php

namespace Src\Algorithm\Patterns\PairSum;

use PHPUnit\Framework\TestCase;

class PairSumTest extends TestCase
{
    private PairSum $pairSum;

    protected function setUp(): void
    {
        $this->pairSum = new PairSum();
    }

    public function testPairSumBruteForce(): void
    {
        // Test case from example
        $nums = [-5, -2, 3, 4, 6];
        $target = 7;
        $expected = [2, 3]; // indexes of 3 and 4
        $this->assertEquals($expected, $this->pairSum->pairSumBruteForce($nums, $target));

        // Test case with no solution
        $nums = [1, 2, 3, 4];
        $target = 10;
        $this->assertEquals([], $this->pairSum->pairSumBruteForce($nums, $target));

        // Test case with negative numbers
        $nums = [-3, -1, 0, 2, 6];
        $target = -4;
        $expected = [0, 1]; // indexes of -3 and -1
        $this->assertEquals($expected, $this->pairSum->pairSumBruteForce($nums, $target));

        // Test case with duplicate numbers
        $nums = [1, 2, 2, 3];
        $target = 4;
        $expected = [1, 2]; // indexes of 2 and 2
        $this->assertEquals($expected, $this->pairSum->pairSumBruteForce($nums, $target));
    }
}