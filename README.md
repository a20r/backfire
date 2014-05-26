Backfire
========

Framework for Automated Robotics Testing

# About
Backfire is an automated testing framework for motion planning algorithms that strives to standardize the format of scenes used for experimentation and allow for easy comparison between algorithms. Backfire will be available to host scenes used for experimentation and provide a RESTful API for retrieval and testing. The idea is to provide bindings that will automatically requests scenes for testing, run your motion planning algorithm on the scenes, and send back the results to the server. We hope to create an easy way to visualize the data that is sent back. This may not only be the amount of time the algorithm takes to complete, but anything you implement in your own statistics package. We hope to make Backfire as extensible as possible so it doesn't disrupt your own testing patterns. This means that adding it to your project will be a simple API call. Also, we will add automatic visualization of the scenes and your algorithm using the ROS package, RViz. Backfire, will be a simple addition to any motion planning project that will allow for automated testing, visualization, and benchmarking of standardized scenes used throughout academia and the industry.

# Authors
- Alexander Wallar :: <aw204@st-andrews.ac.uk>