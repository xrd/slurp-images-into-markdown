Golang tool to slurp in image files and create a markdown story for it. Run it inside the 
directory where you want to create the markdown file and it will copy the files (or resize
them for you). 

```bash
$ export PATH=$PATH:$GOPATH/bin
$ slurp-images-into-markdown --convert --format ~/Desktop/Screen\ Shot\ 2016-12-01
## 0

![](0.resized.png)

## 1

![](1.resized.png)

## 2
...
$ slurp-images-into-markdown --convert --format ~/Desktop/Screen\ Shot\ 2016-12-01 > explanation.md
```

Then, edit the `explanation.md` file and add headers and descriptions of the images inline.

Full usage:

```bash
Usage of slurp-images-into-markdown:
  -convert
    	Resize files
  -format string
    	Filename glob, e.g.: ~/Desktop/Screen\ Shot\ 2016-12-01
  -size int
    	Resize width (default 800)
```
