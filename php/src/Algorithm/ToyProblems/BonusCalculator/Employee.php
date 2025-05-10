<?php

/**
 * Problem 1: Employee Bonus Calculator
 * Task:
 * Calculate yearly bonuses based on:
 *
 * Base bonus = 5% of salary.
 *
 * Tenure ≥ 5 years: +2%.
 *
 * Leadership level (1-3): +1%, +2%, or +3%.
 *
 * Special contracts override rules (e.g., fixed bonus).
 *
 * Example Input (CSV Format):
 * csv
 * Name,Salary,HireDate,LeadershipLevel,SpecialContract
 * Alice Smith,80000,15.03.2018,2,
 * Bob Johnson,60000,10.06.2015,1,"Fixed: 4000"
 * Carol Lee,120000,22.11.2010,3,
 * David Brown,75000,05.01.2021,0,
 * Expected Output:
 * plaintext
 * Alice Smith: $5,600 (5% base + 2% leadership)  
 * Bob Johnson: $4,000 (Special Contract)  
 * Carol Lee: $12,000 (5% base + 2% tenure + 3% leadership)  
 * David Brown: $3,750 (5% base, no tenure/leadership bonus)  
 */

namespace App\Algorithm\ToyProblems\BonusCalculator;;

use DateTime;
class Employee {
    

    public function __construct(
        public string $name, 
        public float $salary,
        public DateTime $hireDate,
        public int $leadershipLevel,
        public float $specialContract = 0,
        ){}


        public function calculateBonus(int $year): float
        {
            $inputYear = new DateTime($year.'-01-01');
            $yearlySalary = $this->salary;
            $bonus = $this->specialContract != 0 ? $this->specialContract  : $yearlySalary * 0.05;

            if($this->specialContract != 0) {
                return $bonus;
            }

            $tenure = $this->hireDate->diff($inputYear)->y;

echo "Tenure is $tenure\n";

            if($tenure >= 5) {
                $bonus += $yearlySalary * 0.02;
            }

           $bonus += match($this->leadershipLevel) {
            1 => $yearlySalary * 0.01,
            2 => $yearlySalary * 0.02,
            3 => $yearlySalary * 0.03,
            default => 0,
           };
           

            return $bonus;

        }


}



