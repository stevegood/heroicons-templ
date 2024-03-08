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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"20\" height=\"20\" viewBox=\"0 0 20 20\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M10 2C10.4142 2 10.75 2.33579 10.75 2.75V8.3401L12.7004 6.23966C12.9823 5.93613 13.4568 5.91855 13.7603 6.20041C14.0639 6.48226 14.0814 6.95681 13.7996 7.26034L10.5496 10.7603C10.4077 10.9132 10.2086 11 10 11C9.79145 11 9.59231 10.9132 9.45041 10.7603L6.20041 7.26034C5.91855 6.95681 5.93613 6.48226 6.23966 6.20041C6.54319 5.91855 7.01774 5.93613 7.29959 6.23966L9.25 8.3401V2.75C9.25 2.33579 9.58579 2 10 2Z\" fill=\"#0F172A\"></path> <path d=\"M5.27298 4.5C4.71065 4.5 4.21753 4.8755 4.06799 5.41759L2.54501 10.9384C2.53938 10.9588 2.53427 10.9794 2.52969 11H6C6.37877 11 6.72504 11.214 6.89443 11.5528L7.34164 12.4472C7.51103 12.786 7.8573 13 8.23607 13H11.674C12.0269 13 12.3537 12.814 12.5339 12.5105L13.1401 11.4895C13.3203 11.186 13.6471 11 14 11H17.4703C17.4657 10.9794 17.4606 10.9588 17.455 10.9384L15.932 5.41759C15.7825 4.8755 15.2894 4.5 14.727 4.5H13.75C13.3358 4.5 13 4.16421 13 3.75C13 3.33579 13.3358 3 13.75 3H14.727C15.9642 3 17.049 3.8261 17.378 5.0187L18.901 10.5395C18.9667 10.7777 19 11.0237 19 11.2708V15C19 16.1046 18.1046 17 17 17H3C1.89543 17 1 16.1046 1 15V11.2708C1 11.0237 1.03331 10.7777 1.09902 10.5395L2.622 5.0187C2.95099 3.82611 4.03584 3 5.27298 3H6.25C6.66421 3 7 3.33579 7 3.75C7 4.16421 6.66421 4.5 6.25 4.5H5.27298Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
