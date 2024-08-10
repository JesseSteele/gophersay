# gophersay
## The talking gopher
Gopher talkback written in Go for Linux

- Using a `.go` script as source for the local compiler
- For alternate ways to build and compile the same `gophersay` package, see `git` repos:
  - [gophersay-tar](https://github.com/JesseSteele/gophersay-tar) (from local tarball)
  - [gophersay-git](https://github.com/JesseSteele/gophersay-git) (from `git` repository)

## Create the simple Linux install package for `gophersay`
This is a guide to create an installer package for the `gophersay` command on:
1. Arch (Manjaro, Black Arch, et al)

Working examples for each already resides in this repository

### Create and install the `gophersay` package directly from this repo

| **Arch** :$ (& Manjaro, Black Arch)

```console
git clone https://github.com/JesseSteele/gophersay.git
cd gophersay/arch
makepkg -si
```

## Detailed instructions per architecture
Instructions explain each in detail to create these packages from scratch...

- These instructions presume you can access [gophersay.go](https://github.com/JesseSteele/gophersay/blob/main/gophersay.go)

### Arch Linux Package (`gophersay-1.0.0-1-x86_64.pkg.tar.zst`)
*Arch package directory structure:*

| **`arch/`** :

```
arch/
├─ gophersay.go
└─ PKGBUILD
```

- Create directory: `arch`
- In `arch/` create file: `PKGBUILD`

| **`arch/PKGBUILD`** :

```
# Maintainer: Jesse Steele <codes@jessesteele.com>
pkgname=gophersay
pkgver=1.0.0
pkgrel=1
pkgdesc="Gopher talkback written in Go for Linux"
url="https://github.com/JesseSteele/gophersay"
arch=('x86_64')     # Go is newer and may not work on older systems, so not 'any'
license=('GPL')
depends=('go')      # Depends on the 'go' package to build the binary
replaces=('gophersay-tar' 'gophersay-git')
source=("$pkgname.go")
sha256sums=('26ebde65cd84a50cbca4146ca8fe78e88da7cb0dc2520768de0c2a1066cc5b4d')  # Hash for gophersay.go only

build() {
  # Get into the root, where our to-be-compiled files are
  cd $srcdir

  # Compile the Go binary
  go build -o "$pkgname" "$pkgname.go"
}

package() {
  install -Dm755 "$srcdir/$pkgname" "$pkgdir/usr/bin/$pkgname"
}
```

- Place file `gophersay.go` in the same directory as `PKGBUILD`
- Build package:
  - Navigate to directory `arch/`
  - Run this, then the package will be built, then installed with `pacman`:

| **Build & install Arch package** :$ (in one command)

```console
makepkg -si
```

- Use this to build and install in two steps:

| **Build, *then* install Arch package** :$ (first line produces the `.pkg.tar.zst` file for repos or manual install)

```console
makepkg
sudo pacman -U gophersay-1.0.0-1-x86_64.pkg.tar.zst
```

- Special notes about Arch:
  - We don't need to resolve any dependencies, we can omit the `-s` flag with `makepkg`
    - This package only needs `bash` as a dependency, which should already be installed merely to execute the script
      - `depends=('bash')` is redundant and only serves as an example in `PKGBUILD`
  - The name of the directory containing the package files does not matter
  - `PKGBUILD` is the instruction file, not a directory as might be expected with other package builders
  - `makepkg` must be run from the same directory containing `PKGBUILD`
  - The `.pkg.tar.zst` file will appear inside the containing directory

| **Remove Arch package** :$ (optional)

```console
sudo pacman -R gophersay
```
