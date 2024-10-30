package models

type File struct {
    ID             int    `json:"id"`
    CodeRepository string `json:"code_repository"`
    CodeFilePath   string `json:"code_file_path"`
    DocsRepository string `json:"docs_repository"`
    DocsFilePath   string `json:"docs_file_path"`
}

