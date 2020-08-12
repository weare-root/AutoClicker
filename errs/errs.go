package errs

import "log"

// ErrorCheck panics when an error occurs
func ErrorCheck(e error) {
	if e != nil {
		// panic for any errors.
		log.Panic(e)
	}
}

// LogOnError logs the error and the message when an error occurs
func LogOnError(e error, msg string) {
	if e != nil {
		log.Println(msg)
		log.Println(e)
	}
}
