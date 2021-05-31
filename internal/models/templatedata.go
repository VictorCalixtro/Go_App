package models

import "github.com/victorcalixtro/Web_App/internal/forms"

//TemplateData holds data send from template
type TemplateData struct{
StringMap map[string]string
IntMap map[string]int
FloatMap map[string]float32
Data map[string]interface{}
CSRFToken string //Cross site request forgery token
Flash string
Warning string
Error string
Form *forms.Form
}