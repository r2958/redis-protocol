<?php
$conn = new redis();
$conn->connect('localhost',6379,0.5);
var_dump($conn->get("xxxx"));
var_dump($conn->set("nnnn","vvv"));


?>

