<?php

$data = [
	[
		"usr_name" => "test name1",
		"age" => 31,
		"address" => "Yokohama Kanagawa.",
	],
	[
		"usr_name" => "test name2",
		"age" => 256,
		"address" => "Atsugi Kanagawa.",
	]

];

$data2 = [
	"usr_name" => "test name",
	"age" => 30,
	"address" => "Yokohama",
];

echo json_encode($data2);
