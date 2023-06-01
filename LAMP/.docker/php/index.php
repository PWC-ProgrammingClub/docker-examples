<?php
require_once("db/getIdeas.php");

$ideas = getIdeas();

require_once("include/styles.php");
?>

<body>

    <p class="info">
        This page is showing all the rows from the 'ideas' table in the mysql database. You can insert a new row with the form underneath the table.

    </p>

    <table>
        <tr>

            <th>id</th>
            <th>category</th>
            <th>description</th>

        </tr>
        <?php
        foreach ($ideas as $idea) {
            echo "<tr><td>";
            echo $idea['id'];
            echo "</td><td>";
            echo $idea['category'];
            echo "</td><td>";
            echo $idea['description'];
            echo "</td></tr>";
        }
        ?>
    </table>
    <h4>Submit new idea!</h4>
    <form action="submit.php" method="POST">
        <div>
            <label for="category">
                Category
            </label>
            <select name="category">
                <option value="projects">projects</option>
                <option value="events">events</option>
                <option value="misc">misc</option>
            </select>
        </div>
        <div>
            <label for="description">Description</label>
        </div>
        <div>
            <textarea name="description" cols="30" rows="10"></textarea>
        </div>
        <div>
            <input type="submit" value="SUBMIT" name="SUBMIT" />
        </div>
    </form>
    <span class="credits">Created by Karl Miller, 6/1/23</span>

</body>