SELECT 'CREATE DATABASE chonho'
    WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'elotus');\gexec

SELECT 'CREATE DATABASE test'
    WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'elotus');\gexec

