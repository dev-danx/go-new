package contentReader

import "embed"

func ReadAsString(res embed.FS, fileName string) string {
	bytes, err := res.ReadFile(fileName)
	if err != nil {
		panic("Could not find File: " + fileName)
	}
	result := string(bytes)
	return result
}
