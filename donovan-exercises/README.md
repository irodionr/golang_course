**Exercise 1.1**: Modify the echo program to also print os.Args[0], the name of the command that invoked it.  
**Exercise 1.2**: Modify the echo program to print the index and value of each of its arguments, one per line.  
**Exercise 1.3**: Experiment to measure the difference in running time between our potentially inefficient versions and the one that uses strings.Join.  
**Exercise 1.4**: Modify dup2 to print the names of all files in which each duplicated line occurs.  
**Exercise 1.5**: Change the Lissajous program’s color palette to green on black, for added authenticity. To create the web color #RRGGBB, use color.RGBA{0xRR, 0xGG, 0xBB, 0xff}, where each pair of hexadecimal digits represents the intensity of the red, green, or blue component of the pixel.  
**Exercise 1.6**: Modify the Lissajous program to produce images in multiple colors by adding more values to palette and then displaying them by changing the third argument of SetColorIndex in some interesting way.  
**Exercise 1.7**: The function call io.Copy(dst, src) reads from src and writes to dst. Use it instead of ioutil.ReadAll to copy the response body to os.Stdout without requiring a buffer large enough to hold the entire stream. Be sure to check the error result of io.Copy.  
**Exercise 1.8**: Modify fetch to add the prefix http:// to each argument URL if it is missing. You might want to use strings.HasPrefix.  
**Exercise 1.9**: Modify fetch to also print the HTTP status code, found in resp.Status.  
**Exercise 1.10**: Find a web site that produces a large amount of data. Investigate caching by running fetchall twice in succession to see whether the reported time changes much. Do you get the same content each time? Modify fetchall to print its output to a file so it can be examined.  
**Exercise 1.11**: Try fetchall with longer argument lists, such as samples from the top million web sites available at alexa.com. How does the program behave if a web site just doesn’t respond?  
**Exercise 1.12**: Modify the Lissajous server to read parameter values from the URL. For example, you might arrange it so that a URL like http://localhost:8000/?cycles=20 sets the number of cycles to 20 instead of the default 5. Use the strconv.Atoi function to convert the string parameter into an integer. You can see its documentation with go doc strconv.Atoi.  
  
**Exercise 2.1**: Add types, constants, and functions to tempconv for processing temperatures in the Kelvin scale, where zero Kelvin is −273.15°C and a difference of 1K has the same magnitude as 1°C.  
**Exercise 2.2**: Write a general-purpose unit-conversion program analogous to cf that reads numbers from its command-line arguments or from the standard input if there are no arguments, and converts each number into units like temperature in Celsius and Fahrenheit, length in feet and meters, weight in pounds and kilograms, and the like.  
**Exercise 2.3**: Rewrite PopCount to use a loop instead of a single expression. Compare the performance of the two versions.  
**Exercise 2.4**: Write a version of PopCount that counts bits by shifting its argument through 64 bit positions, testing the rightmost bit each time. Compare its performance to the table-lookup version.  
**Exercise 2.5**: The expression x&(x-1) clears the rightmost non-zero bit of x. Write a version of PopCount that counts bits by using this fact, and assess its performance.  
  
**Exercise 3.1**: If the function f returns a non-finite float64 value, the SVG file will contain invalid <polygon> elements (although many SVG renderers handle this gracefully). Modify the program to skip invalid polygons.  
**Exercise 3.2**: Experiment with visualizations of other functions from the math package. Can you produce an egg box, moguls, or a saddle?  
**Exercise 3.3**: Color each polygon based on its height, so that the peaks are colored red (#ff0000) and the valleys blue (#0000ff).  
**Exercise 3.4**: Following the approach of the Lissajous example in Section 1.7, construct a web server that computes surfaces and writes SVG data to the client. The server must set the Content-Type header like this: w.Header().Set("Content-Type", "image/svg+xml")  
**Exercise 3.5**: Implement a full-color Mandelbrot set using the function image.NewRGBA and the type color.RGBA or color.YCbCr.  
**Exercise 3.6**: Supersampling is a technique to reduce the effect of pixelation by computing the color value at several points within each pixel and taking the average. The simplest method is to divide each pixel into four "subpixels." Implement it.
