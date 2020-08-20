#pragma once
struct RGB
{
    int r, g, b;
};

class SingleShadeRGB
{
public:
    SingleShadeRGB(int maxIt,RGB outColor, RGB inColor);
    RGB operator()() const;
    RGB operator()(int) const;
private:
    int maxIterations;
    RGB inSetColor;
    RGB colorDelta;

};