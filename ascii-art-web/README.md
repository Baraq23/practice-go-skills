# ascii-art-web-export-file

This is a continuation project of the ascii-art-web project that consists on making sure that it is possible to export the output of the web application through downloading the output in a text file format.

## Project Overview

![](/imgs/preview.png)

## Getting Started

To get started with the project, download and install Go programming language using the link: https://go.dev/doc/install

## Prerequisites

Before interacting with the project, make sure that your local machine has docker installed and the docker engine is running. Use the command below to confirm that the engine is running:
```bash
docker version
```


## Installation

1. Clone the repository from its remote location to your local environment with the command below:
```bash
git clone https://learn.zone01kisumu.ke/git/barraotieno/ascii-art-web-export-file
```

2. Navigate to the project's path
```bash
cd ascii-art-web-export-file
```

## Usage

Before running the program, the user needs to build an image of the docker file in their local machine. An image file can be built using the command below:
```bash
docker build -t [SPECIFY-IMAGE-NAME] .
```

Once an image file has been built, run the container with the command below:
```bash
docker run -p 9000:9000 [IMAGE-NAME]
```

or

```bash
docker run -d -p 9000:9000 [IMAGE-NAME]
```

- Example

1. Build an image out of the docker file using the below command:
```bash
docker build -t ascii-art-web-export-file .
```
![](/imgs/build-image.png)

2. List the images to confirm the image with the specified name has been built using the command below:
```bash
docker images
```
![](/imgs/list-images.png)

3. You will know that the build process was succesfull if the specified name appears on the listed images. You can now run a container from the image using the below command:
```bash
docker run -p 9000:9000 ascii-art-web-export-file
```
![](/imgs/run-container.png)

or

```bash
docker run -d -p 9000:9000 ascii-art-web-export-file
```

**NOTE:** Using the second command to run a container makes the container run on detached mode. This means that the container won't stop running even if the terminal is stopped and may therefore affect perfomance of the host machine if left to run for a long time. Use the command below to view the status of your containers:
```bash
docker ps -a
```

4. If successfully running, a feedback is displayed to the user informing them that the server is successfully running at port `http://localhost:9000` as shown below.

From the illustrated example above, our server is running on port `:9000` and can be accessed by accessing the url `http://localhost:9000` on the web browser or clicking on the link using `CTRL + Left Click` combination.

From this point, the user can interact with the program in a variety of ways including:
- Providing input using the textarea

- Selecting their preferred banner file

- Submitting their input for processing

- Change their preferred display mode e.g light/dark mode

## Examples

1. Type `hello` in the textarea. Click on the select tag and choose a different banner file e.g `Standard` and click on the Submit button. Below should be the expected output.

![](/imgs/preview.png)

### Export File

1. Once an art has been generated, a download button appears below the generated art. Click on the button to have your output written/exported to a file.
**NOTE:** The file name is determined based on `ascii-art-[randomly generated letter]` eg. `ascii-art-rne`. This is to minimalize the chances of 2 or more file bearing the same name. This is as shown in the image below:

![](/imgs/download-preview.png)

## Implementation

The project's implementation uses the same implementation of the original ascii-art project. This therefore affects how input passed by the uses is displayed back to them in the mentioned cases below:

1. Occurrences of `"\\a", "\\b", "\\v", "\\f", "\\r"` or any other character absent in the range of `32 to 126` present in the input will result to an error message displayed to the user.

### HTTP Endpoints

The project uses an HTTP method `GET` to retrieve users input and uses HTTP method `POST` to display information from the web server back to the user using the `ascii-art` path. Incase of any endpoint failure, an HTTP status code will be displayed to the user. Below are some error codes and what may result them:

- **400**

Results from an invalid character passed on the input.

- **404**

Results from an invalid URL accessed or incase a source file is deleted.

- **405**

Results from an inappropriate method accessed.

- **500**

Results if any banner file has been tampered with.

## Authors

The project was a group effort that involved the following members:
- [Godwin Ouma](https://github.com/garveyshah)

- [Barrack Otieno](https://github.com/Baraq23)

- [John Odhiambo](https://github.com/johneliud)

## Acknowledgement

- [Zone01 Kisumu](https://www.zone01kisumu.ke)
