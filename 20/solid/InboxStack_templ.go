// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func InboxStack(classNames string) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" width=\"20\" height=\"20\" viewBox=\"0 0 20 20\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M1.04461 6.95375C1.08721 6.72011 1.16003 6.49228 1.26175 6.27613L2.53096 3.57906C2.98454 2.6152 3.95396 2 5.01921 2H14.9808C16.046 2 17.0155 2.6152 17.469 3.57906L18.7383 6.27613C18.84 6.49228 18.9128 6.72011 18.9554 6.95375C18.9844 7.04736 19 7.14686 19 7.25V8.75C19 10.2688 17.7688 11.5 16.25 11.5H3.75C2.23122 11.5 1 10.2688 1 8.75V7.25C1 7.14686 1.01562 7.04736 1.04461 6.95375ZM3.88819 4.21775C4.09436 3.77964 4.53501 3.5 5.01921 3.5H14.9808C15.465 3.5 15.9056 3.77964 16.1118 4.21775L17.0682 6.25H14C13.6471 6.25 13.3203 6.43601 13.1401 6.73946L12.5339 7.76054C12.3537 8.06399 12.0269 8.25 11.674 8.25H8.23607C7.8573 8.25 7.51103 8.036 7.34164 7.69721L6.89443 6.80279C6.72504 6.464 6.37877 6.25 6 6.25H2.93184L3.88819 4.21775Z\" fill=\"#0F172A\"></path> <path d=\"M1 14C1 13.4477 1.44772 13 2 13H6C6.37877 13 6.72504 13.214 6.89443 13.5528L7.34164 14.4472C7.51103 14.786 7.8573 15 8.23607 15H11.674C12.0269 15 12.3537 14.814 12.5339 14.5105L13.1401 13.4895C13.3203 13.186 13.6471 13 14 13H18C18.5523 13 19 13.4477 19 14V16C19 17.1046 18.1046 18 17 18H3C1.89543 18 1 17.1046 1 16V14Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
