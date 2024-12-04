package main


func DebugLog(args ...string) []string {
	Dump := []string{"[DEBUG]"}
	return append(Dump, args...)
}

func InfoLog(args ...string) []string {
	Dump := []string{"[INFO]"}
	return append(Dump, args...) 
}

func ErrorLog(args ...string) []string {
	Dump := []string{"[ERROR]"}
	return append(Dump, args...) 
}
