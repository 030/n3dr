# configRepository

## Create a repository

### Hosted yum

```bash
n3dr configRepository \
  --configRepoName some-repo \
  --configRepoType yum \
  -p <admin-pass> \
  -u <admin-user> \
  -n=<FQDN-without-http://-or-https>:<port-if-applicable> \
  --https=false
```
