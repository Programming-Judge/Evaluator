package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var paramDefaultValues = map[string]string{
	"id":          "",
	"lang":        "",
	"timelimit":   "1s",
	"memorylimit": "64mb",
}

func validateId(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "Parameter id is required")
	}
}

func validateLang(ctx *gin.Context) {
	lang, msg := ctx.Query("lang"), ""
	if lang == "" {
		msg = "Parameter lang is required"
	} else if _, ok := lang_extension_map[lang]; !ok {
		msg = "Unknown language: " + lang
	}
	if msg != "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, msg)
	}
}

func validateTimelimit(ctx *gin.Context) {
	timelimit := ctx.Query("timelimit")
	if timelimit != "" {
		n, err := strconv.Atoi(timelimit[:len(timelimit)-1])
		if timelimit[len(timelimit)-1] != SECONDS || err != nil || n < DEFAULT_TIME_LIMIT {
			ctx.AbortWithStatusJSON(http.StatusBadRequest,
				"Time limit format: [integer >= 1]s")
		}
	}
}

func validateMemoryLimit(ctx *gin.Context) {
	memorylimit := ctx.Query("memorylimit")
	if memorylimit != "" {
		n, err := strconv.Atoi(memorylimit[:len(memorylimit)-2])
		if memorylimit[len(memorylimit)-2:] != MEGABYTES || err != nil || n < DEFAULT_MEMORY_LIMIT {
			ctx.AbortWithStatusJSON(http.StatusBadRequest,
				"Memory limit format: [integer >= 64]mb")
		}
	}
}

func processRequest(ctx *gin.Context) map[string]string {
	data := make(map[string]string)
	for prop, def_val := range paramDefaultValues {
		g := ctx.Query(prop)
		data[prop] = def_val
		if g != "" {
			data[prop] = g
		}
	}
	return data
}
