package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {


	// create a template cache
	tc, err := createTemplateCache()
	if err != nil {
		fmt.Println("error creating template cache:", err)
	}

	// get the template cache from the map
	tmpl, ok := tc[tmpl]
	if !ok {
		fmt.Println("template not found in cache:", tmpl)
	}

	// render the template
	parsedTemplate, err := template.ParseFiles("./templates/" + tmpl, "./templates/base.layout.tmpl")
	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}

createTemplateCache() (map[string]*template.Template, error	) {
	myCache := map[string]*template.Template{}

	// get all the files named *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return myCache, err
	}

	// range throuth all files ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		fmt.Println("Page is currently:", page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}
