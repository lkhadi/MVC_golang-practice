package views

type Response struct {
	Code int         "json:code"
	Body interface{} "json:body"
}

type SetArtikel struct {
	Judul   string "json:judul"
	Artikel string "json:artikel"
	Author  string "json:author"
	Tanggal string "json:tanggal"
}

type GetArtikel struct {
	IdArtikel int    "json:id_artikel"
	Judul     string "json:judul"
	Artikel   string "json:artikel"
	Author    string "json:author"
	Tanggal   string "json:tanggal"
}
