// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Printer() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M7.875 1.5C6.83947 1.5 6 2.33947 6 3.375V6.36564C5.5736 6.41799 5.1489 6.47583 4.72596 6.53912C3.27191 6.75668 2.25 8.02163 2.25 9.45569V15.75C2.25 17.4069 3.59315 18.75 5.25 18.75H5.51963L5.36461 20.4552C5.26479 21.5533 6.12935 22.5 7.23191 22.5H16.7681C17.8706 22.5 18.7352 21.5533 18.6354 20.4552L18.4804 18.75H18.75C20.4069 18.75 21.75 17.4069 21.75 15.75V9.45569C21.75 8.02163 20.7281 6.75668 19.274 6.53912C18.8511 6.47583 18.4264 6.41799 18 6.36564V3.375C18 2.33947 17.1605 1.5 16.125 1.5H7.875ZM16.5 6.20498V3.375C16.5 3.16789 16.3321 3 16.125 3H7.875C7.66789 3 7.5 3.16789 7.5 3.375V6.20498C8.98198 6.06931 10.483 6 12 6C13.517 6 15.018 6.06931 16.5 6.20498ZM16.2834 14.4696C16.4607 14.4879 16.5996 14.6298 16.6158 14.8073L17.1415 20.591C17.1615 20.8107 16.9886 21 16.7681 21H7.23191C7.0114 21 6.83849 20.8107 6.85845 20.591L7.38425 14.8073C7.40039 14.6298 7.53926 14.4879 7.71659 14.4696C9.12438 14.3244 10.5534 14.25 12 14.25C13.4466 14.25 14.8756 14.3244 16.2834 14.4696ZM17.25 10.5C17.25 10.0858 17.5858 9.75 18 9.75H18.0075C18.4217 9.75 18.7575 10.0858 18.7575 10.5V10.5075C18.7575 10.9217 18.4217 11.2575 18.0075 11.2575H18C17.5858 11.2575 17.25 10.9217 17.25 10.5075V10.5ZM15 9.75C14.5858 9.75 14.25 10.0858 14.25 10.5V10.5075C14.25 10.9217 14.5858 11.2575 15 11.2575H15.0075C15.4217 11.2575 15.7575 10.9217 15.7575 10.5075V10.5C15.7575 10.0858 15.4217 9.75 15.0075 9.75H15Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
