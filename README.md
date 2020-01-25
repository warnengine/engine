# Warnengine

## Introduction

Currently, Warnengine is a project written just to learn the fundamentals of OpenGL. It aims to become a real time strategy game with some touch of the moba genre.

## Build it

Warnengine is working on most linux distro. You just have to install some dependencies first.

On fedora platform :

`sudo dnf install glew-devel freetype-devel make gcc-c++ libXcursor-devel libXinerama-devel libXinerama-devel libXrandr-devel`

On debian platform :

`sudo apt-get install build-essential libgl1-mesa-dev libxcursor-dev libxinerama-dev libxinerama-dev libxrandr-dev`

Fetch the assets :

Next to your `engine` source code folder, clone the warnengine/data repository and zip all folders into `public.zip`.

Then you can build it :

`make install-linux`

`make linux`

## Milestones

* Game engine
    * Visual
        * [x] Open a window
        * [x] Draw mesh
        * [x] Color mesh
        * [x] Texture mesh
        * [x] Camera
        * [x] Camera movement
        * [x] Transform mesh
        * [x] Load from `.obj`
        * [x] Directional light
        * [x] Ambient light
        * [x] Shadow mapping
    * Sound
        * [ ] Play music
        * [ ] Play 3D sound
    * Mods
        * [x] Reading assets from zip
        * [x] Reading from multiple zip
        * [x] Overwrite assets by adding zip
    * UI
        * [x] Draw text
        * [ ] Button
        * [ ] Input text
        * [ ] Check box

## Improvements

* Meshes are currently drawn in the old way. There isn't any indexed VBO.