<?php

namespace App\Algorithm;

class MatrixRotation
{
    public  function rotateMatrix90Clockwise(array $matrix): array
    {
        $n = count($matrix);
        $result = array_fill(0, $n, array_fill(0, $n, 0));

        for ($i = 0; $i < $n; $i++) {

            for ($j = 0; $j < $n; $j++) {
                // Rotate the matrix by 90 degrees clockwise

                $result[$j][$n - 1 - $i] = $matrix[$i][$j];
            }
        }

        return $result;
    }
}
