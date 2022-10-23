# backup

Issuing:

```bash
n3dr repositoriesV2 \
  --backup \
  --zip \
  --directory-prefix /tmp/some-dir \
  --directory-prefix-zip /tmp/some-dir/some-zip
```

will create a zip in: `/tmp/some-dir/some-zip`, e.g.:
`n3dr-backup-10-23-2022T12-30-58.zip`.
