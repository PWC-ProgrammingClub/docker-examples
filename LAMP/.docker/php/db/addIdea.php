<?php

require_once("connect.php");
function addIdea($category, $description)
{
    // prepared statements prevent sqli injection
    $sql = "INSERT INTO ideas (category, description) VALUES (?,?);";
    $mysqli = connect();
    $stmt = $mysqli->prepare($sql);
    $stmt->bind_param("ss", $category, $description);
    $stmt->execute();
    $mysqli->close();
}
