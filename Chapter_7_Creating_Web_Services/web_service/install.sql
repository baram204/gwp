drop database gwp;
create database gwp;
drop user gwp;
create user gwp with password 'gwp';
grant all privileges on all tables in schema public to gwp;
grant all privileges on all sequences in schema public to gwp;