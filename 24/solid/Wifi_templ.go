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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M1.37128 8.1434C7.22915 2.28553 16.7266 2.28553 22.5845 8.1434C22.8774 8.43628 22.8774 8.91114 22.5845 9.20404L22.0542 9.73441C21.9135 9.87507 21.7228 9.9541 21.5238 9.9541C21.3249 9.95411 21.1341 9.87509 20.9935 9.73443C16.0143 4.75525 7.94146 4.75525 2.96227 9.73443C2.82162 9.87509 2.63084 9.95411 2.43193 9.9541C2.23301 9.9541 2.04224 9.87507 1.90159 9.73441L1.37126 9.20404C1.07839 8.91114 1.0784 8.43628 1.37128 8.1434ZM4.55326 11.3254C8.65377 7.22493 15.302 7.22493 19.4025 11.3254C19.6954 11.6183 19.6954 12.0932 19.4025 12.3861L18.8722 12.9164C18.7316 13.0571 18.5408 13.1361 18.3419 13.1361C18.143 13.1361 17.9522 13.0571 17.8115 12.9164C14.5897 9.6946 9.3661 9.6946 6.14427 12.9164C5.85139 13.2093 5.37652 13.2093 5.08362 12.9164L4.55327 12.3861C4.41262 12.2455 4.33359 12.0547 4.33359 11.8558C4.33359 11.6569 4.41261 11.4661 4.55326 11.3254ZM7.75733 14.5074C10.1005 12.1642 13.8995 12.1642 16.2426 14.5074C16.5355 14.8003 16.5355 15.2751 16.2426 15.568L15.7123 16.0983C15.5716 16.239 15.3809 16.318 15.182 16.318C14.983 16.318 14.7923 16.239 14.6516 16.0983C13.1872 14.6339 10.8128 14.6339 9.34832 16.0983C9.05543 16.3912 8.58055 16.3912 8.28766 16.0983L7.75733 15.568C7.46444 15.2751 7.46444 14.8003 7.75733 14.5074ZM10.9393 17.6893C11.5251 17.1036 12.4748 17.1036 13.0606 17.6893C13.3535 17.9822 13.3535 18.4571 13.0606 18.75L12.5303 19.2803C12.3896 19.421 12.1989 19.5 12 19.5C11.8011 19.5 11.6103 19.421 11.4696 19.2803L10.9393 18.75C10.6464 18.4571 10.6464 17.9822 10.9393 17.6893Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
