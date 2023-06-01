<?php

require_once("connect.php");

function getIdeas()
{
    $mysqli = connect();
    $result = $mysqli->query("SELECT * from ideas;", MYSQLI_USE_RESULT);
    $result_array = [];
    while ($row = $result->fetch_array()) {
        $result_array[] = $row;
    }
    $mysqli->close();
    return $result_array;
}
