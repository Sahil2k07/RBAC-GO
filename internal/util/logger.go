package util

import (
	"fmt"
	"net/url"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
)

// ANSI color codes with bold
const (
	reset         = "\033[0m"
	bold          = "\033[1m"
	brightRed     = "\033[91m" // error color
	violet        = "\033[35m"
	green         = "\033[32m"
	cyan          = "\033[36m"
	blue          = "\033[34m" // deep blue for Method key
	brightWhite   = "\033[97m"
	uniqueTeal    = "\033[38;5;37m" // vibrant teal, unique in the palette
	brightYellow  = "\033[93m"
	brightMagenta = "\033[95m"
	brightBlue    = "\033[94m"
	brightCyan    = "\033[96m" // stack trace key color
	maroon        = "\033[38;5;124m"
	gray          = "\033[90m"
)

// colorize for keys: bold + color, for values: normal color
func colorizeKey(text, colorCode string) string {
	return bold + colorCode + text + reset
}

func colorizeValue(text, colorCode string) string {
	return colorCode + text + reset
}

func getCleanStackTrace(skip, maxFrames int) string {
	pc := make([]uintptr, maxFrames+skip)
	n := runtime.Callers(skip, pc)
	frames := runtime.CallersFrames(pc[:n])

	var sb strings.Builder
	count := 0

	for {
		frame, more := frames.Next()
		sb.WriteString(fmt.Sprintf("  %s (%s:%d)\n", frame.Function, filepath.Base(frame.File), frame.Line))
		count++
		if count >= maxFrames {
			break
		}
		if !more {
			break
		}
	}

	return sb.String()
}

func generateLogMessage(c echo.Context, err error) {
	logger := log.Default()

	routeKey := colorizeKey("Route", violet)
	routeVal := colorizeValue(c.Path(), green)

	methodKey := colorizeKey("Method", blue)
	methodVal := colorizeValue(c.Request().Method, brightYellow)

	paramsKey := colorizeKey("Params", brightWhite)
	paramsVal := formatParams(c.ParamNames(), c.ParamValues())

	queryKey := colorizeKey("Query", uniqueTeal)
	queryVal := formatQueryParams(c.QueryParams())

	errKey := colorizeKey("Error", brightRed)
	errMsg := colorizeValue(err.Error(), brightRed)

	stackHeader := colorizeKey("Stack Trace:", brightCyan)
	stackTrace := getCleanStackTrace(4, 10)

	logger.Errorf("\n🐛 API Error\n"+
		"%s : %s\n"+
		"%s : %s\n"+
		"%s : %s\n"+
		"%s : %s\n"+
		"%s : %s\n"+
		"%s\n%s\n",
		routeKey, routeVal,
		methodKey, methodVal,
		paramsKey, paramsVal,
		queryKey, queryVal,
		errKey, errMsg,
		stackHeader,
		stackTrace,
	)
}

func formatParams(keys, values []string) string {
	if len(keys) == 0 {
		return colorizeValue("<none>", gray)
	}

	var pairs []string
	for i, k := range keys {
		keyColored := colorizeKey(k, brightWhite)
		valColored := colorizeValue("<missing>", gray)
		if i < len(values) {
			valColored = colorizeValue(values[i], gray)
		}
		pairs = append(pairs, fmt.Sprintf("%s=%s", keyColored, valColored))
	}

	return strings.Join(pairs, " ")
}

func formatQueryParams(values url.Values) string {
	if len(values) == 0 {
		return colorizeValue("<none>", gray)
	}

	var parts []string
	for key, vals := range values {
		keyColored := colorizeKey(key, uniqueTeal)
		for _, v := range vals {
			valColored := colorizeValue(v, gray)
			parts = append(parts, fmt.Sprintf("%s=%s", keyColored, valColored))
		}
	}

	return strings.Join(parts, " ")
}
