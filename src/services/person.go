package services

// GetPerson gets a person from the database
func GetPerson(id int) (string, error) {
	var firstname, lastname string
	sql := "SELECT firstname, lastname FROM person WHERE id = ?"
	err := dbm.QueryRow(sql, id).Scan(&firstname, &lastname)
	if err != nil {
		return "", err
	}
	return firstname + " " + lastname, nil
}
