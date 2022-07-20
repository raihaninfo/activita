package views

import "text/template"

type View struct {
	Template *template.Template
}

func NewView(files ...string) *View {
	files = append(files, "views/ui/header.gohtml", "views/ui/navbar.gohtml", "views/ui/footer.gohtml")
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
	}
}
