package models

const HTML = `
<div class="insta-images">
	{{range . }} 
	<div class="image-container">
	<img src="{{.URL}}" alt="" width="{{.Width}}" height="{{.Height}}:">
	</div>
{{end}}
</div>
`
