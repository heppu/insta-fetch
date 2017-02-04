package main

const HTML = `
{{range . }}
    <div class="image-container 4u 6u(narrow) 12u(narrower)">
    <img src="{{.URL}}" alt="" width="100%">
    </div>
{{end}}
`
