I AM BUILDING A FILE STORAGE SYSTEM (LIKE S3)

SOME RULES:

- Limit max file size.
- Limit max part size.
- Limit number of parts.
- Never trust user-provided content type fully.
- Store files outside your source directory.
- Generate your own storage keys.
- Do not use the original filename as the storage path.
- Do not allow path traversal like ../../etc/passwd.
- Only allow downloads for ready objects.
- Expire signed URLs.
- Use constant-time signature comparison.