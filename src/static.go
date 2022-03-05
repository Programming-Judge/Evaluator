package main

const (
	defaultTimeLimit   = 1
	defaultMemoryLimit = 64
	unpUser            = "execution-user"
	volume             = "judge-submissions"
	evalMntPath   = "/home/" + unpUser + "/" +evalMntDir
	evalMntDir             = "judge"
	testCasesDir       = "testcases"
	submissionsDir     = "submissions"
	evaluationScript   = "evaluate.sh"
	iface              = "interface"
	opImage            = "op-bash"
	opMntPath          = "/mnt"
)

var (
	langExtensionMap = map[string]string{
		"cpp14":   "cpp",
		"cpp17":   "cpp",
		"cpp20":   "cpp",
		"python3": "py",
		"pypy3":   "py",
		"python2": "py",
		"pypy2":   "py",
		"c":       "c",
		"java":    "java",
	}

	langImageMap = map[string]string{
		"c":       "c-eval",
		"cpp14":   "", // TODO
		"python3": "python3-eval",
		"pypy3":   "pypy3-eval",
		"java":    "java-eval",
	}
)
