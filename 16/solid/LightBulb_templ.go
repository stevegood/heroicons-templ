// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func LightBulb() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M10.6177 10.2608C10.2569 10.483 10 10.8578 10 11.2816C10 11.5075 9.85818 11.7112 9.64083 11.7729C9.11933 11.9208 8.56891 12 8 12C7.43109 12 6.88067 11.9208 6.35917 11.7729C6.14182 11.7112 6 11.5075 6 11.2816C6 10.8578 5.74314 10.483 5.38228 10.2608C3.95293 9.3808 3 7.80168 3 6C3 3.23858 5.23858 1 8 1C10.7614 1 13 3.23858 13 6C13 7.80168 12.0471 9.3808 10.6177 10.2608Z\" fill=\"black\"></path> <path d=\"M6.86696 13.4151C6.45741 13.3531 6.07514 13.6348 6.01313 14.0444C5.95112 14.4539 6.23286 14.8362 6.64241 14.8982C7.08559 14.9653 7.53898 15 8 15C8.46102 15 8.91442 14.9653 9.3576 14.8982C9.76715 14.8362 10.0489 14.4539 9.98688 14.0444C9.92487 13.6348 9.5426 13.3531 9.13305 13.4151C8.76394 13.471 8.3856 13.5 8 13.5C7.6144 13.5 7.23607 13.471 6.86696 13.4151Z\" fill=\"black\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
