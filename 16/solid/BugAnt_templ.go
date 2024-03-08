// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func BugAnt() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M11.9828 1.36435C11.7674 1.01055 11.3059 0.898367 10.9521 1.11378C10.5983 1.32919 10.4862 1.79062 10.7016 2.14441C10.7976 2.30218 10.8859 2.46511 10.966 2.63268C10.7392 2.77825 10.501 2.9076 10.253 3.01909C9.70326 2.39427 8.89767 2 8 2C7.10231 2 6.2967 2.39429 5.74692 3.01913C5.49896 2.90763 5.26075 2.77827 5.0339 2.63268C5.11395 2.46511 5.20226 2.30218 5.29831 2.14441C5.51372 1.79062 5.40154 1.32919 5.04774 1.11378C4.69394 0.898367 4.23251 1.01055 4.0171 1.36435C3.77289 1.76544 3.56777 2.19322 3.40713 2.64236C3.29902 2.94464 3.39529 3.2822 3.64657 3.48199C4.0814 3.82772 4.55849 4.12286 5.06876 4.35844C5.02372 4.56514 5 4.7798 5 5C5 5.12593 5.07175 5.2399 5.18289 5.2991C5.56865 5.50457 5.9795 5.66916 6.40965 5.78707C6.28384 5.95178 6.18303 6.13661 6.1129 6.3359C5.19874 6.17008 4.32623 5.88595 3.51146 5.50007C3.50385 5.33601 3.5 5.17092 3.5 5.00488C3.5 4.92538 3.50088 4.8461 3.50264 4.76705C3.51183 4.35294 3.18357 4.00979 2.76946 4.00059C2.35535 3.9914 2.0122 4.31966 2.00301 4.73377C2.00101 4.82391 2 4.91429 2 5.00488C2 5.35455 2.01498 5.7009 2.04435 6.0433C2.06624 6.29854 2.21691 6.52495 2.44389 6.64372C3.59819 7.24771 4.86282 7.6693 6.19888 7.87051C6.23987 7.95517 6.28666 8.03651 6.33872 8.114C6.26744 8.12386 6.19636 8.13436 6.12549 8.14547L6.09828 8.14996C4.79932 8.35689 3.56918 8.77235 2.44389 9.36116C2.21691 9.47993 2.06624 9.70634 2.04435 9.96159C2.01498 10.304 2 10.6503 2 11C2 11.8228 2.08292 12.627 2.24109 13.4044C2.32367 13.8103 2.71966 14.0724 3.12556 13.9898C3.53146 13.9072 3.79356 13.5112 3.71098 13.1054C3.57271 12.4258 3.5 11.7218 3.5 11C3.5 10.834 3.50385 10.6689 3.51146 10.5048C3.86263 10.3385 4.22454 10.1911 4.59588 10.0639C4.53288 10.3642 4.49951 10.6776 4.49951 11C4.49951 13.2091 6.06651 15 7.99951 15C9.93251 15 11.4995 13.2091 11.4995 11C11.4995 10.6775 11.4661 10.3639 11.4031 10.0635C11.7748 10.1908 12.137 10.3383 12.4885 10.5048C12.4961 10.6689 12.5 10.834 12.5 11C12.5 11.7218 12.4273 12.4258 12.289 13.1054C12.2064 13.5112 12.4685 13.9072 12.8744 13.9898C13.2803 14.0724 13.6763 13.8103 13.7589 13.4044C13.9171 12.627 14 11.8228 14 11C14 10.6503 13.985 10.304 13.9557 9.96159C13.9338 9.70634 13.7831 9.47993 13.5561 9.36116C12.4259 8.76975 11.1898 8.35322 9.8845 8.14723L9.87353 8.14547C9.80295 8.1344 9.73217 8.12395 9.66119 8.11412C9.71329 8.03659 9.76011 7.95522 9.80112 7.87051C11.1372 7.6693 12.4018 7.24771 13.5561 6.64372C13.7831 6.52495 13.9338 6.29854 13.9557 6.0433C13.985 5.7009 14 5.35455 14 5.00488C14 4.91429 13.999 4.82391 13.997 4.73377C13.9878 4.31966 13.6446 3.9914 13.2305 4.00059C12.8164 4.00979 12.4882 4.35294 12.4974 4.76705C12.4991 4.8461 12.5 4.92538 12.5 5.00488C12.5 5.17092 12.4961 5.33601 12.4885 5.50007C11.6738 5.88595 10.8013 6.17008 9.8871 6.3359C9.81694 6.13654 9.71609 5.95164 9.59021 5.78689C10.0206 5.66886 10.4316 5.50411 10.8176 5.29842C10.9284 5.23933 11 5.12562 11 5C11 4.77978 10.9763 4.56511 10.9312 4.35839C11.4415 4.12281 11.9185 3.82769 12.3533 3.48199C12.6046 3.2822 12.7009 2.94464 12.5927 2.64236C12.4321 2.19322 12.227 1.76544 11.9828 1.36435Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
