package models

func Delete(id string) error {
	query, err := con.Query("DELETE FROM artikel WHERE id_artikel=?", id)
	defer query.Close()
	if err != nil {
		return err
	}

	return nil
}
