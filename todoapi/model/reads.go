package model

import "projects/todoapi/views"

func ReadAll() ([]views.PostRequst, error) {
	rows, err := con.Query("SELECT * FROM TODO")
	if err != nil {
		return nil, err
	}
	todos := []views.PostRequst{}
	for rows.Next() {
		data := views.PostRequst{}
		rows.Scan(&data.Name, &data.Todo)
		todos = append(todos, data)
	}
	return todos, nil
}

func ReadByName(name string) ([]views.PostRequst, error) {
	rows, err := con.Query("SELECT * FROM TODO WHERE name=?", name)
	if err != nil {
		return nil, err
	}
	todos := []views.PostRequst{}
	for rows.Next() {
		data := views.PostRequst{}
		rows.Scan(&data.Name, &data.Todo)
		todos = append(todos, data)
	}
	return todos, nil
}