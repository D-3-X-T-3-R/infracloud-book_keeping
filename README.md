# infracloud-book_keeping <br />

## Library
Simple REST API to store list of books and fetch all books using Golang and PostgreSQL <br /> <br />
## Setup <br />
### Postgres <br />
Follow the commands mentioned in [Install Postgres ubuntu](https://www.postgresql.org/download/linux/ubuntu/) <br />
Modify the config.json file according to username password mentioned during installation <br />

### Starting services <br />
Modify the config.json path in run.sh script according to your systems absolute path <br />
give chmod 777 permission to the script <br />
execute the script by typing ./run.sh and the server will start <br /> <br />

##  Server and End points <br />
The services will run on localhost:9000 <br />
The end points available are : <br />
* List all books (/book/) Method GET
* Add books (/book/) Method POST