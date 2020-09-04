#include "../lib/rgb.h"
#include "../lib/mandelbrot.h"

#include "pngwriter.h"

#include <iostream>
#include <chrono>  

int main()
{
    MandelbrotParameters params{30, -2.0, -1.0, 1.0, 1.0, 1920, 1080};
    RGB outOfSetColor{65535, 65535, 65535};
    RGB inSetColor{0, 0, 0};

    SingleShadeRGB blackToWhite{params.maxIterations, std::move(outOfSetColor), std::move(inSetColor)};
    std::cout << "Computing Mandelbrot\n";
    
    auto start {std::chrono::high_resolution_clock::now()};
    auto rgbResult {ComputeMandelbrot<RGB>(params, blackToWhite, blackToWhite)};
    auto stop {std::chrono::high_resolution_clock::now()};
    auto duration {std::chrono::duration_cast<std::chrono::milliseconds>(stop - start)}; 
    
    // time on laptop was 81 milliseconds
    std::cout << "Done computing Mandelbrot. Time was " << duration.count() << " milliseconds. Plotting\n";

    pngwriter outFile(params.sizeX, params.sizeY, 0, "out.png"); 

    for(int y{0}; y < params.sizeY; ++y)
    {
        for(int x{0}; x < params.sizeX; ++x)
        {
            RGB& rgb {rgbResult[y][x]};
            outFile.plot(x, y, rgb.r, rgb.g, rgb.b);
        }
    }

    std::cout << "Done Plotting\n";
    outFile.close();
}