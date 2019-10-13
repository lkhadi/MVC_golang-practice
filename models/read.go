package models

import "github.com/lkhadi/MVC_golang-practice/views"

func ReadAll() ([]views.GetArtikel, error) {
	artikel := []views.GetArtikel{}
	rows, err := con.Query("SELECT id_artikel,judul,artikel,author,tanggal FROM artikel")
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		data := views.GetArtikel{}
		rows.Scan(&data.IdArtikel, &data.Judul, &data.Artikel, &data.Author, &data.Tanggal)
		artikel = append(artikel, data)
	}
	return artikel, nil
}

func ReadOne(id string) ([]views.GetArtikel, error) {
	artikel := []views.GetArtikel{}
	row, err := con.Query("SELECT id_artikel,judul,artikel,author,tanggal FROM artikel WHERE id_artikel=?", id)
	defer row.Close()
	if err != nil {
		return nil, err
	}

	for row.Next() {
		data := views.GetArtikel{}
		row.Scan(&data.IdArtikel, &data.Judul, &data.Artikel, &data.Author, &data.Tanggal)
		artikel = append(artikel, data)
	}
	return artikel, nil
}
