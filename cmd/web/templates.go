package main

import "snipperbox.illia-kornyk/internal/models"

type templateData struct {
	Snippet *models.Snippet
	Snippets []*models.Snippet
}
