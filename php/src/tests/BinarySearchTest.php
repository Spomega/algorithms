<?php
use PHPUnit\Framework\TestCase;
use App\Algorithm\BinarySearch;


class BinarySearchTest extends TestCase{

    public function testBinarySearch()
    {
        $searchParams = [1,2,3,4,5,6];
        
        $binarySearchCheck = new BinarySearch();

        $this->assertEquals("2 is found in position 1",$binarySearchCheck->binarySearch(2,$searchParams));
        
    }
}