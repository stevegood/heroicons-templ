// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func InboxArrowDown() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M8.75 2.75C8.75 2.33579 8.41421 2 8 2C7.58579 2 7.25 2.33579 7.25 2.75V6.43934L6.53033 5.71967C6.23744 5.42678 5.76256 5.42678 5.46967 5.71967C5.17678 6.01256 5.17678 6.48744 5.46967 6.78033L7.46967 8.78033C7.76256 9.07322 8.23744 9.07322 8.53033 8.78033L10.5303 6.78033C10.8232 6.48744 10.8232 6.01256 10.5303 5.71967C10.2374 5.42678 9.76256 5.42678 9.46967 5.71967L8.75 6.43934V2.75Z\" fill=\"black\"></path> <path d=\"M4.78373 4.5C4.47252 4.5 4.19365 4.69219 4.08286 4.983L2.55258 9H4.96482C5.29917 9 5.6114 9.1671 5.79687 9.4453L6.20313 10.0547C6.3886 10.3329 6.70083 10.5 7.03518 10.5H8.96482C9.29917 10.5 9.6114 10.3329 9.79687 10.0547L10.2031 9.4453C10.3886 9.1671 10.7008 9 11.0352 9H13.4474L11.9171 4.983C11.8063 4.69219 11.5275 4.5 11.2163 4.5H10.75C10.3358 4.5 10 4.16421 10 3.75C10 3.33579 10.3358 3 10.75 3H11.2163C12.1499 3 12.9865 3.57656 13.3189 4.44901L14.8526 8.47505C14.95 8.73086 15 9.0023 15 9.27604V10.75C15 11.9926 13.9926 13 12.75 13H3.25C2.00736 13 1 11.9926 1 10.75V9.27604C1 9.0023 1.04995 8.73086 1.1474 8.47505L2.68113 4.44901C3.01349 3.57656 3.85012 3 4.78373 3H5.25C5.66421 3 6 3.33579 6 3.75C6 4.16421 5.66421 4.5 5.25 4.5H4.78373Z\" fill=\"black\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
