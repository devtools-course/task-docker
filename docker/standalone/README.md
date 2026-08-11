# Standalone binary

You have to create a multi-platform image from the [main.c](./main.c) (C/C++ track) or [Main.java](./Main.java) (Java track) file.

The image must support these platforms:
* x86_64
* arm64

Each platform image size must be less than 0.5 MB.

Upload your image to the GitLab Container Registry.
The image name is expected to be `YOUR_REPO_FULLNAME:java` or `YOUR_REPO_FULLNAME:cxx` depending on which language you've chosen.


---
* How to build a standalone binary for C/C++ or Java code?
* How can multi-platform images be built?
* How can an image be uploaded to a registry?
