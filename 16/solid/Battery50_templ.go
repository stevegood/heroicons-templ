// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Battery50() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M1 6.25C1 5.00736 2.00736 4 3.25 4H11.75C12.9926 4 14 5.00736 14 6.25V6.33535C14.5826 6.54127 15 7.09689 15 7.75V8.25C15 8.90311 14.5826 9.45873 14 9.66465V9.75C14 10.9926 12.9926 12 11.75 12H3.25C2.00736 12 1 10.9926 1 9.75V6.25ZM3.25 5.5C2.83579 5.5 2.5 5.83579 2.5 6.25V9.75C2.5 10.1642 2.83579 10.5 3.25 10.5H11.75C12.1642 10.5 12.5 10.1642 12.5 9.75V6.25C12.5 5.83579 12.1642 5.5 11.75 5.5H3.25Z\" fill=\"#0F172A\"></path> <path d=\"M4.75 7C4.33579 7 4 7.33579 4 7.75V8.25C4 8.66421 4.33579 9 4.75 9H6.75C7.16421 9 7.5 8.66421 7.5 8.25V7.75C7.5 7.33579 7.16421 7 6.75 7H4.75Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
