import assert from "node:assert";
import { test } from "node:test";
import Parser from "tree-sitter";

test("can load grammar", async () => {
  const parser = new Parser();
  const { default: language } = await import("./index.js");
  parser.setLanguage(language);
  const tree = parser.parse("(+ 1 2)");
  assert.strictEqual(tree.rootNode.type, "source");
  assert.strictEqual(tree.rootNode.hasError, false);
});
