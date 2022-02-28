package evaluation

//import "github.com/Programming-Judge/Evaluator/main"

const (

	// Default time and memory limits
	DEFAULT_TIME_LIMIT   = 1
	DEFAULT_MEMORY_LIMIT = 64

	// Unprivileged user to execute the
	// submitted code
	unp_user = "execution-user"

	//volume = main.Volume

	// Directory in the container where
	// the volume is mounted
	container_mnt_path = "/home/" + unp_user + "/" + mnt_dir
	mnt_dir            = "judge"

	// Testcases reside in
	// /home/{unp_user}/{mnt_dir}/{testcases_dir}
	testcases_dir = "testcases"

	// Submitted files reside in
	// /home/{unp_user}/{mnt_dir}/{submissions_dir}
	submissions_dir = "submissions"

	// Name of the execution script
	script = "evaluate.sh"
)

var (

	// The file extension for a 
	// submission in a given language
	lang_extension_map = map[string]string{
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

	// The image to be executed for a
	// submisssion in a given language
	lang_image_map = map[string]string{
		"c":       "c-eval",
		"cpp14":   "", // TODO
		"python3": "python3-eval",
		"pypy3":   "pypy3-eval",
		"java":    "java-eval",
	}
)
