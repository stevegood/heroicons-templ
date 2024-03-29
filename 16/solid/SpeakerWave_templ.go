// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func SpeakerWave() templ.Component {
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
		templ_7745c5c3_Err = SpeakerWaveWithAttrs(templ.Attributes{}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func SpeakerWaveWithAttrs(attrs templ.Attributes) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M7.55724 2.06583C7.82667 2.18682 8 2.45467 8 2.75001V13.25C8 13.5454 7.82667 13.8132 7.55724 13.9342C7.28782 14.0552 6.97247 14.0068 6.75173 13.8106L3.58984 11H2C1.44772 11 1 10.5523 1 10V6C1 5.44772 1.44772 5 2 5H3.58986L6.75173 2.18946C6.97247 1.99324 7.28782 1.94484 7.55724 2.06583Z\" fill=\"black\"></path> <path d=\"M12.9497 3.05024C12.6569 2.75734 12.182 2.75734 11.8891 3.05024C11.5962 3.34313 11.5962 3.818 11.8891 4.1109C14.037 6.25878 14.037 9.74119 11.8891 11.8891C11.5962 12.182 11.5962 12.6568 11.8891 12.9497C12.182 13.2426 12.6569 13.2426 12.9497 12.9497C15.6834 10.2161 15.6834 5.78391 12.9497 3.05024Z\" fill=\"black\"></path> <path d=\"M10.8284 5.17156C10.5355 4.87866 10.0607 4.87866 9.76777 5.17156C9.47487 5.46445 9.47487 5.93932 9.76777 6.23222C10.7441 7.20853 10.7441 8.79144 9.76777 9.76775C9.47487 10.0606 9.47487 10.5355 9.76777 10.8284C10.0607 11.1213 10.5355 11.1213 10.8284 10.8284C12.3905 9.26631 12.3905 6.73365 10.8284 5.17156Z\" fill=\"black\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
