# gOCR

*OCR Implementation with Golang*

This project was mainly focus on Thai language and tested with some Thai characters, English characters and Arabic digits

![gOCR](https://i.imgur.com/RjJ9uVg.jpg)

## Implemented functions
- Template generator
- Image adjustment (automatic thresholding and median filter)
- Rows and characters segmentation (horizontal histrogram and blob coloring)
- Character recognition (image resizing and template matching)

## How to use

### Generate the Templates

``` bash
gocr --gentemp <charlist.txt> <fontfile.ttf>
```

Templates of character list in `charlist.txt` with `fontfile.ttf` fontface will be generated in `/templates` directory

### OCR
``` bash
gocr <charlist.txt> <inputImage>
```

The `inputImage` and `charlist.txt` (the same one that use to generate templates) will be use for OCR. The output text file and images of each process will be saved in `/outputs` directory.

## Example
#### Input image
![](https://i.imgur.com/65TqcDN.jpg)

#### Median filter + automatic thresholding
![](https://i.imgur.com/MwzFKSH.jpg)

#### Rows segmentation
![](https://i.imgur.com/fezTbXX.jpg)

#### Characters segmentation
![](https://i.imgur.com/OXrEq5K.jpg)
![](https://i.imgur.com/4usYCcs.jpg)
![](https://i.imgur.com/glv0ziG.jpg)
![](https://i.imgur.com/e4bA7Z2.jpg)
![](https://i.imgur.com/JzwL18K.jpg)

#### Text output
```
234sohappy
sawasdeewansao
ลูกกิดมาล้าว
สุขใจหรรษา
อิอิzaa55บวก
```
NOTE: You can notice some incorrect character recognition here

#### Terminal output
![](https://i.imgur.com/8t6ndrs.png)

## Test cases
No. | Font style | Capture method | Correctness
--- | --- | --- | ---
1 | Same as template | Screen capture | 58/58
2 | Same as template, double font size | Screen capture | 58/58
3 | Same as template | Phone camera | 57/58
4 | Different font | Screen capture | 32/58
5 | Different font | Phone camera | 34/58

## What should be improve
- Support multi-object characters (ex. ะ, ญ)
- Better similar character recognition ( Dictionary? Machine Learning? Character type consideration?)
- A space in between word recognition
- Order recognition for lower and upper alphabet
- Handle system error (ex. can't open file)

## Main Resources
- [Golang](https://golang.org) : an open source programming language 
- [GoCV](https://gocv.io) :  allow Go to access OpenCV 3
- [OpenCV](http://opencv.org/) : a computer vision library

## About the project
Created by Withee Poositasai

Simple Thai OCR Project, CPE489 Image Processing and Computer Vision, KMUTT
