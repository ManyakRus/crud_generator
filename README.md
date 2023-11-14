The crud_generator application is designed to automatically generate source code files
in golang language to perform CRUD operations.

For each table in the Postgres SQL database, files will be created to perform crud operations:
create, read, update, save, delete (or delete + restore)
files are generated:
1. model - struct structure with all fields from the database, and gorm + db + json annotations
2. db - files for performing crud operations, exchange with the database,
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

Code templates are stored in the bin/templates folder.
The code template is a .go file copied from the real project -
so it’s easy to make a template (change it to your own) - just copy your file.
The code generator replaces part of the code in the template with another code:
- name of the database table
- model name
- adding and removing imports

Installation procedure:
1. Compile this repository
make build
the crud_generator file will appear in the bin folder

2. Fill settings in the bin/settings.txt file
- connections to your database
- name of your new service
- URL of your new service
- and etc.

3. Launch crud_generator
A new folder with the name of your service will appear in the bin folder,
with subfolders and finished files inside.

4. Copy the finished files to your service.
(TODO: later I’ll generate a completely ready-made microservice launched with 1 line of code)


P.S.
I generated myself 170,000 lines of code, from 70 tables in the database, for my service.


Source code in Golang language.
Tested on Linux Ubuntu
Readme from 11/14/2023

Made by Alexander Nikitin
https://github.com/ManyakRus/crud_generator