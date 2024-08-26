Name:           gophersay
Version:        1.0.0
Release:        1%{?dist}
Summary:        The talking gopher

License:        GPL
URL:            https://github.com/JesseSteele/gophersay
Source0:        gophersay-1.0.0.tar.xz

BuildArch:      noarch
BuildRequires:  go
Requires:       bash

%description
Gopher talkback written in Go for Linux

%prep
%setup -q

%build
go build -o gophersay gophersay.go

%install
mkdir -p %{buildroot}/usr/bin
install -D -m 0755 gophersay %{buildroot}/usr/bin/gophersay

%files
/usr/bin/gophersay

%changelog
* Thu Jan 01 1970 Jesse Steele <codes@jessesteele.com> - 1.0.0-1
- Something started, probably with v1.0.0-1