CREATE TABLE files (
    id SERIAL PRIMARY KEY,
    code_repository VARCHAR(255) NOT NULL,
    code_file_path VARCHAR(255) NOT NULL,
    docs_repository VARCHAR(255) NOT NULL,
    docs_file_path VARCHAR(255) NOT NULL
);
