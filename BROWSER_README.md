_This is one of the steps in the Truss interview process. If you've
stumbled upon this repository and are interested in a career with
Truss, [check out our jobs page](https://truss.works/jobs)._

# The problem: Display data from an API for human consumption

## Introduction

Approach this solution the way you would a real world problem. Use
libraries where it makes sense to, and be prepared to explain your
thought process if you do. We expect this to take less than 4 hours of actual coding time. Please
submit a working but incomplete solution instead of spending more time
on it.

If you have any questions, please contact [hiring@truss.works](mailto:hiring@truss.works); we're
happy to help.

## How to submit your response

Please submit your solution by emailing [hiring@truss.works](mailto:hiring@truss.works) a link to **one** of the following:

-   A public git repository (Github is fine) that contains your code and a `README.md` that tells us how to build and run it.
-   Your solution running in an online IDE such as [CodePen](https://codepen.io) or [Glitch](https://glitch.com).

Please ensure that your submission includes everything needed to run and view your project:

-   If your solution code requires compilation (such as with a tool like webpack or create-react-app),
    please provide complete instructions for running your code, including
    a `package.json` file with all dependencies included.
-   If your solution is in the form of static assets (such as an index.html file), we will run it
    using a simple no-config HTTP server (<https://github.com/vercel/serve>).
-   If your solution is on a hosted IDE such as Codepen or Glitch, please make sure to
    provide an evergreen link that does not expire.
    
We will review your solution in a recent version of Firefox, Chrome, or Safari running on MacOS.

## Full Description 

Please write a webpage that loads data from `https://swapi.dev/api/planets/`, formats and displays that data in a table. SWAPI documentation can be found [here](https://swapi.dev/documentation)

Name the entrypoint for your application `index.html`.

You may use additional files to organize your code. You may also use frameworks and build tools as long as clear, simple instructions are provided for how to run them. Please ensure your solution can be easily viewed in the browser. Be ready to answer questions about why you chose a given approach.

Don't worry about making the table look fancy; you can style it however you
would like to alongside the requirements below. We will look at the markup you
generate and any code you write to load or display the data.

For each planet in the dataset, please display:

-   The planet's name
    -   The name should be a link that, when clicked, opens the planet's API URL in a new window
-   The planet's climate
-   How many residents the planet has
-   The terrains found on the planet
-   The population
-   The surface area (in km<sup>2</sup>) covered by water
    -   Assume that all planets are perfect spheres.
    -   The radius of a sphere is half its diameter.
    -   The value of `surface_water` from the API is a percentage, so a value of `50` means the planet is 50% covered in water.
    -   Round these values to the nearest integer value.

Additionally, please satisfy these requirements:

-   Show a loading message while loading the data, and hide this message once the data is displayed.
-   Display an error message if the data load fails for some reason. We may test your code against an invalid URL to demonstrate this.
-   Sort the planets alphabetically
-   Any column that is unknown should be displayed as `?`.
-   Format all numbers larger than 999 with spaces to group digits into groups of three.
    For example, ten thousand should be displayed as `10 000`.
-   Cells in the table should be separated by 1 pixel gray lines.
