ALTER TABLE files
ADD CONSTRAINT unique_code_docs_path
UNIQUE (code_repository, code_file_path, docs_repository, docs_file_path);

