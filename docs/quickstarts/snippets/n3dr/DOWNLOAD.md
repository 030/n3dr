# Download

Download the [latest N3DR binary](https://github.com/030/n3dr/releases/tag/7.2.5):

```bash
cd /tmp && \
curl -L https://github.com/030/n3dr/releases/download/7.2.5/n3dr-ubuntu-latest \
  -o n3dr-ubuntu-latest && \
curl -L https://github.com/030/n3dr/releases/download/7.2.5/\
n3dr-ubuntu-latest.sha512.txt \
  -o n3dr-ubuntu-latest.sha512.txt && \
sha512sum -c n3dr-ubuntu-latest.sha512.txt && \
chmod +x n3dr-ubuntu-latest && \
mv n3dr-ubuntu-latest n3dr && \
./n3dr --version
```
