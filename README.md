# gophersay
## The talking gopher
Gopher talkback written in Go for Linux

```
 ------------------------------------------------- 
< I am a talking gopher, written in Go for Linux. >
< (See what I did there?)                         >
 ------------------------------------------------- 
                                                \ \
                                                 \\
                                                  \ (o)—————(o)
                                                     / Q   Q \
                                                    |   o@o   |
                                                    |    W    |
                                                   c|         |ɔ
                                                    |         |
                                                    |         |
                                                     \_______/
                                                     C       Ɔ
```

- Using a `.go` script as source for the local compiler
- For alternate ways to build and compile the same `gophersay` package, see `git` repos:
  - [gophersay-tar](https://github.com/JesseSteele/gophersay-tar) (from local tarball)
  - [gophersay-git](https://github.com/JesseSteele/gophersay-git) (from `git` repository)

## Create the simple Linux install package for `gophersay`
This is a guide to create an installer package for the `gophersay` command on:
1. Arch (Manjaro, Black Arch, et al)
2. Debian (Ubuntu, Kali, Mint, et al)

Working examples for each already resides in this repository

### Create and install the `gophersay` package directly from this repo

| **Arch** :$ (& Manjaro, Black Arch)

```console
git clone https://github.com/JesseSteele/gophersay.git
cd gophersay/arch
makepkg -si
```

| **Debian** :$ (& Ubuntu, Kali, Mint)

```console
git clone https://github.com/JesseSteele/gophersay.git
cd gophersay/deb/build
sudo dpkg-buildpackage -us -uc
cd debian
dpkg-deb --build gophersay
sudo dpkg -i gophersay.deb
```

## Detailed instructions per architecture
Instructions explain each in detail to create these packages from scratch...

- These instructions presume you can access [gophersay.go](https://github.com/JesseSteele/gophersay/blob/main/gophersay.go)

### I. Arch Linux Package (`gophersay-1.0.0-1-x86_64.pkg.tar.zst`)
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
  cd "$srcdir"

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

### II. Debian Package (`gophersay.deb`)
*Debian "**maintainer**" build directory structure:*

| **`deb/`** :

```
deb/
└─ build/
   ├─ debian/
   │  ├─ compat
   │  ├─ control
   │  ├─ copyright
   │  ├─ changelog
   │  ├─ install
   │  └─ rules
   ├─ go.mod
   └─ gophersay.go
```

- Create directories: `deb/build/debian`
- In `debian/` create file: `control`

| **`deb/build/debian/control`** :

```
Source: gophersay
Section: games
Priority: optional
Maintainer: Jesse Steele <codes@jessesteele.com>
Homepage: https://github.com/JesseSteele/gophersay
Build-Depends: debhelper (>= 10), golang-go
Standards-Version: 3.9.6

Package: gophersay
Version: 1.0.0
Architecture: all
Depends: bash (>= 4.0)
Description: Gopher talkback written in Go for Linux
```

- In `debian/` create file: `compat`

| **`deb/build/debian/compat`** : (`debhelper` minimum version)

```
10
```

- In `debian/` create file: `changelog`

| **`deb/build/debian/changelog`** : (optional, for listing changes)

```
gophersay (1.0-1) stable; urgency=low

  * First release

 -- Jesse Steele <codes@jessesteele.com>  Thu, 1 Jan 1970 00:00:00 +0000
```

- In `debian/` create file: `copyright`

| **`deb/build/debian/copyright`** : (optional, may be legally wise)

```
Format: http://www.debian.org/doc/packaging-manuals/copyright-format/1.0/
Upstream-Name: gophersay
Source: https://github.com/JesseSteele/gophersay

Files: *
Copyright: 2024, Jesse Steele <codes@jessesteele.com>
License: GPL-3+
```

- In `debian/` create file: `rules`
  - Make it executable with :$ `chmod +x debian/rules`

| **`deb/build/debian/rules`** : (build compiler)

```
#!/usr/bin/make -f

%:
	dh $@

override_dh_auto_build:
	go build -o gophersay gophersay.go

override_dh_auto_install:
	install -D -m 0755 gophersay $(DESTDIR)/usr/bin/gophersay
```

- In `debian/` create file: `install`

| **`deb/build/debian/install`** : (places files in the `.deb` directory structure)

```
gophersay /usr/bin
```

- Place Go script `gophersay.go` at `deb/build/gophersay.go`
- Install the `dpkg-dev`, `debhelper` & `golang-go` packages

| **Install Debian `dpkg-dev` package** :$

```console
sudo apt-get update
sudo apt-get install dpkg-dev debhelper golang-go
```

- Prepare package builder:
  - Navigate to directory `deb/build/`
  - Run this, then the package builder & repo packages will be created:

| **Prepare the Debian package builder** :$

```console
sudo dpkg-buildpackage -us -uc  # Create the package builder
```

- Note what just happened
  - Everything just done to this point is called "**maintainer**" work in the Debian world
  - Basic repo packages *and also* the package `DEBIAN/` builder structure were greated
  - At this point, one could navigate up one directory to `deb/` and run `sudo dpkg -i gophersay_1.0-1_all.deb` and the package would be installed, *but we won't do this*
  - The command has also been created at `/usr/bin/gophersay`
    - Once installed with `sudo dpkg -i` (later) this can be removed the standard way with `sudo apt-get remove gophersay`
  - This is the new, just-created directory structure for the standard Debian package builder:

| **`deb/build/debian/`** :

```
deb/build/debian/
          └─ gophersay/
             ├─ DEBIAN/
             │  ├─ control
             │  └─ md5sums
             └─ usr/
                └─ bin/
                   └─ gophersay
```

- Build package:
  - Navigate to directory `deb/build/debian/`
    - :$ `cd debian`
  - Run this, then the package will be built, then installed:

| **Build, *then* install Debian package** :$

```console
dpkg-deb --build gophersay  # Create the .deb package
sudo dpkg -i gophersay.deb  # Install the package
```

- Special notes about Debian
  - The `deb/build/` directory can be anything, but we want it for housekeeping...
    - `dpkg-buildpackage` will create a laundry list of files as peers to this directory in `deb/`
  - The `debian/control` file builds `DEBIAN/control`, but uses different fields
    - Fields after the empty line in this `debian/control` example will not be recognized by the `dpkg-buildpackage` builder, but will supply information for `DEBIAN/control` in the final `.deb` package
  - The `rules` script will compile `gophersay` from `gophersay.go`
  - The `install` script will place the compiled `gophersay` binary at `usr/bin/gophersay` inside the package
    - This is why we don't need to place the binary at `usr/bin/gophersay` manually
  - Note `usr/local/bin/` won't work for CLI command files because Debian packages expect binary commands to go in `/usr/bin/`
    - Debian can install *directories*—but *cannot install any **file***—under `usr/local/`
    - Trying to install a file will return an [error from the package manager](https://unix.stackexchange.com/questions/409800/) since it expects directories, but only finds a file
  - The standard package build files (for `dpkg-deb --build`) will appear at `deb/build/debian/gophersay/DEBIAN/`
    - So from `deb/build/debian/` one could run `dpkg-deb --build gophersay` to create the `.deb` package at `deb/build/debian/gophersay.deb`
  - The standard package installer will appear at `deb/gophersay_1.0-1_all.deb`

| **Remove Debian package** :$ (optional)

```console
sudo apt-get remove gophersay
```
