// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func PaperClip() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"20\" height=\"20\" viewBox=\"0 0 20 20\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M15.6213 4.37868C14.4497 3.20711 12.5503 3.20711 11.3787 4.37868L4.37868 11.3787C3.20711 12.5503 3.20711 14.4497 4.37868 15.6213C5.54995 16.7926 7.44878 16.7929 8.62042 15.6222C8.62072 15.6219 8.62102 15.6216 8.62132 15.6213L9.11792 15.1214C9.40985 14.8276 9.88472 14.826 10.1786 15.1179C10.4724 15.4098 10.474 15.8847 10.1821 16.1786L9.68373 16.6802L9.68198 16.682C7.92462 18.4393 5.07538 18.4393 3.31802 16.682C1.56066 14.9246 1.56066 12.0754 3.31802 10.318L10.318 3.31802C12.0754 1.56066 14.9246 1.56066 16.682 3.31802C18.438 5.07407 18.4393 7.92038 16.6859 9.67806L13.2312 13.2312C12.2061 14.2564 10.544 14.2563 9.51885 13.2312C8.49372 12.206 8.49372 10.544 9.51885 9.51886L12.9697 6.06804C13.2626 5.77515 13.7374 5.77515 14.0303 6.06804C14.3232 6.36094 14.3232 6.83581 14.0303 7.1287L10.5795 10.5795C10.1402 11.0189 10.1402 11.7312 10.5795 12.1705C11.0178 12.6088 11.7276 12.6099 12.1672 12.1738L15.6213 8.62127C16.7928 7.4497 16.7929 5.55025 15.6213 4.37868Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
