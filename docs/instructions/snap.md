# Snap

## Installation

```bash
snap install n3dr
```

## Configuration

```bash
~/snap/n3dr/current/.n3dr/config.yml
```

## Usage

Download artifacts:

```bash
n3dr repositoriesV2 --backup --directory-prefix /tmp/some-dir
```

Check the downloaded artifacts:

```bash
sudo ls /tmp/snap-private-tmp/snap.n3dr/tmp/some-dir
sudo cp -r sudo ls /tmp/snap-private-tmp/snap.n3dr/tmp/some-dir /home/${USER}/n3dr-backup
sudo chown $USER:$USER -R /home/${USER}/n3dr-backup
```

Note: if the snap package is used to upload artifacts, one has to ensure that
the folder resides in the /home/$USER folder. Otherwise a:
`lstat <repository-name>: no such file or directory` issue could occur.
