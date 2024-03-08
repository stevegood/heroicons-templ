// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func ArrowPathRoundedSquare(attrs templ.Attributes) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" width=\"20\" height=\"20\" viewBox=\"0 0 20 20\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M10 4.5C11.215 4.5 12.4171 4.55484 13.6038 4.66214C13.9249 4.69118 14.1802 4.93931 14.2185 5.25934C14.3426 6.29702 14.4265 7.34718 14.4685 8.40824L12.7804 6.71973C12.4875 6.4268 12.0126 6.42675 11.7197 6.71961C11.4268 7.01247 11.4267 7.48735 11.7196 7.78027L14.7189 10.7802C15.0117 11.0731 15.4866 11.0732 15.7795 10.7803L18.7803 7.78041C19.0732 7.48755 19.0733 7.01268 18.7804 6.71975C18.4875 6.42681 18.0127 6.42674 17.7197 6.71959L15.9719 8.46698C15.9299 7.32601 15.8413 6.19681 15.7079 5.08123C15.5855 4.05775 14.7654 3.26106 13.7389 3.16824C12.5073 3.05688 11.2602 3 10 3C8.73981 3 7.49271 3.05688 6.26115 3.16824C5.23465 3.26105 4.41449 4.05775 4.2921 5.08123C4.22831 5.61464 4.17477 6.15117 4.13167 6.69061C4.09868 7.10351 4.40666 7.46497 4.81956 7.49796C5.23245 7.53095 5.59391 7.22297 5.6269 6.81007C5.66843 6.29024 5.72003 5.77326 5.78148 5.25934C5.81975 4.93931 6.07514 4.69118 6.39623 4.66214C7.58294 4.55484 8.78497 4.5 10 4.5ZM5.28113 9.22C4.98828 8.9271 4.51345 8.92704 4.22052 9.21987L1.21976 12.2196C0.926819 12.5124 0.926735 12.9873 1.21958 13.2802C1.51242 13.5732 1.98729 13.5733 2.28024 13.2804L4.02814 11.5331C4.0701 12.6741 4.15869 13.8032 4.2921 14.9188C4.41449 15.9422 5.23465 16.7389 6.26115 16.8318C7.49271 16.9431 8.73981 17 10 17C11.2602 17 12.5073 16.9431 13.7389 16.8318C14.7654 16.7389 15.5855 15.9423 15.7079 14.9188C15.7717 14.3849 15.8253 13.848 15.8684 13.3081C15.9014 12.8952 15.5934 12.5338 15.1805 12.5008C14.7676 12.4678 14.4062 12.7758 14.3732 13.1887C14.3316 13.709 14.28 14.2263 14.2185 14.7407C14.1802 15.0607 13.9249 15.3088 13.6038 15.3379C12.4171 15.4452 11.215 15.5 10 15.5C8.78497 15.5 7.58294 15.4452 6.39623 15.3379C6.07514 15.3088 5.81975 15.0607 5.78148 14.7407C5.6574 13.703 5.57352 12.6529 5.53148 11.5919L7.21963 13.2803C7.5125 13.5732 7.98737 13.5732 8.28029 13.2804C8.57321 12.9875 8.57324 12.5126 8.28037 12.2197L5.28113 9.22Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
