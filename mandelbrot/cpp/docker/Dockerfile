FROM ubuntu:latest
ENV TZ=America/Denver
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
RUN apt update && apt -y install build-essential cmake git curl unzip tar zip zlib1g-dev libpng-dev libtbb-dev
RUN git clone https://github.com/Microsoft/vcpkg.git /opt/vcpkg
WORKDIR /opt/vcpkg
RUN ./bootstrap-vcpkg.sh && ./vcpkg integrate install && ./vcpkg integrate bash && echo 'export PATH=$PATH:/opt/vcpkg' >>~/.bashrc
RUN git clone https://github.com/pngwriter/pngwriter.git /opt/pngwriter
WORKDIR /opt/pngwriter
RUN mkdir -p build && cd build
RUN cmake ../pngwriter
RUN make -j && make install
RUN git clone -b v2.13.9 https://github.com/catchorg/Catch2 /opt/catch2
WORKDIR /opt/catch2
RUN cmake -Bbuild -H. -DBUILD_TESTING=OFF && cmake --build build/ --target install
WORKDIR /
CMD ["bash"]