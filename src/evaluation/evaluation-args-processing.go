package evaluation

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// A value of this type represents
// the parameters required by the
// evaluator to evaluate a submission
// Also see static.go
type evaluationArgs struct {
	id          string // mandatory
	lang        string // mandatory
	problemId   string // mandatory
	timeLimit   int // optional
	memoryLimit int // optional
}

// Functions to validate the 
// required parameters follow

func validateId(next gin.HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Query("id")
		if id == "" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, "Parameter id is required")
			return
		}
		next(ctx)
	}
}

func validateProblemID(next gin.HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		problemid, msg := ctx.Query("problemid"), ""
		if problemid == "" {
			msg = "Parameter problemid is required"
		}
		if msg != "" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, msg)
			return
		}
		next(ctx)
	}
}

func validateLang(next gin.HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		lang, msg := ctx.Query("lang"), ""
		if lang == "" {
			msg = "Parameter lang is required"
		} else if _, ok := lang_extension_map[lang]; !ok {
			msg = "Unknown language: " + lang
		}
		if msg != "" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, msg)
			return
		}
		next(ctx)
	}
}

func validateTimelimit(next gin.HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		timelimit := ctx.Query("timelimit")
		if timelimit != "" {
			n, err := strconv.Atoi(timelimit)
			if err != nil || n < DEFAULT_TIME_LIMIT {
				ctx.AbortWithStatusJSON(http.StatusBadRequest,
					"Time limit should be an integer >= "+strconv.FormatInt(DEFAULT_TIME_LIMIT, 10))
				return
			}
		}
		next(ctx)
	}
}

func validateMemoryLimit(next gin.HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		memorylimit := ctx.Query("memorylimit")
		if memorylimit != "" {
			n, err := strconv.Atoi(memorylimit)
			if err != nil || n < DEFAULT_MEMORY_LIMIT {
				ctx.AbortWithStatusJSON(http.StatusBadRequest,
					"Memory limit should be an integer >= "+strconv.FormatInt(DEFAULT_MEMORY_LIMIT, 10))
				return
			}
		}
		next(ctx)
	}
}

// A Middleware is a function that
// takes a gin.HandlerFunc and produces
// another 
type middleware func(gin.HandlerFunc) gin.HandlerFunc

func chainMiddleWareWithDummy(mws ...middleware) gin.HandlerFunc {
	chain := func(ctx *gin.Context) {}
	for j := len(mws) - 1; j >= 0; j-- {
		chain = mws[j](chain)
	}
	return chain
}

// Produces a chained validator 
// Middleware using all the validate...
// functions
func ValidateAll() gin.HandlerFunc {
	return chainMiddleWareWithDummy(validateId, validateLang, validateProblemID,
		validateTimelimit, validateMemoryLimit)
}

// Processes a HTTP request and
// produces an evaluationArgs output
// with default values for the
// missing optional parameters
// Invoked after validation 
func ProcessRequest(ctx *gin.Context) evaluationArgs {
	def := evaluationArgs{
		id:          ctx.Query("id"),
		lang:        ctx.Query("lang"),
		problemId:   ctx.Query("problemid"),
		timeLimit:   DEFAULT_TIME_LIMIT,
		memoryLimit: DEFAULT_MEMORY_LIMIT,
	}
	if tl := ctx.Query("timelimit"); tl != "" {
		def.timeLimit, _ = strconv.Atoi(tl)
	}
	if ml := ctx.Query("memorylimit"); ml != "" {
		def.memoryLimit, _ = strconv.Atoi(ml)
	}
	return def
}
