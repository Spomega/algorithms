<?php

use PHPUnit\Framework\TestCase;
use App\Algorithm\MatrixRotation;

class MatrixRotationTest extends TestCase
{

    private MatrixRotation $matrixRotation;

    public function setUp(): void
    {
        $this->matrixRotation = new MatrixRotation();
    }


    public function testRotate3x3Matrix()
    {
        $matrix = [
            [1, 2, 3],
            [4, 5, 6],
            [7, 8, 9]
        ];

        $expected = [
            [7, 4, 1],
            [8, 5, 2],
            [9, 6, 3]
        ];

        $this->assertSame($expected, $this->matrixRotation->rotateMatrix90Clockwise($matrix));
    }

    public function testRotate4x4Matrix()
    {
        $matrix = [
            [1,  2,  3,  4],
            [5,  6,  7,  8],
            [9, 10, 11, 12],
            [13, 14, 15, 16]
        ];

        $expected = [
            [13,  9, 5, 1],
            [14, 10, 6, 2],
            [15, 11, 7, 3],
            [16, 12, 8, 4]
        ];

        $this->assertSame($expected, $this->matrixRotation->rotateMatrix90Clockwise($matrix));
    }

    public function testRotate1x1Matrix()
    {
        $matrix = [[42]];
        $expected = [[42]]; // Rotating a 1x1 matrix does nothing

        $this->assertSame($expected, $this->matrixRotation->rotateMatrix90Clockwise($matrix));
    }

    public function testRotateEmptyMatrix()
    {
        $matrix = [];
        $expected = []; // Rotating an empty matrix should return empty

        $this->assertSame($expected, $this->matrixRotation->rotateMatrix90Clockwise($matrix));
    }
}
