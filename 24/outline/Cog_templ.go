// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package outline

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Cog() templ.Component {
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
		templ_7745c5c3_Err = CogWithAttrs(templ.Attributes{}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func CogWithAttrs(attrs templ.Attributes) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M4.50073 11.9993C4.50073 16.1414 7.8586 19.4993 12.0007 19.4993C16.1429 19.4993 19.5007 16.1414 19.5007 11.9993M4.50073 11.9993C4.50073 7.85712 7.8586 4.49925 12.0007 4.49925C16.1429 4.49926 19.5007 7.85712 19.5007 11.9993M4.50073 11.9993L3.00073 11.9993M19.5007 11.9993L21.0007 11.9993M19.5007 11.9993L12.0007 11.9993M3.54329 15.0774L4.95283 14.5644M19.0482 9.43411L20.4578 8.92108M5.1062 17.785L6.25527 16.8208M17.7459 7.17897L18.895 6.21479M7.50064 19.7943L8.25064 18.4952M15.7506 5.50484L16.5006 4.2058M10.4378 20.8633L10.6983 19.386M13.303 4.61393L13.5635 3.13672M13.5635 20.8633L13.303 19.3861M10.6983 4.61397L10.4378 3.13676M16.5007 19.7941L15.7507 18.4951M7.50068 4.20565L12.0007 11.9993M18.8952 17.7843L17.7461 16.8202M6.25542 7.17835L5.10635 6.21417M20.458 15.0776L19.0485 14.5646M4.95308 9.43426L3.54354 8.92123M12.0007 11.9993L8.25073 18.4944\" stroke=\"#0F172A\" stroke-width=\"1.5\" stroke-linecap=\"round\" stroke-linejoin=\"round\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
