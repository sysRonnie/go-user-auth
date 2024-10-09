// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package javascript

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func LandingJavascript() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<script>\n        console.log(\"Hello Landing Javascript!\");\n    \n        console.log(\"Hello!\");\n            document.getElementById(\"registerForm\").addEventListener(\"submit\",async function(e){\n\t\t\t\te.PreventDefault();\n\t\t\t\tconst username = document.getElementById(\"username\").value;\n\t\t\t\tconst email = document.getElementById(\"email\").value;\n\t\t\t\tconst password = document.getElementById(\"password\").value;\n\n\t\t\t\tconst response = await fetch(\"/api/v1/register\", {\n\t\t\t\t\tmethod: \"POST\",\n\t\t\t\t\theaders: {\n\t\t\t\t\t\t\"Content-Type\": \"application/json\"\n\t\t\t\t\t},\n\t\t\t\t\tbody: JSON.stringify({ username, email, passwd })\n\t\t\t\t});\n\n\t\t\t\tconst result = await response.json();\n\t\t\t\tconsole.log(result);\n\t\t\t});\n\n\t\t\tdocument.getElementById(\"loginForm\").addEventListener(\"submit\", async function(e) {\n\t\t\t\te.preventDefault();\n\t\t\t\tconst username = document.getElementById(\"loginUsername\").value;\n\t\t\t\tconst password = document.getElementById(\"loginPassword\").value;\n\t\t\t\tconst cookie = true; \n\n\t\t\t\ttry {\n\t\t\t\t\tconst response = await fetch(\"/api/v1/login\", {\n\t\t\t\t\t\tmethod: \"POST\",\n\t\t\t\t\t\theaders: {\n\t\t\t\t\t\t\t\"Content-Type\": \"application/json\"\n\t\t\t\t\t\t},\n\t\t\t\t\t\tbody: JSON.stringify({ username, password, cookie })\n\t\t\t\t\t});\n\n\t\t\t\t\tif (!response.ok) {\n\t\t\t\t\t\t// If the response is not OK, show the error message\n\t\t\t\t\t\tconst errorData = await response.json();\n\t\t\t\t\t\tdisplayErrorMessage(\"loginError\", errorData.error);\n\t\t\t\t\t} else {\n\t\t\t\t\t\tconst result = await response.json();\n\t\t\t\t\t\tconsole.log(result);\n\t\t\t\t\t\tdisplaySuccessMessage(\"loginSuccess\", \"Login successful!\");\n\t\t\t\t\t}\n\t\t\t\t} catch (error) {\n\t\t\t\t\tdisplayErrorMessage(\"loginError\", \"An unexpected error occurred.\");\n\t\t\t\t}\n\t\t\t});\n\n        function displayErrorMessage(elementId, message) {\n\t\t\t\tlet errorElement = document.getElementById(elementId);\n\t\t\t\tif (!errorElement) {\n\t\t\t\t\terrorElement = document.createElement(\"div\");\n\t\t\t\t\terrorElement.id = elementId;\n\t\t\t\t\terrorElement.style.color = \"red\";\n\t\t\t\t\tdocument.body.appendChild(errorElement);  // You can append it wherever appropriate\n\t\t\t\t}\n\t\t\t\terrorElement.textContent = message;\n\t\t\t}\n\n\t\t\tfunction displaySuccessMessage(elementId, message) {\n\t\t\t\tlet successElement = document.getElementById(elementId);\n\t\t\t\tif (!successElement) {\n\t\t\t\t\tsuccessElement = document.createElement(\"div\");\n\t\t\t\t\tsuccessElement.id = elementId;\n\t\t\t\t\tsuccessElement.style.color = \"green\";\n\t\t\t\t\tdocument.body.appendChild(successElement);  // You can append it wherever appropriate\n\t\t\t\t}\n\t\t\t\tsuccessElement.textContent = message;\n\t\t\t}\n\n    \n    </script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
