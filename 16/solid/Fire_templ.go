// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Fire() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M8.07402 0.944846C6.81757 1.85307 6 3.33094 6 4.99976L6.0001 5.0321C6.00387 5.6327 6.11354 6.20818 6.31136 6.74081C6.4705 7.16926 6.10663 7.65174 5.70092 7.44129C4.94479 7.04907 4.30153 6.46942 3.83266 5.76387C3.63093 5.4603 3.18469 5.40126 2.98487 5.7061C2.92216 5.80176 2.86179 5.89977 2.80387 6.00007C1.14702 8.86983 2.13027 12.5394 5.00003 14.1962C7.86978 15.8531 11.5393 14.8698 13.1962 12.0001C14.8527 9.13095 13.8702 5.46237 11.0019 3.805L10.9981 3.79848C9.9539 3.195 9.20965 2.27165 8.81771 1.22373C8.70214 0.914737 8.34139 0.751584 8.07402 0.944846ZM8.8536 7.12299C10.0946 7.49066 11 8.6395 11 9.99981C11 10.9622 10.5468 11.8188 9.84219 12.3678C9.3338 12.7638 8.69401 12.9997 7.99968 12.9997C6.87478 12.9997 5.89449 12.3805 5.38097 11.4645C5.17594 11.0987 5.59569 10.7699 6.00184 10.8742C6.32081 10.9561 6.65515 10.9997 6.99968 10.9997C7.30017 10.9997 7.59292 10.9665 7.87449 10.9037C8.16227 10.8395 8.28666 10.5182 8.19544 10.2378C8.06858 9.84786 8 9.43158 8 8.99929C8 8.41427 8.12559 7.85859 8.35128 7.35773C8.43815 7.16493 8.65084 7.06292 8.8536 7.12299Z\" fill=\"black\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
