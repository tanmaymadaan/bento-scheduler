#!/bin/bash

set -e  # Exit immediately if a command exits with a non-zero status

# Run all .cql files in the migrations directory
for file in /migrations/*.cql
do
    echo "Applying migration: $file"
    if cqlsh cassandra-node1 -f "$file"; then
        echo "Migration $file applied successfully"
    else
        echo "Error applying migration $file"
        exit 1
    fi
done

echo "Migrations completed successfully"