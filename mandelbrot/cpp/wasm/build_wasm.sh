em++ -std=c++17 -O3 -s TOTAL_MEMORY=100MB wasm_driver.cpp ../lib/rgb.cpp ../lib/mandelbrot.cpp -sEXPORTED_FUNCTIONS=_simplified_mandelbrot -o mandelbrot.js