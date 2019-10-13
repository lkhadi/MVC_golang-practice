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

type UpdateArtikel struct {
	IDArtikel string "json:idartikel"
	Judul     string "json:judul"
	Artikel   string "json:artikel"
	Author    string "json:author"
	Tanggal   string "json:tanggal"
}

type GetArtikel struct {
	IDArtikel int    "json:id_artikel"
	Judul     string "json:judul"
	Artikel   string "json:artikel"
	Author    string "json:author"
	Tanggal   string "json:tanggal"
}
