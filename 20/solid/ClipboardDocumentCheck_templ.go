// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func ClipboardDocumentCheck(classNames string) templ.Component {
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
		var templ_7745c5c3_Var2 = []any{classNames}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var2...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ.CSSClasses(templ_7745c5c3_Var2).String()))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" width=\"20\" height=\"20\" viewBox=\"0 0 20 20\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M17.9999 5.25C17.9999 4.08761 17.1185 3.1311 15.9875 3.0124C15.8688 1.88145 14.9123 1 13.7499 1H12.2499C11.0875 1 10.131 1.88145 10.0123 3.0124C9.13721 3.10424 8.4115 3.69769 8.12793 4.5H10.9999C12.3806 4.5 13.4999 5.61929 13.4999 7V14H15.7499C16.9925 14 17.9999 12.9926 17.9999 11.75V5.25ZM12.2499 2.5C11.8357 2.5 11.4999 2.83579 11.4999 3.25V3.5H14.4999V3.25C14.4999 2.83579 14.1641 2.5 13.7499 2.5H12.2499Z\" fill=\"#0F172A\"></path> <path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M3 6C2.44772 6 2 6.44771 2 7V17C2 17.5523 2.44772 18 3 18H11C11.5523 18 12 17.5523 12 17V7C12 6.44772 11.5523 6 11 6H3ZM9.87404 10.166C10.1038 9.82138 10.0107 9.35573 9.66603 9.12596C9.32138 8.8962 8.85573 8.98933 8.62596 9.33397L6.13343 13.0728L5.28033 12.2197C4.98744 11.9268 4.51256 11.9268 4.21967 12.2197C3.92678 12.5126 3.92678 12.9874 4.21967 13.2803L5.71967 14.7803C5.87855 14.9392 6.1003 15.0185 6.3239 14.9963C6.5475 14.9742 6.7494 14.853 6.87404 14.666L9.87404 10.166Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
