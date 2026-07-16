package initiator

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func printPrettyRoutes(server *gin.Engine) {
	if gin.Mode()==gin.ReleaseMode{
		return
	}
	fmt.Printf("\n%sgin mode: %s%s\n", colorBold, gin.Mode(), colorReset)
	fmt.Printf("\n%s🗺️  Registered Routes:%s\n", colorBold+colorCyan, colorReset)
	fmt.Println("----------------------------------------------------------------")

	for _, route := range server.Routes() {
		// Pick an eye-catching color based on the HTTP Method
		var methodColor string
		switch route.Method {
		case "GET":
			methodColor = colorGreen
		case "POST":
			methodColor = colorCyan
		case "PUT", "PATCH":
			methodColor = colorYellow
		case "DELETE":
			methodColor = colorRed
		default:
			methodColor = colorReset
		}

		// Print formatted line: [METHOD]  /path  -> handler
		fmt.Printf("  %s%-7s%s %-35s %s(→ %s)%s\n",
			colorBold+methodColor, route.Method, colorReset,
			route.Path,
			"\033[2m", route.Handler, colorReset, // Made handler name dim for better hierarchy
		)
	}
	fmt.Println("----------------------------------------------------------------")
}
