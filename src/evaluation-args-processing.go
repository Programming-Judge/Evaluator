package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type evaluationArgs struct {
	ID          string // mandatory
	lang        string // mandatory
	problemID   string // mandatory
	timeLimit   int    // optional
	memoryLimit int    // optional
}

func validateEvaluationArgs() gin.HandlerFunc {
	validateID := func(next gin.HandlerFunc) gin.HandlerFunc {
		return func(ctx *gin.Context) {
			id := ctx.Query("ID")
			if id == "" {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, "Parameter ID is required")
				return
			}
			next(ctx)
		}
	}
	validateProblemID := func(next gin.HandlerFunc) gin.HandlerFunc {
		return func(ctx *gin.Context) {
			problemID, msg := ctx.Query("problemID"), ""
			if problemID == "" {
				msg = "Parameter problemID is required"
			}
			if msg != "" {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, msg)
				return
			}
			next(ctx)
		}
	}
	validateLang := func(next gin.HandlerFunc) gin.HandlerFunc {
		return func(ctx *gin.Context) {
			lang, msg := ctx.Query("lang"), ""
			if lang == "" {
				msg = "Parameter lang is required"
			} else if _, ok := langExtensionMap[lang]; !ok {
				msg = "Unknown language: " + lang
			}
			if msg != "" {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, msg)
				return
			}
			next(ctx)
		}
	}
	validateTimeLimit := func(next gin.HandlerFunc) gin.HandlerFunc {
		return func(ctx *gin.Context) {
			timeLimit := ctx.Query("timeLimit")
			if timeLimit != "" {
				n, err := strconv.Atoi(timeLimit)
				if err != nil || n < defaultTimeLimit {
					ctx.AbortWithStatusJSON(http.StatusBadRequest,
						"Time limit should be an integer >= "+strconv.FormatInt(defaultTimeLimit, 10))
					return
				}
			}
			next(ctx)
		}
	}
	validateMemoryLimit := func(next gin.HandlerFunc) gin.HandlerFunc {
		return func(ctx *gin.Context) {
			memoryLimit := ctx.Query("memoryLimit")
			if memoryLimit != "" {
				n, err := strconv.Atoi(memoryLimit)
				if err != nil || n < defaultMemoryLimit {
					ctx.AbortWithStatusJSON(http.StatusBadRequest,
						"Memory limit should be an integer >= "+strconv.FormatInt(defaultMemoryLimit, 10))
					return
				}
			}
			next(ctx)
		}
	}
	return chainMiddleWareWithDummy(validateID, validateLang, validateProblemID,
		validateTimeLimit, validateMemoryLimit)
}

func processEvaluationRequest(ctx *gin.Context) evaluationArgs {
	def := evaluationArgs{
		ID:          ctx.Query("ID"),
		lang:        ctx.Query("lang"),
		problemID:   ctx.Query("problemID"),
		timeLimit:   defaultTimeLimit,
		memoryLimit: defaultMemoryLimit,
	}
	if tl := ctx.Query("timeLimit"); tl != "" {
		def.timeLimit, _ = strconv.Atoi(tl)
	}
	if ml := ctx.Query("memoryLimit"); ml != "" {
		def.memoryLimit, _ = strconv.Atoi(ml)
	}
	return def
}
