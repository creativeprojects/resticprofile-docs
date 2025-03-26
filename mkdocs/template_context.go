package main

type TemplateVersion struct {
	BaseURL  string
	Version  string
	Title    string
	IsLatest bool
}
type TemplateContext struct {
	Current  TemplateVersion
	Versions []TemplateVersion
}
