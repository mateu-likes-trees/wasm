build:
	emcc main.cpp --no-entry -O3 -s EXPORTED_FUNCTIONS=_add,_greet -s EXPORTED_RUNTIME_METHODS=ccall,cwrap,UTF8ToString --extern-post-js main-extend.js -o main.js

serve:
	php -S localhost:8000
