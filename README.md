# tree-sitter-clojure

A tree-sitter grammar for Clojure and ClojureScript, packaged with
WebAssembly support for portable use without requiring a C compiler.

This is a fork of [sogaiu/tree-sitter-clojure](https://github.com/sogaiu/tree-sitter-clojure/)
that adds WASM builds, multi-language bindings (Node.js, Rust, Go, Python, C, Swift),
and CI/CD workflows.

## Installation

```bash
npm install @yogthos/tree-sitter-clojure
```

### Usage with web-tree-sitter (WASM)

```js
const { Parser, Language } = require("web-tree-sitter");

await Parser.init();
const parser = new Parser();
const clojure = await Language.load(
  require.resolve("@yogthos/tree-sitter-clojure/tree-sitter-clojure.wasm")
);
parser.setLanguage(clojure);

const tree = parser.parse("(defn hello [name] (println name))");
console.log(tree.rootNode.toString());
```

## What the Repository Provides

This repository provides some files used to create various artifacts
(e.g. dynamic libraries, WebAssembly modules) used for handling Clojure
and ClojureScript source code via tree-sitter.

Please see the [what and why document](doc/what-and-why.md) for
detailed information.

## Potential Changes

Changes may occur because:

1. There may be unanticipated important use cases we may want to
   account for
2. The grammar depends on tree-sitter which remains in flux (and is
   still pre 1.0)
3. It's possible we missed something or got something wrong about
   Clojure and we might want to remedy that

Note that previously tagged versions may work fine depending on the
use case.  See the [changelog](CHANGELOG.md) for details.

## Other Documents

There are some documents in the [`doc` directory](doc/) covering
topics such as:

* [Scope](doc/scope.md)
* [Limits](doc/limits.md)
* [Testing](doc/testing.md)
* [Uses](doc/use.md)

## Acknowledgments

This project is based on the original [tree-sitter-clojure](https://github.com/sogaiu/tree-sitter-clojure/)
grammar by [sogaiu](https://github.com/sogaiu). All credit for the grammar
design and comprehensive test corpus goes to the original project.

Please see the [credits](doc/credits.md) for additional acknowledgments.

