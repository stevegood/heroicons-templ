// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Signal() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"20\" height=\"20\" viewBox=\"0 0 20 20\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M16.364 3.63605C16.0711 3.34316 15.5962 3.34316 15.3033 3.63605C15.0104 3.92895 15.0104 4.40382 15.3033 4.69671C18.2322 7.62564 18.2322 12.3744 15.3033 15.3033C15.0104 15.5962 15.0104 16.0711 15.3033 16.364C15.5962 16.6569 16.0711 16.6569 16.364 16.364C19.8787 12.8493 19.8787 7.15077 16.364 3.63605Z\" fill=\"#0F172A\"></path> <path d=\"M4.6967 4.69671C4.98959 4.40382 4.98959 3.92895 4.6967 3.63605C4.40381 3.34316 3.92893 3.34316 3.63604 3.63605C0.12132 7.15077 0.12132 12.8493 3.63604 16.364C3.92893 16.6569 4.40381 16.6569 4.6967 16.364C4.98959 16.0711 4.98959 15.5962 4.6967 15.3033C1.76777 12.3744 1.76777 7.62564 4.6967 4.69671Z\" fill=\"#0F172A\"></path> <path d=\"M12.4749 6.46448C12.7678 6.17159 13.2426 6.17159 13.5355 6.46448C15.4882 8.41711 15.4882 11.5829 13.5355 13.5356C13.2426 13.8284 12.7678 13.8284 12.4749 13.5356C12.182 13.2427 12.182 12.7678 12.4749 12.4749C13.8417 11.1081 13.8417 8.89198 12.4749 7.52515C12.182 7.23225 12.182 6.75738 12.4749 6.46448Z\" fill=\"#0F172A\"></path> <path d=\"M7.52513 6.46448C7.81802 6.75738 7.81802 7.23225 7.52513 7.52515C6.15829 8.89198 6.15829 11.1081 7.52513 12.4749C7.81802 12.7678 7.81802 13.2427 7.52513 13.5356C7.23223 13.8284 6.75736 13.8284 6.46447 13.5356C4.51184 11.5829 4.51184 8.41711 6.46447 6.46448C6.75736 6.17159 7.23223 6.17159 7.52513 6.46448Z\" fill=\"#0F172A\"></path> <path d=\"M11 10C11 10.5523 10.5523 11 10 11C9.44772 11 9 10.5523 9 10C9 9.44771 9.44772 9 10 9C10.5523 9 11 9.44771 11 10Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
