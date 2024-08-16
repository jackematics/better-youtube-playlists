CREATE USER module_metadata_api WITH PASSWORD '<password from Secrets Manager>';
GRANT SELECT, UPDATE, INSERT ON module, module_build TO module_metadata_api;
GRANT USAGE, SELECT ON SEQUENCE module_build_id_seq TO module_metadata_api;