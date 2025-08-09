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

// ANSI color codes
const (
	reset         = "\033[0m"
	red           = "\033[31m"
	violet        = "\033[35m"
	green         = "\033[32m"
	cyan          = "\033[36m"
	brightWhite   = "\033[97m"
	seagreen      = "\033[38;5;49m"
	brightYellow  = "\033[93m"
	brightMagenta = "\033[95m"
	brightBlue    = "\033[94m"
	gray          = "\033[90m"
)

// colorize helper
func colorize(text, colorCode string) string {
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

	routeKey := colorize("Route", violet)
	routeVal := colorize(c.Path(), green)

	methodKey := colorize("Method", cyan)
	methodVal := colorize(c.Request().Method, brightYellow)

	paramsKey := colorize("Params", brightWhite)
	paramsVal := formatParams(c.ParamNames(), c.ParamValues())

	queryKey := colorize("Query", brightMagenta)
	queryVal := formatQueryParams(c.QueryParams())

	errKey := colorize("Error", red)
	errMsg := colorize(err.Error(), red)

	stackHeader := colorize("Stack Trace:", brightBlue)
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
		return colorize("<none>", gray)
	}

	var pairs []string
	for i, k := range keys {
		keyColored := colorize(k, brightWhite)
		valColored := colorize("<missing>", gray)
		if i < len(values) {
			valColored = colorize(values[i], gray)
		}
		pairs = append(pairs, fmt.Sprintf("%s=%s", keyColored, valColored))
	}

	return strings.Join(pairs, " ")
}

func formatQueryParams(values url.Values) string {
	if len(values) == 0 {
		return colorize("<none>", gray)
	}

	var parts []string
	for key, vals := range values {
		keyColored := colorize(key, brightMagenta)
		for _, v := range vals {
			valColored := colorize(v, gray)
			parts = append(parts, fmt.Sprintf("%s=%s", keyColored, valColored))
		}
	}

	return strings.Join(parts, " ")
}
