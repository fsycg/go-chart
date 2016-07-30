package chart

// Series is an alias to Renderable.
type Series interface {
	GetName() string
	GetYAxis() yAxisType
	GetStyle() Style
	Render(r Renderer, canvasBox Box, xrange, yrange Range, s Style)
}
