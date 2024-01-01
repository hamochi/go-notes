{
"CreatedAt": "2024-01-01",
"UpdatedAt": "2024-01-01",
"Title": "File Reading Writing - Go notes"
}
---
# File
1. [Reading](#reading)
  1. [Read entire file with os.ReadFile](#read-entire-file-with-osreadfile)
  2. [Read entire file with ioutil.ReadFile](#read-entire-file-with-ioutilreadfile)
  3. [Read file line by line with bufio.Scanner](#read-file-line-by-line-with-bufioscanner)
  4. [Read file in chunks with os.File.Read](#read-file-in-chunks-with-osfileread)
2. [Writing](#writing)
  1. [Write entire file with os.WriteFile](#write-entire-file-with-oswritefile)
  2. [Write entire file with ioutil.WriteFile](#write-entire-file-with-ioutilwritefile)
  3. [Write file line by line with bufio.Writer](#write-file-line-by-line-with-bufiowriter)
  4. [Write file in chunks with os.File.Write](#write-file-in-chunks-with-osfilewrite)
3. [Embeddings](#embeddings)
  1. [Embedding files in binary](#embedding-files-in-binary)
