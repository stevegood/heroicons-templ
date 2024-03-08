// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func ChevronDoubleUp(attrs templ.Attributes) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.RenderAttributes(ctx, templ_7745c5c3_Buffer, attrs)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M7.46967 3.21967C7.76256 2.92678 8.23744 2.92678 8.53033 3.21967L11.7803 6.46967C12.0732 6.76256 12.0732 7.23744 11.7803 7.53033C11.4874 7.82322 11.0126 7.82322 10.7197 7.53033L8 4.81066L5.28033 7.53033C4.98744 7.82322 4.51256 7.82322 4.21967 7.53033C3.92678 7.23744 3.92678 6.76256 4.21967 6.46967L7.46967 3.21967ZM4.21967 11.4697L7.46967 8.21967C7.76256 7.92678 8.23744 7.92678 8.53033 8.21967L11.7803 11.4697C12.0732 11.7626 12.0732 12.2374 11.7803 12.5303C11.4874 12.8232 11.0126 12.8232 10.7197 12.5303L8 9.81066L5.28033 12.5303C4.98744 12.8232 4.51256 12.8232 4.21967 12.5303C3.92678 12.2374 3.92678 11.7626 4.21967 11.4697Z\" fill=\"black\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
