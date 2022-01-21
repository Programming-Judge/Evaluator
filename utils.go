package main

// Folder that will be bind-mounted
// onto the docker container
var bind_mnt_dir = "submissions"

// The extension that a submission in
// a particular language should have
var lang_extension_map = map[string]string{
	"cpp14":   "cpp",
	"cpp17":   "cpp",
	"cpp20":   "cpp",
	"python3": "py",
	"pypy3":   "py",
	"python2": "py",
	"pypy2":   "py",
	"c":       "c",
}

// The name of the image whose container
// is to be made to process a submission
// in the given language
var lang_image_map = map[string]string{
	"cpp14": "", // TODO
	"python3": "python3-eval",
}

func getPaths(id string, lang string) (string, string, string) {

	// We expect the request to contain parameters
	// id and lang.
	// id -> submission ID
	// lang -> language of the submission
	// The corresponding source, input and output
	// should be placed in the "bind_mnt_dir"
	// directory with the following naming convention:
	// source file = id + "-main." + extension
	// input file := id + "-input.txt"
	// output file := id + "-output.txt"
	extension := lang_extension_map[lang]
	common := bind_mnt_dir + "/" + id
	codeFile := common + "-main." + extension
	inputFile := common + "-input.txt"
	outputFile := common + "-output.txt"
	return codeFile, inputFile, outputFile
}
