// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func ChevronDoubleRight() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"20\" height=\"20\" viewBox=\"0 0 20 20\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M15.2803 9.46967C15.5732 9.76256 15.5732 10.2374 15.2803 10.5303L11.0303 14.7803C10.7374 15.0732 10.2626 15.0732 9.96967 14.7803C9.67678 14.4874 9.67678 14.0126 9.96967 13.7197L13.6893 10L9.96967 6.28033C9.67678 5.98744 9.67678 5.51256 9.96967 5.21967C10.2626 4.92678 10.7374 4.92678 11.0303 5.21967L15.2803 9.46967ZM6.03033 5.21967L10.2803 9.46967C10.5732 9.76256 10.5732 10.2374 10.2803 10.5303L6.03033 14.7803C5.73744 15.0732 5.26256 15.0732 4.96967 14.7803C4.67678 14.4874 4.67678 14.0126 4.96967 13.7197L8.68934 10L4.96967 6.28033C4.67678 5.98744 4.67678 5.51256 4.96967 5.21967C5.26256 4.92678 5.73744 4.92678 6.03033 5.21967Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
