// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func CurrencyYen(attrs templ.Attributes) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.RenderAttributes(ctx, templ_7745c5c3_Buffer, attrs)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" width=\"20\" height=\"20\" viewBox=\"0 0 20 20\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M10 18C14.4183 18 18 14.4183 18 10C18 5.58172 14.4183 2 10 2C5.58172 2 2 5.58172 2 10C2 14.4183 5.58172 18 10 18ZM7.34581 5.29439C7.09419 4.96536 6.62349 4.9026 6.29445 5.15421C5.96542 5.40582 5.90266 5.87653 6.15427 6.20557L9.05588 9.99998H6.75004C6.33583 9.99998 6.00004 10.3358 6.00004 10.75C6.00004 11.1642 6.33583 11.5 6.75004 11.5H9.25004V12.5H6.75004C6.33583 12.5 6.00004 12.8358 6.00004 13.25C6.00004 13.6642 6.33583 14 6.75004 14H9.25004V15.25C9.25004 15.6642 9.58583 16 10 16C10.4143 16 10.75 15.6642 10.75 15.25V14H13.25C13.6643 14 14 13.6642 14 13.25C14 12.8358 13.6643 12.5 13.25 12.5H10.75V11.5H13.25C13.6643 11.5 14 11.1642 14 10.75C14 10.3358 13.6643 9.99998 13.25 9.99998H10.9442L13.8458 6.20557C14.0974 5.87653 14.0347 5.40582 13.7056 5.15421C13.3766 4.9026 12.9059 4.96536 12.6543 5.29439L10 8.76531L7.34581 5.29439Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
