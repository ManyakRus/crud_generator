The crud_generator application is designed to automatically generate source code files
in golang language to perform CRUD operations.
A completely ready-made microservice is created that can be launched immediately.

For each table in the Postgres SQL database, files will be created to perform crud operations:
- Create(), Read(), Update(), Delete() (or Delete() + Restore(), if there is an is_deleted field)
- Save() - creating a new record when ID=0, or updating a record when ID<>0
- ReadFromCache() - reading from cache or database
- Update_ColumnName() - changing only one column with the name ColumnName,
separate function for each column of each table

files are generated:
1. table - struct structure with all fields from the database, and gorm + json annotations
   Tables with fields in the database must be created manually in advance.
   The code generator will find all the tables and their fields in the database and use them.
2. model - struct structure, including table, with crud operations functions
2. crud - files for performing crud operations, exchange with the database,
   as well as files with tests
3. grpc server - files for performing crud operations over the network, using the GRPC protocol,
   as well as files with tests
4. grpc client - client files for using GRPC by third-party services,
   as well as files with tests
5. nrpc server - files for performing crud operations over the network, using the NRPC protocol (via the NATS message broker),
   as well as files with tests
6. nrpc client - client files for use by NRPC third-party services,
   as well as files with tests
7. crud_starter - a file with functions for switching to the desired protocol db or grpc or nrpc
8. main.go and other .go files
9. Makefile - script launch configuration file
10. .env - file for filling in the microservice configuration (database connection parameters, etc.)


Code templates are stored in the bin/templates folder.
The code template is a .go file copied from the real project -
so it’s easy to make a template (change it to your own) - just copy your file.
The code generator replaces part of the code in the template with another code:
- name of the database table
- model name
- adding and removing imports

Installation procedure:
1. Compile this repository
>make build
>
the crud_generator file will appear in the bin folder

2. Fill in the settings in the file bin/templates/configs_/settings.txt
- connections to your database
- the name of your new service
- URL of your new service
- and etc.

3. Tables in the database must be created in advance, manually.

4. Launch crud_generator
A new folder with the name of your service will appear in the bin folder,
with subfolders and finished files inside.

5. Fill in the configuration parameters in the .env file
Start microservice:
>make run
>



P.S.
I generated myself 350,000 lines of code, from 70 tables in the database, for my service.


Source code in Golang language.
Tested on Linux Ubuntu
Readme from 11/14/2023

Made by Alexander Nikitin
https://github.com/ManyakRus/crud_generator
