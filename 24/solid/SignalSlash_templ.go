// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func SignalSlash() templ.Component {
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
		templ_7745c5c3_Err = SignalSlashWithAttrs(templ.Attributes{}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func SignalSlashWithAttrs(attrs templ.Attributes) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M2.46968 2.46967C2.76257 2.17678 3.23744 2.17678 3.53034 2.46967L11.9374 10.8767C11.9581 10.8756 11.979 10.875 12 10.875C12.6213 10.875 13.125 11.3787 13.125 12C13.125 12.021 13.1244 12.0419 13.1233 12.0626L14.5848 13.5242C15.2643 12.3737 15.1098 10.8671 14.1213 9.87868C13.8284 9.58579 13.8284 9.11091 14.1213 8.81802C14.4142 8.52513 14.8891 8.52513 15.182 8.81802C16.7587 10.3947 16.9208 12.8503 15.6683 14.6077L16.7402 15.6795C18.572 13.3257 18.4062 9.9209 16.2426 7.75736C15.9498 7.46447 15.9498 6.98959 16.2426 6.6967C16.5355 6.40381 17.0104 6.40381 17.3033 6.6967C20.0533 9.44664 20.2213 13.8008 17.8075 16.7468L18.8725 17.8118C21.8696 14.2758 21.7001 8.97217 18.364 5.63604C18.0711 5.34315 18.0711 4.86827 18.364 4.57538C18.6569 4.28249 19.1317 4.28249 19.4246 4.57538C23.3468 8.49757 23.5174 14.7507 19.9363 18.8756L21.5303 20.4697C21.8232 20.7626 21.8232 21.2374 21.5303 21.5303C21.2374 21.8232 20.7626 21.8232 20.4697 21.5303L18.3644 19.4251L16.2426 17.3033L16.242 17.3027L11.5366 12.5972C11.4896 12.5603 11.4472 12.5179 11.4103 12.4709L2.46968 3.53033C2.17678 3.23744 2.17678 2.76256 2.46968 2.46967ZM3.65854 6.89187C4.02729 7.08054 4.17326 7.53242 3.98459 7.90117C2.25352 11.2844 2.80534 15.5333 5.63605 18.364C5.92894 18.6569 5.92894 19.1317 5.63605 19.4246C5.34315 19.7175 4.86828 19.7175 4.57539 19.4246C1.27115 16.1204 0.630437 11.1635 2.64924 7.21792C2.83791 6.84917 3.28979 6.70319 3.65854 6.89187ZM5.84076 9.13355C6.23339 9.2655 6.44473 9.69075 6.31278 10.0834C5.6094 12.1765 6.09233 14.5776 7.75737 16.2426C8.05026 16.5355 8.05026 17.0104 7.75737 17.3033C7.46447 17.5962 6.9896 17.5962 6.69671 17.3033C4.61367 15.2203 4.01314 12.2177 4.89092 9.60558C5.02286 9.21294 5.44812 9.00161 5.84076 9.13355ZM8.182 11.7875C8.5921 11.7293 8.97176 12.0145 9.02999 12.4246C9.11816 13.0455 9.40043 13.6431 9.87869 14.1213C10.1716 14.4142 10.1716 14.8891 9.87869 15.182C9.58579 15.4749 9.11092 15.4749 8.81803 15.182C8.10217 14.4661 7.67717 13.5671 7.54489 12.6355C7.48666 12.2254 7.7719 11.8458 8.182 11.7875Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
