package datatree

func (m Model) WithStyles(styles Styles) Model {
	m.styles = styles
	m.updateContents()

	return m
}

func (m Model) WithStyleDefault() Model {
	m.styles = styleDefault
	m.updateContents()

	return m
}

func (m Model) WithStyleBlank() Model {
	m.styles = styleBlank
	m.updateContents()

	return m
}
