<?php

require_once __DIR__ . '/../../../../vendor/autoload.php';

use App\Algorithm\ToyProblems\BonusCalculator\Employee;

if($argc < 2){
    echo "Usage: php Employee.php <year>\n";
    exit(1);
}


$year = (int)$argv[1];

$jsonData = file_get_contents('Employee.json');

$data = json_decode($jsonData, true);

$employees = [];

foreach ($data['employees'] as $employee) {
    $employees[] = new Employee(
        $employee['name'],
        $employee['salary'],
        new DateTime($employee['hireDate']),
        $employee['leadershipLevel'],
        $employee['specialContract'] ?? 0
    );
}

foreach($employees as $employee) {
    $bonus = $employee->calculateBonus($year);
    printf("%s: %s\n", $employee->name, number_format($bonus, 2));
}