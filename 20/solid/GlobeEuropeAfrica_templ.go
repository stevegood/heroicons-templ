// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func GlobeEuropeAfrica() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"20\" height=\"20\" viewBox=\"0 0 20 20\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M18 10C18 14.4183 14.4183 18 10 18C5.58172 18 2 14.4183 2 10C2 5.58172 5.58172 2 10 2C14.4183 2 18 5.58172 18 10ZM16.4969 10.2041C16.389 13.6995 13.5216 16.5 10 16.5C6.41015 16.5 3.5 13.5899 3.5 10C3.5 7.12656 5.36451 4.68861 7.94969 3.82998L6.92671 5.62019C6.60171 6.18894 6.69759 6.90485 7.16078 7.36805C7.6029 7.81017 8.27833 7.91977 8.83757 7.64015L9.0123 7.55279C9.08173 7.51807 9.15829 7.5 9.23591 7.5H9.38181C9.5712 7.5 9.74433 7.607 9.82902 7.77639C9.8994 7.91716 9.8994 8.08284 9.82902 8.22361L9.80123 8.27919C9.73357 8.41452 9.59525 8.5 9.44395 8.5H8.9418C8.18614 8.5 7.48048 8.87766 7.06131 9.50641L7.01751 9.57212C6.65828 10.111 6.56795 10.7855 6.77274 11.3998C6.98201 12.0276 7.47787 12.5197 8.10338 12.7282C8.33997 12.8071 8.49984 13.0292 8.49984 13.2749V14.3246C8.49984 14.9738 9.02608 15.5 9.67523 15.5C10.0323 15.5 10.37 15.3377 10.5931 15.0589L12.2044 13.0446C12.3957 12.8056 12.4998 12.5086 12.4998 12.2025C12.4998 11.8866 12.6279 11.5791 12.8531 11.3539C13.3088 10.8982 13.3864 10.1784 13.0257 9.63738L12.5617 8.94145C12.5214 8.88091 12.4998 8.80979 12.4998 8.73703C12.4998 8.43336 12.8465 8.26002 13.0895 8.44222L13.4322 8.69925C13.748 8.93609 14.1704 8.97373 14.523 8.79746C14.7488 8.68456 15.0216 8.72887 15.1996 8.90689L16.4969 10.2041Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
