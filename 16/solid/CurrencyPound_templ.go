// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func CurrencyPound() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M15 8C15 11.866 11.866 15 8 15C4.13401 15 1 11.866 1 8C1 4.13401 4.13401 1 8 1C11.866 1 15 4.13401 15 8ZM7.93952 4.93951C7.64647 5.23255 7.50018 5.61512 7.50018 6.00017V7.25H8.74997C9.16419 7.25 9.49998 7.58579 9.49998 8C9.49998 8.41421 9.16419 8.75 8.74997 8.75H7.50018V9.75018C7.50018 10.0131 7.45506 10.2656 7.37215 10.5002L11.2498 10.5001C11.664 10.5001 11.9998 10.8359 11.9998 11.2501C11.9998 11.6643 11.6641 12.0001 11.2499 12.0001L4.74993 12.0002C4.33572 12.0002 3.9999 11.6644 3.99988 11.2502C3.99986 10.836 4.33562 10.5002 4.74983 10.5002L5.25022 10.5002C5.66442 10.5002 6.00018 10.1644 6.00018 9.75018V8.75H4.74997C4.33576 8.75 3.99997 8.41421 3.99997 8C3.99997 7.58579 4.33576 7.25 4.74997 7.25H6.00018V6.00017C6.00018 5.23356 6.29321 4.46449 6.87885 3.87885C8.05043 2.70728 9.94992 2.70727 11.1215 3.87885C11.4144 4.17174 11.4144 4.64661 11.1215 4.93951C10.8286 5.2324 10.3537 5.2324 10.0608 4.93951C9.47505 4.35372 8.5253 4.35372 7.93952 4.93951Z\" fill=\"black\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
