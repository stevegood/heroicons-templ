// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func PaintBrush() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"20\" height=\"20\" viewBox=\"0 0 20 20\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M15.9932 1.38468C16.3195 1.13518 16.7188 1 17.1295 1C18.1626 1 19.0002 1.83751 19.0002 2.87063C19.0002 3.28136 18.865 3.68067 18.6155 4.00694L14.5856 9.27674C13.4464 10.7666 11.9978 11.9799 10.349 12.839C9.93563 11.2868 8.71335 10.0645 7.16113 9.6512C8.02023 8.00236 9.23359 6.5538 10.7234 5.41453L15.9932 1.38468Z\" fill=\"#0F172A\"></path> <path d=\"M5.99977 11C4.34292 11 2.99977 12.3431 2.99977 14C2.99977 14.2761 2.77591 14.5 2.49977 14.5C2.41931 14.5 2.3453 14.4815 2.2798 14.4493C2.00461 14.314 1.67475 14.3598 1.44686 14.565C1.21898 14.7702 1.139 15.0935 1.2449 15.3813C1.80692 16.9087 3.27495 18 4.99977 18C7.20742 18 8.99735 16.2113 8.99977 14.0042L8.99977 14C8.99977 13.52 8.88647 13.0642 8.68446 12.6601C8.39396 12.0789 7.92089 11.6058 7.33969 11.3153C6.93554 11.1133 6.47981 11 5.99977 11Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
