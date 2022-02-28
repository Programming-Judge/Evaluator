package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/Programming-Judge/Evaluator/evaluation/evaluation"
)

// The evaluation route
// The context is updated with a JSON
// with two fields:
//
// verdict: contains the verdict if
// 'evaluation' was successful, ie
// TLE/MLE/WA/AC; otherwise "Evaluation failed"
//
// internal: data about the error
// that occurred during evaluation, 
// if there was one, otherwise ""
// (mainly for debugging purposes)
func eval(ctx *gin.Context) {
	verdict, _, err := evaluation.evaluate(evaluation.ProcessRequest(ctx))
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"verdict": "Evaluation failed", "internal": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"verdict": verdict, "internal": ""})
}


func uploadSubmission(ctx *gin.Context) {
	
}