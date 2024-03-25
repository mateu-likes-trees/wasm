# WASM

## Concepts

### Sandboxed execution

One of WebAssembly's unique features is that a module cannot access external resources without explicitly importing them, providing strong security guarantees by default. As such, to do anything useful with WebAssembly, it is necessary to wire a module to the host environment, like for example JavaScript and the DOM.

[Source](https://www.assemblyscript.org/concepts.html#sandboxed-execution)

## Sources

### AssemblyScript

Like TypeScript.

#### Setup

node 20, npm 10

```
npm init
npm install --save-dev assemblyscript
npx asinit .
```
Directory structure

```
./assembly
Directory holding the AssemblyScript sources being compiled to WebAssembly.

./assembly/index.ts
Example entry file being compiled to WebAssembly to get you started.

./build
Build artifact directory where compiled WebAssembly files are stored.

./index.html
Starter HTML file that loads the module in a browser.
```

#### Executing JS functions

Create declaration file `assembly/whatever.ts`

```
// This path `../whatever.js` is relative to `build` directory
@external("../whatever.js", "logInteger")
export declare function logInteger(i: i32): void
```

Create implementation file `whatever.js` in root of the repository
```
export function logInteger ( integer ) {
    console.log( 'logIngeger says: ', integer );
}
```

> JavaScript globals in **globalThis** (window - przyp. red.) can be accessed directly via the env module namespace. For example, console.log can be manually imported through:

```
@external("env", "console.log")
declare function consoleLog(s: string): void
```

> Note that this is just an example and console.log is already provided by the standard library when called from an AssemblyScript file. Other global functions not already provided by the standard library may require an import as of this example, though.

### Go

Disadvantages:
- minimal example is 2 MB
- either WASI that doesn't run in the browser
- or the browser version that comes with lots of boilerplate

Worth further investigation: https://tinygo.org/

### Rust

https://developer.mozilla.org/en-US/docs/WebAssembly/Rust_to_Wasm

### C++

https://developer.mozilla.org/en-US/docs/WebAssembly/C_to_Wasm

https://emscripten.org/docs/porting/connecting_cpp_and_javascript/Interacting-with-code.html
https://emscripten.org/docs/tools_reference/settings_reference.html
https://emscripten.org/docs/api_reference/module.html
https://emscripten.org/docs/api_reference/preamble.js.html
https://stackoverflow.com/questions/46332158/assertion-failed-you-need-to-wait-for-the-runtime-to-be-ready-error-when-call
https://stackoverflow.com/questions/58045191/compiling-c-code-to-standalone-wasm-with-emscripten-results-in-wasm-requiring

https://stackoverflow.com/questions/41353389/how-can-i-return-a-javascript-string-from-a-webassembly-function
https://stackoverflow.com/questions/21816960/how-to-pass-strings-between-c-and-javascript-via-emscripten
https://stackoverflow.com/questions/70275690/pass-string-from-c-to-js-using-emscripten
https://www.codeproject.com/Articles/10595/Allocate-your-dynamic-strings-on-the-stack
https://www.reddit.com/r/AskProgramming/comments/107dpaf/c_question_cant_understand_this_error_address_of/

https://stackoverflow.com/questions/70226666/is-there-an-easy-way-to-ask-emscripten-emit-wasm-not-mangling-its-names

https://llama-cpp-wasm.tangledgroup.com/example-single-thread.html

## Runtimes

### Web (browser)

TBD

### Non-web

TBD

## Tools

- https://emscripten.org/
- [w2c2 - WebAssembly modules to portable C](https://github.com/turbolent/w2c2)
- [wasm2c - wasm files to C](https://github.com/WebAssembly/wabt/tree/main/wasm2c)

## Use-cases
- https://webassembly.org/getting-started/developers-guide/
- https://webassembly.org/docs/use-cases/

## Reading list
- https://news.ycombinator.com/item?id=38927960
 - https://github.com/WebAssembly/component-model
 - https://www.fermyon.com/blog/webassembly-component-model
 - https://thenewstack.io/can-webassembly-get-its-act-together-for-a-component-model/


- use same .wasm file for FE and BE for making HTML (components/DS)
- build api with pluggable part with components
