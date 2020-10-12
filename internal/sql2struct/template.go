package sql2struct

const structTpl = `type {{.TableName | ToCamelCase}} struct {
	{{range .Columns}} {{$length := len.Comment}} {{if gt $length 0}}//
	{{.Comment}} {{else}}// {{.Name}} {{end}}
}
`
