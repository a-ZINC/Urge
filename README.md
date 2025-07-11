# Urge

Urge is a high-performance, concurrent command-line tool for batch image processing. Built with Go, it leverages a parallel pipeline architecture to efficiently apply various transformations to images from local files, directories, or remote URLs.

## Features

*   **Concurrent Processing**: Utilizes a multi-stage pipeline with goroutines and channels to process multiple images in parallel, maximizing CPU usage.
*   **Flexible Input**: Process a single image, a comma-separated list of images, an entire directory, or remote images from URLs.
*   **Core Transformations**:
    *   **Flip**: Horizontally or vertically.
    *   **Rotate**: 90, 180, or 270-degree increments.
    *   **Filter**: Apply a grayscale effect.
*   **Dockerized**: A lightweight Docker image is available for easy, dependency-free execution.

## Installation

### Using Docker

A pre-built image is available on Docker Hub. This is the recommended way to run Urge without installing Go or other dependencies.

```bash
docker pull azinc2828/urge:latest
```
*(Note: Replace `latest` with a specific SHA tag for a stable version if needed.)*

### From Source

Ensure you have Go version 1.24+ installed.

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/a-zinc/urge.git
    cd urge
    ```

2.  **Build the binary:**
    ```bash
    go build -o urge .
    ```
    You can now run the tool using `./urge`.

## Usage

The primary command is `process`, which takes an input source and a set of transformation flags.

### Basic Command Structure

```bash
urge process --input <path_or_url> [flags]
```

By default, processed images are saved to a new directory named `transform_<current_directory_name>` in your current working directory. You can specify a different location with the `--output` flag.

### Examples

**1. Rotate and apply a grayscale filter to a single local image:**

```bash
urge process --input ./images/cat.jpg --output ./processed --rotate 90 --filter grayscale
```

**2. Flip multiple local images horizontally:**

```bash
urge process --input "./images/dog.png, ./images/bird.png" --flip h
```

**3. Process an entire directory of images:**

```bash
urge process --input ./source_images/ --output ./flipped_images --flip v
```

**4. Fetch and process a remote image:**

```bash
urge process --input "https://www.golang-book.com/public/img/gopher.png" --rotate 180
```

### Using Docker for Local Files

To process local files with the Docker container, you must mount your input and output directories as volumes.

```bash
# Create input and output directories
mkdir my_input my_output

# Place images in my_input directory...

# Run the container
docker run --rm \
  -v $(pwd)/my_input:/app/input \
  -v $(pwd)/my_output:/app/output \
  azinc2828/urge process --input /app/input --output /app/output --filter grayscale
```

## Command-Line Flags

| Flag                | Shorthand | Description                                                                                             | Accepted Values                                              |
| ------------------- | --------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------ |
| `--input`           | `-i`      | **(Required)** Path to a file, directory, or comma-separated list of files. Can also be an HTTP/S URL.  | `string`                                                     |
| `--output`          | `-o`      | Path to the output directory where processed images will be saved.                                      | `string`                                                     |
| `--flip`            | `-s`      | Flip the image.                                                                                         | `h`, `H`, `horizontal` (horizontal), `v`, `V`, `vertical` (vertical) |
| `--rotate`          | `-t`      | Rotate the image by a specified degree.                                                                 | `90`, `180`, `270` (and negative equivalents like `-90`)       |
| `--filter`          | `-f`      | Apply a filter to the image.                                                                            | `grayscale`                                                  |
| `--resize`          | `-r`      | *Note: This feature is currently a placeholder and does not perform a resize operation.*                | `string`                                                     |
| `--file`            | `-x`      | Path to a file containing a list of image URLs to process.                                              | `string`                                                     |
| `--main`            | `-y`      | Use flags specified alongside URLs in the file provided with `--file`.                                  | `bool`                                                       |

## Architecture

Urge is built on a concurrent pipeline model. When invoked, it sets up a series of stages connected by Go channels.

`Fetch` -> `Flip` -> `Rotate` -> `Resize` -> `Filter` -> `Save`

1.  **Producer**: The `Fetch` stage reads images from the specified input (files, directories, URLs) and places them onto a channel.
2.  **Consumers**: Each subsequent stage (`Flip`, `Rotate`, etc.) runs a pool of worker goroutines. These workers consume images from an upstream channel, apply their specific transformation, and pass the result to the next channel in the pipeline.
3.  **Pipeline Flow**: An image only proceeds to a stage if the corresponding flag is set (e.g., an image only enters the `rotateChannel` if `--rotate` was used). If a transformation is not requested, the image bypasses that stage and is forwarded to the next.
4.  **Save**: The final `Save` stage consumes the fully processed image and writes it to the disk.

This design allows multiple images to be in different stages of processing simultaneously, leading to high throughput for batch operations.
