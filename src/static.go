package main

/*
 * Inside the docker container,
 * the source, input and output
 * files will reside in:
 *
 * /home/{unp_user}/{bind_mnt_dir}
 *
 * and the executor script will
 * reside in:
 *
 * /home/{unp_user}
 *
 * The source, input and output
 * files will reside on the host system
 * at the location:
 *
 * os.Getwd()/../interface/{bind_mnt_dir}
 */
var bind_mnt_dir = "submissions"

var unp_user = "execution_user"

/*
 * The extension that a submission in
 * a particular language should have
 */
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

/*
 * The name of the image whose container
 * is to be made to process a submission
 * in the given language
 */
var lang_image_map = map[string]string{
	"cpp14":   "", // TODO
	"python3": "python3-eval",
	"pypy3":   "pypy3-eval",
}

const DEFAULT_TIME_LIMIT = 1
const SECONDS = "s"
const DEFAULT_MEMORY_LIMIT = 64 //64 MB
