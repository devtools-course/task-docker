# Run options

An image `runconf` is built from [Dockerfile](./Dockerfile) with the [runner](./runner) file as the entry point.

Add `docker run` commands to [solution.properties](./solution.properties):
* `fst` - Run the image with the default settings
* `snd` - Run the image with such a configuration that the output is `IN: nfa ast`
* `trd` - Run the image with such a configuration that the output is ` `
* `frt` - Run the image with such a configuration that it calculates the number of files in the container's `PWD`

---
* What is the difference between `CMD` and `ENTRYPOINT`?
* How can a pipeline command be executed in a container?
