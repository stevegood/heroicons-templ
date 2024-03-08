// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func BellSnooze() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M8 1C10.2091 1 12 2.79086 12 5V7.37868C12 7.7765 12.158 8.15804 12.4393 8.43934L13.7071 9.70711C13.8946 9.89464 14 10.149 14 10.4142V11C14 11.5523 13.5523 12 13 12H11C11 13.6569 9.65685 15 8 15C6.34315 15 5 13.6569 5 12H3C2.44772 12 2 11.5523 2 11V10.4142C2 10.149 2.10536 9.89464 2.29289 9.70711L3.56066 8.43934C3.84197 8.15804 4 7.7765 4 7.37868V5C4 2.79086 5.79086 1 8 1ZM8 13.5C7.17157 13.5 6.5 12.8284 6.5 12H9.5C9.5 12.8284 8.82843 13.5 8 13.5ZM6.75 4C6.33579 4 6 4.33579 6 4.75C6 5.16421 6.33579 5.5 6.75 5.5H7.79261L6.1397 7.81407C5.97641 8.04268 5.95457 8.34338 6.08313 8.59319C6.21168 8.84299 6.46906 9 6.75 9H9.25C9.66422 9 10 8.66421 10 8.25C10 7.83579 9.66422 7.5 9.25 7.5H8.20739L9.8603 5.18593C10.0236 4.95732 10.0454 4.65662 9.91688 4.40681C9.78833 4.15701 9.53094 4 9.25 4H6.75Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
