<?php

$docker_db_container = getenv("DOCKER_DB_CONTAINER_NAME");
$docker_db_pass = getenv("DOCKER_DB_PASS");
$docker_db_user = getenv("DOCKER_DB_USER");
$docker_db_port = getenv("DOCKER_DB_PORT");
$docker_db = getenv("DOCKER_DB");

/** Get a mysqli object connected to the mysql docker container. */
function connect()
{
    global $docker_db_container, $docker_db_pass, $docker_db_user, $docker_db;
    // Set up error reporting
    mysqli_report(MYSQLI_REPORT_ERROR);
    $mysqli = new mysqli($docker_db_container, $docker_db_user, $docker_db_pass, $docker_db);
    if ($mysqli->connect_errno) {
        echo "<h3>Failed to connect to MySQL:</h3>";
        echo $mysqli->connect_error;
    }
    return $mysqli;
}
