## EPIC FDA 

A tool to download and explore FDA datasets from https://api.fda.gov/download.json

This tool is divided in 2 main tools

### epicFDAFetcher (WIP)

Will fetch the entire dataset and store it in the destination folder 


### epicFDAFetcher-server.v1  (WIP)

Will Run an https api currently with the following Enpoints

Both binaries relay on config.json to define Webaddress, storage folder, log file and SSL certificates.


All files will be downloaded in config.datasetfolder destination order by Library / Type / Date / DisplayName 
It has a control (rudimentary) to pick up where it left (in case of failure) So we can download in parts.

It's not using any go routine at this moment so, in case more speed is needed we can maximize download speed.
No Sharding has been setup. 
This is a basic implementation to download the entire dataset to play with.
