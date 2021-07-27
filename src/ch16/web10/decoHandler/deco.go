package decoHandler

import "net/http"

type DecoratorFunc func(w http.ResponseWriter, r *http.Request, h http.Handler)

type DecoHandler struct {
	fn DecoratorFunc
	h http.Handler
}

func (dc *DecoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dc.fn(w, r, dc.h)	
}

func NewDecoHandler(h http.Handler, fn DecoratorFunc) http.Handler {
	return &DecoHandler{
		fn:fn,
		h: h,
	}
}