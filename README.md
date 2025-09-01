# Tech with Privacy 💻 🕵️
## Maintaining privacy in the modern world
[https://www.techwithprivacy.com](https://www.techwithprivacy.com)

A collection of casual and unfiltered notes on how to achieve increased privacy online.

Many of the notes below have varying levels of complexity, compromises and efficiency. Given the ever evolving nature of technology, it's difficult to guarantee that anything is foolproof or failsafe, however, the aim is set some strong principles and foundations to then build on as the world evolves. As with everything, there are tradeoffs and in this case it's often convenience. 

While many of the decisions have the goal of privacy in mind, this is a highly opinionated list of notes and it's encouraged to personalise and adjust it to your specific use-cases and setup.

View the full guide [here](https://techwithprivacy.com)

## Contributing
The content in this repo is for informational purposes only and things change (regularly!) so please do your own research and check what’s best for you. If you notice any inaccuracies, please create an issue on GitHub. Contributions are always welcome!
## Local development
It's a Go app that generates static files. Maybe it's overkill for what it is today, but I have hopes to expand the site and add new things.

For now, to spin it up locally, run the following:
```
# Install deps
go mod tidy

# Run web server
air

# To generate templ templates after an update
templ generate
```

