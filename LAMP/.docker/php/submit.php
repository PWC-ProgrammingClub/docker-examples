<?php
require_once("db/addIdea.php");

if (isset($_POST["SUBMIT"])) {
    $category = filter_input(INPUT_POST, 'category', FILTER_SANITIZE_SPECIAL_CHARS);
    $description = filter_input(INPUT_POST, 'description', FILTER_SANITIZE_SPECIAL_CHARS);
    addIdea($category, $description);
    header("Location: index.php");
} else {
    header("Location: index.php?err=How did you do that");
}
