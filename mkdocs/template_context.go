package main

type TemplateVersion struct {
	Version  string
	IsLatest bool
}
type TemplateContext struct {
	Current  TemplateVersion
	Versions []TemplateVersion
}
