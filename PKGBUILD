# Maintainer: pyndys
pkgname=ggf
pkgver=0.2.0
pkgrel=1
pkgdesc="Great Go Fetch (ggf) - fast system info fetch utility for Linux"
arch=('x86_64' 'aarch64')
url="https://github.com/pyndys/ggf"
license=('MIT')
makedepends=('go')
source=("$pkgname-$pkgver.tar.gz::https://github.com/pyndys/$pkgname/archive/refs/tags/v$pkgver.tar.gz")
b2sums=('ab0fc9fd2a2bcdc5e440dbff0fc28662d93264d7b4977104265063892b67eb365bc81dba5307c60ddf854d8f8b51574d642dd56f8a3e85a8565fb9534421d5f3')

build() {
    cd "$pkgname-$pkgver"
    export CGO_ENABLED=0
    go build -v \
        -ldflags="-s -w" \
        -o "$pkgname" \
        .
}

package() {
    cd "$pkgname-$pkgver"
    install -Dm755 "$pkgname" "$pkgdir/usr/bin/$pkgname"
    install -Dm644 LICENSE "$pkgdir/usr/share/licenses/$pkgname/LICENSE"
}
