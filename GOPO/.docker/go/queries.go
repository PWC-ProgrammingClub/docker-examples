package main

import "fmt"

// database queries

func queryAllIdeas() ([]Idea, error) {
	db := Connect.DB()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM ideas ORDER BY id DESC;")
	fmt.Print(rows)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var ideas []Idea
	for rows.Next() {
		var i Idea
		err = rows.Scan(&i.Id, &i.Category, &i.Description)
		if err != nil {
			return nil, err
		}
		ideas = append(ideas, i)
	}
	return ideas, nil
}

func queryAddIdea(Category string, Description string) error {
	db := Connect.DB()
	defer db.Close()
	sql := fmt.Sprintf("INSERT INTO ideas (category, description) VALUES ('%s', '%s');", Category, Description)
	_, err := db.Query(sql)
	if err != nil {
		return err
	}
	return nil
}
