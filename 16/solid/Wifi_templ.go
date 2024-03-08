// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Wifi() templ.Component {
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
		templ_7745c5c3_Err = WifiWithAttrs(templ.Attributes{}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func WifiWithAttrs(attrs templ.Attributes) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M14.1879 7.06282C10.7708 3.64573 5.23062 3.64573 1.81353 7.06282C1.52064 7.35571 1.04577 7.35571 0.752873 7.06282C0.45998 6.76992 0.45998 6.29505 0.752873 6.00216C4.75575 1.99928 11.2457 1.99928 15.2486 6.00216C15.5415 6.29505 15.5415 6.76992 15.2486 7.06282C14.9557 7.35571 14.4808 7.35571 14.1879 7.06282ZM12.0666 9.18414C9.82107 6.93862 6.18037 6.93862 3.93485 9.18414C3.64196 9.47703 3.16709 9.47703 2.87419 9.18414C2.5813 8.89124 2.5813 8.41637 2.87419 8.12348C5.70549 5.29217 10.2959 5.29217 13.1272 8.12348C13.4201 8.41637 13.4201 8.89124 13.1272 9.18414C12.8343 9.47703 12.3595 9.47703 12.0666 9.18414ZM9.94526 11.3055C8.87132 10.2315 7.13012 10.2315 6.05617 11.3055C5.76328 11.5984 5.28841 11.5984 4.99551 11.3055C4.70262 11.0126 4.70262 10.5377 4.99551 10.2448C6.65524 8.58507 9.34619 8.58507 11.0059 10.2448C11.2988 10.5377 11.2988 11.0126 11.0059 11.3055C10.713 11.5984 10.2382 11.5984 9.94526 11.3055ZM7.11683 12.3661C7.60499 11.878 8.39645 11.878 8.8846 12.3661C9.17749 12.659 9.17749 13.1339 8.8846 13.4268L8.53033 13.781C8.23744 14.0739 7.76256 14.0739 7.46967 13.781L7.11612 13.4275C6.82343 13.1348 6.82319 12.6603 7.11559 12.3674L7.11683 12.3661Z\" fill=\"black\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
