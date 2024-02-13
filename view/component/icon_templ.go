// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package component

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

var pathsByName = map[string][]string{
	"atom": {
		"M12 11a1 1 0 0 1 1 1a1 1 0 0 1-1 1a1 1 0 0 1-1-1a1 1 0 0 1 1-1M4.22 4.22C5.65 2.79 8.75 3.43 12 5.56c3.25-2.13 6.35-2.77 7.78-1.34c1.43 1.43.79 4.53-1.34 7.78c2.13 3.25 2.77 6.35 1.34 7.78c-1.43 1.43-4.53.79-7.78-1.34c-3.25 2.13-6.35 2.77-7.78 1.34c-1.43-1.43-.79-4.53 1.34-7.78c-2.13-3.25-2.77-6.35-1.34-7.78m11.32 4.24c.61.62 1.17 1.25 1.69 1.88c1.38-2.13 1.88-3.96 1.13-4.7c-.74-.75-2.57-.25-4.7 1.13c.63.52 1.26 1.08 1.88 1.69m-7.08 7.08c-.61-.62-1.17-1.25-1.69-1.88c-1.38 2.13-1.88 3.96-1.13 4.7c.74.75 2.57.25 4.7-1.13c-.63-.52-1.26-1.08-1.88-1.69m-2.82-9.9c-.75.74-.25 2.57 1.13 4.7c.52-.63 1.08-1.26 1.69-1.88c.62-.61 1.25-1.17 1.88-1.69c-2.13-1.38-3.96-1.88-4.7-1.13m4.24 8.48c.7.7 1.42 1.34 2.12 1.91c.7-.57 1.42-1.21 2.12-1.91c.7-.7 1.34-1.42 1.91-2.12c-.57-.7-1.21-1.42-1.91-2.12c-.7-.7-1.42-1.34-2.12-1.91c-.7.57-1.42 1.21-2.12 1.91c-.7.7-1.34 1.42-1.91 2.12c.57.7 1.21 1.42 1.91 2.12m8.48 4.24c.75-.74.25-2.57-1.13-4.7c-.52.63-1.08 1.26-1.69 1.88c-.62.61-1.25 1.17-1.88 1.69c2.13 1.38 3.96 1.88 4.7 1.13",
	},
	"github": {
		"M5.88401 18.6533C5.58404 18.4526 5.32587 18.1975 5.0239 17.8369C4.91473 17.7065 4.47283 17.1524 4.55811 17.2583C4.09533 16.6833 3.80296 16.417 3.50156 16.3089C2.9817 16.1225 2.7114 15.5499 2.89784 15.0301C3.08428 14.5102 3.65685 14.2399 4.17672 14.4263C4.92936 14.6963 5.43847 15.1611 6.12425 16.0143C6.03025 15.8974 6.46364 16.441 6.55731 16.5529C6.74784 16.7804 6.88732 16.9182 6.99629 16.9911C7.20118 17.1283 7.58451 17.1874 8.14709 17.1311C8.17065 16.7489 8.24136 16.3783 8.34919 16.0358C5.38097 15.3104 3.70116 13.3952 3.70116 9.63971C3.70116 8.40085 4.0704 7.28393 4.75917 6.3478C4.5415 5.45392 4.57433 4.37284 5.06092 3.15636C5.1725 2.87739 5.40361 2.66338 5.69031 2.57352C5.77242 2.54973 5.81791 2.53915 5.89878 2.52673C6.70167 2.40343 7.83573 2.69705 9.31449 3.62336C10.181 3.41879 11.0885 3.315 12.0012 3.315C12.9129 3.315 13.8196 3.4186 14.6854 3.62277C16.1619 2.69 17.2986 2.39649 18.1072 2.52651C18.1919 2.54013 18.2645 2.55783 18.3249 2.57766C18.6059 2.66991 18.8316 2.88179 18.9414 3.15636C19.4279 4.37256 19.4608 5.45344 19.2433 6.3472C19.9342 7.28337 20.3012 8.39208 20.3012 9.63971C20.3012 13.3968 18.627 15.3048 15.6588 16.032C15.7837 16.447 15.8496 16.9105 15.8496 17.4121C15.8496 18.0765 15.8471 18.711 15.8424 19.4225C15.8412 19.6127 15.8397 19.8159 15.8375 20.1281C16.2129 20.2109 16.5229 20.5077 16.6031 20.9089C16.7114 21.4504 16.3602 21.9773 15.8186 22.0856C14.6794 22.3134 13.8353 21.5538 13.8353 20.5611C13.8353 20.4708 13.836 20.3417 13.8375 20.1145C13.8398 19.8015 13.8412 19.599 13.8425 19.4094C13.8471 18.7019 13.8496 18.0716 13.8496 17.4121C13.8496 16.7148 13.6664 16.2602 13.4237 16.051C12.7627 15.4812 13.0977 14.3973 13.965 14.2999C16.9314 13.9666 18.3012 12.8177 18.3012 9.63971C18.3012 8.68508 17.9893 7.89571 17.3881 7.23559C17.1301 6.95233 17.0567 6.54659 17.199 6.19087C17.3647 5.77663 17.4354 5.23384 17.2941 4.57702L17.2847 4.57968C16.7928 4.71886 16.1744 5.0198 15.4261 5.5285C15.182 5.69438 14.8772 5.74401 14.5932 5.66413C13.7729 5.43343 12.8913 5.315 12.0012 5.315C11.111 5.315 10.2294 5.43343 9.40916 5.66413C9.12662 5.74359 8.82344 5.69492 8.57997 5.53101C7.8274 5.02439 7.2056 4.72379 6.71079 4.58376C6.56735 5.23696 6.63814 5.77782 6.80336 6.19087C6.94565 6.54659 6.87219 6.95233 6.61423 7.23559C6.01715 7.8912 5.70116 8.69376 5.70116 9.63971C5.70116 12.8116 7.07225 13.9683 10.023 14.2999C10.8883 14.3971 11.2246 15.4769 10.5675 16.0482C10.3751 16.2156 10.1384 16.7802 10.1384 17.4121V20.5611C10.1384 21.5474 9.30356 22.2869 8.17878 22.09C7.63476 21.9948 7.27093 21.4766 7.36613 20.9326C7.43827 20.5204 7.75331 20.2116 8.13841 20.1276V19.1381C7.22829 19.1994 6.47656 19.0498 5.88401 18.6533Z",
	},
	"json": {
		"M5 3h2v2H5v5a2 2 0 0 1-2 2a2 2 0 0 1 2 2v5h2v2H5c-1.07-.27-2-.9-2-2v-4a2 2 0 0 0-2-2H0v-2h1a2 2 0 0 0 2-2V5a2 2 0 0 1 2-2m14 0a2 2 0 0 1 2 2v4a2 2 0 0 0 2 2h1v2h-1a2 2 0 0 0-2 2v4a2 2 0 0 1-2 2h-2v-2h2v-5a2 2 0 0 1 2-2a2 2 0 0 1-2-2V5h-2V3zm-7 12a1 1 0 0 1 1 1a1 1 0 0 1-1 1a1 1 0 0 1-1-1a1 1 0 0 1 1-1m-4 0a1 1 0 0 1 1 1a1 1 0 0 1-1 1a1 1 0 0 1-1-1a1 1 0 0 1 1-1m8 0a1 1 0 0 1 1 1a1 1 0 0 1-1 1a1 1 0 0 1-1-1a1 1 0 0 1 1-1",
	},
	"linkedin": {
		"M12.001 9.55005C12.9181 8.61327 14.1121 8 15.501 8C18.5385 8 21.001 10.4624 21.001 13.5V21H19.001V13.5C19.001 11.567 17.434 10 15.501 10C13.568 10 12.001 11.567 12.001 13.5V21H10.001V8.5H12.001V9.55005ZM5.00098 6.5C4.17255 6.5 3.50098 5.82843 3.50098 5C3.50098 4.17157 4.17255 3.5 5.00098 3.5C5.8294 3.5 6.50098 4.17157 6.50098 5C6.50098 5.82843 5.8294 6.5 5.00098 6.5ZM4.00098 8.5H6.00098V21H4.00098V8.5Z",
	},
	"mastodon": {
		"M3.019 12.0075C2.98744 10.7478 3.00692 9.5598 3.00692 8.56644C3.00692 4.22767 5.84954 2.95597 5.84954 2.95597C7.28286 2.29767 9.74238 2.0209 12.2993 2H12.3621C14.919 2.0209 17.3801 2.29767 18.8134 2.95597C18.8134 2.95597 21.656 4.22767 21.656 8.56644C21.656 8.56644 21.6916 11.7674 21.2596 13.9898C20.9852 15.4007 18.8034 16.9446 16.2974 17.2438C14.9906 17.3999 13.7042 17.5431 12.3322 17.4802C10.0885 17.3775 8.31815 16.9446 8.31815 16.9446C8.31815 17.1631 8.33166 17.3711 8.35853 17.5655C8.44182 18.1978 8.65659 18.6604 8.96296 19C9.72944 19.8497 11.0692 19.9301 12.3577 19.9743C14.178 20.0366 15.7986 19.5254 15.7986 19.5254L15.8735 21.1712C15.8735 21.1712 14.6003 21.8548 12.3322 21.9805C11.0815 22.0493 9.52858 21.9491 7.71969 21.4704C6.18802 21.065 5.15153 20.1804 4.45091 19C3.35714 17.1573 3.08191 14.5938 3.019 12.0075ZM6.31815 16.9446V14.3967L8.79316 15.0018C8.8405 15.0134 8.95098 15.0383 9.11692 15.0723C9.40521 15.1313 9.73416 15.1908 10.0959 15.2467C10.8485 15.3628 11.6341 15.4462 12.4237 15.4823C13.4425 15.529 14.3249 15.4652 16.0603 15.2579C17.7233 15.0594 19.208 14.0622 19.2963 13.6082C19.3783 13.1861 19.4472 12.6858 19.5021 12.1261C19.5714 11.4205 19.6155 10.6558 19.6388 9.88068C19.654 9.37026 19.6582 8.93648 19.6564 8.62452L19.656 8.56644C19.656 7.1368 19.2873 6.12756 18.6928 5.40793C18.5008 5.17553 18.3004 4.99408 18.1087 4.85958C18.0183 4.79617 17.9737 4.77136 17.9787 4.77345C16.9662 4.30844 14.8859 4.02069 12.3621 3.99993H12.3156C9.77596 4.02069 7.6969 4.30836 6.66627 4.78161C6.68919 4.77136 6.64459 4.79617 6.55423 4.85958C6.36257 4.99408 6.16214 5.17553 5.97016 5.40793C5.37568 6.12756 5.00692 7.1368 5.00692 8.56644C5.00692 8.7976 5.00628 8.96339 5.00392 9.44137C4.9981 10.6238 5.00004 11.2256 5.01841 11.9589C5.07185 14.156 5.2822 15.7941 5.71797 17C5.93023 17.5874 6.19005 18.0709 6.49741 18.4507C6.37791 18.0162 6.31815 17.5142 6.31815 16.9446ZM8.08576 6.37135C8.71735 6.37135 9.22924 6.88324 9.22924 7.51482C9.22924 8.14626 8.71735 8.6583 8.08576 8.6583C7.45432 8.6583 6.94229 8.14626 6.94229 7.51482C6.94229 6.88324 7.45432 6.37135 8.08576 6.37135Z",
	},
	"rss": {
		"M3 17a4 4 0 0 1 4 4H3zm0-7c6.075 0 11 4.925 11 11h-2a9 9 0 0 0-9-9zm0-7c9.941 0 18 8.059 18 18h-2c0-8.837-7.163-16-16-16z",
	},
}

func Icon(name string) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg class=\"icon\" aria-hidden=\"true\" focusable=\"false\" xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 24 24\" fill-rule=\"evenodd\" clip-rule=\"evenodd\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, pathName := range pathsByName[name] {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<path d=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(pathName))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"></path>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
