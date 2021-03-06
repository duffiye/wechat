package common

import (
	"fmt"
	"testing"
)

func TestAESDecryptWithGCM(t *testing.T) {

	var data string = "jfgt0T+tCPQBCNLLj9VghaeFIJo9qblssBDLBEYdDDQw3xwuseyFrwfbbcf15MBUhsAdQ9GBjX+r80tcA9P5u4A5NNViI67nLpWXOowf9VlP/fSXBiEeq0v08w0ejzPaeZLY8mW/trvdFfTghmaItmg101KSTtbKAN7Lx5+h8ziOi7Tw66IWrZ/cvHfkT5nox3PWivs+mVGXozmH8XF5De56qR0ppLw9+DEUmrTQvgLhHnfrkupdDvxk1z4oaZYfKJeWOK7Rf8DbOnSk0dH+3G66K1lm17ab3IrPBD27hccMcn+Ro6S1AxHFrqxThPgcaaN5z1zSya1i7Vc/eeqqqIZWMuaT4NwvMZgt5cudLDi+SQrebQQ+8OAmks8+heMBOXn6MWpn4jLkWxjkXPoIyWvEgWoAJPsDL6tXsmnnYwqmJwPK1S16h10RcK2FcG3rsIr5H8rfoGy0MTvqFOoWq/mOqPqFVrCQW8rnSp8G+QaQxRjG4Nege3Mqp4HDLS5JsZEqJfSl9+sqdDVcpF6MpcCIow7rURoRAPmleA81VWKo+1jDWjVGqCmNwkh5ZOTqwzzf2eUl6mJR2+aNQzlLIi66SIaOIJitJzLZWiXMTvRmj0ZQivZmhffpSf5+Aha42bEf8HE4ksfutgswlsMo+6AZBQppov9YUFeiCdrxsgHzLYweks/SKCYlBkeMxebX+wQiucXMtc+ZUyAuqgW05ZIqkF3GY3DxOPGksvXXa27bDVbw2QMiAqGVdrbkjAtneX+V+wxWC9GTphSy5WYYwtA4m7dMvExITYLta4zqgUZrgl/Bx/2Igt6nxnyoq2f0ilWbqjWJgm2yRaQayG69WuKEGjtmdOWlxv7whEP9gzQeA8d8BjPml4U1ry5BnnVK6qpWyj1Vc6XpaByrsagytf0aS4z7J5BuNICLf9Ynu9VqjSeaOvLmaJtq9qX/xAGuhqqJqxpH50WqB85GSkuUiEHfkFh8mlx7Xo8GXwidgOAM4gTvFEIcczOhyqRg4jCBe+tPbHsmSUvI4crwhlnE5kAgeRgCLIDFkGLxhukIb5JFUe3PZQfi/Ej3OY++HJjl2njOj0quU1jTDk44ty/mO5HrnsvrIf0MAGumT7aM/ESizK5UjI0pM6+t7GHaxw8+Fkd89jgEd+SfkUB4teybbg9YpmroKJOQLl7P8q/cnMsweI//nl7bzKVdSC6sXYTDi0fhLQE5lxpWEikJ+O0wXQg5j7OfGndRJipx6mW9Lc29Ck3zh+D8949ew5K/PZVnQhiiLNiRjI3g62CeXYULo64wFFNZi3GA3BJrYzENVtAyrQFlldRMYYYGQwWO9IeWQLCmNpMEHw+CzqO5JgClleAdwBXXytWjHfn/N3nx5VmVxFDkgtklw0Jv/f6twqKkIRjmr6CJEP7x7Nh6wq8uzKGEV4G4ZUo8zABgvOTBED0tkvndCZ7i8RAQ7ZMu0cPwJ6rINqSGyhWlcCpb9fLlBfv8g9c1pBUCSKja9M1v3yMMptZHXmrJ/3J90Cr9n9Jeu4tvmVTLIFZDr9jznpTtOrmWZlvipUGBJjbM8EPVWe9ii0QXcNElqZiNOlR0qsOcNKGbtyskHOy9Vq2REcsWsjCPcvL1C+5TEV1d5z8z0KEPKoe2ZmIBC3Eq7sqAmHwpjDb5Ia6nl5Zjm5dx250f0iAfXpyFiiLc/5BosfNaLebX6XBhwmdO7cqp/B3BDjZHrfdgXxCJQOmgNK27KElcNGJR5TAzI6G4/Dy0XrWm5bLuT/suxxbhRwEz4zMAmFHVhtnAcp+vW1rj60Osja8148bJTJZaIAAdUFdGzh0tcTg6p/Ngq/9Pry6R7nYQdjfZ4udlFniddPsY260Xcntvq5TitcMYGw=="
	var key string = "qqwweerrttQQWWEERRTT123456789XFT"
	var nonceStr string = "484f062b9c76"
	var additionalData string = "certificate"
	text := AESDecryptWithGCM(data, key, nonceStr, additionalData)

	fmt.Println(text)
}
