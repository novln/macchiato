# Macchiato

[![Documentation][godoc-img]][godoc-url] ![License][license-img]

A mocha inspired reporter for [Ginkgo](https://onsi.github.io/ginkgo/)

[![Macchiato][macchiato-img]][macchiato-url]

## Example

![Example][example-img]

## Usage

In your Ginkgo Suite, you can import *Macchiato* and use it as follows:

```go
package books_test

import (
	"testing"

	"github.com/novln/macchiato"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

func TestBooks(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	macchiato.RunSpecs(t, "Books Suite")
}
```

## Under the hood

`macchiato.RunSpecs` will run **Ginkgo** with `RunSpecsWithCustomReporters` and the default Ginkgo's `Reporter`. However, the underlying `Stenographer` is replaced with a Macchiato implementation, which handle the rendering on the console output.

## License

This is Free Software, released under the [`MIT License`](LICENSE).

[macchiato-url]: https://github.com/novln/macchiato
[macchiato-img]: https://raw.githubusercontent.com/novln/macchiato/master/macchiato.jpg
[godoc-url]: https://godoc.org/github.com/novln/macchiato
[godoc-img]: https://godoc.org/github.com/novln/macchiato?status.svg
[license-img]: https://img.shields.io/badge/license-MIT-blue.svg
[example-img]: https://raw.githubusercontent.com/novln/macchiato/master/example.jpg
