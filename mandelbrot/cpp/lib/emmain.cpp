#include "rgb.h"
#include "mandelbrot.h"

#include <iostream>
#include <chrono>  
#include <SDL/SDL.h>
#include <emscripten.h>

int main()
{
    MandelbrotParameters params{30, -2.0, -1.0, 1.0, 1.0, 1920, 1080};
    RGB outOfSetColor{255, 255, 255};
    RGB inSetColor{0, 0, 0};

    SDL_Init(SDL_INIT_VIDEO);
    SDL_Surface *screen = SDL_SetVideoMode(params.sizeX, params.sizeY, 32, SDL_SWSURFACE);
    int alpha = 255;

    SingleShadeRGB blackToWhite{params.maxIterations, std::move(outOfSetColor), std::move(inSetColor)};
    std::cout << "Computing Mandelbrot\n";
    
    auto start {std::chrono::high_resolution_clock::now()};
    auto rgbResult {ComputeMandelbrot<RGB>(params, blackToWhite, blackToWhite)};
    auto stop {std::chrono::high_resolution_clock::now()};
    auto duration {std::chrono::duration_cast<std::chrono::milliseconds>(stop - start)}; 
    
    // time on laptop was 81 milliseconds
    std::cout << "Done computing Mandelbrot. Time was " << duration.count() << " milliseconds. Plotting\n";

     if (SDL_MUSTLOCK(screen)) SDL_LockSurface(screen);
    for(int y{0}; y < params.sizeY; ++y)
    {
        for(int x{0}; x < params.sizeX; ++x)
        {
            RGB& rgb {rgbResult[y][x]};
            *((Uint32*)screen->pixels + y * params.sizeX + x) = SDL_MapRGBA(screen->format, rgb.r, rgb.g, rgb.b, alpha);
        }
    }
    if (SDL_MUSTLOCK(screen)) SDL_UnlockSurface(screen);
    SDL_Flip(screen); 
    SDL_Quit();

    std::cout << "Done Plotting\n";
}