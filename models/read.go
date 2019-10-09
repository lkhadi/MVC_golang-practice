package models

import "github.com/lkhadi/MVC_golang-practice/views"

func ReadAll() ([]views.Artikel, error) {
	artikel := []views.Artikel{}
	rows, err := con.Query("SELECT judul,artikel,author,tanggal FROM artikel")
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		data := views.Artikel{}
		rows.Scan(&data.Judul, &data.Artikel, &data.Author, &data.Tanggal)
		artikel = append(artikel, data)
	}
	return artikel, nil
}
