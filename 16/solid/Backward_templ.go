// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Backward() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M8.5 4.75C8.5 4.48569 8.36087 4.24089 8.13379 4.10564C7.90671 3.97038 7.62519 3.96464 7.39279 4.09053L1.39279 7.34053C1.15077 7.47163 1 7.72476 1 8C1 8.27524 1.15077 8.52837 1.39279 8.65947L7.39279 11.9095C7.62519 12.0354 7.90671 12.0296 8.13379 11.8944C8.36087 11.7591 8.5 11.5143 8.5 11.25V8.98838L13.8928 11.9095C14.1252 12.0354 14.4067 12.0296 14.6338 11.8944C14.8609 11.7591 15 11.5143 15 11.25V4.75C15 4.48569 14.8609 4.24089 14.6338 4.10564C14.4067 3.97038 14.1252 3.96464 13.8928 4.09053L8.5 7.01162V4.75Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}